// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"applet-server/internal/biz"
	"applet-server/internal/biz/asr"
	"applet-server/internal/biz/nlp"
	"applet-server/internal/biz/tts"
	"applet-server/internal/conf"
	"applet-server/internal/data"
	"applet-server/internal/data/cache"
	"applet-server/internal/data/minio"
	"applet-server/internal/data/mysql"
	"applet-server/internal/data/s3"
	"applet-server/internal/data/train"
	"applet-server/internal/pkg/log"
	"applet-server/internal/server"
	"applet-server/internal/service"
	"github.com/go-kratos/kratos/v2"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, app *conf.App, confData *conf.Data, confLog *conf.Log) (*kratos.App, func(), error) {
	myLogger := log.NewLogger(confLog)
	s3Service, err := s3.NewS3Service(confData, myLogger)
	if err != nil {
		return nil, nil, err
	}
	client := cache.NewRedisCache(confData)
	db := mysql.NewDataDB(confData)
	minioClient := minio.NewMinioClient(confData)
	trainTrain := train.NewTrain(confData, myLogger)
	dataData, err := data.NewData(s3Service, client, db, minioClient, trainTrain)
	if err != nil {
		return nil, nil, err
	}
	videoUseCase := biz.NewVideoUseCase(dataData, myLogger)
	ttsClient := tts.NewTTSClient(app, dataData, myLogger)
	cloneSpeakerUseCase := biz.NewCloneSpeakerUseCase(dataData, ttsClient)
	voiceDataOperationService := service.NewVoiceDataOperationService(videoUseCase, cloneSpeakerUseCase, myLogger)
	grpcServer := server.NewGRPCServer(confServer, voiceDataOperationService)
	serverOption := server.NewMiddlewares(myLogger, app)
	userUseCase := biz.NewUserUseCase(dataData, myLogger)
	accountService := service.NewAccountService(userUseCase)
	cloneSpeakerService := service.NewCloneSpeakerService(cloneSpeakerUseCase)
	ttsServiceService := service.NewTTSServiceService(ttsClient, cloneSpeakerUseCase)
	asRControllerClient := asr.NewAsRControllerClient(app, dataData, myLogger)
	talkClient := nlp.NewTalkClient(app, dataData, myLogger)
	chatService := service.NewChatService(ttsClient, asRControllerClient, talkClient, myLogger)
	feedbackService := service.NewFeedbackService(app, myLogger)
	httpServer := server.NewHTTPServer(confServer, serverOption, voiceDataOperationService, accountService, cloneSpeakerService, ttsServiceService, chatService, feedbackService, myLogger)
	kratosApp := newApp(myLogger, grpcServer, httpServer)
	return kratosApp, func() {
	}, nil
}
