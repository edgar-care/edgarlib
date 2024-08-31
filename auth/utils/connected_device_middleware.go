package utils

import (
	"encoding/json"
	"fmt"
	"github.com/edgar-care/edgarlib/v2/double_auth"
	"net"
	"net/http"
	"strings"
	"time"
)

type IPInfoResponse struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

type DeviceInfo struct {
	IPAddress     string `json:"ip_address"`
	DeviceName    string `json:"device_name"`
	City          string `json:"city"`
	Country       string `json:"country"`
	OperationTime string `json:"operation_time"`
}

func GetIPAddress(r *http.Request) string {
	// Vérifie l'en-tête X-Forwarded-For
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For peut contenir une liste d'adresses IP séparées par des virgules
		ips := strings.Split(ip, ",")
		for _, addr := range ips {
			addr = strings.TrimSpace(addr)
			// Valide si l'IP est une adresse IPv4 valide
			if net.ParseIP(addr) != nil && strings.Contains(addr, ".") {
				return addr
			}
		}
	}

	// Vérifie l'en-tête X-Real-IP
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		// Valide si l'IP est une adresse IPv4 valide
		if net.ParseIP(ip) != nil && strings.Contains(ip, ".") {
			return ip
		}
	}

	// Utilise l'adresse IP distante si aucune des autres méthodes n'a fonctionné
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	if net.ParseIP(ip) != nil && strings.Contains(ip, ".") {
		return ip
	}

	return ""
}

func getGeoLocationAndDeviceName(ip string) (string, string, string, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json?token=bde9bcbd8cf979", ip)

	resp, err := http.Get(url)
	if err != nil {
		return "Unknown Device", "Unknown City", "Unknown Country", err
	}
	defer resp.Body.Close()

	var ipInfo IPInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return "Unknown Device", "Unknown City", "Unknown Country", err
	}

	deviceName := ipInfo.Hostname
	if deviceName == "" {
		deviceName = "Unknown Device"
	}

	city := ipInfo.City
	if city == "" {
		city = "Unknown City"
	}

	country := ipInfo.Country
	if country == "" {
		country = "Unknown Country"
	}

	return deviceName, city, country, nil
}

func DeviceConnectMiddleware(w http.ResponseWriter, r *http.Request, accountID string) string {

	ip := GetIPAddress(r)
	_, city, country, err := getGeoLocationAndDeviceName(ip)
	if err != nil {
		fmt.Println("Error fetching double_auth information:", err)
		http.Error(w, "Error fetching double_auth information", 400)
		return ""
	}

	operationTime := time.Now().Unix() // Unix timestamp

	userAgent := strings.Join(r.Header["User-Agent"], " ")
	deviceType := getDeviceType(userAgent)
	browser := getBrowser(userAgent)

	getAllDevice := double_auth.GetDeviceConnect(accountID)
	if getAllDevice.Err != nil {
		WriteError(w, getAllDevice.Code, getAllDevice.Err.Error())
		return ""
	}

	var deviceID string
	for _, device := range getAllDevice.DevicesConnect {
		if device.IPAddress == ip && device.Browser == browser {
			deviceID = device.ID
		}
	}

	if deviceID == "" {
		input := double_auth.CreateDeviceConnectInput{
			DeviceType: deviceType,
			Browser:    browser,
			Ip:         ip,
			City:       city,
			Country:    country,
			Date:       int(operationTime),
		}
		response := double_auth.CreateDeviceConnect(input, accountID)
		if response.Err != nil {
			WriteError(w, response.Code, response.Err.Error())
			return ""
		}
	} else {
		updateInput := double_auth.UpdateDeviceConnectInput{
			DeviceType: deviceType,
			Browser:    browser,
			Ip:         ip,
			City:       city,
			Country:    country,
			Date:       int(operationTime),
		}
		update := double_auth.UpdateDeviceConnect(updateInput, deviceID)
		if update.Err != nil {
			WriteError(w, update.Code, update.Err.Error())
			return ""
		}
	}
	return "Successfully fetched double_auth information"
}

func GetDeviceName(ip string) (string, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json?token=bde9bcbd8cf979", ip)

	resp, err := http.Get(url)
	if err != nil {
		return "Unknown Device", err
	}
	defer resp.Body.Close()

	var ipInfo IPInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return "Unknown Device", err
	}

	deviceName := ipInfo.Hostname
	if deviceName == "" {
		deviceName = "Unknown Device"
	}

	return deviceName, nil
}

func GetDeviceNameWithIp(w http.ResponseWriter, r *http.Request) string {

	ip := GetIPAddress(r)
	deviceName, err := GetDeviceName(ip)
	if err != nil {
		fmt.Println("Error fetching device name:", err)
		http.Error(w, "Error fetching device name", 400)
		return ""
	}

	return deviceName
}

func getDeviceType(userAgent string) string {
	if strings.Contains(userAgent, "iPhone") {
		return "iPhone"
	} else if strings.Contains(userAgent, "Android") {
		return "Android"
	} else if strings.Contains(userAgent, "Windows") {
		return "Windows"
	} else if strings.Contains(userAgent, "Macintosh") {
		return "MacOs"
	} else if strings.Contains(userAgent, "Linux") {
		return "Linux"
	} else {
		return "Other"
	}
}

func getBrowser(userAgent string) string {
	userAgent = strings.ToLower(userAgent) // Convert to lowercase for case-insensitive comparison
	switch {
	case strings.Contains(userAgent, "edg"):
		return "Edge"
	case strings.Contains(userAgent, "opr"):
		return "Opera"
	case strings.Contains(userAgent, "chrome"):
		return "Chrome"
	case strings.Contains(userAgent, "firefox"):
		return "Firefox"
	case strings.Contains(userAgent, "safari") && !strings.Contains(userAgent, "chrome"):
		return "Safari"
	default:
		return "Other"
	}
}
