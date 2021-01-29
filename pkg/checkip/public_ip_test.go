package checkip

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	getPublicIpV4HttpCallResponse      *http.Response
	getPublicIpV4HttpCallResponseError error
)

type getPublicIpV4HttpClient struct{}

func (client *getPublicIpV4HttpClient) Do(req *http.Request) (*http.Response, error) {
	return getPublicIpV4HttpCallResponse, getPublicIpV4HttpCallResponseError
}

type ClientObjSuite struct {
	suite.Suite
}

func TestClientObjSuite(t *testing.T) {
	suite.Run(t, new(ClientObjSuite))
}

func (suite *ClientObjSuite) BeforeTest(_, _ string) {
	var emptyResponse *http.Response
	getPublicIpV4HttpCallResponse = emptyResponse

	var emptyError error
	getPublicIpV4HttpCallResponseError = emptyError
}

func (suite *ClientObjSuite) TestGetPublicIpV4() {
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
				httpClient: &getPublicIpV4HttpClient{},
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
				httpClient: &getPublicIpV4HttpClient{},
			},
			httpResp: &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			want:    "",
			wantErr: true,
			errStr:  "Endpoint is not available. Could not detect public IPv4",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			getPublicIpV4HttpCallResponse = tt.httpResp
			getPublicIpV4HttpCallResponseError = tt.httpRespError

			result, err := tt.c.GetPublicIpV4()
			suite.Equal(result, tt.want)

			if tt.wantErr {
				fmt.Println(err)
				suite.Contains(err.Error(), tt.errStr)
			} else {
				suite.NoError(err)
			}
		})
	}
}
