package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
)

var _ = Describe("createOpUseCase", func() {

  fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
  workDirPath := ""
  fakeWorkDirPathGetter.GetReturns(workDirPath)

  Context("Execute", func() {
    It("should invoke bundle.CreateOp with expected args", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)

      expectedOpName := "dummyOpName"

      expectedReq := *opspecModels.NewCreateOpReq(
        path.Join(workDirPath, ".opspec", expectedOpName),
        expectedOpName,
        "dummyOpDescription",
      )

      objectUnderTest := _api{
        bundle:fakeBundle,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.CreateOp(expectedReq.Description, expectedOpName)

      /* assert */

      Expect(fakeBundle.CreateOpArgsForCall(0)).Should(Equal(expectedReq))

    })
    It("should return error from bundle.CreateOp", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)
      expectedError := errors.New("dummyError")
      fakeBundle.CreateOpReturns(expectedError)

      fakeExiter := new(fakeExiter)

      objectUnderTest := _api{
        bundle:fakeBundle,
        exiter:fakeExiter,
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.CreateOp("", "")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:expectedError.Error(), Code:1}))

    })
  })
})
