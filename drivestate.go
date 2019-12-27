package tesla

import (
	"fmt"
	"net/http"
)

// DriveState is the current state of driving for the vehicle.
type DriveState struct {
	GpsAsOf                 int         `json:"gps_as_of"`
	Heading                 int         `json:"heading"`
	Latitude                float64     `json:"latitude"`
	Longitude               float64     `json:"longitude"`
	NativeLatitude          float64     `json:"native_latitude"`
	NativeLocationSupported int         `json:"native_location_supported"`
	NativeLongitude         float64     `json:"native_longitude"`
	NativeType              string      `json:"native_type"`
	Power                   int         `json:"power"`
	ShiftState              interface{} `json:"shift_state"`
	Speed                   interface{} `json:"speed"`
	Timestamp               int64       `json:"timestamp"`
}

// GetDriveState retrieves the driving and position state of the vehicle.
func (c *Conn) GetDriveState(id int) (*DriveState, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response DriveState `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/data_request/drive_state", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
