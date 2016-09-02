package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
)

var _ = Describe("createCollectionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke opspecSdk.CreateCollection with expected args", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)

      expectedCollectionName := "dummyCollectionName"

      expectedReq := *opspecModels.NewCreateCollectionReq(
        path.Join(workDirPath, ".opspec", expectedCollectionName),
        expectedCollectionName,
        "dummyCollectionDescription",
      )

      objectUnderTest := newCreateCollectionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      objectUnderTest.Execute(expectedReq.Description, expectedReq.Name)

      /* assert */
      Expect(fakeOpspecSdk.CreateCollectionArgsForCall(0)).Should(Equal(expectedReq))
    })
    It("should return error from opspecSdk.CreateCollection", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)
      expectedError := errors.New("dummyError")
      fakeOpspecSdk.CreateCollectionReturns(expectedError)

      objectUnderTest := newCreateCollectionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      actualError := objectUnderTest.Execute("", "")

      /* assert */
      Expect(actualError).Should(Equal(expectedError))
    })
  })
})
