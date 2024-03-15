package util

import (
	"applet-server/api/v2/applet/common"
)

const REDIS_KEY_AWS_S3_USER_Prefix = "aws_s3:user"

const MediaApiUrl = "http://172.16.31.96:30247"

const DefaultTrainDuration = 240

func IsSpeakerExist(speakers []string, submittedSpeaker string) bool {
	for _, v := range speakers {
		if v == submittedSpeaker {
			return true
		}
	}
	return false
}

var EnvTypeName = map[common.EnvType]string{
	0: "生成环境",
	1: "86环境",
	2: "251环境",
	3: "85环境",
	4: "87环境",
	5: "134环境",
}
