package service

import (
	"encoding/json"
	"io/ioutil"
	"unisun/api/advisor-listener/src/constants"
	"unisun/api/advisor-listener/src/model"
	"unisun/api/advisor-listener/src/utils"

	"github.com/spf13/viper"
)

type ConsumerServiceAdapter struct {
	Utils utils.UtilsPort
}

func NewConsumerServiceAdapter(utils utils.UtilsPort) *ConsumerServiceAdapter {
	return &ConsumerServiceAdapter{
		Utils: utils,
	}
}

func (svr *ConsumerServiceAdapter) GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) (string, error) {
	var serviceIncomeResponse = model.ServiceIncomeResponse{}
	url := viper.GetString("endpoint.strapi-information-gateway.host") + viper.GetString("endpoint.strapi-information-gateway.path")
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		return "", err
	}
	response, err := svr.Utils.HTTPRequest(url, constants.POST, payload)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		return "", err
	}
	return serviceIncomeResponse.Payload, nil
}
