// Code generated by counterfeiter. DO NOT EDIT.
package internal

import (
	"sync"

	"github.com/mk29142/pooled-reverse-geocode/client"
	"github.com/mk29142/pooled-reverse-geocode/domain"
	"github.com/mk29142/pooled-reverse-geocode/workpool"
)

type FakeClient struct {
	PostcodeStub        func(domain.Coordinates) (client.LatLongPostcode, error)
	postcodeMutex       sync.RWMutex
	postcodeArgsForCall []struct {
		arg1 domain.Coordinates
	}
	postcodeReturns struct {
		result1 client.LatLongPostcode
		result2 error
	}
	postcodeReturnsOnCall map[int]struct {
		result1 client.LatLongPostcode
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Postcode(arg1 domain.Coordinates) (client.LatLongPostcode, error) {
	fake.postcodeMutex.Lock()
	ret, specificReturn := fake.postcodeReturnsOnCall[len(fake.postcodeArgsForCall)]
	fake.postcodeArgsForCall = append(fake.postcodeArgsForCall, struct {
		arg1 domain.Coordinates
	}{arg1})
	stub := fake.PostcodeStub
	fakeReturns := fake.postcodeReturns
	fake.recordInvocation("Postcode", []interface{}{arg1})
	fake.postcodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) PostcodeCallCount() int {
	fake.postcodeMutex.RLock()
	defer fake.postcodeMutex.RUnlock()
	return len(fake.postcodeArgsForCall)
}

func (fake *FakeClient) PostcodeCalls(stub func(domain.Coordinates) (client.LatLongPostcode, error)) {
	fake.postcodeMutex.Lock()
	defer fake.postcodeMutex.Unlock()
	fake.PostcodeStub = stub
}

func (fake *FakeClient) PostcodeArgsForCall(i int) domain.Coordinates {
	fake.postcodeMutex.RLock()
	defer fake.postcodeMutex.RUnlock()
	argsForCall := fake.postcodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) PostcodeReturns(result1 client.LatLongPostcode, result2 error) {
	fake.postcodeMutex.Lock()
	defer fake.postcodeMutex.Unlock()
	fake.PostcodeStub = nil
	fake.postcodeReturns = struct {
		result1 client.LatLongPostcode
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) PostcodeReturnsOnCall(i int, result1 client.LatLongPostcode, result2 error) {
	fake.postcodeMutex.Lock()
	defer fake.postcodeMutex.Unlock()
	fake.PostcodeStub = nil
	if fake.postcodeReturnsOnCall == nil {
		fake.postcodeReturnsOnCall = make(map[int]struct {
			result1 client.LatLongPostcode
			result2 error
		})
	}
	fake.postcodeReturnsOnCall[i] = struct {
		result1 client.LatLongPostcode
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.postcodeMutex.RLock()
	defer fake.postcodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ workpool.Client = new(FakeClient)
