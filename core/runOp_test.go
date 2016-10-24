package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
  "time"
  "fmt"
  "os"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
  "github.com/opspec-io/sdk-golang/pkg/engineclient"
)

var _ = Describe("runOpUseCase", func() {

  Context("Execute", func() {
    It("should call exiter with expected args when bundle.GetOp returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeBundle := new(bundle.FakeBundle)
      fakeBundle.GetOpReturns(models.OpView{}, returnedError)

      fakeEngineClient := new(engineclient.FakeEngineClient)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeEngineClient.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := _api{
        bundle:fakeBundle,
        engineClient:fakeEngineClient,
        exiter:fakeExiter,
        workDirPathGetter:new(fakeWorkDirPathGetter),
      }

      /* act */
      objectUnderTest.RunOp([]string{}, "dummyName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    It("should call bundle.GetOp with expected args", func() {
      /* arrange */
      fakeBundle := new(bundle.FakeBundle)

      fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
      workDirPath := "dummyWorkDirPath"
      fakeWorkDirPathGetter.GetReturns(workDirPath)

      fakeEngineClient := new(engineclient.FakeEngineClient)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeEngineClient.GetEventStreamReturns(eventChannel, nil)

      providedName := "dummyOpName"

      expectedPath := path.Join(workDirPath, ".opspec", providedName)

      objectUnderTest := _api{
        bundle:fakeBundle,
        engineClient:fakeEngineClient,
        exiter:new(fakeExiter),
        workDirPathGetter:fakeWorkDirPathGetter,
      }

      /* act */
      objectUnderTest.RunOp([]string{}, providedName)

      /* assert */
      Expect(fakeBundle.GetOpArgsForCall(0)).Should(Equal(expectedPath))
    })
    It("should call exiter with expected args when bundle.GetEventStream returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeBundle := new(bundle.FakeBundle)
      fakeBundle.GetOpReturns(models.OpView{}, nil)

      fakeEngineClient := new(engineclient.FakeEngineClient)
      fakeEngineClient.GetEventStreamReturns(nil, returnedError)

      objectUnderTest := _api{
        bundle:fakeBundle,
        engineClient:fakeEngineClient,
        exiter:fakeExiter,
        workDirPathGetter:new(fakeWorkDirPathGetter),
      }

      /* act */
      objectUnderTest.RunOp([]string{}, "dummyOpName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    Describe("when op has params defined", func() {
      Describe("and corresponding args are provided explicitly with values", func() {
        It("should call engineClient.StartOpRun with provided arg values", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          fakeBundle := new(bundle.FakeBundle)
          fakeBundle.GetOpReturns(
            models.OpView{
              Inputs:[]models.Param{
                {
                  Name: param1Name,
                  String:&models.StringParam{},
                },
              },
            },
            nil,
          )

          fakeEngineClient := new(engineclient.FakeEngineClient)
          fakeEngineClient.StartOpRunReturns("dummyOpRunId", errors.New(""))

          objectUnderTest := _api{
            bundle:fakeBundle,
            engineClient:fakeEngineClient,
            exiter:new(fakeExiter),
            workDirPathGetter:new(fakeWorkDirPathGetter),
          }

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{fmt.Sprintf("%v=%v", param1Name, param1Value)}

          /* act */
          objectUnderTest.RunOp(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeEngineClient.StartOpRunArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
      Describe("and corresponding args are provided explicitly without values", func() {
        It("should call engineClient.StartOpRun with arg values obtained from the environment", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          os.Setenv(param1Name, param1Value)

          fakeBundle := new(bundle.FakeBundle)
          fakeBundle.GetOpReturns(
            models.OpView{
              Inputs:[]models.Param{
                {
                  Name: param1Name,
                  String:&models.StringParam{},
                },
              },
            },
            nil,
          )

          fakeEngineClient := new(engineclient.FakeEngineClient)
          fakeEngineClient.StartOpRunReturns("dummyOpRunId", errors.New(""))

          objectUnderTest := _api{
            bundle:fakeBundle,
            engineClient:fakeEngineClient,
            exiter:new(fakeExiter),
            workDirPathGetter:new(fakeWorkDirPathGetter),
          }

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{param1Name}

          /* act */
          objectUnderTest.RunOp(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeEngineClient.StartOpRunArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
      Describe("and corresponding args are not provided", func() {
        It("should call bundle.RunOp with arg values obtained from the environment", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          os.Setenv(param1Name, param1Value)

          fakeBundle := new(bundle.FakeBundle)
          fakeBundle.GetOpReturns(
            models.OpView{
              Inputs:[]models.Param{
                {
                  Name: param1Name,
                  String:&models.StringParam{},
                },
              },
            },
            nil,
          )

          fakeEngineClient := new(engineclient.FakeEngineClient)
          fakeEngineClient.StartOpRunReturns("dummyOpRunId", errors.New(""))

          objectUnderTest := _api{
            bundle:fakeBundle,
            engineClient:fakeEngineClient,
            exiter:new(fakeExiter),
            workDirPathGetter:new(fakeWorkDirPathGetter),
          }

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{}

          /* act */
          objectUnderTest.RunOp(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeEngineClient.StartOpRunArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
    })
    It("should call exiter with expected args when engineClient.StartOpRun returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeBundle := new(bundle.FakeBundle)
      fakeBundle.GetOpReturns(models.OpView{}, nil)

      fakeEngineClient := new(engineclient.FakeEngineClient)
      fakeEngineClient.StartOpRunReturns("dummyOpRunId", returnedError)

      objectUnderTest := _api{
        bundle:fakeBundle,
        engineClient:fakeEngineClient,
        exiter:fakeExiter,
        workDirPathGetter:new(fakeWorkDirPathGetter),
      }

      /* act */
      objectUnderTest.RunOp([]string{}, "dummyOpName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    It("should call exiter with expected args when event channel closes unexpectedly", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)

      fakeBundle := new(bundle.FakeBundle)
      fakeBundle.GetOpReturns(models.OpView{}, nil)

      fakeEngineClient := new(engineclient.FakeEngineClient)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeEngineClient.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := _api{
        bundle:fakeBundle,
        engineClient:fakeEngineClient,
        exiter:fakeExiter,
        workDirPathGetter:new(fakeWorkDirPathGetter),
      }

      /* act */
      objectUnderTest.RunOp([]string{}, "dummyOpName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:"Event channel closed unexpectedly", Code:1}))
    })
    Describe("when an OpRunEndedEvent is received for our root op run", func() {
      rootOpRunId := "dummyRootOpRunId"
      It("should call exiter with expected args when it's Outcome is SUCCEEDED", func() {
        /* arrange */
        opRunEndedEvent := models.Event{
          Timestamp:time.Now(),
          OpRunEnded:&models.OpRunEndedEvent{
            OpRunId:rootOpRunId,
            Outcome:models.OpRunOutcomeSucceeded,
            RootOpRunId:rootOpRunId,
          },
        }

        fakeExiter := new(fakeExiter)

        fakeBundle := new(bundle.FakeBundle)
        fakeBundle.GetOpReturns(models.OpView{}, nil)

        fakeEngineClient := new(engineclient.FakeEngineClient)
        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeEngineClient.GetEventStreamReturns(eventChannel, nil)
        fakeEngineClient.StartOpRunReturns(opRunEndedEvent.OpRunEnded.RootOpRunId, nil)

        objectUnderTest := _api{
          bundle:fakeBundle,
          engineClient:fakeEngineClient,
          exiter:fakeExiter,
          workDirPathGetter:new(fakeWorkDirPathGetter),
        }

        /* act/assert */
        objectUnderTest.RunOp([]string{}, "dummyOpName")
        Expect(fakeExiter.ExitArgsForCall(0)).
          Should(Equal(ExitReq{Message:"", Code:0}))
      })
      It("should call exiter with expected args when it's Outcome is KILLED", func() {
        /* arrange */
        opRunEndedEvent := models.Event{
          Timestamp:time.Now(),
          OpRunEnded:&models.OpRunEndedEvent{
            OpRunId:rootOpRunId,
            Outcome:models.OpRunOutcomeKilled,
            RootOpRunId:rootOpRunId,
          },
        }

        fakeExiter := new(fakeExiter)

        fakeBundle := new(bundle.FakeBundle)
        fakeBundle.GetOpReturns(models.OpView{}, nil)

        fakeEngineClient := new(engineclient.FakeEngineClient)
        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeEngineClient.GetEventStreamReturns(eventChannel, nil)
        fakeEngineClient.StartOpRunReturns(opRunEndedEvent.OpRunEnded.RootOpRunId, nil)

        objectUnderTest := _api{
          bundle:fakeBundle,
          engineClient:fakeEngineClient,
          exiter:fakeExiter,
          workDirPathGetter:new(fakeWorkDirPathGetter),
        }

        /* act/assert */
        objectUnderTest.RunOp([]string{}, "dummyOpName")
        Expect(fakeExiter.ExitArgsForCall(0)).
          Should(Equal(ExitReq{Message:"", Code:137}))
      })
      It("should call exiter with expected args when it's Outcome is unexpected", func() {
        /* arrange */
        opRunEndedEvent := models.Event{
          Timestamp:time.Now(),
          OpRunEnded:&models.OpRunEndedEvent{
            OpRunId:rootOpRunId,
            Outcome:"some unexpected outcome",
            RootOpRunId:rootOpRunId,
          },
        }

        fakeExiter := new(fakeExiter)

        fakeBundle := new(bundle.FakeBundle)
        fakeBundle.GetOpReturns(models.OpView{}, nil)

        fakeEngineClient := new(engineclient.FakeEngineClient)
        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeEngineClient.GetEventStreamReturns(eventChannel, nil)
        fakeEngineClient.StartOpRunReturns(opRunEndedEvent.OpRunEnded.RootOpRunId, nil)

        objectUnderTest := _api{
          bundle:fakeBundle,
          engineClient:fakeEngineClient,
          exiter:fakeExiter,
          workDirPathGetter:new(fakeWorkDirPathGetter),
        }

        /* act/assert */
        objectUnderTest.RunOp([]string{}, "dummyOpName")
        Expect(fakeExiter.ExitArgsForCall(0)).
          Should(Equal(ExitReq{Message:fmt.Sprintf("Received unknown outcome `%v`", opRunEndedEvent.OpRunEnded.Outcome), Code:1}))
      })
    })
  })
})
