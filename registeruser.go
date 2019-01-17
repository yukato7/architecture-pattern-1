package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type RegisterUserRequest struct {
	ID string `json:"id"`
	Name string `json:"name"`
	IconURL string `json:"icon_url"`
}

type User struct {
	ID string `dynamo:"id"`
	Name string `dynamo:"name"`
	IconURL string `dynamo:"icon_url"`
}


type AWSCredential struct {
	AccessKeyID string `toml:"access_key_id"`
	SecretKeyID string `toml:"secret_key_id"`
}

type Config struct {
	AWS AWSCredential `toml:"aws_credential"`
}

func RegisterUser(request RegisterUserRequest) error {
	u := &User{
		ID: request.ID,
		Name: request.Name,
		IconURL: request.IconURL,
	}
	var conf Config

	if _, err := toml.DecodeFile("./.config/config.toml", &conf); err != nil {
		return err
	}
	// create aws session
	credentials := credentials.NewStaticCredentials(conf.AWS.AccessKeyID, conf.AWS.SecretKeyID, "")
	region := "ap-northeast-2"
	config := aws.NewConfig().WithCredentials(credentials).WithRegion(region)
	sess, err := session.NewSession(config)
	if err != nil {
		return err
	}
	// connect with dynamodb
	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(u)
	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("Users"),
	}
	// put user item
	if _, err := svc.PutItem(input); err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return nil
}

func main() {
	lambda.Start(RegisterUser)
}