package client_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/Fatiri/fat/common/client"
	mocks "github.com/Fatiri/fat/mocks/common/client"
	"github.com/stretchr/testify/assert"
)

func init() {
	client.Client = &mocks.MockClient{}
}

func TestHttpClient(t *testing.T) {

	type expected struct {
		statusCoce int
		reader     io.ReadCloser
		err        error
		parameter  *client.ParamaterHttpClient
	}

	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	// create a new reader with that JSON
	helperTest := func(ex expected) {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: ex.statusCoce,
				Body:       ex.reader,
			}, ex.err
		}
	}

	tests := []struct {
		name                string
		expectedResult      expected
		funcUseCaseShouldBe func(t *testing.T, output *http.Response, err error)
	}{
		// TODO: Add test cases.
		{
			name: "Success with status code 200 have data",
			expectedResult: expected{
				statusCoce: 200,
				reader:     ioutil.NopCloser(bytes.NewReader([]byte(json))),
				err:        nil,
				parameter: &client.ParamaterHttpClient{
					Headers: []client.RequestDefault{
						{
							Key:   "Auth",
							Value: "token",
						},
					},
					WithAuthorization: true,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output *http.Response, err error) {
				assert.NoError(t, err, "they should be no error")
				assert.NotNil(t, output, "they should be not nil")
				assert.Equal(t, 200, output.StatusCode, "they should be equal")
			},
		},
		{
			name: "Failed with status code 400",
			expectedResult: expected{
				statusCoce: 400,
				reader:     nil,
				err:        errors.New("error"),
				parameter: &client.ParamaterHttpClient{
					Headers: []client.RequestDefault{
						{
							Key:   "Auth",
							Value: "token",
						},
					},
					WithAuthorization: true,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output *http.Response, err error) {
				assert.Error(t, err, "they should be no error")
				assert.Nil(t, output, "they should be nil")
			},
		},
		{
			name: "Failed with status code 400 empty error",
			expectedResult: expected{
				statusCoce: 400,
				reader:     nil,
				err:        errors.New("Bad Request"),
				parameter: &client.ParamaterHttpClient{
					Headers: []client.RequestDefault{
						{
							Key:   "Auth",
							Value: "token",
						},
					},
					WithAuthorization: true,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output *http.Response, err error) {
				assert.Error(t, err, "they should be error")
				assert.Nil(t, output, "they should be nil")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperTest(tt.expectedResult)
			output, err := client.HttpClient(tt.expectedResult.parameter)
			tt.funcUseCaseShouldBe(t, output, err)
		})
	}
}

func TestHttpClientV2(t *testing.T) {
	type expected struct {
		statusCoce int
		req        string
		reader     io.ReadCloser
		err        error
		parameter  *client.ParamaterHttpClient
	}

	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	// create a new reader with that JSON
	helperTest := func(ex expected) {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: ex.statusCoce,
				Body:       ex.reader,
			}, ex.err
		}
	}

	tests := []struct {
		name                string
		expectedResult      expected
		funcUseCaseShouldBe func(t *testing.T, output *http.Response, err error)
	}{
		// TODO: Add test cases.
		{
			name: "Success with status code 200 have data",
			expectedResult: expected{
				statusCoce: 200,
				req:        "test",
				reader:     ioutil.NopCloser(bytes.NewReader([]byte(json))),
				err:        nil,
				parameter: &client.ParamaterHttpClient{
					Headers: []client.RequestDefault{
						{
							Key:   "Auth",
							Value: "token",
						},
					},
					WithAuthorization: true,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output *http.Response, err error) {
				assert.NoError(t, err, "they should be no error")
				assert.NotNil(t, output, "they should be not nil")
				assert.Equal(t, 200, output.StatusCode, "they should be equal")
			},
		},
		{
			name: "Failed with status code 400",
			expectedResult: expected{
				statusCoce: 400,
				reader:     nil,
				err:        errors.New("error"),
				parameter: &client.ParamaterHttpClient{
					Headers: []client.RequestDefault{
						{
							Key:   "Auth",
							Value: "token",
						},
					},
					WithAuthorization: true,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output *http.Response, err error) {
				assert.Error(t, err, "they should be error")
				assert.Nil(t, output, "they should be nil")
			},
		},
		{
			name: "Failed with status code 400 empty error",
			expectedResult: expected{
				statusCoce: 400,
				reader:     nil,
				err:        errors.New("Bad Request"),
				parameter: &client.ParamaterHttpClient{
					Headers: []client.RequestDefault{
						{
							Key:   "Auth",
							Value: "token",
						},
					},
					WithAuthorization: true,
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output *http.Response, err error) {
				assert.Error(t, err, "they should be error")
				assert.Nil(t, output, "they should be nil")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helperTest(tt.expectedResult)
			output, err := client.HttpClientV2(tt.expectedResult.parameter, tt.expectedResult.req)
			tt.funcUseCaseShouldBe(t, output, err)
		})
	}
}

func TestReadHttpResponse(t *testing.T) {
	type expected struct {
		response *http.Response
	}

	tests := []struct {
		name                string
		expected            expected
		result              []byte
		funcUseCaseShouldBe func(t *testing.T, output []byte, result []byte, err error)
	}{
		{
			name: "Success read response body",
			expected: expected{
				response: &http.Response{
					Body: ioutil.NopCloser(bytes.NewReader([]byte("test"))),
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output []byte, result []byte, err error) {
				assert.NoError(t, err, "they should be no error")
				assert.NotNil(t, output, "they should be not nil")
				assert.Equal(t, output, result, "they should be equal")
			},
			result: []byte("test"),
		},
		{
			name: "Success read response body empty",
			expected: expected{
				response: &http.Response{
					Body: ioutil.NopCloser(bytes.NewReader([]byte(""))),
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output []byte, result []byte, err error) {
				assert.NoError(t, err, "they should be no error")
				assert.NotNil(t, output, "they should be not nil")
				assert.Equal(t, output, result, "they should be equal")
			},
			result: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.ReadHttpResponse(tt.expected.response)
			tt.funcUseCaseShouldBe(t, resp, tt.result, err)
		})
	}
}

func TestRequestClientHeaderIndodax(t *testing.T) {
	type args struct {
		sign      string
		urlEncode string
	}
	tests := []struct {
		name                string
		args                args
		want                []client.RequestDefault
		funcUseCaseShouldBe func(t *testing.T, output []client.RequestDefault, want []client.RequestDefault)
	}{
		// TODO: Add test cases.
		{
			name: "Success Assign Array Header",
			args: args{
				sign:      "test",
				urlEncode: "test",
			},
			want: []client.RequestDefault{
				{
					Key:   "Key",
					Value: "1234",
				},
				{
					Key:   "Sign",
					Value: "test",
				},
				{
					Key:   "Content-type",
					Value: "application/x-www-form-urlencoded",
				},
				{
					Key:   "Content-Length",
					Value: strconv.Itoa(len("test")),
				},
			},
			funcUseCaseShouldBe: func(t *testing.T, output, want []client.RequestDefault) {
				assert.Equal(t, output, want, "they should be equal")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := client.RequestClientHeaderIndodax(tt.args.sign, tt.args.urlEncode, "1234")
			tt.funcUseCaseShouldBe(t, output, tt.want)
		})
	}
}
