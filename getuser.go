package main

import (
	"github.com/BurntSushi/toml"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	Name string `json:"user_name"`
	IconURL string `json:"icon_url"`
	Gengo string `json:"gengo"`
	ChargeAmmount string `json:"charge_amount"`
}

func GetUser(request GetUserRequest) (GetUserResponse, error) {
	var conf Config

	if _, err := toml.DecodeFile("./.config/config.toml", &conf); err != nil {
		return GetUserResponse{}, err //Todo Add Error Response
	}
	// create aws session
	credentials := credentials.NewStaticCredentials(conf.AWS.AccessKeyID, conf.AWS.SecretKeyID, "")
	region := "ap-northeast-2"
	config := aws.NewConfig().WithCredentials(credentials).WithRegion(region)
	sess, err := session.NewSession(config)
	if err != nil {
		return GetUserResponse{}, nil //Todo Add Error Response
	}
	// connect with dynamodb
	svc := dynamodb.New(sess)
	// Todo add get process
}

func main() {
	lambda.Start(GetUser)
}