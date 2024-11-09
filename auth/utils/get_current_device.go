package utils

import (
	"github.com/edgar-care/edgarlib/v2/double_auth"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"
	"strings"
)

type DeviceInfoResponse struct {
	Os       string `json:"os"`
	Browser  string `json:"browser"`
	Location string `json:"location"`
}

func GetCurrentUserDevice(w http.ResponseWriter, req *http.Request, accountID string) model.DeviceConnect {
	var device model.DeviceConnect

	ip := GetIPAddress(req)
	userAgent := strings.Join(req.Header["User-Agent"], " ")
	browser := getBrowser(userAgent)

	allDeviceAccount := double_auth.GetDeviceConnect(accountID, 0, 0)
	if allDeviceAccount.Err != nil {
		WriteError(w, allDeviceAccount.Code, allDeviceAccount.Err.Error())
		return model.DeviceConnect{}
	}

	for _, deviceConnected := range allDeviceAccount.DevicesConnect {
		if deviceConnected.IPAddress == ip && deviceConnected.Browser == browser {
			device = deviceConnected
			break
		}
	}
	return device

}

func GetDeviceInfo(req *http.Request) DeviceInfoResponse {
	ip := GetIPAddress(req)
	_, city, _, err := getGeoLocationAndDeviceName(ip)
	if err != nil {
		return DeviceInfoResponse{}
	}

	userAgent := strings.Join(req.Header["User-Agent"], " ")
	deviceType := getDeviceType(userAgent)
	browser := getBrowser(userAgent)

	return DeviceInfoResponse{
		Os:       deviceType,
		Browser:  browser,
		Location: city,
	}
}
