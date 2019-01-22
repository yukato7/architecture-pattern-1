package service

import (
	"context"
	"fmt"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/yutify/gengo-api/domain/model"
	"github.com/yutify/gengo-api/usecase/repository"
	"log"
	"os"
)

type ChargeService interface {
	ChargeMoney(ctx context.Context, chargeInfo *operations.CreateCharge, userID string) error
}

type chargeService struct {
	ChargeRepo repository.ChargeRepository
	UserRepo   repository.UserRepository
}

func NewClient() (*Client, error) {
	oc, e := omise.NewClient(os.Getenv("OMISE_PUBLIC_KEY"), os.Getenv("OMISE_SECREAT_KEY"))
	if e != nil {
		log.Fatal(e)
	}
	client := &Client{oc}
	return client, e
}

func NewChargeService(cr repository.ChargeRepository, ur repository.UserRepository) ChargeService {
	return &chargeService{
		ChargeRepo: cr,
		UserRepo:   ur,
	}
}

type Client struct {
	*omise.Client
}

func (cr *chargeService) ChargeMoney(ctx context.Context, chargeInfo *operations.CreateCharge, userID string) error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	charge, createCharge := &omise.Charge{}, chargeInfo
	if e := client.Do(charge, createCharge); e != nil {
		return err
	}
	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)

	cd := &model.ChargeDetail{
		UserID:   userID,
		Amount:   chargeInfo.Amount,
		Currency: chargeInfo.Currency,
	}
	ok, err := cr.UserRepo.UserExists(ctx, userID)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Errorf("user does not exist")
	}
	if err := cr.ChargeRepo.CreateChargeLog(ctx, cd); err != nil {
		return err
	}
	return nil
}
