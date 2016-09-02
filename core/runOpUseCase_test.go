package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opspec-io/sdk-golang"
  engineModels "github.com/opspec-io/engine/core/models"
  opspecSdkModels "github.com/opspec-io/sdk-golang/models"
  "path"
  "github.com/opspec-io/engine-sdk-golang"
  "github.com/opspec-io/engine-sdk-golang/models"
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
      fakeOpspecSdk.GetOpReturns(opspecSdkModels.OpView{}, returnedError)

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        fakeOpctlEngineSdk,
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

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)

      providedName := "dummyOpName"

      expectedPath := path.Join(workDirPath, ".opspec", providedName)

      objectUnderTest := newRunOpUseCase(
        new(fakeExiter),
        fakeOpspecSdk,
        fakeOpctlEngineSdk,
        fakeWorkDirPathGetter,
      )

      /* act */
      objectUnderTest.Execute([]string{}, providedName)

      /* assert */
      Expect(fakeOpspecSdk.GetOpArgsForCall(0)).Should(Equal(expectedPath))
    })
    It("should call exiter with expected args when opctlEngineSdk.GetEventStream returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.GetOpReturns(opspecSdkModels.OpView{}, nil)

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      fakeOpctlEngineSdk.GetEventStreamReturns(nil, returnedError)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        fakeOpctlEngineSdk,
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
        It("should call opctlEngineSdk.RunOp with provided arg values", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          fakeOpspecSdk := new(opspec.FakeSdk)
          fakeOpspecSdk.GetOpReturns(
            opspecSdkModels.OpView{
              Inputs:[]opspecSdkModels.Parameter{
                {
                  Name:param1Name,
                },
              },
            },
            nil,
          )

          fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
          fakeOpctlEngineSdk.RunOpReturns("dummyOpRunId", "dummyCorrelationId", errors.New(""))

          objectUnderTest := newRunOpUseCase(
            new(fakeExiter),
            fakeOpspecSdk,
            fakeOpctlEngineSdk,
            new(fakeWorkDirPathGetter),
          )

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{fmt.Sprintf("%v=%v", param1Name, param1Value)}

          /* act */
          objectUnderTest.Execute(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeOpctlEngineSdk.RunOpArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
      Describe("and corresponding args are provided explicitly without values", func() {
        It("should call opctlEngineSdk.RunOp with arg values obtained from the environment", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          os.Setenv(param1Name, param1Value)

          fakeOpspecSdk := new(opspec.FakeSdk)
          fakeOpspecSdk.GetOpReturns(
            opspecSdkModels.OpView{
              Inputs:[]opspecSdkModels.Parameter{
                {
                  Name:param1Name,
                },
              },
            },
            nil,
          )

          fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
          fakeOpctlEngineSdk.RunOpReturns("dummyOpRunId", "dummyCorrelationId", errors.New(""))

          objectUnderTest := newRunOpUseCase(
            new(fakeExiter),
            fakeOpspecSdk,
            fakeOpctlEngineSdk,
            new(fakeWorkDirPathGetter),
          )

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{param1Name}

          /* act */
          objectUnderTest.Execute(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeOpctlEngineSdk.RunOpArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
      Describe("and corresponding args are not provided", func() {
        It("should call opctlEngineSdk.RunOp with arg values obtained from the environment", func() {
          /* arrange */
          param1Name := "DUMMY_PARAM1_NAME"
          param1Value := "dummyParam1Value"

          os.Setenv(param1Name, param1Value)

          fakeOpspecSdk := new(opspec.FakeSdk)
          fakeOpspecSdk.GetOpReturns(
            opspecSdkModels.OpView{
              Inputs:[]opspecSdkModels.Parameter{
                {
                  Name:param1Name,
                },
              },
            },
            nil,
          )

          fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
          fakeOpctlEngineSdk.RunOpReturns("dummyOpRunId", "dummyCorrelationId", errors.New(""))

          objectUnderTest := newRunOpUseCase(
            new(fakeExiter),
            fakeOpspecSdk,
            fakeOpctlEngineSdk,
            new(fakeWorkDirPathGetter),
          )

          expectedArgs := map[string]string{param1Name:param1Value}
          providedArgs := []string{}

          /* act */
          objectUnderTest.Execute(providedArgs, "dummyOpName")

          /* assert */
          Expect(fakeOpctlEngineSdk.RunOpArgsForCall(0).Args).To(BeEquivalentTo(expectedArgs))
        })
      })
    })
    It("should call exiter with expected args when opctlEngineSdk.RunOp returns error", func() {
      /* arrange */
      fakeExiter := new(fakeExiter)
      returnedError := errors.New("dummyError")

      fakeOpspecSdk := new(opspec.FakeSdk)
      fakeOpspecSdk.GetOpReturns(opspecSdkModels.OpView{}, nil)

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      fakeOpctlEngineSdk.RunOpReturns("dummyOpRunId", "dummyCorrelationId", returnedError)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        fakeOpctlEngineSdk,
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
      fakeOpspecSdk.GetOpReturns(opspecSdkModels.OpView{}, nil)

      fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
      eventChannel := make(chan models.Event)
      close(eventChannel)
      fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)

      objectUnderTest := newRunOpUseCase(
        fakeExiter,
        fakeOpspecSdk,
        fakeOpctlEngineSdk,
        new(fakeWorkDirPathGetter),
      )

      /* act */
      objectUnderTest.Execute([]string{}, "dummyOpName")

      /* assert */
      Expect(fakeExiter.ExitArgsForCall(0)).
        Should(Equal(ExitReq{Message:"Event channel closed unexpectedly", Code:1}))
    })
    Describe("when a related OpRunEndedEvent is received", func() {
      It("should call exiter with expected args when it's Outcome is SUCCEEDED", func() {
        /* arrange */
        opRunEndedEvent := engineModels.NewOpRunEndedEvent(
          "dummyCorrelationId",
          "dummyRootOpRunId",
          engineModels.OpRunOutcomeSucceeded,
          "dummyRootOpRunId",
          time.Now(),
        )

        fakeExiter := new(fakeExiter)

        fakeOpspecSdk := new(opspec.FakeSdk)
        fakeOpspecSdk.GetOpReturns(opspecSdkModels.OpView{}, nil)

        fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)
        fakeOpctlEngineSdk.RunOpReturns(opRunEndedEvent.RootOpRunId(), "dummyCorrelationId", nil)

        objectUnderTest := newRunOpUseCase(
          fakeExiter,
          fakeOpspecSdk,
          fakeOpctlEngineSdk,
          new(fakeWorkDirPathGetter),
        )

        /* act/assert */
        objectUnderTest.Execute([]string{}, "dummyOpName")
        Expect(fakeExiter.ExitArgsForCall(0)).
          Should(Equal(ExitReq{Message:"", Code:0}))
      })
      It("should call exiter with expected args when it's Outcome is KILLED", func() {
        /* arrange */
        opRunEndedEvent := engineModels.NewOpRunEndedEvent(
          "dummyCorrelationId",
          "dummyRootOpRunId",
          engineModels.OpRunOutcomeKilled,
          "dummyRootOpRunId",
          time.Now(),
        )

        fakeExiter := new(fakeExiter)

        fakeOpspecSdk := new(opspec.FakeSdk)
        fakeOpspecSdk.GetOpReturns(opspecSdkModels.OpView{}, nil)

        fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)
        fakeOpctlEngineSdk.RunOpReturns(opRunEndedEvent.RootOpRunId(), "dummyCorrelationId", nil)

        objectUnderTest := newRunOpUseCase(
          fakeExiter,
          fakeOpspecSdk,
          fakeOpctlEngineSdk,
          new(fakeWorkDirPathGetter),
        )

        /* act/assert */
        objectUnderTest.Execute([]string{}, "dummyOpName")
        Expect(fakeExiter.ExitArgsForCall(0)).
          Should(Equal(ExitReq{Message:"", Code:137}))
      })
      It("should call exiter with expected args when it's Outcome is unexpected", func() {
        /* arrange */
        opRunEndedEvent := engineModels.NewOpRunEndedEvent(
          "dummyCorrelationId",
          "dummyRootOpRunId",
          "some unexpected outcome",
          "dummyRootOpRunId",
          time.Now(),
        )

        fakeExiter := new(fakeExiter)

        fakeOpspecSdk := new(opspec.FakeSdk)
        fakeOpspecSdk.GetOpReturns(opspecSdkModels.OpView{}, nil)

        fakeOpctlEngineSdk := new(opctlengine.FakeSdk)
        eventChannel := make(chan models.Event, 10)
        eventChannel <- opRunEndedEvent
        defer close(eventChannel)
        fakeOpctlEngineSdk.GetEventStreamReturns(eventChannel, nil)
        fakeOpctlEngineSdk.RunOpReturns(opRunEndedEvent.RootOpRunId(), "dummyCorrelationId", nil)

        objectUnderTest := newRunOpUseCase(
          fakeExiter,
          fakeOpspecSdk,
          fakeOpctlEngineSdk,
          new(fakeWorkDirPathGetter),
        )

        /* act/assert */
        objectUnderTest.Execute([]string{}, "dummyOpName")
        Expect(fakeExiter.ExitArgsForCall(0)).
          Should(Equal(ExitReq{Message:fmt.Sprintf("Received unknown outcome `%v`", opRunEndedEvent.Outcome()), Code:1}))
      })
    })
  })
})
