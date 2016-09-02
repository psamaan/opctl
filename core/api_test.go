package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("_api", func() {

  Context("CreateCollection", func() {

    Context("with description", func() {
      It("should call compositionRoot.createCollectionUseCase.Execute with expected args", func() {
        /* arrange */
        fakeCreateCollectionUseCase := new(fakeCreateCollectionUseCase)
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeCompositionRoot.CreateCollectionUseCaseReturns(fakeCreateCollectionUseCase)

        expectedCollectionName := "dummyCollectionName"
        expectedCollectionDescription := "dummyCollectionDescription"

        objectUnderTest := &_api{
          compositionRoot:fakeCompositionRoot,
        }

        /* act */
        objectUnderTest.CreateCollection(expectedCollectionDescription, expectedCollectionName)

        /* assert */
        Expect(fakeCreateCollectionUseCase.ExecuteCallCount()).Should(Equal(1))
        actualCollectionDescription, actualCollectionName := fakeCreateCollectionUseCase.ExecuteArgsForCall(0)
        Expect(actualCollectionName).Should(Equal(expectedCollectionName))
        Expect(actualCollectionDescription).Should(Equal(expectedCollectionDescription))
      })
    })

    Context("with no description", func() {
      It("should call compositionRoot.createCollectionUseCase.Execute with expected args", func() {
        /* arrange */
        fakeCreateCollectionUseCase := new(fakeCreateCollectionUseCase)
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeCompositionRoot.CreateCollectionUseCaseReturns(fakeCreateCollectionUseCase)

        expectedCollectionName := "dummyCollectionName"

        objectUnderTest := &_api{
          compositionRoot:fakeCompositionRoot,
        }

        /* act */
        objectUnderTest.CreateCollection("", expectedCollectionName)

        /* assert */
        Expect(fakeCreateCollectionUseCase.ExecuteCallCount()).Should(Equal(1))
        actualCollectionDescription, actualCollectionName := fakeCreateCollectionUseCase.ExecuteArgsForCall(0)
        Expect(actualCollectionName).Should(Equal(expectedCollectionName))
        Expect(actualCollectionDescription).Should(BeEmpty())
      })
    })
  })

  Context("SetCollectionDescription", func() {
    It("should call compositionRoot.setCollectionDescriptionUseCase.Execute with expected args", func() {
      /* arrange */
      fakeSetCollectionDescriptionUseCase := new(fakeSetCollectionDescriptionUseCase)
      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.SetCollectionDescriptionUseCaseReturns(fakeSetCollectionDescriptionUseCase)

      expectedCollectionDescription := "dummyCollectionDescription"

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      objectUnderTest.SetCollectionDescription(expectedCollectionDescription)

      /* assert */
      Expect(fakeSetCollectionDescriptionUseCase.ExecuteCallCount()).Should(Equal(1))
      Expect(fakeSetCollectionDescriptionUseCase.ExecuteArgsForCall(0)).Should(Equal(expectedCollectionDescription))
    })
  })

  Context("events", func() {
    It("should call compositionRoot.streamEventsUseCase.Execute with expected args", func() {
      /* arrange */
      fakeStreamEventsUseCase := new(fakeStreamEventsUseCase)
      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.StreamEventsUseCaseReturns(fakeStreamEventsUseCase)

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      objectUnderTest.StreamEvents()

      /* assert */
      Expect(fakeStreamEventsUseCase.ExecuteCallCount()).Should(Equal(1))
    })
  })

  Context("kill", func() {
    It("should call compositionRoot.killOpRunUseCase.Execute with expected args", func() {
      /* arrange */
      fakeKillOpRunUseCase := new(fakeKillOpRunUseCase)
      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.KillOpRunUseCaseReturns(fakeKillOpRunUseCase)

      expectedOpRunId := "expectedOpRunId"

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      objectUnderTest.KillOpRun(expectedOpRunId)

      /* assert */
      Expect(fakeKillOpRunUseCase.ExecuteCallCount()).Should(Equal(1))
      Expect(fakeKillOpRunUseCase.ExecuteArgsForCall(0)).Should(Equal(expectedOpRunId))
    })
  })

  Context("ls", func() {
    It("should call compositionRoot.listOpsInCollectionUseCase.Execute with expected args", func() {
      /* arrange */
      fakeListOpsInCollectionUseCase := new(fakeListOpsInCollectionUseCase)
      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.ListOpsInCollectionUseCaseReturns(fakeListOpsInCollectionUseCase)

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      objectUnderTest.ListOpsInCollection()

      /* assert */
      Expect(fakeListOpsInCollectionUseCase.ExecuteCallCount()).Should(Equal(1))
    })
  })

  Context("CreateOp", func() {

    Context("with description", func() {
      It("should call compositionRoot.createOpUseCase.Execute with expected args", func() {
        /* arrange */
        fakeCreateOpUseCase := new(fakeCreateOpUseCase)
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeCompositionRoot.CreateOpUseCaseReturns(fakeCreateOpUseCase)

        expectedOpName := "dummyOpName"
        expectedOpDescription := "dummyOpDescription"

        objectUnderTest := &_api{
          compositionRoot:fakeCompositionRoot,
        }

        /* act */
        objectUnderTest.CreateOp(expectedOpDescription, expectedOpName)

        /* assert */
        Expect(fakeCreateOpUseCase.ExecuteCallCount()).Should(Equal(1))
        actualOpDescription, actualOpName := fakeCreateOpUseCase.ExecuteArgsForCall(0)
        Expect(actualOpName).Should(Equal(expectedOpName))
        Expect(actualOpDescription).Should(Equal(expectedOpDescription))
      })
    })

    Context("with no description", func() {
      It("should call compositionRoot.createOpUseCase.Execute with expected args", func() {
        /* arrange */
        fakeCreateOpUseCase := new(fakeCreateOpUseCase)
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeCompositionRoot.CreateOpUseCaseReturns(fakeCreateOpUseCase)

        expectedOpName := "dummyOpName"

        objectUnderTest := &_api{
          compositionRoot:fakeCompositionRoot,
        }

        /* act */
        objectUnderTest.CreateOp("", expectedOpName)

        /* assert */
        Expect(fakeCreateOpUseCase.ExecuteCallCount()).Should(Equal(1))
        actualOpDescription, actualOpName := fakeCreateOpUseCase.ExecuteArgsForCall(0)
        Expect(actualOpName).Should(Equal(expectedOpName))
        Expect(actualOpDescription).Should(BeEmpty())
      })
    })
  })

  Context("SetOpDescription", func() {
    It("should call compositionRoot.setOpDescriptionUseCase.Execute with expected args", func() {
      /* arrange */
      fakeSetOpDescriptionUseCase := new(fakeSetOpDescriptionUseCase)
      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.SetOpDescriptionUseCaseReturns(fakeSetOpDescriptionUseCase)

      expectedOpName := "dummyOpName"
      expectedOpDescription := "dummyOpDescription"

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      objectUnderTest.SetOpDescription(expectedOpDescription, expectedOpName)

      /* assert */
      Expect(fakeSetOpDescriptionUseCase.ExecuteCallCount()).Should(Equal(1))
      actualOpDescription, actualOpName := fakeSetOpDescriptionUseCase.ExecuteArgsForCall(0)
      Expect(actualOpName).Should(Equal(expectedOpName))
      Expect(actualOpDescription).Should(Equal(expectedOpDescription))
    })
  })

  Context("RunOp", func() {

    Context("with two op run args", func() {
      It("should call compositionRoot.runOpUseCase.Execute with expected args", func() {
        /* arrange */
        fakeRunOpUseCase := new(fakeRunOpUseCase)
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeCompositionRoot.RunOpUseCaseReturns(fakeRunOpUseCase)

        expectedOpUrl := "dummyOpName"
        expectedOpRunArgs := []string{"arg1Name=arg1Value", "arg2Name=arg2Value"}

        objectUnderTest := &_api{
          compositionRoot:fakeCompositionRoot,
        }

        /* act */
        objectUnderTest.RunOp(expectedOpRunArgs, expectedOpUrl)

        /* assert */
        Expect(fakeRunOpUseCase.ExecuteCallCount()).Should(Equal(1))

        actualOpRunArgs, actualOpUrl := fakeRunOpUseCase.ExecuteArgsForCall(0)
        Expect(actualOpUrl).Should(Equal(expectedOpUrl))
        Expect(actualOpRunArgs).Should(Equal(expectedOpRunArgs))
      })
    })

    Context("with zero op run args", func() {
      It("should call compositionRoot.runOpUseCase.Execute with expected args", func() {
        /* arrange */
        fakeRunOpUseCase := new(fakeRunOpUseCase)
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeCompositionRoot.RunOpUseCaseReturns(fakeRunOpUseCase)

        expectedOpUrl := "dummyOpName"

        objectUnderTest := &_api{
          compositionRoot:fakeCompositionRoot,
        }

        /* act */
        objectUnderTest.RunOp([]string{}, expectedOpUrl)

        /* assert */
        Expect(fakeRunOpUseCase.ExecuteCallCount()).Should(Equal(1))

        actualOpRunArgs, actualOpUrl := fakeRunOpUseCase.ExecuteArgsForCall(0)
        Expect(actualOpUrl).Should(Equal(expectedOpUrl))
        Expect(actualOpRunArgs).Should(BeEmpty())
      })
    })
  })

})
