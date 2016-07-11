package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
)

var _ = Describe("setCollectionDescriptionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke opspecSdk.SetCollectionDescription with expected args", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)

      expectedReq := *opspecModels.NewSetCollectionDescriptionReq(
        path.Join(workDirPath, ".opspec"),
        "dummyOpDescription",
      )

      objectUnderTest := newSetCollectionDescriptionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      objectUnderTest.Execute(expectedReq.Description)

      /* assert */

      Expect(fakeOpspecSdk.SetCollectionDescriptionArgsForCall(0)).Should(Equal(expectedReq))
    })
    It("should return errors from opspecSdk.SetCollectionDescription", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)
      expectedError := errors.New("dummyError")
      fakeOpspecSdk.SetCollectionDescriptionReturns(expectedError)

      objectUnderTest := newSetCollectionDescriptionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter)

      /* act */
      actualError := objectUnderTest.Execute("")

      /* assert */

      Expect(actualError).Should(Equal(expectedError))
    })
  })
})
