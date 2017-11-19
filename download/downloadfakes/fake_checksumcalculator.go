// This file was generated by counterfeiter
package downloadfakes

import (
	"sync"

	"github.com/pivotal-cf/spring-cloud-dataflow-for-pcf-cli-plugin/download/cache"
)

type FakeChecksumCalculator struct {
	CalculateChecksumStub        func(filePath string) (string, error)
	calculateChecksumMutex       sync.RWMutex
	calculateChecksumArgsForCall []struct {
		filePath string
	}
	calculateChecksumReturns struct {
		result1 string
		result2 error
	}
	calculateChecksumReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeChecksumCalculator) CalculateChecksum(filePath string) (string, error) {
	fake.calculateChecksumMutex.Lock()
	ret, specificReturn := fake.calculateChecksumReturnsOnCall[len(fake.calculateChecksumArgsForCall)]
	fake.calculateChecksumArgsForCall = append(fake.calculateChecksumArgsForCall, struct {
		filePath string
	}{filePath})
	fake.recordInvocation("CalculateChecksum", []interface{}{filePath})
	fake.calculateChecksumMutex.Unlock()
	if fake.CalculateChecksumStub != nil {
		return fake.CalculateChecksumStub(filePath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.calculateChecksumReturns.result1, fake.calculateChecksumReturns.result2
}

func (fake *FakeChecksumCalculator) CalculateChecksumCallCount() int {
	fake.calculateChecksumMutex.RLock()
	defer fake.calculateChecksumMutex.RUnlock()
	return len(fake.calculateChecksumArgsForCall)
}

func (fake *FakeChecksumCalculator) CalculateChecksumArgsForCall(i int) string {
	fake.calculateChecksumMutex.RLock()
	defer fake.calculateChecksumMutex.RUnlock()
	return fake.calculateChecksumArgsForCall[i].filePath
}

func (fake *FakeChecksumCalculator) CalculateChecksumReturns(result1 string, result2 error) {
	fake.CalculateChecksumStub = nil
	fake.calculateChecksumReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeChecksumCalculator) CalculateChecksumReturnsOnCall(i int, result1 string, result2 error) {
	fake.CalculateChecksumStub = nil
	if fake.calculateChecksumReturnsOnCall == nil {
		fake.calculateChecksumReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.calculateChecksumReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeChecksumCalculator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.calculateChecksumMutex.RLock()
	defer fake.calculateChecksumMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeChecksumCalculator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cache.ChecksumCalculator = new(FakeChecksumCalculator)
