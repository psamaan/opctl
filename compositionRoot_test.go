package main

import (
  . "github.com/onsi/ginkgo"
  "github.com/opctl/engine-sdk-golang"
  "github.com/opspec-io/sdk-golang"
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
      _, ok := actualOpSpecSdk.(opspec.Sdk)
      if !ok {
        Fail("result not assignable to opspec.Sdk")
      }

    })

  })

  Context("Sdk()", func() {

    It("should return a opctlengine.Sdk instance", func() {

      /* arrange */
      objectUnderTest, err := newCompositionRoot()
      if (nil != err) {
        Fail(err.Error())
      }

      /* act */
      actualOpCtlEngineSdk := objectUnderTest.OpCtlEngineSdk()

      /* assert */
      _, ok := actualOpCtlEngineSdk.(opctlengine.Sdk)
      if !ok {
        Fail("result not assignable to opctlengine.Sdk")
      }

    })

  })

})
