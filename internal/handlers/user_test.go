package handlers

import (
	"inter/config"
	"inter/internal/models"
	"inter/internal/repositories"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoUserMock = repositories.RepoUserMock{}
var reqBody = `{
	"email": "bujang123@gmail.com",
	"pass": "12345678",
	"phone_number": "0821111",
	"role": "user"
}`

var getData struct {
	Id_user    string `db:"id_user" form:"id_user" json:"id_user" uri:"id_user" valid:"-"`
	Email_user string `db:"email" form:"email" json:"email" valid:"email, required~please input in type email"`
	Pass_user  string `db:"pass" form:"pass" json:"pass" valid:"alphanum, stringlength(8|15)~password harus lebih sama dengan dari 8"`
	Role       string `db:"role" json:"role" form:"role" valid:"-"`
	Phone_user string `db:"phone_number" form:"phone_number" json:"phone_number" valid:"numeric"`
}

type TestUser struct {
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user" uri:"id_user" valid:"-"`
	Email_user string     `db:"email" form:"email" json:"email" valid:"email, required~please input in type email"`
	Pass_user  string     `db:"pass" form:"pass" json:"pass" valid:"alphanum, stringlength(8|15)~password harus lebih sama dengan dari 8"`
	Role       string     `db:"role" json:"role" form:"role" valid:"-"`
	Phone_user string     `db:"phone_number" form:"phone_number" json:"phone_number" valid:"numeric"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

func TestPostData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectResult := &config.Result{
		Data: "User created",
	}
	repoUserMock.On("CreateUser", mock.Anything).Return(expectResult, nil)

	r.POST("/create", handler.PostData)
	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"code":200, "description":"User created" ,"status":"OK"}`, w.Body.String())

}

func TestUpdateData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectResult := "1 data has been updated"
	repoUserMock.On("UpdateUser", mock.Anything).Return(expectResult, nil)

	r.PATCH("/update/:id", handler.UpdateData)
	req := httptest.NewRequest("PATCH", "/update/123", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"description":"OK", "message":"1 data has been updated", "status":200}`, w.Body.String())

}

func TestDeleteData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectResult := "1 data has been deleted"
	repoUserMock.On("DeleteUser", mock.Anything).Return(expectResult, nil)

	r.DELETE("/delete/:id", handler.DeleteData)
	req := httptest.NewRequest("DELETE", "/delete/123", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"description":"OK", "message":"1 data has been deleted", "status":200}`, w.Body.String())
}

func TestGetData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	getData.Id_user = "123"
	getData.Email_user = "1cak@gmail.com"
	getData.Pass_user = "12345679"
	getData.Role = "081111111"
	getData.Role = "user"

	expectResult := getData
	repoUserMock.On("GetUser", mock.Anything).Return(expectResult, nil)

	r.GET("/id", handler.GetDataUser)
	req := httptest.NewRequest("GET", "/id", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"email":"1cak@gmail.com", "id_user":"123", "pass":"12345679", "phone_number":"", "role":"user"}`, w.Body.String())
}

func TestGetAllUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)

	expectResult := []models.User{{
		Id_user:    "94f1cbe9-e73a-4a45-854f-95075ca927a7",
		Email_user: "user50@gmail.com",
		Pass_user:  "$2a$10$j3lnmQYvDhrp5rW8PRj44emyUb2A87WOfHnQVRJ.RN8nVIjItI5Zq",
		Phone_user: "0821897869071",
		Role:       "admin",
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}, {
		Id_user:    "4e4b277c-eb5d-481e-991a-ee0b090754c8",
		Email_user: "user40@gmail.com",
		Pass_user:  "$2a$10$o4P4Yt.tebth8qr4u8B/LOSFYevQ/PSXvOlV5m0AwtZNBILm7wwI2",
		Phone_user: "0821897869071",
		Role:       "user",
		CreatedAt:  nil,
		UpdatedAt:  nil,
	},
	}

	repoUserMock.On("GetAllUser", mock.Anything).Return(expectResult, nil)

	r.GET("/all", handler.GetAllData)
	req := httptest.NewRequest("GET", "/all", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.JSONEq(t, `[
		{
		"id_user": "94f1cbe9-e73a-4a45-854f-95075ca927a7",
		"email": "user50@gmail.com",
		"pass": "$2a$10$j3lnmQYvDhrp5rW8PRj44emyUb2A87WOfHnQVRJ.RN8nVIjItI5Zq",
		"role": "admin",
		"phone_number": "0821897869071",
		"created_at": null,
		"updated_at": null},    
	{
        "id_user": "4e4b277c-eb5d-481e-991a-ee0b090754c8",
        "email": "user40@gmail.com",
        "pass": "$2a$10$o4P4Yt.tebth8qr4u8B/LOSFYevQ/PSXvOlV5m0AwtZNBILm7wwI2",
        "role": "user",
        "phone_number": "0821897869071",
        "created_at": null,
        "updated_at": null
    }]`, w.Body.String())
}
