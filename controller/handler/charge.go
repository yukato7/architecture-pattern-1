package handler

import (
	"github.com/omise/omise-go/operations"
	"github.com/yutify/gengo-api/usecase/service"
	"log"
	"net/http"
)

type ChargeHandler interface {
	CreateCharge(w http.ResponseWriter, r *http.Request)
}

type chargeRequest struct {
	tokenID  string `json:"token_id"`
	userID   string `json:"user_id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}
type chargeHandler struct {
	ChargeApp service.ChargeService
}

func NewChargeHandler(ca service.ChargeService) ChargeHandler {
	return &chargeHandler{
		ChargeApp: ca,
	}
}

func (ch *chargeHandler) CreateCharge(w http.ResponseWriter, r *http.Request) {
	//authHeader := r.Header.Get("Authorization")
	//tokenID := strings.Replace(authHeader, "Bearer ", "", 1)
	//if tokenID == "" {
	//	rendering.JSON(w, http.StatusUnauthorized, "required authorization")
	//}
	ctx := r.Context()
	var cr chargeRequest
	if err := decodeRequest(r.Body, &cr); err != nil {
		log.Fatal(err)
	}
	cc := &operations.CreateCharge{
		Amount:   cr.Amount,
		Currency: cr.Currency,
		Card:     cr.tokenID,
	}
	ch.ChargeApp.ChargeMoney(ctx, cc, cr.userID)
	rendering.JSON(w, http.StatusOK, "ok.")
}
