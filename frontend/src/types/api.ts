export interface HouseResponse {
  id: number;
  name: string;
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
