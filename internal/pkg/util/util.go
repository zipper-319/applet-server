package util

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
