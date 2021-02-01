package checkip

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	getPublicIPV4HttpCallResponse      *http.Response
	getPublicIPV4HttpCallResponseError error
)

type getPublicIPV4HttpClient struct{}

func (client *getPublicIPV4HttpClient) Do(req *http.Request) (*http.Response, error) {
	return getPublicIPV4HttpCallResponse, getPublicIPV4HttpCallResponseError
}

type ClientObjSuite struct {
	suite.Suite
}

func TestClientObjSuite(t *testing.T) {
	suite.Run(t, new(ClientObjSuite))
}

func (suite *ClientObjSuite) BeforeTest(_, _ string) {
	var emptyResponse *http.Response
	getPublicIPV4HttpCallResponse = emptyResponse

	var emptyError error
	getPublicIPV4HttpCallResponseError = emptyError
}

func (suite *ClientObjSuite) TestGetPublicIPV4() {
	tests := []struct {
		name          string
		c             *ClientObj
		httpResp      *http.Response
		httpRespError error
		want          string
		wantErr       bool
		errStr        string
	}{
		{
			name: "with default happy path",
			c: &ClientObj{
				httpClient: &getPublicIPV4HttpClient{},
			},
			httpResp: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader("172.0.0.1")),
			},
			want:    "172.0.0.1",
			wantErr: false,
		},
		{
			name: "with http not found status",
			c: &ClientObj{
				httpClient: &getPublicIPV4HttpClient{},
			},
			httpResp: &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			want:    "",
			wantErr: true,
			errStr:  "Endpoint is not available. Could not detect public IPv4",
		},
		{
			name: "with http request error",
			c: &ClientObj{
				httpClient: &getPublicIPV4HttpClient{},
			},
			httpResp:      &http.Response{},
			httpRespError: errors.New("Random HTTP error"),
			want:          "",
			wantErr:       true,
			errStr:        "Random HTTP error",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			getPublicIPV4HttpCallResponse = tt.httpResp
			getPublicIPV4HttpCallResponseError = tt.httpRespError

			result, err := tt.c.GetPublicIPV4()
			suite.Equal(result, tt.want)

			if tt.wantErr {
				suite.Error(err)
				suite.Contains(err.Error(), tt.errStr)
			} else {
				suite.NoError(err)
			}
		})
	}
}
