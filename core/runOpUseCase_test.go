package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang"
  "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
  "time"
  "fmt"
  "os"
)

var _ = Describe("runOpUseCase", func() {

  Context("Execute", func() {
    It("should call exiter with expected args when opspecSdk.GetOp returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.GetOpReturns(models.OpView{}, returnedError)

      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        new(fakeWorkDirPathGetter),
      )

      /* act */
      objectUnderTest.Execute([]string{}, "dummyName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    It("should call opspecSdk.GetOp with expected args", func() {
      /* arrange */
      fakeOpspecSdk := new(opspec.FakeSdk)

      fakeWorkDirPathGetter := new(fakeWorkDirPathGetter)
      workDirPath := "dummyWorkDirPath"
      fakeWorkDirPathGetter.GetReturns(workDirPath)

      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)

      providedName := "dummyOpName"

      expectedPath := path.Join(workDirPath, ".opspec", providedName)

      objectUnderTest := newRunOpUseCase(
        new(fakeExiter),
        fakeOpspecSdk,
        fakeWorkDirPathGetter,
      )

      /* act */
      objectUnderTest.Execute([]string{}, providedName)

      /* assert */
      Expect(fakeOpspecSdk.GetOpArgsForCall(0)).Should(Equal(expectedPath))
    })
    It("should call exiter with expected args when opspecSdk.GetEventStream returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.GetOpReturns(models.OpView{}, nil)

      fakeOpspecSdk.GetEventStreamReturns(nil, returnedError)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        new(fakeWorkDirPathGetter),
      )

      /* act */
      objectUnderTest.Execute([]string{}, "dummyOpName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    Describe("when op has params defined", func() {
      Describe("and corresponding args are provided explicitly with values", func() {
        It("should call opspecSdk.StartOpRun with provided arg values", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          fakeOpspecSdk := new(opspec.FakeSdk)
          fakeOpspecSdk.GetOpReturns(
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

          fakeOpspecSdk.StartOpRunReturns("dummyOpRunId", errors.New(""))

          objectUnderTest := newRunOpUseCase(
            new(fakeExiter),
            fakeOpspecSdk,
            new(fakeWorkDirPathGetter),
          )

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{fmt.Sprintf("%v=%v", param1Name, param1Value)}

          /* act */
          objectUnderTest.Execute(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeOpspecSdk.StartOpRunArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
      Describe("and corresponding args are provided explicitly without values", func() {
        It("should call opspecSdk.StartOpRun with arg values obtained from the environment", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          os.Setenv(param1Name, param1Value)

          fakeOpspecSdk := new(opspec.FakeSdk)
          fakeOpspecSdk.GetOpReturns(
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

          fakeOpspecSdk.StartOpRunReturns("dummyOpRunId", errors.New(""))

          objectUnderTest := newRunOpUseCase(
            new(fakeExiter),
            fakeOpspecSdk,
            new(fakeWorkDirPathGetter),
          )

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{param1Name}

          /* act */
          objectUnderTest.Execute(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeOpspecSdk.StartOpRunArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
      Describe("and corresponding args are not provided", func() {
        It("should call opspecSdk.RunOp with arg values obtained from the environment", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          os.Setenv(param1Name, param1Value)

          fakeOpspecSdk := new(opspec.FakeSdk)
          fakeOpspecSdk.GetOpReturns(
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

          fakeOpspecSdk.StartOpRunReturns("dummyOpRunId", errors.New(""))

          objectUnderTest := newRunOpUseCase(
            new(fakeExiter),
            fakeOpspecSdk,
            new(fakeWorkDirPathGetter),
          )

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{}

          /* act */
          objectUnderTest.Execute(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeOpspecSdk.StartOpRunArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
    })
    It("should call exiter with expected args when opspecSdk.StartOpRun returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.GetOpReturns(models.OpView{}, nil)

      fakeOpspecSdk.StartOpRunReturns("dummyOpRunId", returnedError)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        new(fakeWorkDirPathGetter),
      )

      /* act */
      objectUnderTest.Execute([]string{}, "dummyOpName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:returnedError.Error(), Code:1}))
    })
    It("should call exiter with expected args when event channel closes unexpectedly", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)

      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.GetOpReturns(models.OpView{}, nil)

      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        new(fakeWorkDirPathGetter),
      )

      /* act */
      objectUnderTest.Execute([]string{}, "dummyOpName")

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

        fakeOpspecSdk := new(opspec.FakeSdk)
        fakeOpspecSdk.GetOpReturns(models.OpView{}, nil)

        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)
        fakeOpspecSdk.StartOpRunReturns(opRunEndedEvent.OpRunEnded.RootOpRunId, nil)

        objectUnderTest := newRunOpUseCase(
          fakeExiter,
          fakeOpspecSdk,
          new(fakeWorkDirPathGetter),
        )

        /* act/assert */
        objectUnderTest.Execute([]string{}, "dummyOpName")
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

        fakeOpspecSdk := new(opspec.FakeSdk)
        fakeOpspecSdk.GetOpReturns(models.OpView{}, nil)

        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)
        fakeOpspecSdk.StartOpRunReturns(opRunEndedEvent.OpRunEnded.RootOpRunId, nil)

        objectUnderTest := newRunOpUseCase(
          fakeExiter,
          fakeOpspecSdk,
          new(fakeWorkDirPathGetter),
        )

        /* act/assert */
        objectUnderTest.Execute([]string{}, "dummyOpName")
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

        fakeOpspecSdk := new(opspec.FakeSdk)
        fakeOpspecSdk.GetOpReturns(models.OpView{}, nil)

        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeOpspecSdk.GetEventStreamReturns(eventChannel, nil)
        fakeOpspecSdk.StartOpRunReturns(opRunEndedEvent.OpRunEnded.RootOpRunId, nil)

        objectUnderTest := newRunOpUseCase(
          fakeExiter,
          fakeOpspecSdk,
          new(fakeWorkDirPathGetter),
        )

        /* act/assert */
        objectUnderTest.Execute([]string{}, "dummyOpName")
        Expect(fakeExiter.ExitArgsForCall(0)).
          Should(Equal(ExitReq{Message:fmt.Sprintf("Received unknown outcome `%v`", opRunEndedEvent.OpRunEnded.Outcome), Code:1}))
      })
    })
  })
})
