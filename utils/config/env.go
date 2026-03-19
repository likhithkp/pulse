package config

import (
	"errors"
	_const "pulse/utils/const"

	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	DeploymentEnv string
	MongodbUri    string
	Database      string
	Addr          string
	WebAddr       string
	// AwsRegion          string
	// AwsAccessKey       string
	// AwsSecretAccessKey string
	JwtSecretKey string
	// S3BucketName       string
	// S3BucketUrl        string
	// SenderEmail        string
}

func NewEnv() (*Env, error) {
	deploymentEnv := strings.TrimSpace(os.Getenv("DEPLOYMENT_ENV"))
	if deploymentEnv != _const.Deployment_Production {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}
	mongodbUri := os.Getenv("MONGODB_URI")
	if len(mongodbUri) == 0 {
		return nil, errors.New("MONGODB_URI is empty")
	}
	database := os.Getenv("DATABASE")
	if len(database) == 0 {
		return nil, errors.New("DATABASE is empty")
	}
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		return nil, errors.New("ADDR is empty")
	}
	webAddr := os.Getenv("WEB_ADDR")
	if len(webAddr) == 0 {
		return nil, errors.New("webAddr is empty")
	}
	// awsRegion := os.Getenv("AWS_REGION")
	// if len(awsRegion) == 0 {
	// 	return nil, errors.New("AWS_REGION is empty")
	// }
	// awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	// if len(awsAccessKey) == 0 {
	// 	return nil, errors.New("AWS_ACCESS_KEY is empty")
	// }
	// awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	// if len(awsSecretAccessKey) == 0 {
	// 	return nil, errors.New("AWS_SECRET_ACCESS_KEY is empty")
	// }
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if len(jwtSecretKey) == 0 {
		return nil, errors.New("JWT_SECRET_KEY is empty")
	}
	// s3BucketName := os.Getenv("S3_BUCKET_NAME")
	// if len(s3BucketName) == 0 {
	// 	return nil, errors.New("S3_BUCKET_NAME is empty")
	// }
	// senderEmail := os.Getenv("SENDER_EMAIL")
	// if len(senderEmail) == 0 {
	// 	return nil, errors.New("SENDER_EMAIL is empty")
	// }

	return &Env{
		DeploymentEnv: deploymentEnv,
		MongodbUri:    mongodbUri,
		Database:      database,
		Addr:          addr,
		WebAddr:       webAddr,
		// AwsRegion:          awsRegion,
		// AwsAccessKey:       awsAccessKey,
		// AwsSecretAccessKey: awsSecretAccessKey,
		JwtSecretKey: jwtSecretKey,
		// S3BucketName:       s3BucketName,
	}, nil
}
