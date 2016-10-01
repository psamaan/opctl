package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang/adapters"
)

var _ = Describe("compositionRoot", func() {

  fakeEngineHost := new(adapters.FakeEngineHost)

  Context("CreateCollectionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.CreateCollectionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("CreateOpUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.CreateOpUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("KillOpRunUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.KillOpRunUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("ListOpsInCollectionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.ListOpsInCollectionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("RunOpUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.RunOpUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("SetCollectionDescriptionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.SetCollectionDescriptionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("SetOpDescriptionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.SetOpDescriptionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("StreamEventsUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot(fakeEngineHost)

      /* act */
      actualCli := objectUnderTest.StreamEventsUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })
})
