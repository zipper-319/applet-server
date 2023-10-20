package vad

import (
	"applet-server/internal/pkg/pointer"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"unsafe"
)

import "C"

//export goVoiceStart
func goVoiceStart(pUserData unsafe.Pointer) {
	pDataInfo, ok := pointer.Load(pUserData).(*DataInfo)
	if !ok {
		log.Error("goVoiceStart; pUserData not found")
		return
	}
	if pDataInfo.LastCancel != nil {
		pDataInfo.LastCancel()
		err := pDataInfo.Server.CancelHandlerDHuman(context.Background(), pDataInfo.Session)
		log.Debugf("goVoiceStart; CancelHandlerDHuman; err=%v", err)
	}
	pDataInfo.OutputCh = make(chan []byte, 50)
	ctx, cancel := context.WithCancel(context.Background())
	pDataInfo.LastCancel = cancel
	session := pDataInfo.Session
	session.TraceId = uuid.New().String()
	go pDataInfo.Server.HandlerVoice(ctx, pDataInfo.OutputCh, pDataInfo.Conn, session)
	log.Infof("goOnStart:%d", pUserData)
}

//export goVoiceEnd
func goVoiceEnd(pUserData unsafe.Pointer, info *C.char) {

	pDataInfo, ok := pointer.Load(pUserData).(*DataInfo)
	if !ok {
		log.Error("goVoiceEnd; pUserData not found")
		return
	}
	close(pDataInfo.OutputCh)
	log.Debugf("goVoiceEnd; send receive end flag; pUserData:%d", pUserData)

}

//export goVoiceData
func goVoiceData(pUserData unsafe.Pointer, voice *C.char, size C.uint, flags C.uint) {

	pDataInfo, ok := pointer.Load(pUserData).(*DataInfo)
	if !ok {
		log.Error("goVoiceEnd; pUserData not found")
		return
	}

	voiceData := make([]byte, size)
	temp := voice
	for i := C.uint(0); i < size; i++ {
		voiceData[i] = byte(*temp)
		temp = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(temp)) + 1))
	}
	log.Debugf("goVoiceData；send verbal voice; the length:%d, pUserData:%d ", len(voiceData), pUserData)
	pDataInfo.OutputCh <- voiceData

}

//export goNonVoiceData
func goNonVoiceData(pUserData unsafe.Pointer, data *C.char, size C.uint) {

}
