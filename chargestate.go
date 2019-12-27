package tesla

import (
	"fmt"
	"net/http"
)

// ChargeState is the current state of charging for the vehicle.
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

// GetChargeState gets information on the state of charge in the battery and its various settings.
func (c *Conn) GetChargeState(id int) (*ChargeState, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response ChargeState `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/data_request/charge_state", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
