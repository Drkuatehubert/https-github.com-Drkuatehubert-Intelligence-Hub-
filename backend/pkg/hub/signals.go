package hub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Device struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	SSID      string  `json:"ssid,omitempty"`
	BSSID     string  `json:"bssid,omitempty"`
	Vendor    string  `json:"vendor,omitempty"`
	Signal    float64 `json:"signal,omitempty"`
	Timestamp string  `json:"timestamp,omitempty"`
	Type      string  `json:"type"`
	IP        string  `json:"ip,omitempty"`
	Info      string  `json:"info,omitempty"`
}

func ClassifyDevice(name string, originalType string) string {
	if name == "" {
		return originalType
	}
	nameUpper := strings.ToUpper(name)

	cars := []string{"CAR", "FORD", "TOYOTA", "BMW", "TESLA", "SYNC", "MAZDA", "HONDA", "UCONNECT", "HYUNDAI", "LEXUS", "NISSAN"}
	for _, k := range cars {
		if strings.Contains(nameUpper, k) {
			return "car"
		}
	}

	tvs := []string{"TV", "BRAVIA", "VIZIO", "SAMSUNG", "LG", "ROKU", "FIRE", "SMARTVIEW", "KDL-"}
	for _, k := range tvs {
		if strings.Contains(nameUpper, k) {
			return "tv"
		}
	}

	audio := []string{"HEADPHONE", "EARBUD", "BOSE", "SONY", "BEATS", "AUDIO", "AIRPOD", "JBL", "SENNHEISER"}
	for _, k := range audio {
		if strings.Contains(nameUpper, k) {
			return "headphone"
		}
	}

	if strings.Contains(nameUpper, "CAM") || strings.Contains(nameUpper, "SURVEILLANCE") {
		return "camera"
	}

	return originalType
}

func GetNearbyDevices(lat, lon float64, mode string) ([]Device, error) {
	var devices []Device

	// Simulated Local Hardware Scan (Wi-Fi/BT)
	if os.Getenv("ENABLE_HARDWARE_SCAN") == "true" {
		devices = append(devices, Device{
			Lat:    lat + 0.0001,
			Lon:    lon + 0.0001,
			SSID:   "LOCAL_HARDWARE_WIFI",
			Type:   "router",
			Vendor: "Ubiquiti",
			Signal: -30,
			Info:   "Detected via local interface",
		})
	}

	wigleName := os.Getenv("WIGLE_API_NAME")
	wigleToken := os.Getenv("WIGLE_API_TOKEN")

	if wigleName != "" && wigleToken != "" {
		endpoint := "https://api.wigle.net/api/v2/network/search"
		if mode == "bluetooth" {
			endpoint = "https://api.wigle.net/api/v2/bluetooth/search"
		}

		req, _ := http.NewRequest("GET", endpoint, nil)
		q := req.URL.Query()
		q.Add("latrange1", fmt.Sprintf("%f", lat-0.01))
		q.Add("latrange2", fmt.Sprintf("%f", lat+0.01))
		q.Add("longrange1", fmt.Sprintf("%f", lon-0.01))
		q.Add("longrange2", fmt.Sprintf("%f", lon+0.01))
		req.URL.RawQuery = q.Encode()
		req.SetBasicAuth(wigleName, wigleToken)

		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			var result struct {
				Results []map[string]interface{} `json:"results"`
			}
			json.NewDecoder(resp.Body).Decode(&result)
			for _, r := range result.Results {
				dev := Device{
					Lat:    r["trilat"].(float64),
					Lon:    r["trilong"].(float64),
					Signal: r["level"].(float64),
				}
				if mode == "bluetooth" {
					dev.SSID, _ = r["name"].(string)
					dev.BSSID, _ = r["netid"].(string)
					dev.Type = ClassifyDevice(dev.SSID, "bluetooth")
				} else {
					dev.SSID, _ = r["ssid"].(string)
					dev.BSSID, _ = r["netid"].(string)
					dev.Vendor, _ = r["vendor"].(string)
					dev.Type = ClassifyDevice(dev.SSID, "router")
				}
				devices = append(devices, dev)
			}
		}
	}

	// Fallback to mock data if no APIs are configured or return nothing
	if len(devices) == 0 {
		devices = []Device{
			{Lat: lat + 0.0005, Lon: lon + 0.0005, SSID: "Mock_Secure_Net", Type: "router", Vendor: "Cisco", Signal: -42},
			{Lat: lat - 0.0005, Lon: lon - 0.0005, SSID: "Mock_BT_Device", Type: "bluetooth", Vendor: "Apple", Signal: -68},
		}
	}

	return devices, nil
}
