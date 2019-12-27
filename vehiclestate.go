package tesla

import (
	"fmt"
	"net/http"
)

// VehicleState represents the current vehicle state.
type VehicleState struct {
	APIVersion          int    `json:"api_version"`
	AutoparkStateV3     string `json:"autopark_state_v3"`
	AutoparkStyle       string `json:"autopark_style"`
	CalendarSupported   bool   `json:"calendar_supported"`
	CarVersion          string `json:"car_version"`
	CenterDisplayState  int    `json:"center_display_state"`
	Df                  int    `json:"df"`
	Dr                  int    `json:"dr"`
	FdWindow            int    `json:"fd_window"`
	FpWindow            int    `json:"fp_window"`
	Ft                  int    `json:"ft"`
	HomelinkDeviceCount int    `json:"homelink_device_count"`
	HomelinkNearby      bool   `json:"homelink_nearby"`
	IsUserPresent       bool   `json:"is_user_present"`
	LastAutoparkError   string `json:"last_autopark_error"`
	Locked              bool   `json:"locked"`
	MediaState          struct {
		RemoteControlEnabled bool `json:"remote_control_enabled"`
	} `json:"media_state"`
	NotificationsSupported  bool    `json:"notifications_supported"`
	Odometer                float64 `json:"odometer"`
	ParsedCalendarSupported bool    `json:"parsed_calendar_supported"`
	Pf                      int     `json:"pf"`
	Pr                      int     `json:"pr"`
	RdWindow                int     `json:"rd_window"`
	RemoteStart             bool    `json:"remote_start"`
	RemoteStartEnabled      bool    `json:"remote_start_enabled"`
	RemoteStartSupported    bool    `json:"remote_start_supported"`
	RpWindow                int     `json:"rp_window"`
	Rt                      int     `json:"rt"`
	SentryMode              bool    `json:"sentry_mode"`
	SentryModeAvailable     bool    `json:"sentry_mode_available"`
	SmartSummonAvailable    bool    `json:"smart_summon_available"`
	SoftwareUpdate          struct {
		DownloadPerc        int    `json:"download_perc"`
		ExpectedDurationSec int    `json:"expected_duration_sec"`
		InstallPerc         int    `json:"install_perc"`
		ScheduledTimeMs     int64  `json:"scheduled_time_ms"`
		Status              string `json:"status"`
		Version             string `json:"version"`
	} `json:"software_update"`
	SpeedLimitMode struct {
		Active          bool    `json:"active"`
		CurrentLimitMph float64 `json:"current_limit_mph"`
		MaxLimitMph     int     `json:"max_limit_mph"`
		MinLimitMph     int     `json:"min_limit_mph"`
		PinCodeSet      bool    `json:"pin_code_set"`
	} `json:"speed_limit_mode"`
	SummonStandbyModeEnabled bool   `json:"summon_standby_mode_enabled"`
	SunRoofPercentOpen       int    `json:"sun_roof_percent_open"`
	SunRoofState             string `json:"sun_roof_state"`
	Timestamp                int64  `json:"timestamp"`
	ValetMode                bool   `json:"valet_mode"`
	ValetPinNeeded           bool   `json:"valet_pin_needed"`
	VehicleName              string `json:"vehicle_name"`
}

// GetVehicleState retrieves the given vehicles current state.
func (c *Conn) GetVehicleState(id int) (*VehicleState, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response VehicleState `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/data_request/vehicle_state", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
