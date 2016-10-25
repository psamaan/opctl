package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "errors"
  "github.com/opspec-io/sdk-golang/models"
  "github.com/opspec-io/sdk-golang/pkg/engineclient"
)

var _ = Describe("killOpRunUseCase", func() {
  Context("Execute", func() {
    It("should invoke engineClient.KillOpRun with expected args", func() {
      /* arrange */
      fakeEngineClient := new(engineclient.FakeEngineClient)

      expectedReq := models.KillOpRunReq{
        OpRunId:"dummyOpRunId",
      }

      objectUnderTest := _api{
        engineClient:fakeEngineClient,
      }

      /* act */
      objectUnderTest.KillOpRun(expectedReq.OpRunId)

      /* assert */

      Expect(fakeEngineClient.KillOpRunArgsForCall(0)).Should(BeEquivalentTo(expectedReq))
    })
    It("should return error from bundle.KillOpRun", func() {
      /* arrange */
      fakeEngineClient := new(engineclient.FakeEngineClient)
      expectedError := errors.New("dummyError")
      fakeEngineClient.KillOpRunReturns(expectedError)

      fakeExiter := new(fakeExiter)

      objectUnderTest := _api{
        engineClient:fakeEngineClient,
        exiter:fakeExiter,
      }

      /* act */
      objectUnderTest.KillOpRun("")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:expectedError.Error(), Code:1}))
    })
  })
})
