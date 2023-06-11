package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/kerimkhanov/sites-duration/internal/delivery"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthReady(t *testing.T) {
	var h delivery.SiteHandler
	gin.SetMode(gin.TestMode)
	router := h.InitRoutes()
	w := httptest.NewRecorder()

	wanted := `{"data":{"data":"Server is up and running"}}`
	wanted2 := `{"data":{"data":"Server is up and running"}}`

	req2, _ := http.NewRequest(http.MethodGet, "/health/live", nil)
	router.ServeHTTP(w, req2)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, wanted, w.Body.String())
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/health/ready", nil)
	router.ServeHTTP(w2, req)

	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, wanted2, w2.Body.String())
}
