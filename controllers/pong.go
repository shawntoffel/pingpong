package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shawntoffel/pingpong/models"
	"github.com/shawntoffel/pingpong/utilities/httpUtils"
	"github.com/shawntoffel/pingpong/utilities/responseUtils"
)

type (
	PongController struct{}
)

func NewPongController() *PongController {
	return &PongController{}
}

// GetPong returns a Pong containing the client IP address.
func (pongController PongController) GetPong(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	ip, err := httpUtils.GetRemoteIp(req.RemoteAddr)

	if err != nil {
		responseUtils.JsonResponse(w, err)
		return
	}

	pong := models.Pong{
		Pong: ip,
	}

	responseUtils.JsonResponse(w, pong)
}
