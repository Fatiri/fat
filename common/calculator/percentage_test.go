package calculator_test

import (
	"testing"

	"github.com/FAT/common/calculator"
	"github.com/stretchr/testify/assert"
)

func TestCalCulatePercenatageFromString(t *testing.T) {
	type args struct {
		calculator calculator.PayloadPercentage
	}
	tests := []struct {
		name                string
		args                args
		funcUseCaseShouldBe func(t *testing.T, output, output2 float64, err error)
	}{
		// TODO: Add test cases.
		{
			name: "Success calculate percentage from string",
			args: args{
				calculator: calculator.PayloadPercentage{
					OldNumber: 10,
					NewNumber: 1,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output, output2 float64, err error) {
				assert.Equal(t, float64(900), output, "they sould equal")
				assert.Equal(t, float64(90), output2, "they sould equal")
				assert.NoError(t, err, "they shoul no error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, output2, err := tt.args.calculator.CalculatePercenatgeIncreaseDecrease()
			tt.funcUseCaseShouldBe(t, output, output2, err)
		})
	}
}

func TestCalculator_PercentageChange(t *testing.T) {
	type args struct {
		calculator calculator.PayloadPercentage
	}
	tests := []struct {
		name                string
		args                args
		funcUseCaseShouldBe func(t *testing.T, output float64)
	}{
		// TODO: Add test cases.
		{
			name: "Failed calculate percentage from string new Number",
			args: args{
				calculator: calculator.PayloadPercentage{
					OldNumber: 1,
					NewNumber: 2,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output float64) {
				assert.Equal(t, float64(100), output, "they sould equal")
			},
		},
		{
			name: "Failed calculate percentage from string new Number",
			args: args{
				calculator: calculator.PayloadPercentage{},
			},
			funcUseCaseShouldBe: func(t *testing.T, output float64) {
				assert.Equal(t, float64(0), output, "they sould equal")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.args.calculator.PercentageChange()
			tt.funcUseCaseShouldBe(t, output)
		})
	}
}
