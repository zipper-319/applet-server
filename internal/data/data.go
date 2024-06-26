package data

import (
	"applet-server/api/v2/applet"
	"applet-server/api/v2/applet/common"
	"applet-server/internal/conf"
	"applet-server/internal/data/cache"
	"applet-server/internal/data/minio"
	"applet-server/internal/data/mysql"
	"applet-server/internal/data/s3"
	"applet-server/internal/data/train"
	"applet-server/internal/pkg/log"
	"applet-server/internal/pkg/ws"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"go.uber.org/atomic"
	"gorm.io/gorm"
	"io"
	"os"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, s3.NewS3Service, cache.NewRedisCache, mysql.NewDataDB, minio.NewMinioClient, train.NewTrain)

// Data .
type Data struct {
	S3          *s3.S3Service
	RedisClient *redis.Client
	*gorm.DB
	*minio.Client
	Train *train.Train
	*conf.App
	envMap map[string]ServiceAddr
}

type ServiceAddr struct {
	TTS      string `json:"tts"`
	ASR      string `json:"asr"`
	NLP      string `json:"nlp"`
	Feedback string `json:"feedback"`
}

// NewData .
func NewData(s3 *s3.S3Service, rdb *redis.Client, db *gorm.DB, minioClient *minio.Client, trainObject *train.Train, app *conf.App) (*Data, error) {

	var envMap map[string]ServiceAddr

	if envContent, err := os.ReadFile("configs/env.json"); err != nil {
		log.Error(err.Error())
	} else {
		if err = json.Unmarshal(envContent, &envMap); err != nil {
			panic(err)
		}
	}
	log.Debugf("envMap: %v", envMap)
	return &Data{
		S3:          s3,
		RedisClient: rdb,
		DB:          db,
		Client:      minioClient,
		Train:       trainObject,
		envMap:      envMap,
		App:         app,
	}, nil
}

func (d *Data) GetASRAddr(env string) string {
	result := d.Asr.GetAddr()
	if tmp, ok := d.envMap[env]; ok {
		result = tmp.ASR
	}
	log.Debugf("envMap:%v, env: %s;asr addr:%s", d.envMap, env, result)
	return result
}

func (d *Data) GetTTSAddr(env string) string {
	result := d.Tts.GetAddr()
	if tmp, ok := d.envMap[env]; ok {
		result = tmp.TTS
	}
	log.Debugf("envMap:%v, env: %s; tts addr:%s", d.envMap, env, result)
	return result
}

func (d *Data) GetNLPAddr(env string) string {
	result := d.Nlp.GetAddr()
	if tmp, ok := d.envMap[env]; ok {
		result = tmp.NLP
	}
	log.Debugf("envMap:%v, env: %s; nlp addr:%s", d.envMap, env, result)
	return result
}

func (d *Data) GetFeedbackAddr(env string) string {
	result := d.Feedback.GetAddr()
	if tmp, ok := d.envMap[env]; ok {
		result = tmp.Feedback
	}
	log.Debugf("envMap:%v, env: %s; feedback addr:%s", d.envMap, env, result)
	return result
}

type Session struct {
	Id       string
	Username string
	RobotId  int32
	Position string
	AgentId  int
	Env      common.EnvType
	Language *atomic.String
	*ws.WsClient
	applet.MethodType
	TtsParam atomic.Value
	AsrParam atomic.Value
	NlpParam atomic.Value
}

//func NewSession(robotId int32, position string, agentId int, language string) *Session {
//	id := uuid.New().String()
//	return &Session{TraceId: id, RobotId: robotId, Position: position, AgentId: agentId, Language: language}
//}

type TalkResp struct {
	// 问题
	Question string `json:"question"`
	// 答案
	AnsItem []*Answer `json:"answer"`
	// 技能
	Source string `json:"source"`
	// 技能领域
	Domain string `json:"domain"`
	// 意图
	Intent string                          `json:"intent"`
	Entity []*applet.CollectLikeReq_Entity `json:"entity"`
}

type Answer struct {
	// 文本
	Text string `json:"text"`
	// 语言
	Lang string `json:"lang"`
	// 音频文件URL
	MusicUrl string `json:"musicUrl"`
	// 图片文件URL
	PicUrl string `json:"picUrl"`
	// 视频文件URL
	VideoUrl string `json:"videoUrl"`
}

type FileObject struct {
	File     io.Reader
	FileName string
	Username string
}

//func NewSession(robotId int32, position string, agentId int, language string) *Session {
//	id := uuid.New().String()
//	return &Session{TraceId: id, RobotId: robotId, Position: position, AgentId: agentId, Language: language}
//}

func GenSession(req applet.ChatWSReq, username, sessionId string, conn *websocket.Conn, logger *log.MyLogger) *Session {
	return &Session{
		Id:         sessionId,
		Username:   username,
		Position:   req.Position,
		AgentId:    int(req.AgentId),
		Language:   atomic.NewString(req.Language),
		Env:        req.EnvType,
		MethodType: req.Method,
		WsClient:   ws.NewWsClient(conn, logger),
	}
}

type ASRParam struct {
	AsrDomain string
}
type NlpParam struct {
	RobotId string
}

type TTSParam struct {
	Speed    string
	Volume   string
	Pitch    string
	Emotions string
	Speaker  string
	IsClone  bool
}

type ChatParameter struct {

	// tts
	Speed   string `json:"speed"`
	Volume  string ` json:"volume"`
	Pitch   string ` json:"pitch"`
	Speaker string ` json:"speaker"`
	IsClone int32  ` json:"is_clone"` // 1: 不是克隆， 2：是克隆
	// asr
	ServiceProvider int32  `json:"service_provider"` // 1:达闼  2：Google  3：科大讯飞  4:CPAll
	LanguageType    int32  `json:"language_type"`    //语言类型：    1:CH  2:EN 3:TCH 4:JA 5:ES 6:TH
	ServiceType     int32  `json:"service_type"`
	RobotId         string `json:"robot_id"`
	Agent           int32  `json:"agent"`
	Language        string ` json:"language"`
	Position        string ` json:"position"`
}
