package web

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jayvib/golog"
	"gophr/v2/gophr.api/user"
	"net/http"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Method  string      `json:"method,omitempty"`
}

type Parameters struct {
	UserService user.Service `validation:"required"`
}

func (p *Parameters) Validate() error {
	validation := validator.New()
	return validation.Struct(p)
}

func New(param *Parameters) *Handler {
	if err := param.Validate(); err != nil {
		panic(err)
	}

	return &Handler{
		svc: param.UserService,
	}
}


type Handler struct {
	svc user.Service
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request)       {
	vars := mux.Vars(r)
	golog.Debug("vars:", vars)
	userId := vars["id"]
	golog.Debug("id:", userId)

	usr, _ := h.svc.GetByID(r.Context(), userId)

	golog.Debug("user:", usr)
	response := &Response{
		Message: "OK",
		Data: usr,
		Method: r.Method,
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		golog.Error(err)
	}

}
func (h *Handler) GetByEmail(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) GetByUsername(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Save(w http.ResponseWriter, r *http.Request)          {}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request)         {}
