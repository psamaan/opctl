package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/engine-sdk-golang"
  "errors"
  "github.com/opspec-io/engine/core/models"
)

var _ = Describe("killOpRunUseCase", func() {
  Context("Execute", func() {
    It("should invoke opspecSdk.KillOpRun with expected args", func() {
      /* arrange */
      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)

      expectedReq := models.NewKillOpRunReq("dummyOprunId")

      objectUnderTest := newKillOpRunUseCase(fakeOpctlEngineSdk)

      /* act */
      objectUnderTest.Execute(expectedReq.OpRunId)

      /* assert */

      Expect(fakeOpctlEngineSdk.KillOpRunArgsForCall(0)).Should(BeEquivalentTo(*expectedReq))
    })
    It("should return error from opspecSdk.KillOpRun", func() {
      /* arrange */
      expectedError := errors.New("dummyError")
      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      fakeOpctlEngineSdk.KillOpRunReturns("", expectedError)

      objectUnderTest := newKillOpRunUseCase(fakeOpctlEngineSdk)

      /* act */
      actualError := objectUnderTest.Execute("")

      /* assert */

      Expect(actualError).Should(Equal(expectedError))
    })
  })
})
