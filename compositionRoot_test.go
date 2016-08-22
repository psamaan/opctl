package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("compositionRoot", func() {

  Context("CreateCollectionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.CreateCollectionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("CreateOpUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.CreateOpUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("KillOpRunUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.KillOpRunUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("ListOpsInCollectionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.ListOpsInCollectionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("RunOpUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.RunOpUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("SetCollectionDescriptionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.SetCollectionDescriptionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("SetOpDescriptionUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.SetOpDescriptionUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })

  Context("StreamEventsUseCase()", func() {

    It("should not return nil", func() {

      /* arrange */
      objectUnderTest := newCompositionRoot()

      /* act */
      actualCli := objectUnderTest.StreamEventsUseCase()

      /* assert */
      Expect(actualCli).ShouldNot(BeNil())

    })

  })
})
