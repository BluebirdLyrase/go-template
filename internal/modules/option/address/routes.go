package address

import (
	"my-api/internal/modules/option/address/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *handlers.AddressHandler) {
	auth := router.Group("/options/address")
	{
		auth.GET("/countries", handler.GetAllCountries)
		auth.GET("/countries/:code", handler.GetCountryByCode)
		auth.GET("/cities", handler.GetAllCities)
		auth.GET("/cities/:id", handler.GetCityByID)
		auth.GET("/countries/:code/cities", handler.GetCitiesByCountryCode)
	}
}
