package action

import (
	"fmt"
	"github.com/golang/glog"
	sdkprobe "github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"time"
	"k8s.io/client-go/rest"
)

type turboActionType struct {
	actionType       proto.ActionItemDTO_ActionType
	targetEntityType proto.EntityDTO_EntityType
}

var (
	turboActionClientMove turboActionType = turboActionType{proto.ActionItemDTO_MOVE, proto.EntityDTO_VIRTUAL_APPLICATION}
)

type ActionHandler struct {
	kubeConfig   *rest.Config
	actionExecutors map[turboActionType]TurboActionExecutor
}

type TurboActionExecutorInput struct {
	ActionItem *proto.ActionItemDTO
}

type TurboActionExecutorOutput struct {
	Succeeded bool
}

type TurboActionExecutor interface {
	Execute(input *TurboActionExecutorInput) (*TurboActionExecutorOutput, error)
}

func NewActionHandler(kubeConfig  *rest.Config) *ActionHandler {
	handler := &ActionHandler{
		actionExecutors: make(map[turboActionType]TurboActionExecutor),
		kubeConfig:   kubeConfig,
	}

	handler.registerActionExecutors(kubeConfig)
	return handler
}

// Register supported action executor.
// As action executor is stateless, they can be safely reused.
func (h *ActionHandler) registerActionExecutors(kubeConfig  *rest.Config) {
	ae := NewClientMoveActionExecutor(kubeConfig)

	h.actionExecutors[turboActionClientMove] = ae
}

// Implement ActionExecutorClient interface defined in Go SDK.
// Execute the current action and return the action result to SDK.
func (h *ActionHandler) ExecuteAction(actionExecutionDTO *proto.ActionExecutionDTO, accountValues []*proto.AccountValue,
	progressTracker sdkprobe.ActionProgressTracker) (*proto.ActionResult, error) {
	// 1. get the action, NOTE: only deal with one action item in current implementation.
	// Check if the action execution DTO is valid, including if the action is supported or not
	if err := h.checkActionExecutionDTO(actionExecutionDTO); err != nil {
		err := fmt.Errorf("Action is not valid: %v", err.Error())
		glog.Errorf(err.Error())
		return h.failedResult(err.Error()), err
	}

	actionItemDTO := actionExecutionDTO.GetActionItem()[0]
	fmt.Printf("############ ActionItemDTO ----> \n %++v\n", actionItemDTO)
	fmt.Printf("##### NewSE -----> \n %++v\n", actionItemDTO.GetNewSE())
	fmt.Printf("##### TargetSE ------> \n %++v\n", actionItemDTO.GetTargetSE())

	// 2. keep sending fake progress to prevent timeout
	stop := make(chan struct{})
	defer close(stop)
	go keepAlive(progressTracker, stop)

	// 3. execute the action
	glog.V(2).Infof("Waiting for action result ...")
	err := h.execute(actionItemDTO)
	if err != nil {
		return h.failedResult(err.Error()), nil
	}

	return h.goodResult(), nil
	//err = fmt.Errorf("DEBUG IN PROGRESS....")
	//return h.failedResult(err.Error()), nil		//h.goodResult(), nil
}

func (h *ActionHandler) execute(actionItem *proto.ActionItemDTO) error {
	input := &TurboActionExecutorInput{
		ActionItem: actionItem,
	}
	actionType := getTurboActionType(actionItem)
	worker := h.actionExecutors[actionType]
	_, err := worker.Execute(input)

	if err != nil {
		msg := fmt.Errorf("Action %v on %s failed.", actionType, actionItem.GetTargetSE().GetEntityType())
		glog.Errorf(msg.Error())
		return err
	}

	return nil
}

// Get the associated turbo action type of the action item dto
func getTurboActionType(ai *proto.ActionItemDTO) turboActionType {
	return turboActionType{ai.GetActionType(), ai.GetTargetSE().GetEntityType()}
}

// Checks if the action execution DTO includes action item and the target SE. Also, check if
// the action type is supported by faas-istio probe.
func (h *ActionHandler) checkActionExecutionDTO(actionExecutionDTO *proto.ActionExecutionDTO) error {
	actionItems := actionExecutionDTO.GetActionItem()

	if actionItems == nil || len(actionItems) == 0 || actionItems[0] == nil {
		return fmt.Errorf("Action execution (%v) validation failed: no action item found", actionExecutionDTO)

	}

	ai := actionItems[0]

	if ai.GetTargetSE() == nil {
		return fmt.Errorf("Action execution (%v) validation failed: no target SE found", actionExecutionDTO)
	}

	// Check for the currentSE
	if ai.GetCurrentSE() == nil {
		return fmt.Errorf("Null current SE.")
	}

	// Check for the newSE
	if ai.GetNewSE() == nil {
		return fmt.Errorf("Null new SE.")
	}

	actionType := turboActionType{ai.GetActionType(), ai.GetTargetSE().GetEntityType()}
	glog.V(2).Infof("Receive a action request of type: %++v", actionType)

	// Check if action is supported
	if _, supported := h.actionExecutors[actionType]; !supported {
		return fmt.Errorf("Action execution (%v) validation failed: not supported type %++v", actionExecutionDTO, actionType)
	}

	return nil
}

func keepAlive(tracker sdkprobe.ActionProgressTracker, stop chan struct{}) {

	//TODO: add timeout
	go func() {
		var progress int32 = 0
		state := proto.ActionResponseState_IN_PROGRESS

		for {
			progress = progress + 1
			if progress > 99 {
				progress = 99
			}

			tracker.UpdateProgress(state, "in progress", progress)

			t := time.NewTimer(time.Second * 3)
			select {
			case <-stop:
				glog.V(2).Infof("action keep alive goroutine exit.")
				return
			case <-t.C:
				glog.V(2).Infof("running faas istio action keep alive goroutine ...")
			}
		}
	}()
}

func (h *ActionHandler) goodResult() *proto.ActionResult {

	state := proto.ActionResponseState_SUCCEEDED
	progress := int32(100)
	msg := "Faas Istio Action Success"

	res := &proto.ActionResponse{
		ActionResponseState: &state,
		Progress:            &progress,
		ResponseDescription: &msg,
	}

	return &proto.ActionResult{
		Response: res,
	}
}

func (h *ActionHandler) failedResult(msg string) *proto.ActionResult {

	state := proto.ActionResponseState_FAILED
	progress := int32(0)
	msg = "Faas Istio Action Failed"

	res := &proto.ActionResponse{
		ActionResponseState: &state,
		Progress:            &progress,
		ResponseDescription: &msg,
	}

	return &proto.ActionResult{
		Response: res,
	}
}
