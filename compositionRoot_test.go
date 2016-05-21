package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("compositionRoot", func() {
  Context("Sdk", func() {
    It("should return a sdk.Client instance", func() {

      /* arrange */
      objectUnderTest, err := newCompositionRoot()
      if (nil != err) {
        Fail(err.Error())
      }

      /* act */
      actualDevOpSpecSdk := objectUnderTest.DevOpSpecSdk()

      /* assert */
      Expect(actualDevOpSpecSdk).ToNot(BeNil())

    })
  })
})
