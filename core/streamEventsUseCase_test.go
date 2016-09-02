package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/engine-sdk-golang"
  "github.com/opspec-io/engine-sdk-golang/models"
  "errors"
)

var _ = Describe("streamEventsUseCase", func() {
  Context("Execute", func() {
    It("should invoke opctlEngineSdk.GetEventStream", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newStreamEventsUseCase(fakeExiter, fakeOpctlEngineSdk)

      /* act */
      objectUnderTest.Execute()

      /* assert */
      Expect(fakeOpctlEngineSdk.GetEventStreamCallCount()).Should(Equal(1))

    })
    It("should call exiter with expected args when opctlEngineSdk.GetEventStream returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      fakeOpctlEngineSdk.GetEventStreamReturns(nil, returnedError)

      objectUnderTest := newStreamEventsUseCase(fakeExiter, fakeOpctlEngineSdk)

      /* act */
      objectUnderTest.Execute()

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    It("should call exiter with expected args when event channel closes unexpectedly", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newStreamEventsUseCase(fakeExiter, fakeOpctlEngineSdk)

      /* act */
      objectUnderTest.Execute()

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:"Event channel closed unexpectedly", Code:1}))
    })
  })
})