package restapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/internal/service/mocks"
	"github.com/akrovv/filmlibrary/pkg/logger"
	"github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	us := mocks.NewMockUserService(ctrl)
	ss := mocks.NewMockSessionService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	userHandler := NewUserController(logger, us, ss)

	body := `{
		"username": "user",
		  "password": "user"
	  }`

	createUser := &domain.CRUser{
		Username: "user",
		Password: "user",
	}
	createSession := &domain.CreateSession{
		Username: "user",
	}

	req := httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	us.EXPECT().Register(createUser).Return(nil)
	ss.EXPECT().Create(createSession).Return("id", nil)

	userHandler.Register(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got: %d", w.Code)
		return
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	userHandler.Register(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
		return
	}

	// Incorrect JSON
	req = httptest.NewRequest("POST", "/register", strings.NewReader(`{"username": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	userHandler.Register(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("POST", "/register", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	userHandler.Register(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}

	// Register returned error
	req = httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	us.EXPECT().Register(createUser).Return(domain.ErrTest)
	userHandler.Register(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}

	// Create Session returned error
	req = httptest.NewRequest("POST", "/register", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	us.EXPECT().Register(createUser).Return(nil)
	ss.EXPECT().Create(createSession).Return("id", domain.ErrTest)

	userHandler.Register(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	us := mocks.NewMockUserService(ctrl)
	ss := mocks.NewMockSessionService(ctrl)

	logger, err := logger.NewLogger()
	if err != nil {
		t.Fatalf("can't create logger: %s", err)
	}

	userHandler := NewUserController(logger, us, ss)

	body := `{
		"username": "user",
		  "password": "user"
	  }`

	user := &domain.User{
		Username: "user",
	}

	createUser := &domain.CRUser{
		Username: "user",
		Password: "user",
	}
	createSession := &domain.CreateSession{
		Username: "user",
	}

	req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w := httptest.NewRecorder()

	// OK
	us.EXPECT().Login(createUser).Return(user, nil)
	ss.EXPECT().Create(createSession).Return("id", nil)

	userHandler.Login(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got: %d", w.Code)
		return
	}

	// Missed Content-Type
	req.Header.Del("Content-type")
	w = httptest.NewRecorder()
	userHandler.Login(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", w.Code)
		return
	}

	// Incorrect JSON
	req = httptest.NewRequest("POST", "/login", strings.NewReader(`{"username": {`))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	userHandler.Login(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}

	// io.ReadAll returned error
	req = httptest.NewRequest("POST", "/login", &BadReader{})
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	userHandler.Login(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}

	// Login returned error
	req = httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	w = httptest.NewRecorder()

	us.EXPECT().Login(createUser).Return(nil, domain.ErrTest)
	userHandler.Login(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}

	// Create Session returned error
	req = httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	us.EXPECT().Login(createUser).Return(user, nil)
	ss.EXPECT().Create(createSession).Return("id", domain.ErrTest)

	userHandler.Login(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got: %d", w.Code)
		return
	}
}
