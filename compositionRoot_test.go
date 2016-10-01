package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang/adapters"
)

var _ = Describe("compositionRoot", func() {

  fakeEngineHost := new(adapters.FakeEngineHost)

  Context("CoreApi()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualTcpApi := objectUnderTest.CoreApi()

      /* assert */
      Expect(actualTcpApi).ShouldNot(BeNil())

    })

  })

})
