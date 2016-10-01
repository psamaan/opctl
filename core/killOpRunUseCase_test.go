package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "errors"
  "github.com/opspec-io/sdk-golang"
  "github.com/opspec-io/sdk-golang/models"
)

var _ = Describe("killOpRunUseCase", func() {
  Context("Execute", func() {
    It("should invoke opspecSdk.KillOpRun with expected args", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)

      expectedReq := models.KillOpRunReq{
        OpRunId:"dummyOpRunId",
      }

      objectUnderTest := newKillOpRunUseCase(fakeOpspecSdk)

      /* act */
      objectUnderTest.Execute(expectedReq.OpRunId)

      /* assert */

      Expect(fakeOpspecSdk.KillOpRunArgsForCall(0)).Should(BeEquivalentTo(expectedReq))
    })
    It("should return error from opspecSdk.KillOpRun", func() {
      /* arrange */
      expectedError := errors.New("dummyError")
      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.KillOpRunReturns(expectedError)

      objectUnderTest := newKillOpRunUseCase(fakeOpspecSdk)

      /* act */
      actualError := objectUnderTest.Execute("")

      /* assert */

      Expect(actualError).Should(Equal(expectedError))
    })
  })
})
