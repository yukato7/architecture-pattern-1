package main

import "github.com/aws/aws-lambda-go/lambda"

type PostGengoRequest struct {
	UserID string `json:"user_id"`
	Name string `json:"name"` //Todo add validate
}

type PostGengoResponse struct {
	Message string `json:"message"`
	ReceiptNumber uint `json:"receipt_number"`
}

func PostGengo(request PostGengoRequest) (PostGengoResponse, error) {
	// Insert Data
	return PostGengoResponse{}, nil
}

func main() {
	lambda.Start(PostGengo)
}