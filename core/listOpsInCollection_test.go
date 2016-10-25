package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "path"
  "errors"
  "github.com/opspec-io/sdk-golang/models"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
  "os"
)

var _ = Describe("listOpsInCollectionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke bundle.GetCollection with expected args", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)

      expectedPath := path.Join(workDirPath, ".opspec")

      objectUnderTest := _api{
        bundle:fakeBundle,
        workDirPathGetter:fakeWorkDirPathGetter,
        writer:os.Stdout,
      }

      /* act */
      objectUnderTest.ListOpsInCollection()

      /* assert */

      Expect(fakeBundle.GetCollectionArgsForCall(0)).Should(Equal(expectedPath))
    })
    It("should return errors from bundle.GetCollection", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)
      expectedError := errors.New("dummyError")
      fakeBundle.GetCollectionReturns(models.CollectionView{}, expectedError)

      fakeExiter := new(fakeExiter)

      objectUnderTest := _api{
        bundle:fakeBundle,
        exiter:fakeExiter,
        workDirPathGetter:fakeWorkDirPathGetter,
        writer:os.Stdout,
      }

      /* act */
      objectUnderTest.ListOpsInCollection()

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:expectedError.Error(), Code:1}))
    })
  })
})
