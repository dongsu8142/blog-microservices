package main

import (
	"errors"
	"net/http"

	common "github.com/dongsu8142/blog-common"
	pb "github.com/dongsu8142/blog-common/api"
	"github.com/dongsu8142/blog-gateway/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	gateway gateway.UserGateway
}

func NewHandler(gateway gateway.UserGateway) *handler {
	return &handler{gateway}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/register", h.HandleRegisterUser)
	mux.HandleFunc("POST /api/auth/login", h.HandleLoginUser)
}

func (h *handler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var user *pb.RegisterUserRequest
	if err := common.ReadJSON(r, &user); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateUser(user); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	u, err := h.gateway.RegisterUser(r.Context(), user)

	if rStatus := status.Convert(err); rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, u)
}

func (h *handler) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	var user *pb.LoginUserRequest
	if err := common.ReadJSON(r, &user); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	u, err := h.gateway.LoginUser(r.Context(), user)

	if rStatus := status.Convert(err); rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, u)
}

func validateUser(user *pb.RegisterUserRequest) error {
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
