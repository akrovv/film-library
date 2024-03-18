package restapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/pkg/logger"
	"github.com/akrovv/filmlibrary/pkg/sender"
)

type actorController struct {
	logger  logger.Logger
	service ActorService
}

func NewActorController(logger logger.Logger, service ActorService) *actorController {
	return &actorController{
		logger:  logger,
		service: service,
	}
}

func (c *actorController) ManagePath(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "POST":
		c.Create(w, r)
	case "PUT":
		c.Update(w, r)
	case "DELETE":
		c.Delete(w, r)
	case "GET":
		c.GetList(w, r)
	}
}

// @Summary Create
// @Description Create an actor in the film library
// @Tags		 actor
// @Accept       json
// @Produce      json
// @Param request body domain.CreateActor true "request"
// @Success 201 {object} sender.JSONResponse
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /actor [post]
func (c *actorController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		c.logger.Infof("request didnt contain application/json")
		sender.ErrorJSON(w, domain.ErrRequest, http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		c.logger.Infof("io.ReadAll error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	createActorDTO := domain.CreateActor{}
	err = json.Unmarshal(data, &createActorDTO)

	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = c.service.Create(&createActorDTO); err != nil {
		c.logger.Infof("c.ActorService.Create error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = sender.WriteJSON(w, http.StatusCreated, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "actor was created",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}

// @Summary Update
// @Description  Update details of an existing actor in the film library
// @Tags		 actor
// @Accept       json
// @Produce      json
// @Param request body domain.Actor true "request"
// @Success 200 {object} sender.JSONResponse
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /actor [put]
func (c *actorController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		c.logger.Infof("request didnt contain application/json")
		sender.ErrorJSON(w, domain.ErrRequest, http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		c.logger.Infof("io.ReadAll error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	actor := domain.Actor{}
	err = json.Unmarshal(data, &actor)

	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = c.service.Update(&actor); err != nil {
		c.logger.Infof("c.ActorService.Update error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = sender.WriteJSON(w, http.StatusOK, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "actor was updated",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}

// @Summary Delete
// @Description  Delete an actor from the film library
// @Tags		 actor
// @Accept       json
// @Produce      json
// @Param request body domain.DeleteActor true "request"
// @Success 200 {object} sender.JSONResponse
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /actor [delete]
func (c *actorController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		c.logger.Infof("request didnt contain application/json")
		sender.ErrorJSON(w, domain.ErrRequest, http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		c.logger.Infof("io.ReadAll error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	deleteActorDTO := domain.DeleteActor{}
	err = json.Unmarshal(data, &deleteActorDTO)

	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = c.service.Delete(&deleteActorDTO); err != nil {
		c.logger.Infof("c.ActorService.Delete error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = sender.WriteJSON(w, http.StatusOK, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "actor was deleted",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}

// @Summary GetList
// @Description  Get a list of all actors available in the film library
// @Tags		 actor
// @Accept       json
// @Produce      json
// @Success 200 {object} []domain.ActorWithMovie
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /actor [get]
func (c *actorController) GetList(w http.ResponseWriter, r *http.Request) {
	actors, err := c.service.GetList()

	if err != nil {
		c.logger.Infof("c.ActorService.GetList error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	jsonResult, err := json.Marshal(actors)
	if err != nil {
		c.logger.Infof("json.Marshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResult)
	if err != nil {
		c.logger.Infof("Write %w", err)
		return
	}
}
