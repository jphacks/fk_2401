import {
  ClimateDataResponse,
  HouseResponse,
  JoinedDeviceResponse,
} from "@/types/api";

export function getHouses(): HouseResponse[] {
  const houses: HouseResponse[] = [
    { id: 1, name: "トマトハウス" },
    { id: 2, name: "ナスハウス" },
    { id: 3, name: "ほうれん草ハウス" },
  ];

  return houses;
}

export function getDevices(houseID: number): JoinedDeviceResponse[] {
  const tomatoDevices: JoinedDeviceResponse[] = [
    {
      id: 1,
      name: "ヒーター",
      house_id: 1,
      set_point: 25,
      duration: 3,
      climate_data: "気温",
      unit: "℃",
    },
    {
      id: 2,
      name: "ミスト",
      house_id: 1,
      set_point: 70,
      duration: 1,
      climate_data: "湿度",
      unit: "%",
    },
    {
      id: 3,
      name: "二酸化炭素供給装置",
      house_id: 1,
      set_point: 420,
      duration: 5,
      climate_data: "二酸化炭素量",
      unit: "ppm",
    },
    {
      id: 4,
      name: "窓開閉装置",
      house_id: 1,
      set_point: 30,
      duration: 1,
      climate_data: "気温",
      unit: "℃",
    },
  ];
  const nasuDevices: JoinedDeviceResponse[] = [
    {
      id: 5,
      name: "ヒーター",
      house_id: 2,
      set_point: 25,
      duration: undefined,
      climate_data: "気温",
      unit: "℃",
    },
    {
      id: 6,
      name: "ミスト",
      house_id: 2,
      set_point: 60,
      duration: undefined,
      climate_data: "湿度",
      unit: "%",
    },
    {
      id: 7,
      name: "二酸化炭素供給装置",
      house_id: 2,
      set_point: undefined,
      duration: undefined,
      climate_data: "二酸化炭素量",
      unit: "ppm",
    },
    {
      id: 8,
      name: "窓開閉装置",
      house_id: 2,
      set_point: undefined,
      duration: undefined,
      climate_data: "気温",
      unit: "℃",
    },
  ];
  const hourensouDevices: JoinedDeviceResponse[] = [
    {
      id: 9,
      name: "テストデバイス1",
      house_id: 3,
      set_point: 25,
      duration: 3,
      climate_data: "気温",
      unit: "℃",
    },
    {
      id: 10,
      name: "テストデバイス2",
      house_id: 3,
      set_point: 70,
      duration: 1,
      climate_data: "湿度",
      unit: "%",
    },
    {
      id: 11,
      name: "テストデバイス3",
      house_id: 3,
      set_point: 420,
      duration: 5,
      climate_data: "二酸化炭素量",
      unit: "ppm",
    },
    {
      id: 12,
      name: "テストデバイス4",
      house_id: 3,
      set_point: 30,
      duration: 1,
      climate_data: "気温",
      unit: "℃",
    },
    {
      id: 13,
      name: "テストデバイス5",
      house_id: 3,
      set_point: 25,
      duration: 3,
      climate_data: "気温",
      unit: "℃",
    },
    {
      id: 14,
      name: "テストデバイス6",
      house_id: 3,
      set_point: 70,
      duration: 1,
      climate_data: "湿度",
      unit: "%",
    },
    {
      id: 15,
      name: "テストデバイス7",
      house_id: 3,
      set_point: 420,
      duration: 5,
      climate_data: "二酸化炭素量",
      unit: "ppm",
    },
    {
      id: 16,
      name: "テストデバイス8",
      house_id: 3,
      set_point: 30,
      duration: 1,
      climate_data: "気温",
      unit: "℃",
    },
  ];

  const deviceMap: Map<number, JoinedDeviceResponse[]> = new Map();

  deviceMap.set(1, tomatoDevices);
  deviceMap.set(2, nasuDevices);
  deviceMap.set(3, hourensouDevices);

  const result: JoinedDeviceResponse[] = deviceMap.get(houseID) || [];

  return result;
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
