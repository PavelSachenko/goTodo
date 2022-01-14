package v1

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"newExp/internal/model"
	"newExp/internal/usecase"
	"newExp/internal/usecase/user/mock"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	testCases := []struct {
		name                string
		inputBody           string
		inputUser           *model.User
		mockAuth            *mock.MockAuthorization
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"test", "password":"qwerty"}`,
			inputUser: &model.User{
				Username: "Test",
				Password: "qwerty",
			},
			mockAuth:            &mock.MockAuthorization{},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: `{"bearerToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDIxNDIwMDQsImlhdCI6MTY0MjA5ODgwNCwidXNlcklkIjozfQ.Tb6xJnaTBJ03Kpffkdvt3LIjVIfqGGDLazUUfmTE3WY"}`,
		},
		{
			name:                "Empty password",
			inputBody:           `{"password":"qwerty"}`,
			mockAuth:            &mock.MockAuthorization{},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:                "Empty username",
			inputBody:           `{"username":"username"}`,
			mockAuth:            &mock.MockAuthorization{},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:                "Empty body",
			inputBody:           ``,
			mockAuth:            &mock.MockAuthorization{},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			services := &usecase.SuperService{Auth: testCase.mockAuth}
			handler := NewHandler(services)

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/api/v1/user/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/user/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	testCases := []struct {
		name                string
		inputBody           string
		mockAuth            *mock.MockAuthorization
		inputUser           *model.User
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"pavel", "password":"qwerty"}`,
			inputUser: &model.User{
				Username: "Test",
				Password: "qwerty",
			},
			mockAuth:            &mock.MockAuthorization{},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"bearerToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDIxNDIwMDQsImlhdCI6MTY0MjA5ODgwNCwidXNlcklkIjozfQ.Tb6xJnaTBJ03Kpffkdvt3LIjVIfqGGDLazUUfmTE3WY"}`,
		},
		{
			name:      "OK",
			inputBody: `{"password":"qwerty"}`,
			inputUser: &model.User{
				Username: "Test",
				Password: "qwerty",
			},
			mockAuth:            &mock.MockAuthorization{},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			services := &usecase.SuperService{Auth: testCase.mockAuth}
			handler := NewHandler(services)

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/api/v1/user/sign-in", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/user/sign-in", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
