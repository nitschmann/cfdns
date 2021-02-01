package publicip

import (
	"errors"
	"testing"

	"github.com/nitschmann/cfdns/pkg/util/httpclient"
	"github.com/stretchr/testify/suite"
)

var (
	getPublicIPV4UtilResponse string
	getPublicIPV4UtilError    error
)

type UtilMock struct{}

func (m *UtilMock) GetPublicIPV4() (string, error) {
	return getPublicIPV4UtilResponse, getPublicIPV4UtilError
}

type ServiceObjSuite struct {
	suite.Suite
}

func TestServiceObjSuite(t *testing.T) {
	suite.Run(t, new(ServiceObjSuite))
}

func (suite *ServiceObjSuite) BeforeTest(_, _ string) {
	var emptyString string
	getPublicIPV4UtilResponse = emptyString

	var emptyError error
	getPublicIPV4UtilError = emptyError
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

func (suite *ServiceObjSuite) TestFetchPublicIPV4() {
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
				checkIPClient: &UtilMock{},
			},
			want:     "192.168.1.2",
			wantErr:  false,
			utilResp: "192.168.1.2",
		},
		{
			name: "with util returning an error",
			serv: &ServiceObj{
				checkIPClient: &UtilMock{},
			},
			want:     "",
			wantErr:  true,
			utilResp: "",
			utilErr:  errors.New("could not fetch public IPv4"),
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			getPublicIPV4UtilResponse = tt.utilResp
			getPublicIPV4UtilError = tt.utilErr

			result, err := tt.serv.FetchPublicIPV4()
			suite.Equal(result, tt.want)

			if tt.wantErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}
