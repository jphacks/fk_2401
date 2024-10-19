import { CreateDeviceRequest } from "@/types/api";
import { apiClient } from "@/lib/api-client";

// 戻り値は作成したレコードのID
export async function createDevice(
  houseID: number,
  data: CreateDeviceRequest
): Promise<number> {
  try {
    const response = await apiClient.post(`houses/${houseID}/devices`, data);
    return response.data;
  } catch (error) {
    console.error("Error creating device:", error);
    throw error;
  }
}
