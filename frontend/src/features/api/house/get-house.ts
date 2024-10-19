import { HouseResponse } from "@/types/api";
import { apiClient } from "@/lib/api-client";

export async function getHouses(): Promise<HouseResponse[]> {
  try {
    const response = await apiClient.get(`/houses`);
    return response.data;
  } catch (error) {
    console.error("Error getting houses:", error);
    throw error;
  }
}
