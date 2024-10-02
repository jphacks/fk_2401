import { HouseResponse, JoinedDeviceResponse } from "../types/api";

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
        { id: 1, name: "ヒーター", houseID: 1, setPoint: 25, duration: 3, climateData: "気温", unit: "℃"},
        { id: 2, name: "ミスト", houseID: 1, setPoint: 70, duration: 1, climateData: "湿度", unit: "%"},
        { id: 3, name: "二酸化炭素供給装置", houseID: 1, setPoint: 420, duration: 5, climateData: "二酸化炭素量", unit: "ppm"},
        { id: 4, name: "窓開閉装置", houseID: 1, setPoint: 30, duration: 1, climateData: "気温", unit: "℃"},
    ];
    const nasuDevices: JoinedDeviceResponse[] = [
        { id: 5, name: "ヒーター", houseID: 2, setPoint: 25, duration: 3, climateData: "気温", unit: "℃"},
        { id: 6, name: "ミスト", houseID: 2, setPoint: 70, duration: 1, climateData: "湿度", unit: "%"},
        { id: 7, name: "二酸化炭素供給装置", houseID: 2, setPoint: 420, duration: 5, climateData: "二酸化炭素量", unit: "ppm"},
        { id: 8, name: "窓開閉装置", houseID: 2, setPoint: 30, duration: 1, climateData: "気温", unit: "℃"},
    ];
    const hourensouDevices: JoinedDeviceResponse[] = [
        { id: 9, name: "テストデバイス1", houseID: 3, setPoint: 25, duration: 3, climateData: "気温", unit: "℃"},
        { id: 10, name: "テストデバイス2", houseID: 3, setPoint: 70, duration: 1, climateData: "湿度", unit: "%"},
        { id: 11, name: "テストデバイス3", houseID: 3, setPoint: 420, duration: 5, climateData: "二酸化炭素量", unit: "ppm"},
        { id: 12, name: "テストデバイス4", houseID: 3, setPoint: 30, duration: 1, climateData: "気温", unit: "℃"},
        { id: 13, name: "テストデバイス5", houseID: 3, setPoint: 25, duration: 3, climateData: "気温", unit: "℃"},
        { id: 14, name: "テストデバイス6", houseID: 3, setPoint: 70, duration: 1, climateData: "湿度", unit: "%"},
        { id: 15, name: "テストデバイス7", houseID: 3, setPoint: 420, duration: 5, climateData: "二酸化炭素量", unit: "ppm"},
        { id: 16, name: "テストデバイス8", houseID: 3, setPoint: 30, duration: 1, climateData: "気温", unit: "℃"},
    ];

    const deviceMap: Map<number, JoinedDeviceResponse[]> = new Map();

    deviceMap.set(1, tomatoDevices);
    deviceMap.set(2, nasuDevices);
    deviceMap.set(3, hourensouDevices);

    const result: JoinedDeviceResponse[] = deviceMap.get(houseID) || [];

    return result
} 
