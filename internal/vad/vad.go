package vad

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L${SRCDIR}/libs -lCmVadWrapper -ldl -lstdc++
#include <stdio.h>
#include "vad_wrapper.h"
#include "vad_wrapper_x86_extend.h"
extern void goVoiceStart(void*);
extern void goVoiceEnd(void*,int);
extern void goVoiceData(void*, char*,int,int);
extern void goNonVoiceData(void*, short*,int);

typedef  void (*tpyVoiceStart)(void* pUserData);//语音开始
typedef void (*tpyVoiceEnd)(void* pUserData, int flags);//语音结束
typedef void (*tpyVoiceData)(void* pUserData, char *data, int size, int flags);
typedef void (*tpyNonVoiceData)(void* pUserData, char *data, int size);//非语音数据
*/
import "C"
import (
	"applet-server/internal/data"
	"applet-server/internal/pkg/pointer"
	"applet-server/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"unsafe"
)

type DesiredSsl struct {
	Dimension int
	Size      int
}

type VadConfig struct {
	IsEnableSslFilter    int        // 是否使能声源方位（SSL）过滤
	DesiredSslRange      DesiredSsl // 期望的声源方位范围
	SslThreshold         float32    // 判决门限，0~1之间.
	PauseMode            int        // 断句模式. 0--极速模式（默认）1--交互模式（600ms）2--听写模式（1200ms） 只有当vad mode等于3时，断句模式的设置才有效.
	IsEnableEnergyFilter int        // 是否使能能量过滤
}

type DataInfo struct {
	InputCh    chan []byte
	OutputCh   chan []byte
	IsEnd      chan struct{}
	Conn       *websocket.Conn
	Session    *data.Session
	Server     service.ChatWebsocketServer
	LastCancel context.CancelFunc
	log.Logger
}

var callback C.VAD_Callback

var vadConfig VadConfig

func init() {
	callback = C.VAD_Callback{}
	callback.VoiceStart = C.tpyVoiceStart(C.goVoiceStart)
	callback.VoiceEnd = C.tpyVoiceEnd(C.goVoiceEnd)
	callback.VoiceData = C.tpyVoiceData(C.goVoiceData)
	callback.NonVoiceData = C.tpyNonVoiceData(C.goNonVoiceData)
	vadConfig = VadConfig{
		IsEnableSslFilter: 0,
		DesiredSslRange: DesiredSsl{
			Dimension: 1,
			Size:      1,
		},
		SslThreshold:         0.5,
		IsEnableEnergyFilter: 0,
		PauseMode:            0,
	}
}
func FeedAudio(ctx context.Context, vadDateInfo *DataInfo) {
	pData, err := pointer.Save(vadDateInfo)
	if err != nil {
		panic(err)
	}
	handle := C.VAD_Wrapper_Create(&callback, unsafe.Pointer(uintptr(pData)), C.int(3))
	//// 打开一个文件保存音频
	//file, err := os.OpenFile(fmt.Sprintf("asset/%s.pcm", vadDateInfo.Session.Id), os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Info(err)
	//}

	defer func() {
		C.VAD_Wrapper_Destory(handle)
		pointer.Unref(pData)
		//file.Close()
		log.Debugf("finish to destroy vad")
	}()

	desireSsl := vadConfig.DesiredSslRange
	length := desireSsl.Size * desireSsl.Dimension * 2 * int(unsafe.Sizeof(vadConfig.SslThreshold))
	desireSslFloat := make([]float32, length)
	C.VAD_Wrapper_SetVadConfig(handle, C.int(vadConfig.IsEnableSslFilter), C.int(desireSsl.Dimension), C.int(desireSsl.Size),
		(*C.float)(unsafe.Pointer(&desireSslFloat)), C.float(vadConfig.SslThreshold), C.int(vadConfig.PauseMode), C.int(vadConfig.IsEnableEnergyFilter))

	for {
		select {
		case voiceData, isOpen := <-vadDateInfo.InputCh:
			if isOpen {
				size := len(voiceData)
				if size <= 0 {
					break
				}

				ok := C.VAD_Wrapper_FeedAudio(handle, (*C.char)(unsafe.Pointer(&voiceData[0])), C.int(size), nil)
				log.Debugf("voiceData length: %d; VAD_FeedAudio result:%d", size, int(ok))
				//file.Write(voiceData)
			} else {
				log.Debugf("vadDateInfo channel closed")
				return
			}
			//防止未接收到结束帧，而没有回收资源；超时取消
		//case <-vadTimer.C:
		//	log.Debugf("vad timer out")
		//	return
		case <-ctx.Done():
			log.Debugf("vad context done")
			return
		}

	}

}
