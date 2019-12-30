package web

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"github.com/jayvib/golog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gophr/v2/gophr.api/user"
	"gophr/v2/gophr.api/user/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var debugTest = flag.Bool("debug", false, "debugging")

type responseTest struct {
	Message string     `json:"message"`
	Data    *user.User `json:"data"`
	Method  string     `json:"method"`
}

func TestMain(m *testing.M) {
	flag.Parse()
	if *debugTest {
		golog.SetLevel(golog.DebugLevel)
	}
	os.Exit(m.Run())
}

func TestHandle_GetByID(t *testing.T) {

	t.Run("OK", func(t *testing.T) {

		service := new(mocks.Service)

		mockReturn := &user.User{
			ID:       "mock1",
			Username: "luffy.monkey",
			Email:    "luffy.monkey@gmail.com",
			Password: "1234567890",
		}

		want := responseTest{
			Message: "OK",
			Data:    mockReturn,
			Method:  http.MethodGet,
		}

		service.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(mockReturn, nil).Once()

		param := &Parameters{
			UserService: service,
		}

		h := New(param)

		router := mux.NewRouter()
		router.HandleFunc("/users/{id}", h.GetByID).Methods(http.MethodGet)

		w := performRequest(router, http.MethodGet, "/users/mock1", nil)
		assertGetByIDResponse(t, want, w)
		service.AssertExpectations(t)
	})
}

func assertGetByIDResponse(t *testing.T, want responseTest, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusOK, w.Code)

	var got responseTest

	err := json.NewDecoder(w.Body).Decode(&got)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}

func performRequest(h http.Handler, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/users/mock1", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
}