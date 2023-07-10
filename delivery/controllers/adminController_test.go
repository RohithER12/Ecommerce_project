package controllers

import (
	"70_Off/services/mockServices/mockAdminServices"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func setupAdminControllerTest(t *testing.T) (*gin.Engine, *gomock.Controller, *mockAdminServices.MockAdminService, *AdminController) {
	ctrl := gomock.NewController(t)
	mockAdminService := mockAdminServices.NewMockAdminService(ctrl)
	adminController := NewAdminController(mockAdminService)
	router := gin.Default()
	return router, ctrl, mockAdminService, adminController
}

func TestAdminController_AdminSignUp(t *testing.T) {
	router, ctrl, mockAdminService, adminController := setupAdminControllerTest(t)
	defer ctrl.Finish()

	router.POST("/adminSignup", adminController.AdminSignUp)

	requestBody := `{
		"name": "John Doe",
		"email": "john@example.com",
		"phonenumber": "1234567890",
		"password": "password123"
	}`

	req, err := http.NewRequest(http.MethodPost, "/adminSignup", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	mockAdminService.EXPECT().Create(gomock.Any()).Return(nil)

	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, res.Code)
	}

	expectedResponseBody := `{"name":"John Doe","email":"john@example.com","phonenumber":"1234567890","password":"password123"}`
	if res.Body.String() != expectedResponseBody {
		t.Errorf("expected response body %s, got %s", expectedResponseBody, res.Body.String())
	}
}

// func TestAdminController_AdminLogin(t *testing.T) {
// 	router, ctrl, mockAdminService, adminController := setupAdminControllerTest(t)
// 	defer ctrl.Finish()

// 	router.POST("/adminLogin", adminController.AdminLogin)

// 	requestBody := `{
// 		"email": "john@example.com",
// 		"password": "password123"
// 	}`

// 	req, err := http.NewRequest(http.MethodPost, "/adminLogin", strings.NewReader(requestBody))
// 	if err != nil {
// 		t.Fatalf("failed to create HTTP request: %v", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	res := httptest.NewRecorder()

// 	mockAdminService.EXPECT().AdminValidateLogin(gomock.Any()).Return(&entity.Admin{Email: "john@example.com"}, nil)

// 	router.ServeHTTP(res, req)

// 	if res.Code != http.StatusOK {
// 		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
// 	}

// 	expectedResponseBody := `{"message":"Login successfully"}`
// 	if res.Body.String() != expectedResponseBody {
// 		t.Errorf("expected response body %s, got %s", expectedResponseBody, res.Body.String())
// 	}

// 	// Verify the token cookie is set
// 	tokenCookie := res.Result().Cookies()[0]
// 	expectedTokenCookieName := "token"
// 	expectedTokenCookieValue := "mocked_token"
// 	if tokenCookie.Name != expectedTokenCookieName || tokenCookie.Value != expectedTokenCookieValue {
// 		t.Errorf("expected token cookie %s=%s, got %s=%s", expectedTokenCookieName, expectedTokenCookieValue, tokenCookie.Name, tokenCookie.Value)
// 	}
// }
