package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
)

var _ = Describe("setOpDescriptionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke opspecSdk.SetOpDescription with expected args", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)

      expectedOpName := "dummyOpName"

      expectedReq := *opspecModels.NewSetOpDescriptionReq(
        path.Join(workDirPath, ".opspec", expectedOpName),
        "dummyOpDescription",
      )

      objectUnderTest := newSetOpDescriptionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      objectUnderTest.Execute(expectedReq.Description, expectedOpName)

      /* assert */

      Expect(fakeOpspecSdk.SetOpDescriptionArgsForCall(0)).Should(Equal(expectedReq))
    })
    It("should return errors from opspecSdk.SetOpDescription", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)
      expectedError := errors.New("dummyError")
      fakeOpspecSdk.SetOpDescriptionReturns(expectedError)

      objectUnderTest := newSetOpDescriptionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      actualError := objectUnderTest.Execute("", "")

      /* assert */

      Expect(actualError).Should(Equal(expectedError))
    })
  })
})
