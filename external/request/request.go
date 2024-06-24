package request

import (
	"fmt"

	"github.com/SineChat/notification-ms/external/mocks"
	"github.com/SineChat/notification-ms/external/thirdparty/ipstack"
	"github.com/SineChat/notification-ms/external/thirdparty/termii"
	"github.com/SineChat/notification-ms/internal/config"
	"github.com/SineChat/notification-ms/utility"
)

type ExternalRequest struct {
	Logger *utility.Logger
	Test   bool
}

var (
	JsonDecodeMethod    string = "json"
	PhpSerializerMethod string = "phpserializer"

	// microservice

	// third party
	IpstackResolveIp string = "ipstack_resolve_ip"
	TermiiSendSMS    string = "termii_send_sms"
)

func (er ExternalRequest) SendExternalRequest(name string, data interface{}) (interface{}, error) {
	var (
		config = config.GetConfig()
	)
	if !er.Test {
		switch name {
		case IpstackResolveIp:
			obj := ipstack.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v", config.IPStack.BaseUrl),
				Method:       "GET",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.IpstackResolveIp()
		case TermiiSendSMS:
			obj := termii.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v/api/sms/send", config.Termii.BaseUrl),
				Method:       "POST",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.TermiiSendSMS()
		default:
			return nil, fmt.Errorf("request not found")
		}

	} else {
		mer := mocks.ExternalRequest{Logger: er.Logger, Test: true}
		return mer.SendExternalRequest(name, data)
	}
}
