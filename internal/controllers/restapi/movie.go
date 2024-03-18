package restapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/pkg/logger"
	"github.com/akrovv/filmlibrary/pkg/sender"
)

type movieController struct {
	logger  logger.Logger
	service MovieService
}

func NewMovieController(logger logger.Logger, service MovieService) *movieController {
	return &movieController{
		logger:  logger,
		service: service,
	}
}

func (c *movieController) ManagePath(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "POST":
		c.Create(w, r)
	case "PUT":
		c.Update(w, r)
	case "DELETE":
		c.Delete(w, r)
	case "GET":
		c.Get(w, r)
	}
}

// @Summary Create
// @Description  Create a new movie in the film library
// @Tags		 movie
// @Accept       json
// @Produce      json
// @Param request body domain.CreateMovie true "request"
// @Success 200 {object} sender.JSONResponse
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /movie [post]
func (c *movieController) Create(w http.ResponseWriter, r *http.Request) {
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

	createMovieDTO := domain.CreateMovie{}
	err = json.Unmarshal(data, &createMovieDTO)
	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = c.service.Create(&createMovieDTO); err != nil {
		c.logger.Infof("c.MovieService.Create error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = sender.WriteJSON(w, http.StatusCreated, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "movie was created",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}

// @Summary Update
// @Description  Update details of an existing movie in the film library
// @Tags		 movie
// @Accept       json
// @Produce      json
// @Param request body domain.Movie true "request"
// @Success 200
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /movie [put]
func (c *movieController) Update(w http.ResponseWriter, r *http.Request) {
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

	movie := domain.Movie{}
	err = json.Unmarshal(data, &movie)

	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = c.service.Update(&movie); err != nil {
		c.logger.Infof("c.MovieService.Update error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = sender.WriteJSON(w, http.StatusOK, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "movie was updated",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}

// @Summary Delete
// @Description  Delete a movie from the film library
// @Tags		 movie
// @Accept       json
// @Produce      json
// @Param request body domain.Movie true "request"
// @Success 200
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /movie [delete]
func (c *movieController) Delete(w http.ResponseWriter, r *http.Request) {
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

	deleteMove := domain.DeleteMovie{}
	err = json.Unmarshal(data, &deleteMove)

	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = c.service.Delete(&deleteMove); err != nil {
		c.logger.Infof("c.MovieService.Delete error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err = sender.WriteJSON(w, http.StatusOK, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "movie was deleted",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}

// @Summary Get
// @Description  Delete a movie from the film library
// @Tags		 movie
// @Accept       json
// @Produce      json
// @Param title query string true "Movie title"
// @Param actor query string true "Actor name"
// @Success 200 {object} domain.MovieWithoudID
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /movie [get]
func (c *movieController) Get(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	name := r.URL.Query().Get("actor")

	if title == "" && name == "" {
		c.logger.Infof("empty query: %w", domain.ErrRequest)
		sender.ErrorJSON(w, domain.ErrRequest, http.StatusBadRequest)
		return
	}

	var searchType string
	if title != "" {
		searchType = "title"
	} else {
		searchType = "actor"
	}

	getMovieDTO := domain.GetMovie{
		Title:      title,
		ActorName:  name,
		SearchType: searchType,
	}

	movie, err := c.service.Get(&getMovieDTO)
	if err != nil {
		c.logger.Infof("c.MovieService.Get error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	jsonResult, err := json.Marshal(movie)
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

// @Summary GetList
// @Description  Delete a movie from the film library
// @Tags		 movie
// @Accept       json
// @Produce      json
// @Param order query string false "Order"
// @Success 200 {object} domain.GetOrderedMovie
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /movie/all [get]
func (c *movieController) GetOrderedList(w http.ResponseWriter, r *http.Request) {
	order := r.URL.Query().Get("order")

	if order == "" {
		order = "rating"
	}

	getOrderedDTO := domain.GetOrderedMovie{
		Order: order,
	}

	movies, err := c.service.GetOrderedList(&getOrderedDTO)
	if err != nil {
		c.logger.Infof("c.MovieService.GetOrderedList error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	jsonResult, err := json.Marshal(movies)
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
