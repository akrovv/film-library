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

func TestActorCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	as := mocks.NewMockActorService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	actorHandler := NewActorController(logger, as)

	body := `{
		"actor_name": "user",
		"gender": "Male", 
		"date_of_birth": "2007-02-02T00:00:00Z"
	  }`

	expectedTime, err := time.Parse(time.RFC3339, "2007-02-02T00:00:00Z")

	if err != nil {
		t.Fatalf("can't parse time: %s", err)
	}

	createActorDTO := domain.CreateActor{
		Name:      "user",
		Gender:    "Male",
		DateBirth: expectedTime,
	}

	req := httptest.NewRequest("POST", "/actor", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	as.EXPECT().Create(&createActorDTO).Return(nil)

	actorHandler.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got: %d", w.Code)
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	actorHandler.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
	}

	// Incorrect JSON
	req = httptest.NewRequest("POST", "/actor", strings.NewReader(`{"actor_name": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	actorHandler.Create(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("POST", "/actor", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	actorHandler.Create(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// Register returned error
	req = httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	as.EXPECT().Create(&createActorDTO).Return(domain.ErrTest)
	actorHandler.Create(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}
}

func TestActorUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	as := mocks.NewMockActorService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	actorHandler := NewActorController(logger, as)

	body := `{
		"actor_id": 1,
		"actor_name": "user",
		"gender": "Male", 
		"date_of_birth": "2007-02-02T00:00:00Z"
	  }`

	expectedTime, err := time.Parse(time.RFC3339, "2007-02-02T00:00:00Z")

	if err != nil {
		t.Fatalf("can't parse time: %s", err)
	}

	actor := domain.Actor{
		ID:        1,
		Name:      "user",
		Gender:    "Male",
		DateBirth: expectedTime,
	}

	req := httptest.NewRequest("PUT", "/actor", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	as.EXPECT().Update(&actor).Return(nil)
	actorHandler.Update(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	actorHandler.Update(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
	}

	// Incorrect JSON
	req = httptest.NewRequest("PUT", "/actor", strings.NewReader(`{"actor_id": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	actorHandler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("PUT", "/actor", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	actorHandler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// Register returned error
	req = httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	as.EXPECT().Update(&actor).Return(domain.ErrTest)
	actorHandler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}
}

func TestActorDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	as := mocks.NewMockActorService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	actorHandler := NewActorController(logger, as)

	body := `{
		"actor_id": 1
	  }`

	deleteActor := domain.DeleteActor{
		ID: 1,
	}

	req := httptest.NewRequest("DELETE", "/actor", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	as.EXPECT().Delete(&deleteActor).Return(nil)
	actorHandler.Delete(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	actorHandler.Delete(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
	}

	// Incorrect JSON
	req = httptest.NewRequest("DELETE", "/actor", strings.NewReader(`{"actor_id": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	actorHandler.Delete(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("DELETE", "/actor", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	actorHandler.Delete(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
	}

	// Delete returned error
	req = httptest.NewRequest("DELETE", "/actor", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	as.EXPECT().Delete(&deleteActor).Return(domain.ErrTest)
	actorHandler.Delete(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}
}

func TestActorGetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	as := mocks.NewMockActorService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	actorHandler := NewActorController(logger, as)

	actors := []domain.ActorWithMovie{}
	req := httptest.NewRequest("GET", "/actor", nil)
	w := httptest.NewRecorder()

	// OK
	as.EXPECT().GetList().Return(actors, nil)
	actorHandler.GetList(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
	}

	// GetList returned error
	w = httptest.NewRecorder()

	as.EXPECT().GetList().Return(nil, domain.ErrTest)
	actorHandler.GetList(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}
}

func TestActorManagePath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	as := mocks.NewMockActorService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	actorHandler := NewActorController(logger, as)

	// POST
	req := httptest.NewRequest("POST", "/actor", nil)
	w := httptest.NewRecorder()
	actorHandler.ManagePath(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 200, got: %d", w.Code)
		return
	}

	// PUT
	req = httptest.NewRequest("PUT", "/actor", nil)
	w = httptest.NewRecorder()
	actorHandler.ManagePath(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 200, got: %d", w.Code)
		return
	}

	// DELETE
	req = httptest.NewRequest("DELETE", "/actor", nil)
	w = httptest.NewRecorder()
	actorHandler.ManagePath(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 200, got: %d", w.Code)
		return
	}

	// GET
	as.EXPECT().GetList().Return(nil, nil)
	req = httptest.NewRequest("GET", "/actor", nil)
	w = httptest.NewRecorder()

	actorHandler.ManagePath(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got: %d", w.Code)
		return
	}
}
