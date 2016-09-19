package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "errors"
  "github.com/opspec-io/sdk-golang"
  "github.com/opspec-io/sdk-golang/models"
)

var _ = Describe("streamEventsUseCase", func() {
  Context("Execute", func() {
    It("should invoke opspecSdk.GetEventStream", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)

      fakeOpspecSdk := new(opspec.FakeSdk)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newStreamEventsUseCase(fakeExiter, fakeOpspecSdk)

      /* act */
      objectUnderTest.Execute()

      /* assert */
      Expect(fakeOpspecSdk.GetEventStreamCallCount()).Should(Equal(1))

    })
    It("should call exiter with expected args when opspecSdk.GetEventStream returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.GetEventStreamReturns(nil, returnedError)

      objectUnderTest := newStreamEventsUseCase(fakeExiter, fakeOpspecSdk)

      /* act */
      objectUnderTest.Execute()

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    It("should call exiter with expected args when event channel closes unexpectedly", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)

      fakeOpspecSdk := new(opspec.FakeSdk)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newStreamEventsUseCase(fakeExiter, fakeOpspecSdk)

      /* act */
      objectUnderTest.Execute()

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:"Event channel closed unexpectedly", Code:1}))
    })
  })
})
