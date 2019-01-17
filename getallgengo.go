package main

import "github.com/aws/aws-lambda-go/lambda"

type GetAllGengoRequest struct {
	MaxLimit uint `json:"max_limit"`
}

type GetAllGengoResponse struct {
	Gengoes []Gengo `json:"gengoes"`
}

type Gengo struct {
	Name string `json:"name"`
	Voters []Voter `json:"voters"`
	//Todo add Rank order by charge amount
}

type Voter struct {
	Name string `json:"name"`
	IconURL string `json:"icon_url"`
	ChargeAmount uint `json:"charge_amount"`
}

func GetAllGengo(request GetAllGengoRequest) (GetAllGengoResponse, error) {
	// fetch AllGengo limited by MaxLimit

	return GetAllGengoResponse{}, nil
}

func main() {
	lambda.Start(GetAllGengo)
}