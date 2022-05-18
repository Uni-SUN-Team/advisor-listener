package config

import (
	"os"
	"unisun/api/advisor-listener/src/constants"
)

func ConfigENV() {
	os.Setenv(constants.PORT, "8080")
	os.Setenv(constants.CONTEXT_PATH, "/advisor-listener")
	// Strapi information gateway
	os.Setenv(constants.HOST_STRAPI_SERVICE, "https://api.unisun.dynu.com")
	os.Setenv(constants.PATH_STRAPI_INFORMATION_GATEWAY, "/strapi-information-gateway/api/strapi")
	// Path
	os.Setenv(constants.PATH_STRAPI_ADVISOR, "/api/advisors")
}
