import { WorkflowWithUIResponse } from "@/types/api";
import { apiClient } from "@/lib/api-client";

export async function getWorkflows(): Promise<WorkflowWithUIResponse[]> {
  try {
    const response = await apiClient.get(`/workflows-with-ui`);
    return response.data;
  } catch (error) {
    console.error("Error getting workflows:", error);
    throw error;
  }
}
