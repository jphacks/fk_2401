package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
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

}

// func (h Handler) GetDevice(c *gin.Context, houseId int) {
// 	devices, err := h.deviceService.GetJoinedDevices(houseId)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	devicesRes := make([]*DeviceResponse, len(devices))
// 	for i, v := range devices {
// 		devicesRes[i] = NewDeviceResponse(
// 			v.ID,
// 			v.HouseID,
// 			*v.DeviceName,
// 			*v.SetPoint,
// 			*v.Duration,
// 			v.ClimateData,
// 			v.Unit,
// 		)
// 	}

// 	c.JSON(http.StatusOK, devicesRes)
// }

func (h Handler) CreateDevice(c *gin.Context, houseId int) {
}

// func (h Handler) CreateDevice(c *gin.Context, houseId int) {
// 	var json CreateDeviceController
// 	if err := c.BindJSON(&json); err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
// 		return
// 	}
// 	device := domain.Device{
// 		HouseID:       houseId,
// 		ClimateDataID: json.ClimateDataID,
// 		DeviceName:    &json.DeviceName,
// 		SetPoint:      &json.SetPoint,
// 		Duration:      &json.Duration,
// 	}
// 	id, err := h.deviceService.CreateDevice(device)

// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, id)
// }

func (h Handler) GetClimateData(c *gin.Context) {
	climateData, err := h.climateDataService.GetAllClimateData()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalSeverError"})
		return
	}

	climateDataRes := make([]*ClimateDataResponse, len(climateData))
	for i, v := range climateData {
		climateDataRes[i] = NewClimateDataResponse(v.ID, v.ClimateData, v.Unit)
	}

	c.JSON(http.StatusOK, climateDataRes)
}

func (h Handler) CreateAndBuildTimeSchedule(c *gin.Context) {
	var json TimeScheduleRequest
	if err := c.BindJSON(&json); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	setStartTime := json.TimeSchedule[0].StartTime
	setEndTime := json.TimeSchedule[0].EndTime
	StartTime := strings.Split(setStartTime, ":")
	EndTime := strings.Split(setEndTime, ":")
	sthr, err := strconv.Atoi(StartTime[0])
	if err != nil {
		fmt.Println("Error:", err)
	}
	stmn, err := strconv.Atoi(StartTime[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	edhr, err := strconv.Atoi(EndTime[0])
	if err != nil {
		fmt.Println("Error:", err)
	}
	edmn, err := strconv.Atoi(EndTime[0])
	if err != nil {
		fmt.Println("Error:", err)
	}
	timeSchedules := json.TimeSchedule
	var deviceIDs []int
	for _, v := range timeSchedules {
		workFlows := v.Workflows
		for _, w := range workFlows {
			deviceIDs = append(deviceIDs, w.DeviceId)
		}
	}
	var rly []int
	for _, v := range deviceIDs {
		device, err := h.deviceService.GetDeviceFromID(v)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		rly = append(rly, *device.Rly)
	}

	unique := make(map[int]bool) // ユニークな要素を記録
	var result []int             // 最終的な結果

	for _, num := range rly {
		// 0~7の範囲内かつまだ登録されていない場合に追加
		if num >= 0 && num <= 7 && !unique[num] {
			unique[num] = true
			result = append(result, num)
		}
	}
	demoData := service.DemoData{
		IpAddr: json.IpAddress,
		StHr:   sthr,
		StMn:   stmn,
		EdHr:   edhr,
		EdMn:   edmn,
		InMn:   1,
		DuMn:   1,
		Rly:    result,
	}
	id, err := service.BuildDemoM304(&demoData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h Handler) GetWorkflows(c *gin.Context) {
	ctx := context.Background()

	workFlows, err := mysqlc.GetAllWorkflows(ctx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalSeverError"})
		return
	}

	workFlowsRes := make([]*WorkflowResponse, len(workFlows))
	for i, v := range workFlows {
		workFlowsRes[i] = NewWorkflowResponse(int(v.ID), v.Name)
	}

	c.JSON(http.StatusOK, workFlowsRes)
}

func (h Handler) GetWorkflowsWithUI(c *gin.Context) {
	ctx := context.Background()

	workFlows, err := mysqlc.GetAllWorkflows(ctx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalSeverError"})
		return
	}

	workFlowsUIRes := make([]*WorkflowWithUIResponse, len(workFlows))
	for i, v := range workFlows {
		edges, err := mysqlc.GetEdgesFromWorkflow(ctx, int32(v.ID))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalSeverError"})
			return
		}
		edgeRes := make([]EdgeResponse, len(edges))
		for j, w := range edges {
			edgeRes[j] = *NewEdgeResponse(w.Id, w.SourceNodeID, w.TargetNodeID, int(w.WorkflowId))
		}
		nodes, err := mysqlc.GetNodesFromWorkflow(ctx, int32(v.ID))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalSeverError"})
			return
		}
		nodeRes := make([]NodeResponse, len(nodes))
		for j, w := range nodes {
			nodeRes[j] = *NewNodeResponse(w.Data, w.Id, w.PositionX, w.PositionY, w.NodeType, int(w.WorkflowId), w.WorkflowNodeID)
		}
		workFlowUI := NewWorkFlowUIResponse(edgeRes, nodeRes)
		workFlowsUIRes[i] = NewWorkFlowWithUIResponse(v, *workFlowUI)
	}

	c.JSON(http.StatusOK, workFlowsUIRes)
}

func (h Handler) CreateWorkflowWithUI(c *gin.Context) {
	ctx := context.Background()

	var jsonWorkflow WorkflowWithUIRequest
	if err := c.BindJSON(&jsonWorkflow); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	workFlowName := jsonWorkflow.Workflow
	workFlowID, err := mysqlc.CreateWorkflow(ctx, workFlowName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	workFlowUIEdges := jsonWorkflow.WorkflowUI.Edges
	workFlowUINodes := jsonWorkflow.WorkflowUI.Nodes
	for _, v := range workFlowUIEdges {
		edgeReq := mysqlc.CreateEdgeParams{
			WorkflowID:   int32(workFlowID),
			SourceNodeID: v.SourceNodeId,
			TargetNodeID: v.TargetNodeId,
		}
		_, err := mysqlc.CreateEdge(ctx, edgeReq)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
	}
	for _, v := range workFlowUINodes {
		rawData, err := json.Marshal(v.Data)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
		nodeReq := mysqlc.CreateNodeParams{
			WorkflowID: int32(workFlowID),
			Type:       v.Type,
			Data:       json.RawMessage(rawData),
			PositionX:  float64(v.PositionX),
			PositionY:  float64(v.PositionY),
		}
		_, err = mysqlc.CreateNode(ctx, nodeReq)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
	}
	c.JSON(http.StatusOK, workFlowID)
}

func (h Handler) CreateWorkflow(c *gin.Context)
