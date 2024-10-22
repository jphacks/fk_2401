package controller

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/generated"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/db/mysql"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository/mocks"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initHandler() (*Handler, error) {
	db, err := mysql.ConnectDB()
	if err != nil {
		return nil, err
	}

	query := mysqlc.New(db)

	dr := repository.NewDeviceRepository(query)
	hr := repository.NewHouseRepository(query)
	cdr := repository.NewClimateDataRepository(query)

	ds := service.NewDeviceService(dr)
	hs := service.NewHouseService(hr)
	cds := service.NewClimateDataService(cdr)

	h := NewHandler(ds, hs, cds)

	return h, nil
}

func initMockHandler() (*Handler, error) {
	dr := mocks.NewMockDeviceRepository()
	hr := mocks.NewMockHouseRepository()
	cdr := mocks.NewMockClimateDataRepository()

	ds := service.NewDeviceService(dr)
	hs := service.NewHouseService(hr)
	cds := service.NewClimateDataService(cdr)

	h := NewHandler(ds, hs, cds)

	return h, nil
}

func StartServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	h, err := initHandler()
	if err != nil {
		log.Fatalf("Failed to initialize handler: %v", err)
	}

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
