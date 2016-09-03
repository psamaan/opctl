package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/cli/core"
)

var _ = Describe("cli", func() {
  Context("Run", func() {

    Context("collection", func() {

      Context("create", func() {

        Context("with description", func() {
          It("should call compositionRoot.CoreApi.CreateCollection with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeApi := new(core.FakeApi)
            fakeCompositionRoot.CoreApiReturns(fakeApi)

            expectedCollectionName := "dummyCollectionName"
            expectedCollectionDescription := "dummyCollectionDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "collection", "create", "-d", expectedCollectionDescription, expectedCollectionName})

            /* assert */
            Expect(fakeApi.CreateCollectionCallCount()).Should(Equal(1))
            actualCollectionDescription, actualCollectionName := fakeApi.CreateCollectionArgsForCall(0)
            Expect(actualCollectionName).Should(Equal(expectedCollectionName))
            Expect(actualCollectionDescription).Should(Equal(expectedCollectionDescription))
          })
        })

        Context("with no description", func() {
          It("should call compositionRoot.CoreApi.CreateCollectionUseCase with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeApi := new(core.FakeApi)
            fakeCompositionRoot.CoreApiReturns(fakeApi)

            expectedCollectionName := "dummyCollectionName"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "collection", "create", expectedCollectionName})

            /* assert */
            Expect(fakeApi.CreateCollectionCallCount()).Should(Equal(1))
            actualCollectionDescription, actualCollectionName := fakeApi.CreateCollectionArgsForCall(0)
            Expect(actualCollectionName).Should(Equal(expectedCollectionName))
            Expect(actualCollectionDescription).Should(BeEmpty())
          })
        })
      })

      Context("set", func() {

        Context("description", func() {
          It("should call compositionRoot.CoreApi.SetCollectionDescription with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeApi := new(core.FakeApi)
            fakeCompositionRoot.CoreApiReturns(fakeApi)

            expectedCollectionDescription := "dummyCollectionDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "collection", "set", "description", expectedCollectionDescription})

            /* assert */
            Expect(fakeApi.SetCollectionDescriptionCallCount()).Should(Equal(1))
            Expect(fakeApi.SetCollectionDescriptionArgsForCall(0)).Should(Equal(expectedCollectionDescription))
          })
        })

      })

    })

    Context("events", func() {
      It("should call compositionRoot.CoreApi.StreamEvents with expected args", func() {
        /* arrange */
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeApi := new(core.FakeApi)
        fakeCompositionRoot.CoreApiReturns(fakeApi)

        objectUnderTest := newCli(fakeCompositionRoot)

        /* act */
        objectUnderTest.Run([]string{"opctl", "events"})

        /* assert */
        Expect(fakeApi.StreamEventsCallCount()).Should(Equal(1))
      })
    })

    Context("kill", func() {
      It("should call compositionRoot.CoreApi.KillOpRun with expected args", func() {
        /* arrange */
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeApi := new(core.FakeApi)
        fakeCompositionRoot.CoreApiReturns(fakeApi)

        expectedOpRunId := "dummyOpRunId"

        objectUnderTest := newCli(fakeCompositionRoot)

        /* act */
        objectUnderTest.Run([]string{"opctl", "kill", expectedOpRunId})

        /* assert */
        Expect(fakeApi.KillOpRunCallCount()).Should(Equal(1))
        Expect(fakeApi.KillOpRunArgsForCall(0)).Should(Equal(expectedOpRunId))
      })
    })

    Context("ls", func() {
      It("should call compositionRoot.CoreApi.ListOpsInCollection with expected args", func() {
        /* arrange */
        fakeCompositionRoot := new(fakeCompositionRoot)
        fakeApi := new(core.FakeApi)
        fakeCompositionRoot.CoreApiReturns(fakeApi)

        objectUnderTest := newCli(fakeCompositionRoot)

        /* act */
        objectUnderTest.Run([]string{"opctl", "ls"})

        /* assert */
        Expect(fakeApi.ListOpsInCollectionCallCount()).Should(Equal(1))
      })
    })

    Context("op", func() {

      Context("create", func() {

        Context("with description", func() {
          It("should call compositionRoot.CoreApi.CreateOp with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeApi := new(core.FakeApi)
            fakeCompositionRoot.CoreApiReturns(fakeApi)

            expectedOpName := "dummyOpName"
            expectedOpDescription := "dummyOpDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "op", "create", "-d", expectedOpDescription, expectedOpName})

            /* assert */
            Expect(fakeApi.CreateOpCallCount()).Should(Equal(1))
            actualOpDescription, actualOpName := fakeApi.CreateOpArgsForCall(0)
            Expect(actualOpName).Should(Equal(expectedOpName))
            Expect(actualOpDescription).Should(Equal(expectedOpDescription))
          })
        })

        Context("with no description", func() {
          It("should call compositionRoot.CoreApi.CreateOp with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeApi := new(core.FakeApi)
            fakeCompositionRoot.CoreApiReturns(fakeApi)

            expectedOpName := "dummyOpName"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "op", "create", expectedOpName})

            /* assert */
            Expect(fakeApi.CreateOpCallCount()).Should(Equal(1))
            actualOpDescription, actualOpName := fakeApi.CreateOpArgsForCall(0)
            Expect(actualOpName).Should(Equal(expectedOpName))
            Expect(actualOpDescription).Should(BeEmpty())
          })
        })
      })

      Context("set", func() {

        Context("description", func() {
          It("should call compositionRoot.CoreApi.SetOpDescription with expected args", func() {
            /* arrange */
            fakeCompositionRoot := new(fakeCompositionRoot)
            fakeApi := new(core.FakeApi)
            fakeCompositionRoot.CoreApiReturns(fakeApi)

            expectedOpName := "dummyOpName"
            expectedOpDescription := "dummyOpDescription"

            objectUnderTest := newCli(fakeCompositionRoot)

            /* act */
            objectUnderTest.Run([]string{"opctl", "op", "set", "description", expectedOpDescription, expectedOpName})

            /* assert */
            Expect(fakeApi.SetOpDescriptionCallCount()).Should(Equal(1))
            actualOpDescription, actualOpName := fakeApi.SetOpDescriptionArgsForCall(0)
            Expect(actualOpName).Should(Equal(expectedOpName))
            Expect(actualOpDescription).Should(Equal(expectedOpDescription))
          })
        })

      })

    })

    Context("run", func() {

      Context("with two op run args", func() {
        It("should call compositionRoot.CoreApi.RunOp with expected args", func() {
          /* arrange */
          fakeCompositionRoot := new(fakeCompositionRoot)
          fakeApi := new(core.FakeApi)
          fakeCompositionRoot.CoreApiReturns(fakeApi)

          expectedOpUrl := "dummyOpUrl"
          expectedOpRunArgs := []string{"arg1Name=arg1Value", "arg2Name=arg2Value"}

          objectUnderTest := newCli(fakeCompositionRoot)

          /* act */
          objectUnderTest.Run([]string{"opctl", "run", "-a", expectedOpRunArgs[0], "-a", expectedOpRunArgs[1], expectedOpUrl})

          /* assert */
          Expect(fakeApi.RunOpCallCount()).Should(Equal(1))
          actualOpRunArgs, actualOpUrl := fakeApi.RunOpArgsForCall(0)
          Expect(actualOpUrl).Should(Equal(expectedOpUrl))
          Expect(actualOpRunArgs).Should(Equal(expectedOpRunArgs))
        })
      })

      Context("with zero op run args", func() {
        It("should call compositionRoot.CoreApi.RunOp with expected args", func() {
          /* arrange */
          fakeCompositionRoot := new(fakeCompositionRoot)
          fakeApi := new(core.FakeApi)
          fakeCompositionRoot.CoreApiReturns(fakeApi)

          expectedOpUrl := "dummyOpUrl"

          objectUnderTest := newCli(fakeCompositionRoot)

          /* act */
          objectUnderTest.Run([]string{"opctl", "run", expectedOpUrl})

          /* assert */
          Expect(fakeApi.RunOpCallCount()).Should(Equal(1))

          actualOpRunArgs, actualOpUrl := fakeApi.RunOpArgsForCall(0)
          Expect(actualOpUrl).Should(Equal(expectedOpUrl))
          Expect(actualOpRunArgs).Should(BeEmpty())
        })
      })
    })
  })
})
