package httpcontroller

import (
	"client-info-service/internal/model"
	"client-info-service/internal/service"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"io"
	"log/slog"
	"net/http"
)

type ClientController struct {
	log      *slog.Logger
	services *service.Service
}

func NewClientController(log *slog.Logger, services *service.Service) *ClientController {
	return &ClientController{log: log, services: services}
}

func (ch ClientController) GetClients(w http.ResponseWriter, r *http.Request) {
	log := ch.log
	infos, err := ch.services.GetClients()
	if err != nil {
		log.Error(fmt.Sprintf("Could not get clients info: %s", err))
		w.WriteHeader(http.StatusServiceUnavailable)
		render.JSON(w, r, "could not get clients info")
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, infos)
}

func (ch ClientController) AddClient(w http.ResponseWriter, r *http.Request) {
	log := ch.log
	var client model.Client
	err := render.DecodeJSON(r.Body, &client)
	if errors.Is(err, io.EOF) {
		// Такую ошибку встретим, если получили запрос с пустым телом.
		// Обработаем её отдельно
		log.Error("request body is empty")

		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "empty request")
		return
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(client); err != nil {
		//validateErr := err.(validator.ValidationErrors)

		log.Error("invalid request", slog.Any("request", client))

		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Request is invalid")
		return
	}

	log.Info("Request body is decoded", slog.Any("request", client))

	clientId, err := ch.services.PutClientInQueue(client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, "Some problem adding user!")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	render.JSON(w, r, clientId)
}

func (ch ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "clientId")
	ch.log.Debug(idParam)
	clientId, err := uuid.Parse(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Id is not in uuid format!")
		return
	}
	err = ch.services.DeleteClientById(clientId)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		render.JSON(w, r, "Could not delete user")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (ch ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) {

}
