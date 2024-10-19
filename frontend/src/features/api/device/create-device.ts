import { CreateDeviceRequest } from "@/types/api";
import { api } from "@/lib/api-client";

// 戻り値は作成したレコードのID
export async function createDevice(data: CreateDeviceRequest): Promise<number> {
  try {
    const response = await api.post("/devices", data);
    return response.data;
  } catch (error) {
    console.error("Error creating device:", error);
    throw error;
  }
}
