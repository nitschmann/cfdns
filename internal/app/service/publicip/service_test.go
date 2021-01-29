package publicip

import (
	"errors"
	"testing"

	"github.com/nitschmann/cfd/pkg/util/httpclient"
	"github.com/stretchr/testify/suite"
)

var (
	getPublicIpV4UtilResponse string
	getPublicIpV4UtilError    error
)

type UtilMock struct{}

func (m *UtilMock) GetPublicIpV4() (string, error) {
	return getPublicIpV4UtilResponse, getPublicIpV4UtilError
}

type ServiceObjSuite struct {
	suite.Suite
}

func TestServiceObjSuite(t *testing.T) {
	suite.Run(t, new(ServiceObjSuite))
}

func (suite *ServiceObjSuite) BeforeTest(_, _ string) {
	var emptyString string
	getPublicIpV4UtilResponse = emptyString

	var emptyError error
	getPublicIpV4UtilError = emptyError
}

func (suite *ServiceObjSuite) TestNew() {
	type args struct {
		httpClient httpclient.Client
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "default",
			args: args{httpClient: httpclient.New()},
			want: New(httpclient.New()),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.want, New(tt.args.httpClient))
		})
	}
}

func (suite *ServiceObjSuite) TestFetchPublicIpV4() {
	tests := []struct {
		name     string
		serv     Service
		want     string
		wantErr  bool
		utilResp string
		utilErr  error
	}{
		{
			name: "with default happy path",
			serv: &ServiceObj{
				checkIpClient: &UtilMock{},
			},
			want:     "192.168.1.2",
			wantErr:  false,
			utilResp: "192.168.1.2",
		},
		{
			name: "with util returning an error",
			serv: &ServiceObj{
				checkIpClient: &UtilMock{},
			},
			want:     "",
			wantErr:  true,
			utilResp: "",
			utilErr:  errors.New("could not fetch public IPv4"),
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			getPublicIpV4UtilResponse = tt.utilResp
			getPublicIpV4UtilError = tt.utilErr

			result, err := tt.serv.FetchPublicIpV4()
			suite.Equal(result, tt.want)

			if tt.wantErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}
