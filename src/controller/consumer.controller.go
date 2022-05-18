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

func Advisors(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
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
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "result": advisors.Data, "panigation": advisors.Meta.Pagination})
}

func AdvisorById(c *gin.Context) {
	id := c.Param("id")
	var advisor model.Advisor
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ADVISOR) + "/" + id + "?populate[attachment]=*&populate[categories]=*&populate[thumnail]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &advisor)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "result": advisor})
}

func AdvisorBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var advisors model.AdvisorsSlug
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ADVISOR) + "?populate[attachment]=*&populate[categories]=*&populate[thumnail]=*&filters[slug][$eq]=" + slug
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &advisors)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	advisor := advisors.Data[0]
	c.JSON(http.StatusOK, gin.H{"error": err, "result": advisor})
}
