package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"library/internal/application"
	"library/internal/domain/service"
	"library/internal/infrastructure/app"
	"library/internal/infrastructure/controller"
	mockAuthorRepo "library/internal/test/mock"

	"github.com/gin-gonic/gin"
)

func TestNewHandlers_SetsAuthor(t *testing.T) {
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	useCase := application.NewAuthorApp(svc)
	ctrl := controller.NewAuthorController(useCase)

	h := app.NewHandlers(ctrl)

	if h.Author != ctrl {
		t.Fatalf("se esperaba que Handlers.Author fuera el pasado a NewHandlers")
	}
}

func TestRegisterapp_NoAuthor_POSTAuthors_Returns404(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	useCase := application.NewAuthorApp(svc)
	ctrl := controller.NewAuthorController(useCase)
	h := app.NewHandlers(ctrl)
	app.RegisterRoutes(r, h)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/library/authors", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusBadRequest; got != want {
		t.Fatalf("POST /api/v1/library/authors => status %d, se esperaba %d", got, want)
	}
}

func TestRegisterapp_NoAuthor_GETAuthors_Returns404(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := app.NewHandlers(nil)
	app.RegisterRoutes(r, h)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/library/authors", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusNotFound; got != want {
		t.Fatalf("GET /api/v1/library/authors => status %d, se esperaba %d", got, want)
	}
}

func TestRegisterapp_NoAuthor_GETAuthorByID_Returns404(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := app.NewHandlers(nil)
	app.RegisterRoutes(r, h)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/library/authors/123", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusNotFound; got != want {
		t.Fatalf("GET /api/v1/library/authors/123 => status %d, se esperaba %d", got, want)
	}
}

func TestRegisterapp_NoAuthor_PUTAuthorByID_Returns404(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := app.NewHandlers(nil)
	app.RegisterRoutes(r, h)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/library/authors/123", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusNotFound; got != want {
		t.Fatalf("PUT /api/v1/library/authors/123 => status %d, se esperaba %d", got, want)
	}
}
