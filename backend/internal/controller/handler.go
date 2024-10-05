package controller

import (
	"fmt"
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/generated"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	deviceService *service.DeviceService
	houseService  *service.HouseService
}

func NewHandler(ds *service.DeviceService, hs *service.HouseService) *Handler {
	return &Handler{
		deviceService: ds,
		houseService:  hs,
	}
}

func (h Handler) GetHouses(c *gin.Context) {
	houses, err := h.houseService.GetHouses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, houses)
}

type CreateHouseController struct {
	Name string `json:"name" binding:"required,min=1,max=10,alphanum"`
}

func (h Handler) CreateHouse(c *gin.Context) {
	var json CreateHouseController
	if err := c.BindJSON(&json); err != nil {

		var errorMessages []string
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				errorMessages = append(errorMessages, fmt.Sprintf("Field %s is %s", e.Field(), e.ActualTag()))
			}
		} else {
			errorMessages = append(errorMessages, err.Error())
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	name := json.Name
	id, err := h.houseService.CreateHouse(name)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, id)
}

func (h Handler) GetDevice(c *gin.Context, houseId int) {
	devices, err := h.deviceService.GetDevices(houseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, devices)
}

func (h Handler) CreateDevice(c *gin.Context, houseId generated.HouseId) {

}
