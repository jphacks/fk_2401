import {
  ClimateDataResponse,
  DeviceResponse,
  OperationResponse,
} from "@/types/api";

export function getDevices(): DeviceResponse[] {
  const devices: DeviceResponse[] = [
    { id: 1, device_name: "ヒーター" },
    { id: 2, device_name: "ミスト" },
    { id: 3, device_name: "二酸化炭素供給装置" },
    { id: 4, device_name: "照明" },
  ];

  return devices;
}

export function getClimateDatas(): ClimateDataResponse[] {
  const climateDatas: ClimateDataResponse[] = [
    { id: 1, climate_data: "気温", unit: "℃" },
    { id: 2, climate_data: "湿度", unit: "%" },
    { id: 3, climate_data: "二酸化炭素量", unit: "ppm" },
    { id: 4, climate_data: "照度", unit: "lx" },
  ];

  return climateDatas;
}

export function getOperations(): OperationResponse[] {
  const operations: OperationResponse[] = [
    { id: 1, device_id: 1, name: "加温", rly_on: 1 },
    { id: 2, device_id: 1, name: "送風", rly_on: 1 },
    { id: 3, device_id: 1, name: "加温＆送風", rly_on: 1 },
    { id: 4, device_id: 2, name: "加湿 弱", rly_on: 1 },
    { id: 5, device_id: 2, name: "加湿 強", rly_on: 1 },
    { id: 6, device_id: 3, name: "二酸化炭素供給 弱", rly_on: 1 },
    { id: 7, device_id: 3, name: "二酸化炭素供給 強", rly_on: 1 },
    { id: 8, device_id: 4, name: "点灯 弱", rly_on: 1 },
    { id: 9, device_id: 4, name: "点灯 強", rly_on: 1 },
    { id: 10, device_id: 4, name: "消灯", rly_on: 1 },
  ];

  return operations;
}
