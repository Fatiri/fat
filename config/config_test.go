package config_test

import (
	"errors"
	"testing"

	"github.com/Fatiri/fat/config"
	mocksConfig "github.com/Fatiri/fat/mocks/config"
	"github.com/Fatiri/fat/models"
	"github.com/Fatiri/fat/repository"
	"github.com/stretchr/testify/assert"
)

type mock struct {
	env     *mocksConfig.Environment
	storage *mocksConfig.Storage
}

func provideMockConfig() *mock {
	return &mock{
		&mocksConfig.Environment{},
		&mocksConfig.Storage{},
	}
}

func provideUsecaseMockConfig(m *mock) config.Config {
	return config.NewConfig(m.env, m.storage)
}

func TestConfigCtx_InitConfig(t *testing.T) {
	type expectedData struct {
		env        *models.Environment
		repo       *repository.Queries
		errEnv     error
		errStorage error
	}

	helperTest := func(ex expectedData) *mock {
		mc := provideMockConfig()
		mc.env.On("InitEnvironment").Return(ex.env, ex.errEnv).Once()
		mc.storage.On("Postgres").Return(ex.repo, ex.errStorage).Once()
		return mc
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcUseCaseShouldBe func(t *testing.T, output *models.Config, err error)
	}{
		// TODO: Add test cases.
		{
			name: "Should get error environment",
			expected: expectedData{
				env:        &models.Environment{},
				repo:       &repository.Queries{},
				errEnv:     errors.New("env file not found"),
				errStorage: nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, output *models.Config, err error) {
				assert.Nil(t, output)
				assert.Error(t, err)
			},
		},
		{
			name: "Should get error storage",
			expected: expectedData{
				env:        &models.Environment{},
				repo:       &repository.Queries{},
				errEnv:     nil,
				errStorage: errors.New("failed setup config storage"),
			},
			funcUseCaseShouldBe: func(t *testing.T, output *models.Config, err error) {
				assert.Nil(t, output)
				assert.Error(t, err)
			},
		},
		{
			name: "Should return config data",
			expected: expectedData{
				env:        &models.Environment{},
				repo:       &repository.Queries{},
				errEnv:     nil,
				errStorage: nil,
			},
			funcUseCaseShouldBe: func(t *testing.T, output *models.Config, err error) {
				assert.NotNil(t, output)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		mc := helperTest(tt.expected)
		output, err := provideUsecaseMockConfig(mc).InitConfig()
		tt.funcUseCaseShouldBe(t, output, err)
	}
}
