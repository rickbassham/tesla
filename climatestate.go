package tesla

import (
	"fmt"
	"net/http"
)

// ClimateState represents the current state of climate control for the vehicle.
type ClimateState struct {
	BatteryHeater              bool    `json:"battery_heater"`
	BatteryHeaterNoPower       bool    `json:"battery_heater_no_power"`
	ClimateKeeperMode          string  `json:"climate_keeper_mode"`
	DefrostMode                int     `json:"defrost_mode"`
	DriverTempSetting          float64 `json:"driver_temp_setting"`
	FanStatus                  int     `json:"fan_status"`
	InsideTemp                 float64 `json:"inside_temp"`
	IsAutoConditioningOn       bool    `json:"is_auto_conditioning_on"`
	IsClimateOn                bool    `json:"is_climate_on"`
	IsFrontDefrosterOn         bool    `json:"is_front_defroster_on"`
	IsPreconditioning          bool    `json:"is_preconditioning"`
	IsRearDefrosterOn          bool    `json:"is_rear_defroster_on"`
	LeftTempDirection          int     `json:"left_temp_direction"`
	MaxAvailTemp               float64 `json:"max_avail_temp"`
	MinAvailTemp               float64 `json:"min_avail_temp"`
	OutsideTemp                float64 `json:"outside_temp"`
	PassengerTempSetting       float64 `json:"passenger_temp_setting"`
	RemoteHeaterControlEnabled bool    `json:"remote_heater_control_enabled"`
	RightTempDirection         int     `json:"right_temp_direction"`
	SeatHeaterLeft             int     `json:"seat_heater_left"`
	SeatHeaterRearCenter       int     `json:"seat_heater_rear_center"`
	SeatHeaterRearLeft         int     `json:"seat_heater_rear_left"`
	SeatHeaterRearLeftBack     int     `json:"seat_heater_rear_left_back"`
	SeatHeaterRearRight        int     `json:"seat_heater_rear_right"`
	SeatHeaterRearRightBack    int     `json:"seat_heater_rear_right_back"`
	SeatHeaterRight            int     `json:"seat_heater_right"`
	SideMirrorHeaters          bool    `json:"side_mirror_heaters"`
	SteeringWheelHeater        bool    `json:"steering_wheel_heater"`
	Timestamp                  int64   `json:"timestamp"`
	WiperBladeHeater           bool    `json:"wiper_blade_heater"`
}

// GetClimateState retrieves information on the current internal temperature and climate control
// system.
func (c *Conn) GetClimateState(id int) (*ClimateState, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response ClimateState `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/data_request/climate_state", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
