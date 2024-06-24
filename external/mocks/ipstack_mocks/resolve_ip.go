package ipstack_mocks

import (
	"fmt"

	"github.com/SineChat/notification-ms/external/external_models"
	"github.com/SineChat/notification-ms/internal/config"
	"github.com/SineChat/notification-ms/utility"
)

func IpstackResolveIp(logger *utility.Logger, idata interface{}) (external_models.IPStackResolveIPResponse, error) {

	var (
		key              = config.GetConfig().IPStack.Key
		outBoundResponse external_models.IPStackResolveIPResponse
	)

	ip, ok := idata.(string)
	if !ok {
		logger.Error("ipstack resolve ip", idata, "request data format error")
		return outBoundResponse, fmt.Errorf("request data format error")
	}
	outBoundResponse.Ip = ip
	outBoundResponse.City = "city"
	outBoundResponse.CountryName = "name"

	path := "/" + ip + "?access_key=" + key

	logger.Info("ipstack resolve ip", ip, path)

	return outBoundResponse, nil
}
