package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "os"
)

var _ = Describe("workDirPathGetter", func() {
  Context("Get", func() {
    It("should return current work dir path", func() {
      /* arrange */
      expectedWorkDirPath, err := os.Getwd()
      if (nil != err) {
        Fail(err.Error())
      }

      objectUnderTest := newWorkDirPathGetter()

      /* act */
      actualWorkDirPath := objectUnderTest.Get()

      /* assert */
      Expect(actualWorkDirPath).Should(Equal(expectedWorkDirPath))

    })
  })
})
