package tesla

import (
	"fmt"
	"net/http"
)

// Vehicle represents the basic data about the vehicle.
type Vehicle struct {
	ID                     int      `json:"id"`
	VehicleID              int      `json:"vehicle_id"`
	VIN                    string   `json:"vin"`
	DisplayName            string   `json:"display_name"`
	OptionCodes            string   `json:"option_codes"`
	Color                  *string  `json:"color"`
	Tokens                 []string `json:"tokens"`
	State                  string   `json:"state"`
	InService              bool     `json:"in_service"`
	CalendarEnabled        bool     `json:"calendar_enabled"`
	APIVersion             int      `json:"api_version"`
	BackseatToken          *string  `json:"backseat_token"`
	BackseatTokenUpdatedAt *int     `json:"backseat_token_updated_at"`
}

// GetVehicles retrieves a list of vehicles for the currently authenticated account.
func (c *Conn) GetVehicles() ([]Vehicle, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response []Vehicle `json:"response"`
		Count    int       `json:"count"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, "/api/1/vehicles", nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Response, nil
}

// GetVehicle returns a vehicle by id.
func (c *Conn) GetVehicle(id int) (*Vehicle, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response Vehicle `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
