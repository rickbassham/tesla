package tesla

import (
	"fmt"
	"net/http"
)

// ChargingSites represents nearby Tesla-operated charging stations.
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

// GetNearbyChargingSites returns a list of nearby Tesla-operated charging stations. (Requires car
// software version 2018.48 or higher.)
func (c *Conn) GetNearbyChargingSites(id int) (*ChargingSites, error) {
	if c.accessToken == "" {
		return nil, fmt.Errorf("%w", ErrMissingAccessToken)
	}

	type response struct {
		Response ChargingSites `json:"response"`
	}

	var respBody response

	err := c.doRequest(http.MethodGet, fmt.Sprintf("/api/1/vehicles/%d/nearby_charging_sites", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody.Response, nil
}
