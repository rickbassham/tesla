package tesla

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/text/language"
)

// WakeUp will wake up the vehicle to make it available to receive other commands.
func (c *Conn) WakeUp(id int) (*Vehicle, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response Vehicle `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodPost, fmt.Sprintf("/api/1/vehicles/%d/wake_up", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}

func (c *Conn) doCommand(url string, reqBody interface{}) error {
	if c.accessToken == "" {
		return fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Reason string `json:"reason"`
		Result bool   `json:"result"`
	}

	var respBody response

	err := c.doRequest(http.MethodPost, url, reqBody, &respBody)
	if err != nil {
		return err
	}

	if !respBody.Result {
		return fmt.Errorf("%s: %w", respBody.Reason, ErrCommandError)
	}

	return nil
}

// HonkHorn honks the horn twice.
func (c *Conn) HonkHorn(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/honk_horn", id), nil)
}

// FlashLights flashes the headlights once.
func (c *Conn) FlashLights(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/flash_lights", id), nil)
}

// RemoteStart enables keyless driving. There is a two minute window after issuing the command to
// start driving the car. The password provided is the password for the authenticated tesla.com
// account.
func (c *Conn) RemoteStart(id int, password string) error {
	type request struct {
		Password string `json:"password"`
	}

	reqBody := request{
		Password: password,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/remote_start_drive", id), &reqBody)
}

// TriggerHomelink opens or closes the primary Homelink device. The provided location must be in
// proximity of stored location of the Homelink device.
func (c *Conn) TriggerHomelink(id int, latitude, longitude float64) error {
	type request struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	}

	reqBody := request{
		Latitude:  latitude,
		Longitude: longitude,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/trigger_homelink", id), &reqBody)
}

// SpeedLimitSetLimit sets the maximum speed allowed when Speed Limit Mode is active.
func (c *Conn) SpeedLimitSetLimit(id int, limitMPH int) error {
	type request struct {
		LimitMPH int `json:"limit_mph"`
	}

	reqBody := request{
		LimitMPH: limitMPH,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/speed_limit_set_limit", id), &reqBody)
}

// SpeedLimitActivate activates Speed Limit Mode at the currently set speed.
func (c *Conn) SpeedLimitActivate(id int, pin string) error {
	type request struct {
		Pin string `json:"pin"`
	}

	reqBody := request{
		Pin: pin,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/speed_limit_activate", id), &reqBody)
}

// SpeedLimitDeactivate deactivates Speed Limit Mode if it is currently active.
func (c *Conn) SpeedLimitDeactivate(id int, pin string) error {
	type request struct {
		Pin string `json:"pin"`
	}

	reqBody := request{
		Pin: pin,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/speed_limit_deactivate", id), &reqBody)
}

// SpeedLimitClearPin clears the currently set PIN for Speed Limit Mode.
func (c *Conn) SpeedLimitClearPin(id int, pin string) error {
	type request struct {
		Pin string `json:"pin"`
	}

	reqBody := request{
		Pin: pin,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/speed_limit_clear_pin", id), &reqBody)
}

// SetValetMode activates or deactivates Valet Mode.
//
// Valet Mode limits the car's top speed to 70MPH and 80kW of acceleration power. It also disables
// Homelink, Bluetooth and Wifi settings, and the ability to disable mobile access to the car. It
// also hides your favorites, home, and work locations in navigation.
func (c *Conn) SetValetMode(id int, on bool, pin string) error {
	type request struct {
		On  bool   `json:"on"`
		Pin string `json:"password"`
	}

	reqBody := request{
		On:  on,
		Pin: pin,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/set_valet_mode", id), &reqBody)
}

// ResetValetPin clears the currently set PIN for Valet Mode when deactivated. A new PIN will be
// required when activating again.
func (c *Conn) ResetValetPin(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/reset_valet_pin", id), nil)
}

// SetSentryMode turns sentry mode on or off.
func (c *Conn) SetSentryMode(id int, on bool) error {
	type request struct {
		On bool `json:"on"`
	}

	reqBody := request{
		On: on,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/set_sentry_mode", id), &reqBody)
}

// UnlockDoors unlocks the doors to the car. Extends the handles on the S and X.
func (c *Conn) UnlockDoors(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/door_unlock", id), nil)
}

// LockDoors locks the doors to the car. Retracts the handles on the S and X, if they are extended.
func (c *Conn) LockDoors(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/door_lock", id), nil)
}

// Trunk is used to identify which trunk you are using.
type Trunk string

var (
	// TrunkFront is the front trunk, or frunk.
	TrunkFront = Trunk("front")
	// TrunkRear is the rear trunk.
	TrunkRear = Trunk("rear")
)

// OpenTrunk opens either the front or rear trunk. On the Model S and X, it will also close the rear
// trunk.
func (c *Conn) OpenTrunk(id int, trunk Trunk) error {
	type request struct {
		Trunk Trunk `json:"which_trunk"`
	}

	reqBody := request{
		Trunk: trunk,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/actuate_trunk", id), &reqBody)
}

// WindowCommand is used to identify which direction the window should move.
type WindowCommand string

var (
	// WindowCommandVent will move the windows down.
	WindowCommandVent = WindowCommand("vent")
	// WindowCommandClose will close the windows.
	WindowCommandClose = WindowCommand("close")
)

// ActuateWindows controls the windows. Will vent or close all windows simultaneously.
//
// Location must be near the current location of the car for close operation to succeed.
// For vent, the lat and lon values are ignored, and may both be 0 (which has been observed from
// the app itself).
func (c *Conn) ActuateWindows(id int, cmd WindowCommand, latitude, longitude float64) error {
	type request struct {
		WindowCommand WindowCommand `json:"command"`
		Latitude      float64       `json:"lat"`
		Longitude     float64       `json:"lon"`
	}

	reqBody := request{
		WindowCommand: cmd,
		Latitude:      latitude,
		Longitude:     longitude,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/window_control", id), &reqBody)
}

// SunroofCommand is used to identify which direction the sunroof should move.
type SunroofCommand string

var (
	// SunroofCommandVent will open the sunroof.
	SunroofCommandVent = SunroofCommand("vent")
	// SunroofCommandClose will close the sunroof.
	SunroofCommandClose = SunroofCommand("close")
)

// ActuateSunroof controls the panoramic sunroof on the Model S.
func (c *Conn) ActuateSunroof(id int, cmd SunroofCommand) error {
	type request struct {
		SunroofCommand SunroofCommand `json:"state"`
	}

	reqBody := request{
		SunroofCommand: cmd,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/sun_roof_control", id), &reqBody)
}

// OpenChargePortDoor opens the charge port.
func (c *Conn) OpenChargePortDoor(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/charge_port_door_open", id), nil)
}

// CloseChargePortDoor closes the charge port for vehicles with a motorized charge port door.
func (c *Conn) CloseChargePortDoor(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/charge_port_door_close", id), nil)
}

// StartCharging will start the vehicle charging if the vehicle is plugged in but not currently
// charging.
func (c *Conn) StartCharging(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/charge_start", id), nil)
}

// StopCharging will stop the vehicle charging if the vehicle is currently charging.
func (c *Conn) StopCharging(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/charge_stop", id), nil)
}

// SetChargeLimitStandard sets the charge limit to "standard" or ~90%.
func (c *Conn) SetChargeLimitStandard(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/charge_standard", id), nil)
}

// SetChargeLimitMaxRange sets the charge limit to "max range" or 100%.
func (c *Conn) SetChargeLimitMaxRange(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/charge_max_range", id), nil)
}

// SetChargeLimit sets the charge limit to the given value.
func (c *Conn) SetChargeLimit(id int, percent int) error {
	type request struct {
		Percent int `json:"percent"`
	}

	reqBody := request{
		Percent: percent,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/set_charge_limit", id), &reqBody)
}

// AutoConditioningStart will start the climate control (HVAC) system. Will cool or heat
// automatically, depending on set temperature.
func (c *Conn) AutoConditioningStart(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/auto_conditioning_start", id), nil)
}

// AutoConditioningStop will stop the climate control (HVAC) system.
func (c *Conn) AutoConditioningStop(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/auto_conditioning_stop", id), nil)
}

// SetTemperatures sets the target temperature for the climate control (HVAC) system.
//
// Note: The parameters are always in Celsius, regardless of the region the car is in or the
// display settings of the car.
func (c *Conn) SetTemperatures(id int, driver, passenger float64) error {
	type request struct {
		Driver    float64 `json:"driver_temp"`
		Passenger float64 `json:"passenger_temp"`
	}

	reqBody := request{
		Driver:    driver,
		Passenger: passenger,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/set_temps", id), &reqBody)
}

// SetPreconditioningMax toggles the climate controls between Max Defrost and the previous setting.
func (c *Conn) SetPreconditioningMax(id int, on bool) error {
	type request struct {
		On bool `json:"on"`
	}

	reqBody := request{
		On: on,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/set_preconditioning_max", id), &reqBody)
}

// Seat represents a seat in the vehicle.
type Seat int

var (
	// SeatFrontDriver is the driver's seat.
	SeatFrontDriver = Seat(0)
	// SeatFrontPassenger is the front passenger seat.
	SeatFrontPassenger = Seat(1)
	// SeatRearDriver is the rear seat on the driver's side.
	SeatRearDriver = Seat(2)
	// SeatRearCenter is the rear seat in the center.
	SeatRearCenter = Seat(3)
	// SeatRearPassenger is the rear seat on the passenger's side.
	SeatRearPassenger = Seat(4)
)

// SeatHeatLevel represents a heat level for a seat. 0 is off, 3 is max.
type SeatHeatLevel int

var (
	// SeatHeatLevelZero turns off the heated seat.
	SeatHeatLevelZero = SeatHeatLevel(0)
	// SeatHeatLevelOne is the lowest heat setting.
	SeatHeatLevelOne = SeatHeatLevel(1)
	// SeatHeatLevelTwo is the middle heat setting.
	SeatHeatLevelTwo = SeatHeatLevel(2)
	// SeatHeatLevelThree is the highest heat setting.
	SeatHeatLevelThree = SeatHeatLevel(3)
)

// SetSeatHeater sets the specified seat's heater level.
func (c *Conn) SetSeatHeater(id int, seat Seat, heatLevel SeatHeatLevel) error {
	type request struct {
		Seat      Seat          `json:"heater"`
		HeatLevel SeatHeatLevel `json:"level"`
	}

	reqBody := request{
		Seat:      seat,
		HeatLevel: heatLevel,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/remote_seat_heater_request", id), &reqBody)
}

// SetHeatedSteeringWheel turns steering wheel heater on or off.
func (c *Conn) SetHeatedSteeringWheel(id int, on bool) error {
	type request struct {
		On bool `json:"on"`
	}

	reqBody := request{
		On: on,
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/remote_steering_wheel_heater_request", id), &reqBody)
}

// MediaTogglePlayback toggles the media between playing and paused. For the radio, this mutes or
// unmutes the audio.
func (c *Conn) MediaTogglePlayback(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/media_toggle_playback", id), nil)
}

// MediaNextTrack skips to the next track in the current playlist.
func (c *Conn) MediaNextTrack(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/media_next_track", id), nil)
}

// MediaPreviousTrack skips to the previous track in the current playlist. Does nothing for
// streaming from Stitcher.
func (c *Conn) MediaPreviousTrack(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/media_prev_track", id), nil)
}

// MediaNextFavorite skips to the next saved favorite in the media system.
func (c *Conn) MediaNextFavorite(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/media_next_fav", id), nil)
}

// MediaPreviousFavorite skips to the previous saved favorite in the media system.
func (c *Conn) MediaPreviousFavorite(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/media_prev_fav", id), nil)
}

// MediaVolumeUp turns up the volume of the media system.
func (c *Conn) MediaVolumeUp(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/media_volume_up", id), nil)
}

// MediaVolumeDown turns down the volume of the media system.
func (c *Conn) MediaVolumeDown(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/media_volume_down", id), nil)
}

// Share sends a location for the car to start navigation or play a video in theatre mode.
func (c *Conn) Share(id int, tag language.Tag, text string) error {
	type request struct {
		Type  string `json:"on"`
		Value struct {
			Text string `json:"android.intent.extra.TEXT"`
		} `json:"value"`
		Locale    string `json:"locale"`
		Timestamp int64  `json:"timestamp_ms"`
	}

	reqBody := request{
		Type:      "share_ext_content_raw",
		Locale:    tag.String(),
		Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
	}

	reqBody.Value.Text = text

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/share", id), &reqBody)
}

// ScheduleSoftwareUpdate schedules a software update to be installed, if one is available.
//
// The offset given is how long to delay installing the update.
func (c *Conn) ScheduleSoftwareUpdate(id int, offset time.Duration) error {
	type request struct {
		Offset int `json:"offset_sec"`
	}

	reqBody := request{
		Offset: int(offset / time.Second),
	}

	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/schedule_software_update", id), &reqBody)
}

// CancelSoftwareUpdate cancels a software update, if one is scheduled and has not yet started.
func (c *Conn) CancelSoftwareUpdate(id int) error {
	return c.doCommand(fmt.Sprintf("/api/1/vehicles/%d/command/cancel_software_update", id), nil)
}
