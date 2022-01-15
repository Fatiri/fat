package config_test

import (
	"testing"

	"github.com/FAT/config"
	"github.com/FAT/models"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentCtx_InitEnvironment(t *testing.T) {
	type expectedData struct {
		envName, dirFile string
		
	}

	tests := []struct {
		name                string
		expected            expectedData
		funcUseCaseShouldBe func(t *testing.T, output *models.Environment, err error)
	}{
		// TODO: Add test cases.
		{
			name: "Shold success get env file and return data environment",
			expected: expectedData{
				envName: "dev",
				dirFile: "./test",
			},
			funcUseCaseShouldBe: func(t *testing.T, output *models.Environment, err error) {
				assert.Nil(t, output)
				assert.Error(t, err)
			},
		},
		{
			name: "Shold success get env file and return data environment",
			expected: expectedData{
				envName: "dev",
				dirFile: ".",
			},
			funcUseCaseShouldBe: func(t *testing.T, output *models.Environment, err error) {
				assert.NotNil(t, output)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := config.NewEnvironment(tt.expected.envName, tt.expected.dirFile).InitEnvironment()
			tt.funcUseCaseShouldBe(t, output, err)
		})
	}
}
