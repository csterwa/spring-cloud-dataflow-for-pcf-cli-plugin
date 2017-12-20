package dataflow_test

import (
	. "github.com/pivotal-cf/spring-cloud-dataflow-for-pcf-cli-plugin/dataflow"

	"bytes"
	"io/ioutil"

	"net/http"

	"hash"

	"errors"

	"fmt"

	"crypto/sha256"

	"crypto/sha1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/spring-cloud-dataflow-for-pcf-cli-plugin/httpclient/httpclientfakes"
)

var _ = Describe("DataflowShellDownloadUrl", func() {
	const (
		dataflowServerUrl  = "https://data.flow.server"
		dataflowShellUrl   = "https://data.flow.shell"
		testAccessToken    = "someaccesstoken"
		errMessage         = "Apparently failure was an option after all."
		testSha1Checksum   = "cf23df2207d99a74fbe169e3eba035e633b65d94"
		testSha256Checksum = "9dec3eab5740cb087d7842bcb6bf924f9e008638dedeca16c5336bbc3c0e4453"
	)

	var (
		fakeAuthClient *httpclientfakes.FakeAuthenticatedClient
		payload        string
		testError      error
		getErr         error
		getStatus      int
		downloadUrl    string
		checksum       string
		hashFunc       hash.Hash
		err            error
	)

	BeforeEach(func() {
		fakeAuthClient = &httpclientfakes.FakeAuthenticatedClient{}
		getErr = nil
		getStatus = http.StatusOK
		testError = errors.New(errMessage)
	})

	JustBeforeEach(func() {
		fakeAuthClient.DoAuthenticatedGetReturns(ioutil.NopCloser(bytes.NewBufferString(payload)), getStatus, http.Header{}, getErr)
		downloadUrl, checksum, hashFunc, err = DataflowShellDownloadUrl(dataflowServerUrl, fakeAuthClient, testAccessToken)
	})

	It("should drive the /about endpoint with the supplied access token", func() {
		Expect(fakeAuthClient.DoAuthenticatedGetCallCount()).To(Equal(1))
		aboutUrl, accessToken := fakeAuthClient.DoAuthenticatedGetArgsForCall(0)
		Expect(aboutUrl).To(Equal(dataflowServerUrl + "/about"))
		Expect(accessToken).To(Equal(testAccessToken))
	})

	Context("when driving the /about endpoint returns an error", func() {
		BeforeEach(func() {
			getErr = testError
		})

		It("should propagate the error", func() {
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(": %s", errMessage)))
		})
	})

	Context("when driving the /about endpoint returns an invalid HTTP status", func() {
		BeforeEach(func() {
			getStatus = http.StatusBadGateway
		})

		It("should propagate the error", func() {
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(fmt.Sprintf(": %d", http.StatusBadGateway)))
		})
	})

	Context("when the /about endpoint returns a response reader which cannot be read", func() {
		JustBeforeEach(func() {
			fakeAuthClient.DoAuthenticatedGetReturns(ioutil.NopCloser(badReader{}), getStatus, http.Header{}, getErr)
			downloadUrl, checksum, hashFunc, err = DataflowShellDownloadUrl(dataflowServerUrl, fakeAuthClient, testAccessToken)
		})

		It("should return a suitable error", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("Cannot read dataflow server response body: read error"))
		})
	})

	Context("when the /about endpoint returns invalid JSON", func() {
		BeforeEach(func() {
			payload = "{"
		})

		It("should return a suitable error", func() {
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("Invalid dataflow server response JSON: unexpected end of JSON input, response body: '{'"))
		})
	})

	Context("when the /about endpoint returns a shell download URL", func() {
		BeforeEach(func() {
			payload = fmt.Sprintf(`
				{"versionInfo":
					{"shell":
						{"url": "%s"
						}
					}
				}`, dataflowShellUrl)
		})

		It("should succeed", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the shell download URL", func() {
			Expect(downloadUrl).To(Equal(dataflowShellUrl))
		})
	})

	Context("when the /about endpoint returns a SHA-1 shell checksum", func() {
		BeforeEach(func() {
			payload = fmt.Sprintf(`
				{"versionInfo":
					{"shell":
						{"checksumSha1": "%s"
						}
					}
				}`, testSha1Checksum)
		})

		It("should succeed", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the SHA-1 checksum", func() {
			Expect(checksum).To(Equal(testSha1Checksum))
		})

		It("should return a SHA-1 hash function", func() {
			Expect(hashFunc).To(BeAssignableToTypeOf(sha1.New()))
		})
	})

	Context("when the /about endpoint returns a SHA-256 shell checksum", func() {
		BeforeEach(func() {
			payload = fmt.Sprintf(`
				{"versionInfo":
					{"shell":
						{"checksumSha256": "%s"
						}
					}
				}`, testSha256Checksum)
		})

		It("should succeed", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the SHA-256 checksum", func() {
			Expect(checksum).To(Equal(testSha256Checksum))
		})

		It("should return a SHA-256 hash function", func() {
			Expect(hashFunc).To(BeAssignableToTypeOf(sha256.New()))
		})
	})

	Context("when the /about endpoint returns SHA-1 and SHA-256 shell checksums", func() {
		BeforeEach(func() {
			payload = fmt.Sprintf(`
				{"versionInfo":
					{"shell":
						{"checksumSha1": "%s",
						 "checksumSha256": "%s"
						}
					}
				}`, testSha1Checksum, testSha256Checksum)
		})

		It("should succeed", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the SHA-256 checksum", func() {
			Expect(checksum).To(Equal(testSha256Checksum))
		})

		It("should return a SHA-256 hash function", func() {
			Expect(hashFunc).To(BeAssignableToTypeOf(sha256.New()))
		})
	})
})

type badReader struct{}

func (b badReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}
