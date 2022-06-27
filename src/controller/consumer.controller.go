package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unisun/api/advisor-listener/src/constants"
	"unisun/api/advisor-listener/src/model"
	"unisun/api/advisor-listener/src/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ConsumerControllerAdapter struct {
	Service service.ServicePort
}

func NewConsumerControllerAdapter(serviceValue service.ServicePort) *ConsumerControllerAdapter {
	return &ConsumerControllerAdapter{
		Service: serviceValue,
	}
}

// Advisor godoc
// @summary      Advisor Listener
// @description  Advisor Listener for the service
// @id           AdvisorListenerHandler
// @tags         advisor
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseAdvisors    "OK"
// @failure      400    {object}    model.ResponseFail    "Bad Request"
// @response     500    {object}    model.ResponseFail    "Internal Server Error"
// @router       /advisor-listener/api/advisors [get]
func (srv *ConsumerControllerAdapter) Advisors(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseAdvisors{}
	var responseError = model.ResponseFail{}
	payloadRequest.Path = strings.Join([]string{viper.GetString("endpoint.advisor.path"), viper.GetString("endpoint.advisor.query.value")}, "")
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		payloadRequest.Path += "&" + strings.Trim(query, "?")
	}
	data, err := srv.Service.GetInformationFormStrapi(payloadRequest)
	if err != nil {
		responseError.Error.Status = http.StatusBadRequest
		responseError.Error.Message = err.Error()
		responseError.Error.Name = ""
		responseError.Error.Details = err
		c.AbortWithStatusJSON(http.StatusBadRequest, responseError)
		log.Panic(err)
		return
	}
	if err = json.Unmarshal([]byte(data), &response); err != nil {
		responseError.Error.Status = http.StatusBadRequest
		responseError.Error.Message = err.Error()
		responseError.Error.Name = "Change byte to json article"
		responseError.Error.Details = err
		c.AbortWithStatusJSON(http.StatusBadRequest, responseError)
		log.Panic(err)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}

// Advisor godoc
// @summary      Advisor Listener
// @description  Advisor Listener for the service
// @id           AdvisorListenerIdHandler
// @tags         advisor
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseAdvisor    "OK"
// @failure      400    {object}    model.ResponseFail   "Bad Request"
// @response     500    {object}    model.ResponseFail   "Internal Server Error"
// @router       /advisor-listener/api/advisors/:id [get]
func (srv *ConsumerControllerAdapter) AdvisorById(c *gin.Context) {
	id := c.Param("id")
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseAdvisor{}
	var responseError = model.ResponseFail{}
	payloadRequest.Path = strings.Join([]string{viper.GetString("endpoint.advisor.path"), "/", id, viper.GetString("endpoint.advisor.query.value")}, "")
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data, err := srv.Service.GetInformationFormStrapi(payloadRequest)
	if err != nil {
		responseError.Error.Status = http.StatusBadRequest
		responseError.Error.Message = err.Error()
		responseError.Error.Name = ""
		responseError.Error.Details = err
		c.AbortWithStatusJSON(http.StatusBadRequest, responseError)
		log.Panic(err)
		return
	}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		responseError.Error.Status = http.StatusBadRequest
		responseError.Error.Message = err.Error()
		responseError.Error.Name = "Change byte to json article"
		responseError.Error.Details = err
		c.AbortWithStatusJSON(http.StatusBadRequest, responseError)
		log.Panic("Change byte to json article", err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

// Advisor godoc
// @summary      Advisor Listener
// @description  Advisor Listener for the service
// @id           AdvisorListenerSlugHandler
// @tags         advisor
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseAdvisors    "OK"
// @failure      400    {object}    model.ResponseFail    "Bad Request"
// @response     500    {object}    model.ResponseFail    "Internal Server Error"
// @router       /advisor-listener/api/advisors/slug/:slug [get]
func (srv *ConsumerControllerAdapter) AdvisorBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseAdvisors{}
	var responseError = model.ResponseFail{}
	payloadRequest.Path = strings.Join([]string{viper.GetString("endpoint.advisor.path"), viper.GetString("endpoint.advisor.query.value"), "&filters[slug][$eq]=", slug}, "")
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data, err := srv.Service.GetInformationFormStrapi(payloadRequest)
	if err != nil {
		responseError.Error.Status = http.StatusBadRequest
		responseError.Error.Message = err.Error()
		responseError.Error.Name = ""
		responseError.Error.Details = err
		c.AbortWithStatusJSON(http.StatusBadRequest, responseError)
		log.Panic(err)
		return
	}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		responseError.Error.Status = http.StatusBadRequest
		responseError.Error.Message = err.Error()
		responseError.Error.Name = "Change byte to json article"
		responseError.Error.Details = err
		c.AbortWithStatusJSON(http.StatusBadRequest, responseError)
		log.Panic("Change byte to json article", err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}
