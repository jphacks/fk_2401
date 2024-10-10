export interface HouseResponse {
    id: number;
    name: string;
}

export interface JoinedDeviceResponse {
    id: number;
    name: string;
    houseID: number;
    setPoint?: number;
    duration?: number;
    climateData: string;
    unit: string;
}
