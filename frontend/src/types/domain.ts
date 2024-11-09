export interface ClimateData {
  id: number;
  climate_data: string;
  unit: string;
}

export interface device {
  id: number;
  house_id: number;
  climate_data_id: number;
  device_name: string;
  set_point: number;
  duration: number;
}

export interface house {
  id: number;
  name: string;
}
