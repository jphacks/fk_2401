package controller

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/generated"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository/mocks"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func StartServer() {
	r := gin.Default()
	dr := mocks.NewMockDeviceRepository()
	hr := mocks.NewMockHouseRepository()
	ds := service.NewDeviceService(dr)
	hs := service.NewHouseService(hr)
	h := NewHandler(ds, hs)
	generated.RegisterHandlers(r, h)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.Run() // listen and serve on 0.0.0.0:8080
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
