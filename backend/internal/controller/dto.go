package controller

type CreateHouseController struct {
	Name string `json:"name" binding:"required,min=1,max=12,alphanum"`
}

type CreateDeviceController struct {
	ClimateDataID int     `json:"climate_data_id" binding:"required,number"`
	DeviceName    string  `json:"device_name" binding:"required,min=1,max=12"`
	SetPoint      float64 `json:"set_point" binding:"number"`
	Duration      int     `json:"duration" binding:"number"`
}

type HouseResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DeviceResponse struct {
	ID          int     `json:"id,omitempty"`
	HouseID     int     `json:"house_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	SetPoint    float64 `json:"set_point,omitempty"`
	Duration    int     `json:"duration,omitempty"`
	ClimateData string  `json:"climate_data,omitempty"`
	Unit        string  `json:"unit,omitempty"`
}

type ClimateDataResponse struct {
	ID          int    `json:"id,omitempty"`
	ClimateData string `json:"climate_data,omitempty"`
	Unit        string `json:"unit,omitempty"`
}

// EdgeRequest defines model for EdgeRequest.
type EdgeRequest struct {
	SourceNodeId string `json:"source_node_id,omitempty"`
	TargetNodeId string `json:"target_node_id,omitempty"`
}

// EdgeResponse defines model for EdgeResponse.
type EdgeResponse struct {
	Id           int    `json:"id,omitempty"`
	SourceNodeId string `json:"source_node_id,omitempty"`
	TargetNodeId string `json:"target_node_id,omitempty"`
	WorkflowId   int    `json:"workflow_id,omitempty"`
}

// NodeRequest defines model for NodeRequest.
type NodeRequest struct {
	Data           map[string]interface{} `json:"data,omitempty"`
	PositionX      float32                `json:"position_x,omitempty"`
	PositionY      float32                `json:"position_y,omitempty"`
	Type           string                 `json:"type,omitempty"`
	WorkflowNodeId string                 `json:"workflow_node_id,omitempty"`
}

// NodeResponse defines model for NodeResponse.
type NodeResponse struct {
	Data           map[string]interface{} `json:"data,omitempty"`
	Id             int                    `json:"id,omitempty"`
	PositionX      float32                `json:"position_x,omitempty"`
	PositionY      float32                `json:"position_y,omitempty"`
	NodeType       string                 `json:"type,omitempty"`
	WorkflowId     int                    `json:"workflow_id,omitempty"`
	WorkflowNodeId string                 `json:"workflow_node_id,omitempty"`
}

// TimeSchedule defines model for TimeSchedule.
type TimeSchedule struct {
	// EndTime 終了時刻（HH:mm形式）
	EndTime string `json:"end_time,omitempty"`

	// StartTime 開始時刻（HH:mm形式）
	StartTime string            `json:"start_time,omitempty"`
	Workflows []WorkflowRequest `json:"workflows,omitempty"`
}

// TimeScheduleRequest defines model for TimeScheduleRequest.
type TimeScheduleRequest struct {
	IpAddress    string         `json:"ip_address,omitempty"`
	TimeSchedule []TimeSchedule `json:"time_schedule,omitempty"`
}

// WorkflowRequest defines model for WorkflowRequest.
type WorkflowRequest struct {
	DeviceId int    `json:"device_id,omitempty"`
	Name     string `json:"name,omitempty"`
}

// WorkflowResponse defines model for WorkflowResponse.
type WorkflowResponse struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// WorkflowUIRequest defines model for WorkflowUIRequest.
type WorkflowUIRequest struct {
	Edges []EdgeRequest `json:"edges,omitempty"`
	Nodes []NodeRequest `json:"nodes,omitempty"`
}

// WorkflowUIResponse defines model for WorkflowUIResponse.
type WorkflowUIResponse struct {
	Edges []EdgeResponse `json:"edges,omitempty"`
	Nodes []NodeResponse `json:"nodes,omitempty"`
}

// WorkflowWithUIRequest defines model for WorkflowWithUIRequest.
type WorkflowWithUIRequest struct {
	Workflow   WorkflowRequest   `json:"workflow,omitempty"`
	WorkflowUI WorkflowUIRequest `json:"workflow_ui,omitempty"`
}

// WorkflowWithUIResponse defines model for WorkflowWithUIResponse.
type WorkflowWithUIResponse struct {
	Workflow   WorkflowResponse   `json:"workflow,omitempty"`
	WorkflowUI WorkflowUIResponse `json:"workflow_ui,omitempty"`
}

func NewHouseResponse(id int, name string) *HouseResponse {
	return &HouseResponse{
		ID:   id,
		Name: name,
	}
}

func NewDeviceResponse(
	id int,
	houseID int,
	deviceName string,
	setPoint float64,
	duration int,
	climateData string,
	unit string,
) *DeviceResponse {
	return &DeviceResponse{
		ID:          id,
		HouseID:     houseID,
		Name:        deviceName,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: climateData,
		Unit:        unit,
	}
}

func NewClimateDataResponse(id int, climateData, unit string) *ClimateDataResponse {
	return &ClimateDataResponse{
		ID:          id,
		ClimateData: climateData,
		Unit:        unit,
	}
}

func NewWorkflowResponse(id int, name string) *WorkflowResponse {
	return &WorkflowResponse{
		Id:   id,
		Name: name,
	}
}

func NewWorkFlowWithUIResponse(workFlow WorkflowResponse, workFlowUI WorkflowUIResponse) *WorkflowWithUIResponse {
	return &WorkflowWithUIResponse{
		Workflow:   workFlow,
		WorkflowUI: workFlowUI,
	}
}

func NewWorkFlowUIResponse(edges []EdgeResponse, nodes []NodeResponse) *WorkflowUIResponse {
	return &WorkflowUIResponse{
		Edges: edges,
		Nodes: nodes,
	}
}

func NewEdgeResponse(id int, sourceNodeID string, targetNodeID string, workFlowID int) *EdgeResponse {
	return &EdgeResponse{
		Id:           id,
		SourceNodeId: sourceNodeID,
		TargetNodeId: targetNodeID,
		WorkflowId:   workFlowID,
	}
}

func NewNodeResponse(
	data map[string]interface{},
	id int, positionX float32,
	positionY float32,
	nodeType string,
	workFlowID int,
	workFlowNodeID string,
) *NodeResponse {
	return &NodeResponse{
		Data:           data,
		Id:             id,
		PositionX:      positionX,
		PositionY:      positionY,
		NodeType:       nodeType,
		WorkflowId:     workFlowID,
		WorkflowNodeId: workFlowNodeID,
	}
}
