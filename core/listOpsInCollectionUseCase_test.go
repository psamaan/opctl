package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang"
  "path"
  "errors"
  "github.com/opspec-io/sdk-golang/models"
)

var _ = Describe("listOpsInCollectionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke opspecSdk.GetCollection with expected args", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)

      expectedPath := path.Join(workDirPath, ".opspec")

      objectUnderTest := newListOpsInCollectionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter, new(fakeWriter))

      /* act */
      objectUnderTest.Execute()

      /* assert */

      Expect(fakeOpspecSdk.GetCollectionArgsForCall(0)).Should(Equal(expectedPath))
    })
    It("should return errors from opspecSdk.GetCollection", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)
      expectedError := errors.New("dummyError")
      fakeOpspecSdk.GetCollectionReturns(models.CollectionView{}, expectedError)

      objectUnderTest := newListOpsInCollectionUseCase(fakeOpspecSdk, fakeWorkDirPathGetter, new(fakeWriter))

      /* act */
      actualError := objectUnderTest.Execute()

      /* assert */

      Expect(actualError).Should(Equal(expectedError))
    })
  })
})
