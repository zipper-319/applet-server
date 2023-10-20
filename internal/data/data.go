package data

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data/cache"
	"applet-server/internal/data/mysql"
	"applet-server/internal/data/s3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/atomic"
	"gorm.io/gorm"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, s3.NewS3Service, cache.NewRedisCache, mysql.NewDataDB)

// Data .
type Data struct {
	S3          *s3.S3Service
	RedisClient *redis.Client
	*gorm.DB
}

// NewData .
func NewData(s3 *s3.S3Service, rdb *redis.Client, db *gorm.DB) (*Data, error) {
	return &Data{
		S3:          s3,
		RedisClient: rdb,
		DB:          db,
	}, nil
}

type Session struct {
	Id       string
	TraceId  string
	RobotId  int32
	Position string
	AgentId  int
	Language *atomic.String
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
	Intent string `json:"intent"`
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

type ChatServerMessage struct {
	ServiceType applet.ServiceType ` json:"service_type,omitempty"`
	Content     interface{}        ` json:"content,omitempty"`
	IsEnd       bool               `son:"is_end,omitempty"`
	IsSuccess   bool               `json:"is_success,omitempty"`
	ErrMsg      string             `json:"err_msg,omitempty"`
}
