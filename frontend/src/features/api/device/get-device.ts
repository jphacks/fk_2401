import { JoinedDeviceResponse } from "../../../types/api";
import { api } from "../../../lib/api-client";

export async function getDevices(
  houseID: number
): Promise<JoinedDeviceResponse[]> {
  try {
    const response = await api.get(`/house/${houseID}/devices`);
    return response.data;
  } catch (error) {
    console.error("Error getting device:", error);
    throw error;
  }
}
