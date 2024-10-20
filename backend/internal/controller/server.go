package controller

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/generated"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository/mocks"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	dr := mocks.NewMockDeviceRepository()
	hr := mocks.NewMockHouseRepository()
	cdr := mocks.NewMockClimateDataRepository()

	ds := service.NewDeviceService(dr)
	hs := service.NewHouseService(hr)
	cds := service.NewClimateDataService(cdr)

	h := NewHandler(ds, hs, cds)
	generated.RegisterHandlers(r, h)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
