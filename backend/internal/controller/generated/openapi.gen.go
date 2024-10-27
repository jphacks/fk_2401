// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package generated

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ClimateDatasResponse defines model for ClimateDatasResponse.
type ClimateDatasResponse struct {
	ClimateData *string `json:"climate_data,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Unit        *string `json:"unit,omitempty"`
}

// DeviceRequest defines model for DeviceRequest.
type DeviceRequest struct {
	ClimateDataId *int     `json:"climate_data_id,omitempty"`
	DeviceName    *string  `json:"device_name,omitempty"`
	Duration      *int     `json:"duration"`
	SetPoint      *float64 `json:"set_point,omitempty"`
	Unit          *string  `json:"unit,omitempty"`
}

// DeviceResponse defines model for DeviceResponse.
type DeviceResponse struct {
	ClimateData *string  `json:"climate_data,omitempty"`
	DeviceName  *string  `json:"device_name,omitempty"`
	Duration    *int     `json:"duration,omitempty"`
	Id          *int     `json:"id,omitempty"`
	SetPoint    *float64 `json:"set_point,omitempty"`
	Unit        *string  `json:"unit,omitempty"`
}

// HousesRequest defines model for HousesRequest.
type HousesRequest = string

// HousesResponse defines model for HousesResponse.
type HousesResponse struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// HouseId defines model for house-id.
type HouseId = int

// CreateHouseJSONRequestBody defines body for CreateHouse for application/json ContentType.
type CreateHouseJSONRequestBody = HousesRequest

// CreateDeviceJSONRequestBody defines body for CreateDevice for application/json ContentType.
type CreateDeviceJSONRequestBody = DeviceRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get climate data list
	// (GET /climate-datas)
	GetClimateData(c *gin.Context)
	// Get houses list
	// (GET /houses)
	GetHouses(c *gin.Context)
	// Create a house
	// (POST /houses)
	CreateHouse(c *gin.Context)
	// Get devices list
	// (GET /houses/{house-id})
	GetDevice(c *gin.Context, houseId HouseId)
	// Create device
	// (POST /houses/{house-id}/devices)
	CreateDevice(c *gin.Context, houseId HouseId)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetClimateData operation middleware
func (siw *ServerInterfaceWrapper) GetClimateData(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetClimateData(c)
}

// GetHouses operation middleware
func (siw *ServerInterfaceWrapper) GetHouses(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetHouses(c)
}

// CreateHouse operation middleware
func (siw *ServerInterfaceWrapper) CreateHouse(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateHouse(c)
}

// GetDevice operation middleware
func (siw *ServerInterfaceWrapper) GetDevice(c *gin.Context) {

	var err error

	// ------------- Path parameter "house-id" -------------
	var houseId HouseId

	err = runtime.BindStyledParameterWithOptions("simple", "house-id", c.Param("house-id"), &houseId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter house-id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetDevice(c, houseId)
}

// CreateDevice operation middleware
func (siw *ServerInterfaceWrapper) CreateDevice(c *gin.Context) {

	var err error

	// ------------- Path parameter "house-id" -------------
	var houseId HouseId

	err = runtime.BindStyledParameterWithOptions("simple", "house-id", c.Param("house-id"), &houseId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter house-id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateDevice(c, houseId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/climate-datas", wrapper.GetClimateData)
	router.GET(options.BaseURL+"/houses", wrapper.GetHouses)
	router.POST(options.BaseURL+"/houses", wrapper.CreateHouse)
	router.GET(options.BaseURL+"/houses/:house-id", wrapper.GetDevice)
	router.POST(options.BaseURL+"/houses/:house-id/devices", wrapper.CreateDevice)
}