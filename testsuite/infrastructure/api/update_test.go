package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"github.com/markitos/markitos-svc-boilerplate/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanUpdateABoilerplate(t *testing.T) {
	var boiler *domain.Boilerplate = createPersistedRandomBoilerplate()

	name := boiler.Name + " UPDATED"
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateUpdateRequest{
		Name: name,
	})
	request, _ := http.NewRequest(http.MethodPatch, "/v1/boilerplates/"+boiler.Id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, http.StatusOK, recorder.Code)

	deletePersisteRandomBoilerplate(boiler.Id)
}
