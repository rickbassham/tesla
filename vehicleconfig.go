package tesla

import (
	"fmt"
	"net/http"
)

// VehicleConfig represents the current capabilities of the vehicle.
type VehicleConfig struct {
	CanAcceptNavigationRequests bool   `json:"can_accept_navigation_requests"`
	CanActuateTrunks            bool   `json:"can_actuate_trunks"`
	CarSpecialType              string `json:"car_special_type"`
	CarType                     string `json:"car_type"`
	ChargePortType              string `json:"charge_port_type"`
	EuVehicle                   bool   `json:"eu_vehicle"`
	ExteriorColor               string `json:"exterior_color"`
	HasAirSuspension            bool   `json:"has_air_suspension"`
	HasLudicrousMode            bool   `json:"has_ludicrous_mode"`
	KeyVersion                  int    `json:"key_version"`
	MotorizedChargePort         bool   `json:"motorized_charge_port"`
	PerfConfig                  string `json:"perf_config"`
	Plg                         bool   `json:"plg"`
	RearSeatHeaters             int    `json:"rear_seat_heaters"`
	RearSeatType                int    `json:"rear_seat_type"`
	Rhd                         bool   `json:"rhd"`
	RoofColor                   string `json:"roof_color"`
	SeatType                    int    `json:"seat_type"`
	SpoilerType                 string `json:"spoiler_type"`
	SunRoofInstalled            int    `json:"sun_roof_installed"`
	ThirdRowSeats               string `json:"third_row_seats"`
	Timestamp                   int64  `json:"timestamp"`
	TrimBadging                 string `json:"trim_badging"`
	WheelType                   string `json:"wheel_type"`
}

// GetVehicleConfig retrieves the vehicles config.
func (c *Conn) GetVehicleConfig(id int) (*VehicleConfig, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response VehicleConfig `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/data_request/vehicle_config", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
