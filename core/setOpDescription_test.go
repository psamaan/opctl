package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
)

var _ = Describe("setOpDescriptionUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke bundle.SetOpDescription with expected args", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)

      expectedOpName := "dummyOpName"

      expectedReq := *opspecModels.NewSetOpDescriptionReq(
        path.Join(workDirPath, ".opspec", expectedOpName),
        "dummyOpDescription",
      )

      objectUnderTest := _api{
        bundle:fakeBundle,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.SetOpDescription(expectedReq.Description, expectedOpName)

      /* assert */

      Expect(fakeBundle.SetOpDescriptionArgsForCall(0)).Should(Equal(expectedReq))
    })
    It("should return errors from bundle.SetOpDescription", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)
      expectedError := errors.New("dummyError")
      fakeBundle.SetOpDescriptionReturns(expectedError)

      fakeExiter := new(fakeExiter)

      objectUnderTest := _api{
        bundle:fakeBundle,
        exiter:fakeExiter,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.SetOpDescription("", "")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:expectedError.Error(), Code:1}))
    })
  })
})
