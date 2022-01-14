package mock

import "newExp/internal/model"

type MockAuthorization struct {
}

func (m *MockAuthorization) CreateUser(user *model.User) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDIxNDIwMDQsImlhdCI6MTY0MjA5ODgwNCwidXNlcklkIjozfQ.Tb6xJnaTBJ03Kpffkdvt3LIjVIfqGGDLazUUfmTE3WY", nil
}

func (m *MockAuthorization) SignIn(username, password string) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDIxNDIwMDQsImlhdCI6MTY0MjA5ODgwNCwidXNlcklkIjozfQ.Tb6xJnaTBJ03Kpffkdvt3LIjVIfqGGDLazUUfmTE3WY", nil
}

func (m *MockAuthorization) ParseToke(token string) (uint64, error) {
	return 1, nil
}

func (m *MockAuthorization) GetUser(id uint64) (*model.User, error) {
	return &model.User{
		Username: "test",
		Password: "48b619efaadca769fe9134e07c83dda9740c23cb",
	}, nil
}
