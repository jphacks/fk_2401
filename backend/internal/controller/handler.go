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
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	housesRes := make([]*HouseResponse, len(houses))
	for i, v := range houses {
		housesRes[i] = NewHouseResponse(v.ID, v.Name)
	}

	c.JSON(http.StatusOK, housesRes)
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
	devices, err := h.deviceService.GetJoinedDevices(houseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	devicesRes := make([]*DeviceResponse, len(devices))
	for i, v := range devices {
		devicesRes[i] = NewDeviceResponse(
			v.ID,
			v.HouseID,
			v.DeviceName,
			v.SetPoint,
			v.Duration,
			v.ClimateData,
			v.Unit,
		)
	}

	c.JSON(http.StatusOK, devicesRes)
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
	climateData, err := h.climateDataService.GetAllClimateData()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalSeverError"})
	}

	climateDataRes := make([]*ClimateDataResponse, len(climateData))
	for i, v := range climateData {
		climateDataRes[i] = NewClimateDataResponse(v.ID, v.ClimateData, v.Unit)
	}

	c.JSON(http.StatusOK, climateDataRes)
}
