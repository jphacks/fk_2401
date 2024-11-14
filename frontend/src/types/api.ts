export interface HouseResponse {
  id: number;
  name: string;
}

export interface DeviceResponse {
  id: number;
  device_name: string;
}

export interface JoinedDeviceResponse {
  id: number;
  name: string;
  house_id: number;
  set_point?: number;
  duration?: number;
  climate_data: string;
  unit: string;
}

export interface ClimateDataResponse {
  id: number;
  climate_data: string;
  unit: string;
}

export interface CreateDeviceRequest {
  device_name: string;
  climate_data_id: number;
  set_point?: number;
  duration?: number;
}

export interface CreateHouseRequest {
  name: string;
}

export interface WorkflowRequest {
  device_id: number;
  house_id: number;
  climate_data_id: number;
  uecs_device_id: number;
  valid: boolean;
  set_point: number;
  duration: number;
  operator: number;
}

export interface OperationResponse {
  id: number;
  device_id: number;
  name: string;
  rly_on: number;
}

export interface Node {
  id: number;
  workflow_id: number;
  workflow_node_id: string;
  type: string;
  data: object;
  position_x: number;
  position_y: number;
}

export interface Edge {
  id: number;
  workflow_id: number;
  source_node_id: string;
  target_node_id: string;
}

export interface WorkflowResponse {
  id: number;
  name: string;
  nodes: Node[];
  edges: Edge[];
}
