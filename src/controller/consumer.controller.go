package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"unisun/api/advisor-listener/src/constants"
	"unisun/api/advisor-listener/src/model"
	"unisun/api/advisor-listener/src/service"

	"github.com/gin-gonic/gin"
)

// Advisor godoc
// @summary      Advisor Listener
// @description  Advisor Listener for the service
// @id           AdvisorListenerHandler
// @tags         advisor
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseAdvisorsAll    "OK"
// @failure      400    {object}    model.ResponseAdvisorsAll    "Bad Request"
// @response     500    {object}    model.ResponseAdvisorsAll    "Internal Server Error"
// @router       /advisor-listener/api/advisors [get]
func Advisors(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseAdvisorsAll{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ADVISOR) + "?populate=thumnail"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		payloadRequest.Path += "&" + strings.Trim(query, "?")
	}
	var advisors model.Advisors
	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &advisors)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
		response.Error = err
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		err = nil
	}
	response.Error = err
	response.Result = advisors.Data
	response.Pagination = advisors.Meta.Pagination
	c.JSON(http.StatusOK, response)
}

// Advisor godoc
// @summary      Advisor Listener
// @description  Advisor Listener for the service
// @id           AdvisorListenerIdHandler
// @tags         advisor
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseAdvisor    "OK"
// @failure      400    {object}    model.ResponseAdvisor    "Bad Request"
// @response     500    {object}    model.ResponseAdvisor    "Internal Server Error"
// @router       /advisor-listener/api/advisors/:id [get]
func AdvisorById(c *gin.Context) {
	id := c.Param("id")
	var advisor model.Advisor
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseAdvisor{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ADVISOR) + "/" + id + "?populate[attachment]=*&populate[categories]=*&populate[thumnail]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &advisor)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
		response.Error = err
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		err = nil
	}
	response.Error = err
	response.Result = advisor.Data
	c.JSON(http.StatusOK, response)
}

// Advisor godoc
// @summary      Advisor Listener
// @description  Advisor Listener for the service
// @id           AdvisorListenerSlugHandler
// @tags         advisor
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseAdvisor    "OK"
// @failure      400    {object}    model.ResponseAdvisor    "Bad Request"
// @response     500    {object}    model.ResponseAdvisor    "Internal Server Error"
// @router       /advisor-listener/api/advisors/slug/:slug [get]
func AdvisorBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var advisors model.AdvisorsSlug
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseAdvisor{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ADVISOR) + "?populate[attachment]=*&populate[categories]=*&populate[thumnail]=*&filters[slug][$eq]=" + slug
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &advisors)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
		response.Error = err
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		err = nil
	}
	advisor := advisors.Data[0]
	response.Error = err
	response.Result = advisor
	c.JSON(http.StatusOK, response)
}
