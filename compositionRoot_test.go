package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("compositionRoot", func() {
  Context("CoreApi()", func() {
    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualTcpApi := objectUnderTest.CoreApi()

      /* assert */
      Expect(actualTcpApi).ShouldNot(BeNil())

    })
  })
})
