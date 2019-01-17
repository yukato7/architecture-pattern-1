package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"log"
	"os"
)

type chargeRequest struct {
	Gengo string `json:"gengo"`
	UserID string `json:"user_id"`
	TokenID string `json:"token_id"`
	Amount int64 `json:"amount"`
	Currency string `json:"currency"`
}

type Client struct {
	*omise.Client
}

func NewClient() (*Client, error) {
	oc , e := omise.NewClient(os.Getenv("OMISE_PUBLIC_KEY"), os.Getenv("OMISE_SECREAT_KEY"))
	if e != nil {
		log.Fatal(e)
	}
	client := &Client{oc}
	return  client, e
}


func (c *Client) ChargeMoney(chargeInfo *operations.CreateCharge) {
	charge, createCharge := &omise.Charge{}, chargeInfo
	if e := c.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}
	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
}

func CreateCharge(cr chargeRequest) {
	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	cc := &operations.CreateCharge{
		Amount: cr.Amount,
		Currency: cr.Currency,
		Card: cr.TokenID,
	}
	client.ChargeMoney(cc)
	// ユーザ課金テーブルに課金データを永続化
}

func main() {
	lambda.Start(CreateCharge)
}