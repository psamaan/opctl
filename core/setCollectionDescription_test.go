package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
)

var _ = Describe("setCollectionDescriptionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke bundle.SetCollectionDescription with expected args", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)

      expectedReq := *opspecModels.NewSetCollectionDescriptionReq(
        path.Join(workDirPath, ".opspec"),
        "dummyOpDescription",
      )

      objectUnderTest := _api{
        bundle:fakeBundle,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.SetCollectionDescription(expectedReq.Description)

      /* assert */

      Expect(fakeBundle.SetCollectionDescriptionArgsForCall(0)).Should(Equal(expectedReq))
    })
    It("should return errors from bundle.SetCollectionDescription", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)
      expectedError := errors.New("dummyError")
      fakeBundle.SetCollectionDescriptionReturns(expectedError)

      fakeExiter := new(fakeExiter)

      objectUnderTest := _api{
        bundle:fakeBundle,
        exiter:fakeExiter,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.SetCollectionDescription("")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:expectedError.Error(), Code:1}))
    })
  })
})
