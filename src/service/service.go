package service

import "unisun/api/advisor-listener/src/model"

type ServicePort interface {
	GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) (string, error)
}
