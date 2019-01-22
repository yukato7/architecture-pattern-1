package handler

import (
	"encoding/json"
	"fmt"
	"github.com/unrolled/render"
	"github.com/yutify/gengo-api/usecase/service"
	"io"
	"log"
)

var rendering = render.New(render.Options{})

type ApiHandler interface {
	UserHandler
	ChargeHandler
}

type apiHandler struct {
	UserHandler
	ChargeHandler
}

func NewHandler(ua service.UserService, ca service.ChargeService) ApiHandler {
	uh := NewUserHandler(ua)
	ch := NewChargeHandler(ca)
	handler := &apiHandler{uh, ch}
	return handler
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
