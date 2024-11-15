import { WorkflowResponse } from "@/types/api";
import { apiClient } from "@/lib/api-client";

export async function getWorkflows(): Promise<WorkflowResponse[]> {
  try {
    const response = await apiClient.get(`/workflows`);
    return response.data;
  } catch (error) {
    console.error("Error getting workflows:", error);
    throw error;
  }
}
