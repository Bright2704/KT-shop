package entities

import (
	"github.com/Bright2704/KT-shop-tutorial/pkg/ktlogger"
	"github.com/gofiber/fiber/v2"
)

type IResponse interface {
	Success(code int, data any) IResponse
	Error(code int, tractId, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data  	   any
	ErrorRes   *ErrorResponse
	Context    *fiber.Ctx
	IsError    bool
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg	 string `json:"msg"`
}

func NewResponse(c *fiber.Ctx) IResponse {
	return &Response{
		Context: c,
	}
}

func (r *Response) Success(code int, data any) IResponse{
	r.StatusCode = code
	r.Data = data
	ktlogger.InitKtlogger(r.Context, &r.Data).Print().Save()
	return r
}
func (r *Response) Error(code int, tractId, msg string) IResponse {
	r.StatusCode = code 
	r.ErrorRes = &ErrorResponse{
		TraceId: tractId,
		Msg: 	 msg,
	}
	r.IsError = true
	ktlogger.InitKtlogger(r.Context, &r.ErrorRes).Print().Save()
	return r
}
func (r *Response) Res() error {
	
		return r.Context.Status(r.StatusCode).JSON(func() any {
			if r.IsError {
				return &r.ErrorRes
			}
			return &r.Data
		}())
	
	
}