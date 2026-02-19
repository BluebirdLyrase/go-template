package handlers

import (
	"my-api/internal/modules/option/address/service"
	"my-api/internal/shared/errors"
	"my-api/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	service *service.AddressService
}

func NewAddressHandler(service *service.AddressService) *AddressHandler {
	return &AddressHandler{service: service}
}

func (h *AddressHandler) GetAllCountries(c *gin.Context) {
	countries, err := h.service.GetAllCountries()
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, countries)
}

func (h *AddressHandler) GetCountryByCode(c *gin.Context) {
	code := c.Param("code")

	country, err := h.service.GetCountryByCode(code)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if country == nil {
		c.JSON(404, gin.H{"error": "country not found"})
		return
	}

	c.JSON(200, country)
}

func (h *AddressHandler) GetAllCities(c *gin.Context) {
	cities, err := h.service.GetAllCities()
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, cities)
}

func (h *AddressHandler) GetCityByID(c *gin.Context) {
	id, err := utils.ParseUintParam(c, "id")
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid city id"})
		return
	}

	city, err := h.service.GetCityByID(id)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if city == nil {
		c.JSON(404, gin.H{"error": "city not found"})
		return
	}

	c.JSON(200, city)
}

func (h *AddressHandler) GetCitiesByCountryCode(c *gin.Context) {
	code := c.Param("code")

	cities, err := h.service.GetCitiesByCountryCode(code)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, cities)
}
