package repositories

import (
	"inter/config"
	"inter/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoUserMock struct {
	mock.Mock
}

func (rp *RepoUserMock) CreateUser(data *models.User) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
func (rp *RepoUserMock) UpdateUser(data *models.User) (string, error) {
	args := rp.Mock.Called(data)
	return "1 data has been updated", args.Error(1)
}
func (rp *RepoUserMock) GetUser(data *models.User) (interface{}, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(interface{}), args.Error(1)
}
func (rp *RepoUserMock) GetAllUser(data *models.User) ([]models.User, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).([]models.User), args.Error(1)
}
func (rp *RepoUserMock) DeleteUser(data *models.User) (string, error) {
	args := rp.Mock.Called(data)
	return "1 data has been deleted", args.Error(1)
}
func (rp *RepoUserMock) GetAuthData(user string) (*models.User, error) {
	args := rp.Mock.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}
