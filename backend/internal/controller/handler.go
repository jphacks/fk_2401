package controller

import (
	"log"
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	deviceService      *service.DeviceService
	houseService       *service.HouseService
	climateDataService *service.ClimateDataService
}

func NewHandler(ds *service.DeviceService, hs *service.HouseService, cds *service.ClimateDataService) *Handler {
	return &Handler{
		deviceService:      ds,
		houseService:       hs,
		climateDataService: cds,
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

func (h Handler) GetClimateData(c *gin.Context) {
	climateData, err := h.climateDataService.GetClimateData()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalSeverError"})
	}

	c.JSON(http.StatusOK, climateData)
}
