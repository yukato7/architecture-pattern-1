package handler

import (
	"encoding/json"
	"fmt"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/unrolled/render"
	"io"
	"log"
	"net/http"
	"os"
)

type chargeRequest struct {
	ID string `json:"id"`
	Amount int64 `json:"amount"`
	Currency string `json:"currency"`
}

type Client struct {
	*omise.Client
}

var rendering = render.New(render.Options{})

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

func CreateCharge(w http.ResponseWriter, r *http.Request) {
	//authHeader := r.Header.Get("Authorization")
	//tokenID := strings.Replace(authHeader, "Bearer ", "", 1)
	//if tokenID == "" {
	//	rendering.JSON(w, http.StatusUnauthorized, "required authorization")
	//}
	var cr chargeRequest
	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	if err := decodeRequest(r.Body, &cr); err != nil {
		log.Fatal(err)
	}
	cc := &operations.CreateCharge{
		Amount: cr.Amount,
		Currency: cr.Currency,
		Card: cr.ID,
	}
	client.ChargeMoney(cc)
	rendering.JSON(w, http.StatusOK, "ok.")
}

func decodeRequest(rb io.ReadCloser, dst interface{}) error {
	if err := json.NewDecoder(rb).Decode(dst); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("decode error")
	}
	// バリデーション
	//if err := validate.Struct(dst); err != nil {
	//	log.Println(err.Error())
	//	return fmt.Errorf("validate error")
	//}
	return nil
}