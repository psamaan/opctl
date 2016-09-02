package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
)

var _ = Describe("createOpUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke opspecSdk.CreateOp with expected args", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)

      expectedOpName := "dummyOpName"

      expectedReq := *opspecModels.NewCreateOpReq(
        path.Join(workDirPath, ".opspec", expectedOpName),
        expectedOpName,
        "dummyOpDescription",
      )

      objectUnderTest := newCreateOpUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      objectUnderTest.Execute(expectedReq.Description, expectedOpName)

      /* assert */

      Expect(fakeOpspecSdk.CreateOpArgsForCall(0)).Should(Equal(expectedReq))

    })
    It("should return error from opspecSdk.CreateOp", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)
      expectedError := errors.New("dummyError")
      fakeOpspecSdk.CreateOpReturns(expectedError)

      objectUnderTest := newCreateOpUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      actualError := objectUnderTest.Execute("", "")

      /* assert */

      Expect(actualError).Should(Equal(expectedError))

    })
  })
})
