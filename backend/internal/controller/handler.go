package controller

import (
	"log"
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
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

type CreateHouseController struct {
	Name string `json:"name" binding:"required,min=1,max=12,alphanum"`
}

func (h Handler) CreateHouse(c *gin.Context) {
	var json CreateHouseController
	if err := c.BindJSON(&json); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	name := json.Name
	id, err := h.houseService.CreateHouse(name)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
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

func (h Handler) CreateDevice(c *gin.Context, houseId int) {
	var json CreateDeviceController
	if err := c.BindJSON(&json); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	device := domain.Device{
		HouseID:       houseId,
		ClimateDataID: json.ClimateDataID,
		DeviceName:    json.DeviceName,
		SetPoint:      json.SetPoint,
		Duration:      json.Duration,
	}
	id, err := h.deviceService.CreateDevice(device)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	c.JSON(http.StatusOK, id)
}
