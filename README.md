# tesla
--
    import "."

Package tesla is a client for interacting with the Tesla Owner's API. Unofficial
documentation can be found at https://tesla-api.timdorr.com. This package
implements all of the documented endpoints, as well as websocket streaming for
live data while driving.

## Usage

```go
const (
	// DefaultBaseURL is the URL for the Tesla owner's API.
	DefaultBaseURL = "https://owner-api.teslamotors.com"
)
```

```go
var (
	// TrunkFront is the front trunk, or frunk.
	TrunkFront = Trunk("front")
	// TrunkRear is the rear trunk.
	TrunkRear = Trunk("rear")
)
```

```go
var (
	// WindowCommandVent will move the windows down.
	WindowCommandVent = WindowCommand("vent")
	// WindowCommandClose will close the windows.
	WindowCommandClose = WindowCommand("close")
)
```

```go
var (
	// SunroofCommandVent will open the sunroof.
	SunroofCommandVent = SunroofCommand("vent")
	// SunroofCommandClose will close the sunroof.
	SunroofCommandClose = SunroofCommand("close")
)
```

```go
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
```

```go
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
```

```go
var (
	// ErrMissingRefreshToken is returned when an API call is made without the required refresh token.
	ErrMissingRefreshToken = errors.New("missing refresh token")
	// ErrMissingAccessToken is returned when an API call is made without the required access token.
	ErrMissingAccessToken = errors.New("missing access token, authenticate first")
	// ErrCommandError is returned with executing a command against the vehicle and the Tesla API returns an error message.
	ErrCommandError = errors.New("error executing command")
)
```

#### type ChargeState

```go
type ChargeState struct {
	BatteryHeaterOn             bool        `json:"battery_heater_on"`
	BatteryLevel                int         `json:"battery_level"`
	BatteryRange                float64     `json:"battery_range"`
	ChargeCurrentRequest        int         `json:"charge_current_request"`
	ChargeCurrentRequestMax     int         `json:"charge_current_request_max"`
	ChargeEnableRequest         bool        `json:"charge_enable_request"`
	ChargeEnergyAdded           float64     `json:"charge_energy_added"`
	ChargeLimitSoc              int         `json:"charge_limit_soc"`
	ChargeLimitSocMax           int         `json:"charge_limit_soc_max"`
	ChargeLimitSocMin           int         `json:"charge_limit_soc_min"`
	ChargeLimitSocStd           int         `json:"charge_limit_soc_std"`
	ChargeMilesAddedIdeal       float64     `json:"charge_miles_added_ideal"`
	ChargeMilesAddedRated       float64     `json:"charge_miles_added_rated"`
	ChargePortColdWeatherMode   bool        `json:"charge_port_cold_weather_mode"`
	ChargePortDoorOpen          bool        `json:"charge_port_door_open"`
	ChargePortLatch             string      `json:"charge_port_latch"`
	ChargeRate                  float64     `json:"charge_rate"`
	ChargeToMaxRange            bool        `json:"charge_to_max_range"`
	ChargerActualCurrent        int         `json:"charger_actual_current"`
	ChargerPhases               interface{} `json:"charger_phases"`
	ChargerPilotCurrent         int         `json:"charger_pilot_current"`
	ChargerPower                int         `json:"charger_power"`
	ChargerVoltage              int         `json:"charger_voltage"`
	ChargingState               string      `json:"charging_state"`
	ConnChargeCable             string      `json:"conn_charge_cable"`
	EstBatteryRange             float64     `json:"est_battery_range"`
	FastChargerBrand            string      `json:"fast_charger_brand"`
	FastChargerPresent          bool        `json:"fast_charger_present"`
	FastChargerType             string      `json:"fast_charger_type"`
	IdealBatteryRange           float64     `json:"ideal_battery_range"`
	ManagedChargingActive       bool        `json:"managed_charging_active"`
	ManagedChargingStartTime    *int        `json:"managed_charging_start_time"`
	ManagedChargingUserCanceled bool        `json:"managed_charging_user_canceled"`
	MaxRangeChargeCounter       int         `json:"max_range_charge_counter"`
	MinutesToFullCharge         int         `json:"minutes_to_full_charge"`
	NotEnoughPowerToHeat        bool        `json:"not_enough_power_to_heat"`
	ScheduledChargingPending    bool        `json:"scheduled_charging_pending"`
	ScheduledChargingStartTime  *int        `json:"scheduled_charging_start_time"`
	ScheduledDepartureTime      *int        `json:"scheduled_departure_time"`
	TimeToFullCharge            float64     `json:"time_to_full_charge"`
	Timestamp                   int64       `json:"timestamp"`
	TripCharging                bool        `json:"trip_charging"`
	UsableBatteryLevel          int         `json:"usable_battery_level"`
	UserChargeEnableRequest     interface{} `json:"user_charge_enable_request"`
}
```

ChargeState is the current state of charging for the vehicle.

#### type ChargingSites

```go
type ChargingSites struct {
	CongestionSyncTimeUtcSecs int `json:"congestion_sync_time_utc_secs"`
	DestinationCharging       []struct {
		Location struct {
			Lat  float64 `json:"lat"`
			Long float64 `json:"long"`
		} `json:"location"`
		Name          string  `json:"name"`
		Type          string  `json:"type"`
		DistanceMiles float64 `json:"distance_miles"`
	} `json:"destination_charging"`
	Superchargers []struct {
		Location struct {
			Lat  float64 `json:"lat"`
			Long float64 `json:"long"`
		} `json:"location"`
		Name            string  `json:"name"`
		Type            string  `json:"type"`
		DistanceMiles   float64 `json:"distance_miles"`
		AvailableStalls int     `json:"available_stalls"`
		TotalStalls     int     `json:"total_stalls"`
		SiteClosed      bool    `json:"site_closed"`
	} `json:"superchargers"`
	Timestamp int64 `json:"timestamp"`
}
```

ChargingSites represents nearby Tesla-operated charging stations.

#### type ClimateState

```go
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
```

ClimateState represents the current state of climate control for the vehicle.

#### type Conn

```go
type Conn struct {
}
```

Conn represents a connection to the Tesla owner's API.

#### func  NewConn

```go
func NewConn(rt http.RoundTripper, baseURL, clientID, clientSecret string) *Conn
```
NewConn creates a new connection.

#### func (*Conn) ActuateSunroof

```go
func (c *Conn) ActuateSunroof(id int, cmd SunroofCommand) error
```
ActuateSunroof controls the panoramic sunroof on the Model S.

#### func (*Conn) ActuateWindows

```go
func (c *Conn) ActuateWindows(id int, cmd WindowCommand, latitude, longitude float64) error
```
ActuateWindows controls the windows. Will vent or close all windows
simultaneously.

Location must be near the current location of the car for close operation to
succeed. For vent, the lat and lon values are ignored, and may both be 0 (which
has been observed from the app itself).

#### func (*Conn) Authenticate

```go
func (c *Conn) Authenticate(email, password string) error
```
Authenticate starts the initial authentication process via an OAuth 2.0 Password
Grant with the same credentials used for tesla.com and the mobile apps.

The current client ID and secret are available at https://pastebin.com/pS7Z6yyP.

We will get back an access_token which is treated as an OAuth 2.0 Bearer Token.
This token is passed along in an Authorization header with all future requests:

#### func (*Conn) AutoConditioningStart

```go
func (c *Conn) AutoConditioningStart(id int) error
```
AutoConditioningStart will start the climate control (HVAC) system. Will cool or
heat automatically, depending on set temperature.

#### func (*Conn) AutoConditioningStop

```go
func (c *Conn) AutoConditioningStop(id int) error
```
AutoConditioningStop will stop the climate control (HVAC) system.

#### func (*Conn) CancelSoftwareUpdate

```go
func (c *Conn) CancelSoftwareUpdate(id int) error
```
CancelSoftwareUpdate cancels a software update, if one is scheduled and has not
yet started.

#### func (*Conn) CloseChargePortDoor

```go
func (c *Conn) CloseChargePortDoor(id int) error
```
CloseChargePortDoor closes the charge port for vehicles with a motorized charge
port door.

#### func (*Conn) FlashLights

```go
func (c *Conn) FlashLights(id int) error
```
FlashLights flashes the headlights once.

#### func (*Conn) GetChargeState

```go
func (c *Conn) GetChargeState(id int) (*ChargeState, error)
```
GetChargeState gets information on the state of charge in the battery and its
various settings.

#### func (*Conn) GetClimateState

```go
func (c *Conn) GetClimateState(id int) (*ClimateState, error)
```
GetClimateState retrieves information on the current internal temperature and
climate control system.

#### func (*Conn) GetDriveState

```go
func (c *Conn) GetDriveState(id int) (*DriveState, error)
```
GetDriveState retrieves the driving and position state of the vehicle.

#### func (*Conn) GetGUISettings

```go
func (c *Conn) GetGUISettings(id int) (*GUISettings, error)
```
GetGUISettings retrieves the current GUI settings for the vehicle.

#### func (*Conn) GetMobileEnabled

```go
func (c *Conn) GetMobileEnabled(id int) (bool, error)
```
GetMobileEnabled returns whether or not the Mobile Access setting is enabled in
the vehicle.

#### func (*Conn) GetNearbyChargingSites

```go
func (c *Conn) GetNearbyChargingSites(id int) (*ChargingSites, error)
```
GetNearbyChargingSites returns a list of nearby Tesla-operated charging
stations. (Requires car software version 2018.48 or higher.)

#### func (*Conn) GetVehicle

```go
func (c *Conn) GetVehicle(id int) (*Vehicle, error)
```
GetVehicle returns a vehicle by id.

#### func (*Conn) GetVehicleConfig

```go
func (c *Conn) GetVehicleConfig(id int) (*VehicleConfig, error)
```
GetVehicleConfig retrieves the vehicles config.

#### func (*Conn) GetVehicleState

```go
func (c *Conn) GetVehicleState(id int) (*VehicleState, error)
```
GetVehicleState retrieves the given vehicles current state.

#### func (*Conn) GetVehicles

```go
func (c *Conn) GetVehicles() ([]Vehicle, error)
```
GetVehicles retrieves a list of vehicles for the currently authenticated
account.

#### func (*Conn) HonkHorn

```go
func (c *Conn) HonkHorn(id int) error
```
HonkHorn honks the horn twice.

#### func (*Conn) LockDoors

```go
func (c *Conn) LockDoors(id int) error
```
LockDoors locks the doors to the car. Retracts the handles on the S and X, if
they are extended.

#### func (*Conn) MediaNextFavorite

```go
func (c *Conn) MediaNextFavorite(id int) error
```
MediaNextFavorite skips to the next saved favorite in the media system.

#### func (*Conn) MediaNextTrack

```go
func (c *Conn) MediaNextTrack(id int) error
```
MediaNextTrack skips to the next track in the current playlist.

#### func (*Conn) MediaPreviousFavorite

```go
func (c *Conn) MediaPreviousFavorite(id int) error
```
MediaPreviousFavorite skips to the previous saved favorite in the media system.

#### func (*Conn) MediaPreviousTrack

```go
func (c *Conn) MediaPreviousTrack(id int) error
```
MediaPreviousTrack skips to the previous track in the current playlist. Does
nothing for streaming from Stitcher.

#### func (*Conn) MediaTogglePlayback

```go
func (c *Conn) MediaTogglePlayback(id int) error
```
MediaTogglePlayback toggles the media between playing and paused. For the radio,
this mutes or unmutes the audio.

#### func (*Conn) MediaVolumeDown

```go
func (c *Conn) MediaVolumeDown(id int) error
```
MediaVolumeDown turns down the volume of the media system.

#### func (*Conn) MediaVolumeUp

```go
func (c *Conn) MediaVolumeUp(id int) error
```
MediaVolumeUp turns up the volume of the media system.

#### func (*Conn) OpenChargePortDoor

```go
func (c *Conn) OpenChargePortDoor(id int) error
```
OpenChargePortDoor opens the charge port.

#### func (*Conn) OpenTrunk

```go
func (c *Conn) OpenTrunk(id int, trunk Trunk) error
```
OpenTrunk opens either the front or rear trunk. On the Model S and X, it will
also close the rear trunk.

#### func (*Conn) RemoteStart

```go
func (c *Conn) RemoteStart(id int, password string) error
```
RemoteStart enables keyless driving. There is a two minute window after issuing
the command to start driving the car. The password provided is the password for
the authenticated tesla.com account.

#### func (*Conn) ResetValetPin

```go
func (c *Conn) ResetValetPin(id int) error
```
ResetValetPin clears the currently set PIN for Valet Mode when deactivated. A
new PIN will be required when activating again.

#### func (*Conn) ScheduleSoftwareUpdate

```go
func (c *Conn) ScheduleSoftwareUpdate(id int, offset time.Duration) error
```
ScheduleSoftwareUpdate schedules a software update to be installed, if one is
available.

The offset given is how long to delay installing the update.

#### func (*Conn) SetAccessToken

```go
func (c *Conn) SetAccessToken(accessToken string)
```
SetAccessToken allows you to override the access token received from
Authenticate.

#### func (*Conn) SetChargeLimit

```go
func (c *Conn) SetChargeLimit(id int, percent int) error
```
SetChargeLimit sets the charge limit to the given value.

#### func (*Conn) SetChargeLimitMaxRange

```go
func (c *Conn) SetChargeLimitMaxRange(id int) error
```
SetChargeLimitMaxRange sets the charge limit to "max range" or 100%.

#### func (*Conn) SetChargeLimitStandard

```go
func (c *Conn) SetChargeLimitStandard(id int) error
```
SetChargeLimitStandard sets the charge limit to "standard" or ~90%.

#### func (*Conn) SetDebugMode

```go
func (c *Conn) SetDebugMode(debug bool)
```
SetDebugMode turns debug mode on or off. If on, all requests and responses are
dumped in their raw state to stdout.

#### func (*Conn) SetHeatedSteeringWheel

```go
func (c *Conn) SetHeatedSteeringWheel(id int, on bool) error
```
SetHeatedSteeringWheel turns steering wheel heater on or off.

#### func (*Conn) SetPreconditioningMax

```go
func (c *Conn) SetPreconditioningMax(id int, on bool) error
```
SetPreconditioningMax toggles the climate controls between Max Defrost and the
previous setting.

#### func (*Conn) SetRefreshToken

```go
func (c *Conn) SetRefreshToken(refreshToken string)
```
SetRefreshToken allows you to override the refresh token received from
Authenticate.

#### func (*Conn) SetSeatHeater

```go
func (c *Conn) SetSeatHeater(id int, seat Seat, heatLevel SeatHeatLevel) error
```
SetSeatHeater sets the specified seat's heater level.

#### func (*Conn) SetSentryMode

```go
func (c *Conn) SetSentryMode(id int, on bool) error
```
SetSentryMode turns sentry mode on or off.

#### func (*Conn) SetTemperatures

```go
func (c *Conn) SetTemperatures(id int, driver, passenger float64) error
```
SetTemperatures sets the target temperature for the climate control (HVAC)
system.

Note: The parameters are always in Celsius, regardless of the region the car is
in or the display settings of the car.

#### func (*Conn) SetValetMode

```go
func (c *Conn) SetValetMode(id int, on bool, pin string) error
```
SetValetMode activates or deactivates Valet Mode.

Valet Mode limits the car's top speed to 70MPH and 80kW of acceleration power.
It also disables Homelink, Bluetooth and Wifi settings, and the ability to
disable mobile access to the car. It also hides your favorites, home, and work
locations in navigation.

#### func (*Conn) Share

```go
func (c *Conn) Share(id int, tag language.Tag, text string) error
```
Share sends a location for the car to start navigation or play a video in
theatre mode.

#### func (*Conn) SpeedLimitActivate

```go
func (c *Conn) SpeedLimitActivate(id int, pin string) error
```
SpeedLimitActivate activates Speed Limit Mode at the currently set speed.

#### func (*Conn) SpeedLimitClearPin

```go
func (c *Conn) SpeedLimitClearPin(id int, pin string) error
```
SpeedLimitClearPin clears the currently set PIN for Speed Limit Mode.

#### func (*Conn) SpeedLimitDeactivate

```go
func (c *Conn) SpeedLimitDeactivate(id int, pin string) error
```
SpeedLimitDeactivate deactivates Speed Limit Mode if it is currently active.

#### func (*Conn) SpeedLimitSetLimit

```go
func (c *Conn) SpeedLimitSetLimit(id int, limitMPH int) error
```
SpeedLimitSetLimit sets the maximum speed allowed when Speed Limit Mode is
active.

#### func (*Conn) StartCharging

```go
func (c *Conn) StartCharging(id int) error
```
StartCharging will start the vehicle charging if the vehicle is plugged in but
not currently charging.

#### func (*Conn) StopCharging

```go
func (c *Conn) StopCharging(id int) error
```
StopCharging will stop the vehicle charging if the vehicle is currently
charging.

#### func (*Conn) Stream

```go
func (c *Conn) Stream(id int, token string) (*Stream, error)
```
Stream will initiate a stream of data from the car, with updates going to the
returned channel. New messages are received approximately every 250ms, but that
is not reliable. If the API closes the stream, the returned channel will be
closed as well.

#### func (*Conn) TriggerHomelink

```go
func (c *Conn) TriggerHomelink(id int, latitude, longitude float64) error
```
TriggerHomelink opens or closes the primary Homelink device. The provided
location must be in proximity of stored location of the Homelink device.

#### func (*Conn) UnlockDoors

```go
func (c *Conn) UnlockDoors(id int) error
```
UnlockDoors unlocks the doors to the car. Extends the handles on the S and X.

#### func (*Conn) UpdateRefreshToken

```go
func (c *Conn) UpdateRefreshToken() error
```
UpdateRefreshToken will do an OAuth 2.0 Refresh Token Grant and obtain a new
access token. Note: This will invalidate the previous access token.

#### func (*Conn) WakeUp

```go
func (c *Conn) WakeUp(id int) (*Vehicle, error)
```
WakeUp will wake up the vehicle to make it available to receive other commands.

#### type DriveState

```go
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
```

DriveState is the current state of driving for the vehicle.

#### type GUISettings

```go
type GUISettings struct {
	GUI24HourTime       bool   `json:"gui_24_hour_time"`
	GUIChargeRateUnits  string `json:"gui_charge_rate_units"`
	GUIDistanceUnits    string `json:"gui_distance_units"`
	GUIRangeDisplay     string `json:"gui_range_display"`
	GUITemperatureUnits string `json:"gui_temperature_units"`
	ShowRangeUnits      bool   `json:"show_range_units"`
	Timestamp           int64  `json:"timestamp"`
}
```

GUISettings represents the configured settings in the vehicle GUI.

#### type HTTPStatusError

```go
type HTTPStatusError struct {
}
```

HTTPStatusError is returned if the response code received from the API is
non-200.

#### func (HTTPStatusError) Error

```go
func (err HTTPStatusError) Error() string
```

#### type Seat

```go
type Seat int
```

Seat represents a seat in the vehicle.

#### type SeatHeatLevel

```go
type SeatHeatLevel int
```

SeatHeatLevel represents a heat level for a seat. 0 is off, 3 is max.

#### type Stream

```go
type Stream struct {
}
```


#### func (*Stream) Close

```go
func (s *Stream) Close()
```

#### func (*Stream) Data

```go
func (s *Stream) Data() <-chan StreamingMessage
```

#### func (*Stream) Err

```go
func (s *Stream) Err() error
```

#### type StreamingMessage

```go
type StreamingMessage struct {
	Timestamp    time.Time
	Speed        int
	Odometer     float64
	SOC          int
	Elevation    int
	EstHeading   int
	EstLatitude  float64
	EstLongitude float64
	Power        int
	ShiftState   int
	Range        int
	EstRange     int
	Heading      int
}
```

StreamingMessage represents the current state of the car.

#### type SunroofCommand

```go
type SunroofCommand string
```

SunroofCommand is used to identify which direction the sunroof should move.

#### type Trunk

```go
type Trunk string
```

Trunk is used to identify which trunk you are using.

#### type Vehicle

```go
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
```

Vehicle represents the basic data about the vehicle.

#### type VehicleConfig

```go
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
```

VehicleConfig represents the current capabilities of the vehicle.

#### type VehicleState

```go
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
```

VehicleState represents the current vehicle state.

#### type WindowCommand

```go
type WindowCommand string
```

WindowCommand is used to identify which direction the window should move.
