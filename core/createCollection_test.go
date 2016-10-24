package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
)

var _ = Describe("createCollectionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke bundle.CreateCollection with expected args", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)

      expectedCollectionName := "dummyCollectionName"

      expectedReq := *opspecModels.NewCreateCollectionReq(
        path.Join(workDirPath, expectedCollectionName),
        expectedCollectionName,
        "dummyCollectionDescription",
      )

      objectUnderTest := _api{
        bundle:fakeBundle,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.CreateCollection(expectedReq.Description, expectedReq.Name)

      /* assert */
      Expect(fakeBundle.CreateCollectionArgsForCall(0)).Should(Equal(expectedReq))
    })
    It("should return error from bundle.CreateCollection", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)
      expectedError := errors.New("dummyError")
      fakeBundle.CreateCollectionReturns(expectedError)

      fakeExiter := new(fakeExiter)

      objectUnderTest := _api{
        bundle:fakeBundle,
        exiter:fakeExiter,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.CreateCollection("", "")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:expectedError.Error(), Code:1}))
    })
  })
})
