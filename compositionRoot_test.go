package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang/pkg/engineprovider/providers/fake"
)

var _ = Describe("compositionRoot", func() {

  fakeEngineProvider := new(fake.EngineProvider)

  Context("CoreApi()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineProvider)

      /* act */
      actualTcpApi := objectUnderTest.CoreApi()

      /* assert */
      Expect(actualTcpApi).ShouldNot(BeNil())

    })

  })

})
