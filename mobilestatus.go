package tesla

import (
	"fmt"
	"net/http"
)

// GetMobileEnabled returns whether or not the Mobile Access setting is enabled in the vehicle.
func (c *Conn) GetMobileEnabled(id int) (bool, error) {
	if c.accessToken == "" {
		return false, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response bool `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/mobile_enabled", id), nil, &respBody)
	if err != nil {
		return false, err
	}

	return respBody.Response, nil
}
