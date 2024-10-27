import { ClimateDataResponse } from "@/types/api";
import { apiClient } from "@/lib/api-client";

export async function getClimateDatas(): Promise<ClimateDataResponse[]> {
  try {
    const response = await apiClient.get("/climate-datas");
    return response.data;
  } catch (error) {
    console.error("Error getting climate datas:", error);
    throw error;
  }
}
