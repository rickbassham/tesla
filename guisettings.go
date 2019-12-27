package tesla

import (
	"fmt"
	"net/http"
)

// GUISettings represents the configured settings in the vehicle GUI.
type GUISettings struct {
	GUI24HourTime       bool   `json:"gui_24_hour_time"`
	GUIChargeRateUnits  string `json:"gui_charge_rate_units"`
	GUIDistanceUnits    string `json:"gui_distance_units"`
	GUIRangeDisplay     string `json:"gui_range_display"`
	GUITemperatureUnits string `json:"gui_temperature_units"`
	ShowRangeUnits      bool   `json:"show_range_units"`
	Timestamp           int64  `json:"timestamp"`
}

// GetGUISettings retrieves the current GUI settings for the vehicle.
func (c *Conn) GetGUISettings(id int) (*GUISettings, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response GUISettings `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/data_request/gui_settings", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
