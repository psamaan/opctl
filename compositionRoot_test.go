package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("compositionRoot", func() {
  Context("OpSpecSdk()", func() {
    It("should return an opspec.Sdk instance", func() {

      /* arrange */
      objectUnderTest, err := newCompositionRoot()
      if (nil != err) {
        Fail(err.Error())
      }

      /* act */
      actualOpSpecSdk := objectUnderTest.OpSpecSdk()

      /* assert */
      Expect(actualOpSpecSdk).ToNot(BeNil())

    })
  })
  Context("Sdk()", func() {
    It("should return a opctlengine.OpCtlEngineSdk instance", func() {

      /* arrange */
      objectUnderTest, err := newCompositionRoot()
      if (nil != err) {
        Fail(err.Error())
      }

      /* act */
      actualOpCtlEngineSdk := objectUnderTest.OpCtlEngineSdk()

      /* assert */
      Expect(actualOpCtlEngineSdk).ToNot(BeNil())

    })
  })
})
