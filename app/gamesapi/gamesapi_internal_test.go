package gamesapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamesApiServer_Healthcheck(t *testing.T) {
	s:= New(NewConfig())
	rec:= httptest.NewRecorder()
	req, _:= http.NewRequest(http.MethodGet, "/healthcheck", nil)
	s.healthcheckHandler().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "ok")

}
