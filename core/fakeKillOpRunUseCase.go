// This file was generated by counterfeiter
package core

import "sync"

type fakeKillOpRunUseCase struct {
  ExecuteStub        func(opRunId string) error
  executeMutex       sync.RWMutex
  executeArgsForCall []struct {
    opRunId string
  }
  executeReturns     struct {
                       result1 error
                     }
  invocations        map[string][][]interface{}
  invocationsMutex   sync.RWMutex
}

func (fake *fakeKillOpRunUseCase) Execute(opRunId string) error {
  fake.executeMutex.Lock()
  fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
    opRunId string
  }{opRunId})
  fake.recordInvocation("Execute", []interface{}{opRunId})
  fake.executeMutex.Unlock()
  if fake.ExecuteStub != nil {
    return fake.ExecuteStub(opRunId)
  } else {
    return fake.executeReturns.result1
  }
}

func (fake *fakeKillOpRunUseCase) ExecuteCallCount() int {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return len(fake.executeArgsForCall)
}

func (fake *fakeKillOpRunUseCase) ExecuteArgsForCall(i int) string {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.executeArgsForCall[i].opRunId
}

func (fake *fakeKillOpRunUseCase) ExecuteReturns(result1 error) {
  fake.ExecuteStub = nil
  fake.executeReturns = struct {
    result1 error
  }{result1}
}

func (fake *fakeKillOpRunUseCase) Invocations() map[string][][]interface{} {
  fake.invocationsMutex.RLock()
  defer fake.invocationsMutex.RUnlock()
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.invocations
}

func (fake *fakeKillOpRunUseCase) recordInvocation(key string, args []interface{}) {
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