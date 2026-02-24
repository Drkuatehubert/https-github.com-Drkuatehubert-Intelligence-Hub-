package hub

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Flight struct {
	ICAO string     `json:"icao"`
	Pos  Coordinate `json:"pos"`
}

type Vessel struct {
	MMSI string     `json:"mmsi"`
	Pos  Coordinate `json:"pos"`
}

func GetLiveFlights() ([]Flight, error) {
	// Use OpenSky Network API
	resp, err := http.Get("https://opensky-network.org/api/states/all")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch OpenSky states: %s", resp.Status)
	}

	var data struct {
		States [][]interface{} `json:"states"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var flights []Flight
	// Only take the first 50 flights to avoid overwhelming the client
	limit := 50
	if len(data.States) < limit {
		limit = len(data.States)
	}

	for i := 0; i < limit; i++ {
		s := data.States[i]
		if len(s) < 7 {
			continue
		}
		icao, _ := s[0].(string)
		lon, _ := s[5].(float64)
		lat, _ := s[6].(float64)

		flights = append(flights, Flight{
			ICAO: icao,
			Pos:  Coordinate{Lat: lat, Lng: lon},
		})
	}

	return flights, nil
}

func GetLiveVessels() ([]Vessel, error) {
	// Placeholder for actual API call
	return []Vessel{
		{MMSI: "227123456", Pos: Coordinate{Lat: 43.2965, Lng: 5.3698}},
	}, nil
}
