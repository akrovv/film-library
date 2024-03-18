package restapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/internal/service/mocks"
	"github.com/akrovv/filmlibrary/pkg/logger"
	"github.com/golang/mock/gomock"
)

func TestMovieCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ms := mocks.NewMockMovieService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	movieHandler := NewMovieController(logger, ms)

	body := `{
		"movie_title": "title",
		"description": "description", 
		"release_date": "2007-02-02T00:00:00Z",
		"rating": 5,
		"actors": [1, 2]
	  }`

	expectedTime, err := time.Parse(time.RFC3339, "2007-02-02T00:00:00Z")

	if err != nil {
		t.Fatalf("can't parse time: %s", err)
	}

	createMovieDTO := domain.CreateMovie{
		Title:       "title",
		Description: "description",
		ReleaseDate: expectedTime,
		Rating:      5,
		Actors:      []int64{1, 2},
	}

	req := httptest.NewRequest("POST", "/movie", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	ms.EXPECT().Create(&createMovieDTO).Return(nil)
	movieHandler.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got: %d", w.Code)
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	movieHandler.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
	}

	// Incorrect JSON
	req = httptest.NewRequest("POST", "/movie", strings.NewReader(`{"movie_title": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	movieHandler.Create(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("POST", "/movie", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	movieHandler.Create(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// Create returned error
	req = httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	ms.EXPECT().Create(&createMovieDTO).Return(domain.ErrTest)
	movieHandler.Create(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}
}

func TestMovieUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ms := mocks.NewMockMovieService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	movieHandler := NewMovieController(logger, ms)

	body := `{
		"movie_id": 1,
		"movie_title": "title",
		"description": "description", 
		"release_date": "2007-02-02T00:00:00Z",
		"rating": 5,
		"actors": [1, 2]
	  }`

	expectedTime, err := time.Parse(time.RFC3339, "2007-02-02T00:00:00Z")

	if err != nil {
		t.Fatalf("can't parse time: %s", err)
	}

	movie := domain.Movie{
		ID:          1,
		Title:       "title",
		Description: "description",
		ReleaseDate: expectedTime,
		Rating:      5,
		Actors:      []int64{1, 2},
	}

	req := httptest.NewRequest("PUT", "/movie", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	ms.EXPECT().Update(&movie).Return(nil)
	movieHandler.Update(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	movieHandler.Update(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
	}

	// Incorrect JSON
	req = httptest.NewRequest("PUT", "/movie", strings.NewReader(`{"movie_title": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	movieHandler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("PUT", "/movie", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	movieHandler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// Update returned error
	req = httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	ms.EXPECT().Update(&movie).Return(domain.ErrTest)
	movieHandler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}
}

func TestMovieDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ms := mocks.NewMockMovieService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	movieHandler := NewMovieController(logger, ms)

	body := `{
		"movie_id": 1
	  }`

	deleteMovie := domain.DeleteMovie{
		ID: 1,
	}

	req := httptest.NewRequest("DELETE", "/movie", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	ms.EXPECT().Delete(&deleteMovie).Return(nil)
	movieHandler.Delete(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	movieHandler.Delete(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
	}

	// Incorrect JSON
	req = httptest.NewRequest("DELETE", "/movie", strings.NewReader(`{"movie_title": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	movieHandler.Delete(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("DELETE", "/movie", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	movieHandler.Delete(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// Delete returned error
	req = httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	ms.EXPECT().Delete(&deleteMovie).Return(domain.ErrTest)
	movieHandler.Delete(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}
}

func TestMovieGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ms := mocks.NewMockMovieService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	movieHandler := NewMovieController(logger, ms)
	getMovieTitle := domain.GetMovie{
		Title:      "br",
		SearchType: "title",
	}

	getMovieActor := domain.GetMovie{
		ActorName:  "ti",
		SearchType: "actor",
	}

	req := httptest.NewRequest("GET", "/movie?title=br", nil)
	w := httptest.NewRecorder()

	// OK. Title
	ms.EXPECT().Get(&getMovieTitle).Return(&domain.MovieWithoudID{}, nil)
	movieHandler.Get(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// OK. ActorName
	req = httptest.NewRequest("GET", "/movie?actor=ti", nil)
	w = httptest.NewRecorder()

	ms.EXPECT().Get(&getMovieActor).Return(&domain.MovieWithoudID{}, nil)
	movieHandler.Get(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// Get returned error
	req = httptest.NewRequest("GET", "/movie?title=br", nil)
	w = httptest.NewRecorder()

	ms.EXPECT().Get(&getMovieTitle).Return(nil, domain.ErrTest)
	movieHandler.Get(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// Empty query
	req = httptest.NewRequest("GET", "/movie", nil)
	w = httptest.NewRecorder()

	movieHandler.Get(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
	}
}

func TestMovieGetOrderedList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ms := mocks.NewMockMovieService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	movieHandler := NewMovieController(logger, ms)

	orderByMovieTitle := domain.GetOrderedMovie{
		Order: "movie_title",
	}

	orderByDefault := domain.GetOrderedMovie{
		Order: "rating",
	}

	req := httptest.NewRequest("GET", "/movie/all?order=movie_title", nil)
	w := httptest.NewRecorder()

	// OK. Order by movie title
	ms.EXPECT().GetOrderedList(&orderByMovieTitle).Return([]domain.MovieWithoudID{}, nil)
	movieHandler.GetOrderedList(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// OK. Missed order
	req = httptest.NewRequest("GET", "/movie/all", nil)
	w = httptest.NewRecorder()

	ms.EXPECT().GetOrderedList(&orderByDefault).Return([]domain.MovieWithoudID{}, nil)
	movieHandler.GetOrderedList(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// GetOrderedList returned error
	req = httptest.NewRequest("GET", "/movie?title=br", nil)
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	ms.EXPECT().GetOrderedList(&orderByDefault).Return(nil, domain.ErrTest)
	movieHandler.GetOrderedList(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}
}
