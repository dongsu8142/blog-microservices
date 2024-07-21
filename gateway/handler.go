package main

import (
	"errors"
	"net/http"

	common "github.com/dongsu8142/blog-common"
	pb "github.com/dongsu8142/blog-common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.UserServiceClient
}

func NewHandler(client pb.UserServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/register", h.HandleRegisterUser)
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

	u, err := h.client.RegisterUser(r.Context(), user)
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
		return errors.New("Username is required")
	}
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if user.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}
