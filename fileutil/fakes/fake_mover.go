// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/bluebosh/bosh-utils/fileutil"
)

type FakeMover struct {
	MoveStub        func(string, string) error
	moveMutex       sync.RWMutex
	moveArgsForCall []struct {
		arg1 string
		arg2 string
	}
	moveReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMover) Move(arg1 string, arg2 string) error {
	fake.moveMutex.Lock()
	fake.moveArgsForCall = append(fake.moveArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Move", []interface{}{arg1, arg2})
	fake.moveMutex.Unlock()
	if fake.MoveStub != nil {
		return fake.MoveStub(arg1, arg2)
	} else {
		return fake.moveReturns.result1
	}
}

func (fake *FakeMover) MoveCallCount() int {
	fake.moveMutex.RLock()
	defer fake.moveMutex.RUnlock()
	return len(fake.moveArgsForCall)
}

func (fake *FakeMover) MoveArgsForCall(i int) (string, string) {
	fake.moveMutex.RLock()
	defer fake.moveMutex.RUnlock()
	return fake.moveArgsForCall[i].arg1, fake.moveArgsForCall[i].arg2
}

func (fake *FakeMover) MoveReturns(result1 error) {
	fake.MoveStub = nil
	fake.moveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMover) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.moveMutex.RLock()
	defer fake.moveMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMover) recordInvocation(key string, args []interface{}) {
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

var _ fileutil.Mover = new(FakeMover)
