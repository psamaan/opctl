package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("cli", func() {
  Context("Run", func() {

    Context("collection", func() {

      Context("create", func() {

        Context("with description", func() {
          It("should call createCollectionUseCase.Execute with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeCreateCollectionUseCase := new(fakeCreateCollectionUseCase)
            fakeCompositionRoot.CreateCollectionUseCaseReturns(fakeCreateCollectionUseCase)

            expectedCollectionName := "dummyCollectionName"
            expectedCollectionDescription := "dummyCollectionDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "collection", "create", "-d", expectedCollectionDescription, expectedCollectionName})

            /* assert */
            Expect(fakeCreateCollectionUseCase.ExecuteCallCount()).Should(Equal(1))
            actualCollectionDescription, actualCollectionName := fakeCreateCollectionUseCase.ExecuteArgsForCall(0)
            Expect(actualCollectionName).Should(Equal(expectedCollectionName))
            Expect(actualCollectionDescription).Should(Equal(expectedCollectionDescription))
          })
        })

        Context("with no description", func() {
          It("should call createCollectionUseCase.Execute with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeCreateCollectionUseCase := new(fakeCreateCollectionUseCase)
            fakeCompositionRoot.CreateCollectionUseCaseReturns(fakeCreateCollectionUseCase)

            expectedCollectionName := "dummyCollectionName"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "collection", "create", expectedCollectionName})

            /* assert */
            Expect(fakeCreateCollectionUseCase.ExecuteCallCount()).Should(Equal(1))
            actualCollectionDescription, actualCollectionName := fakeCreateCollectionUseCase.ExecuteArgsForCall(0)
            Expect(actualCollectionName).Should(Equal(expectedCollectionName))
            Expect(actualCollectionDescription).Should(BeEmpty())
          })
        })
      })

      Context("set", func() {

        Context("description", func() {
          It("should call setCollectionDescriptionUseCase.Execute with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeSetCollectionDescriptionUseCase := new(fakeSetCollectionDescriptionUseCase)
            fakeCompositionRoot.SetCollectionDescriptionUseCaseReturns(fakeSetCollectionDescriptionUseCase)

            expectedCollectionDescription := "dummyCollectionDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "collection", "set", "description", expectedCollectionDescription})

            /* assert */
            Expect(fakeSetCollectionDescriptionUseCase.ExecuteCallCount()).Should(Equal(1))
            Expect(fakeSetCollectionDescriptionUseCase.ExecuteArgsForCall(0)).Should(Equal(expectedCollectionDescription))
          })
        })

      })

    })

    Context("events", func() {
      It("should call streamEventsUseCase.Execute with expected args", func() {
        /* arrange */
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeStreamEventsUseCase := new(fakeStreamEventsUseCase)
        fakeCompositionRoot.StreamEventsUseCaseReturns(fakeStreamEventsUseCase)

        objectUnderTest := newCli(fakeCompositionRoot)

        /* act */
        objectUnderTest.Run([]string{"opctl", "events"})

        /* assert */
        Expect(fakeStreamEventsUseCase.ExecuteCallCount()).Should(Equal(1))
      })
    })

    Context("kill", func() {
      It("should call killOpRunUseCase.Execute with expected args", func() {
        /* arrange */
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeKillOpRunUseCase := new(fakeKillOpRunUseCase)
        fakeCompositionRoot.KillOpRunUseCaseReturns(fakeKillOpRunUseCase)

        expectedOpRunId := "dummyOpRunId"

        objectUnderTest := newCli(fakeCompositionRoot)

        /* act */
        objectUnderTest.Run([]string{"opctl", "kill", expectedOpRunId})

        /* assert */
        Expect(fakeKillOpRunUseCase.ExecuteCallCount()).Should(Equal(1))
        Expect(fakeKillOpRunUseCase.ExecuteArgsForCall(0)).Should(Equal(expectedOpRunId))
      })
    })

    Context("ls", func() {
      It("should call listOpsInCollectionUseCase.Execute with expected args", func() {
        /* arrange */
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeListOpsInCollectionUseCase := new(fakeListOpsInCollectionUseCase)
        fakeCompositionRoot.ListOpsInCollectionUseCaseReturns(fakeListOpsInCollectionUseCase)

        objectUnderTest := newCli(fakeCompositionRoot)

        /* act */
        objectUnderTest.Run([]string{"opctl", "ls"})

        /* assert */
        Expect(fakeListOpsInCollectionUseCase.ExecuteCallCount()).Should(Equal(1))
      })
    })

    Context("op", func() {

      Context("create", func() {

        Context("with description", func() {
          It("should call createOpUseCase.Execute with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeCreateOpUseCase := new(fakeCreateOpUseCase)
            fakeCompositionRoot.CreateOpUseCaseReturns(fakeCreateOpUseCase)

            expectedOpName := "dummyOpName"
            expectedOpDescription := "dummyOpDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "op", "create", "-d", expectedOpDescription, expectedOpName})

            /* assert */
            Expect(fakeCreateOpUseCase.ExecuteCallCount()).Should(Equal(1))
            actualOpDescription, actualOpName := fakeCreateOpUseCase.ExecuteArgsForCall(0)
            Expect(actualOpName).Should(Equal(expectedOpName))
            Expect(actualOpDescription).Should(Equal(expectedOpDescription))
          })
        })

        Context("with no description", func() {
          It("should call createOpUseCase.Execute with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeCreateOpUseCase := new(fakeCreateOpUseCase)
            fakeCompositionRoot.CreateOpUseCaseReturns(fakeCreateOpUseCase)

            expectedOpName := "dummyOpName"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "op", "create", expectedOpName})

            /* assert */
            Expect(fakeCreateOpUseCase.ExecuteCallCount()).Should(Equal(1))
            actualOpDescription, actualOpName := fakeCreateOpUseCase.ExecuteArgsForCall(0)
            Expect(actualOpName).Should(Equal(expectedOpName))
            Expect(actualOpDescription).Should(BeEmpty())
          })
        })
      })

      Context("set", func() {

        Context("description", func() {
          It("should call setOpDescriptionUseCase.Execute with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeSetOpDescriptionUseCase := new(fakeSetOpDescriptionUseCase)
            fakeCompositionRoot.SetOpDescriptionUseCaseReturns(fakeSetOpDescriptionUseCase)

            expectedOpName := "dummyOpName"
            expectedOpDescription := "dummyOpDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "op", "set", "description", expectedOpDescription, expectedOpName})

            /* assert */
            Expect(fakeSetOpDescriptionUseCase.ExecuteCallCount()).Should(Equal(1))
            actualOpDescription, actualOpName := fakeSetOpDescriptionUseCase.ExecuteArgsForCall(0)
            Expect(actualOpName).Should(Equal(expectedOpName))
            Expect(actualOpDescription).Should(Equal(expectedOpDescription))
          })
        })

      })

    })

    Context("run", func() {

      Context("with two op run args", func() {
        It("should call runOpUseCase.Execute with expected args", func() {
          /* arrange */
          fakeCompositionRoot := new(fakeCompositionRoot)
          fakeRunOpUseCase := new(fakeRunOpUseCase)
          fakeCompositionRoot.RunOpUseCaseReturns(fakeRunOpUseCase)

          expectedOpUrl := "dummyOpUrl"
          expectedOpRunArgs := []string{"arg1Name=arg1Value", "arg2Name=arg2Value"}

          objectUnderTest := newCli(fakeCompositionRoot)

          /* act */
          objectUnderTest.Run([]string{"opctl", "run", "-a", expectedOpRunArgs[0], "-a", expectedOpRunArgs[1], expectedOpUrl})

          /* assert */
          Expect(fakeRunOpUseCase.ExecuteCallCount()).Should(Equal(1))
          actualOpRunArgs, actualOpUrl := fakeRunOpUseCase.ExecuteArgsForCall(0)
          Expect(actualOpUrl).Should(Equal(expectedOpUrl))
          Expect(actualOpRunArgs).Should(Equal(expectedOpRunArgs))
        })
      })

      Context("with zero op run args", func() {
        It("should call runOpUseCase.Execute with expected args", func() {
          /* arrange */
          fakeCompositionRoot := new(fakeCompositionRoot)
          fakeRunOpUseCase := new(fakeRunOpUseCase)
          fakeCompositionRoot.RunOpUseCaseReturns(fakeRunOpUseCase)

          expectedOpUrl := "dummyOpUrl"

          objectUnderTest := newCli(fakeCompositionRoot)

          /* act */
          objectUnderTest.Run([]string{"opctl", "run", expectedOpUrl})

          /* assert */
          Expect(fakeRunOpUseCase.ExecuteCallCount()).Should(Equal(1))

          actualOpRunArgs, actualOpUrl := fakeRunOpUseCase.ExecuteArgsForCall(0)
          Expect(actualOpUrl).Should(Equal(expectedOpUrl))
          Expect(actualOpRunArgs).Should(BeEmpty())
        })
      })
    })
  })
})
