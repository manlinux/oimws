package core_func

import (
	"fmt"

	"github.com/OpenIMSDK/tools/errs"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
)

/*
	定义 RespMessage 结构体，用于发送各种消息
*/

type RespMessage struct {
	respMessagesChan chan *EventData
	//	type EventData struct {
	//	Event       string `json:"event"`
	//	ErrCode     int32  `json:"errCode"`
	//	ErrMsg      string `json:"errMsg"`
	//	Data        string `json:"data"`
	//	OperationID string `json:"operationID"`
	//}
}

func NewRespMessage(respMessagesChan chan *EventData) *RespMessage {
	return &RespMessage{respMessagesChan: respMessagesChan}
}

func (r *RespMessage) sendOnSuccessResp(operationID string, data string) {
	r.respMessagesChan <- &EventData{
		Event:       Success,
		OperationID: operationID,
		Data:        data,
	}
}

func (r *RespMessage) sendOnErrorResp(operationID string, err error) {
	resp := &EventData{
		Event:       Failed,
		OperationID: operationID,
	}
	if code, ok := err.(errs.CodeError); ok {
		resp.ErrCode = int32(code.Code())
		resp.ErrMsg = code.Error()
	} else {
		resp.ErrCode = sdkerrs.UnknownCode
		resp.ErrMsg = fmt.Sprintf("error %T not implement CodeError: %s", err, err)
	}
	r.respMessagesChan <- resp
}

func (r *RespMessage) sendEventSuccessRespNoData(event string) {
	r.respMessagesChan <- &EventData{
		Event: event,
	}
}

func (r *RespMessage) sendEventSuccessRespWithData(event string, data string) {
	r.respMessagesChan <- &EventData{
		Event: event,
		Data:  data,
	}
}

func (r *RespMessage) sendEventFailedRespNoData(event string, errCode int32, errMsg string) {
	r.respMessagesChan <- &EventData{
		Event:   event,
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}

func (r *RespMessage) sendEventFailedREspNoErr(event string) {
	r.respMessagesChan <- &EventData{
		Event: event,
	}
}