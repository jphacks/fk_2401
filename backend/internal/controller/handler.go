package controller

import (
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/generated"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service"
	"github.com/gin-gonic/gin"
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

func (h Handler) CreateHouse(c *gin.Context) {

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
