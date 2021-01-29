package checkip

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

// IPv4CheckURL is the endpoint to detect the public IPv4
const IPv4CheckURL string = "https://checkip.amazonaws.com"

// GetPublicIpV4 fetches the current public IPv4 of this machine and network from an endpoint
func (c *ClientObj) GetPublicIpV4() (string, error) {
	var ip string

	req, err := http.NewRequest("GET", IPv4CheckURL, nil)
	if err != nil {
		return ip, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ip, err
	}

	if resp.StatusCode != 200 {
		return ip, errors.New("Endpoint is not available. Could not detect public IPv4")
	}

	body, err := ioutil.ReadAll(resp.Body)
	re := regexp.MustCompile("\n")
	ip = re.ReplaceAllString(string(body), "")

	return ip, nil
}
