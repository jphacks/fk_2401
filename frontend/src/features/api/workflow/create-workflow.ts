import { WorkflowWithUIRequest } from "@/types/api";
import { apiClient } from "@/lib/api-client";

// 戻り値は作成したレコードのID
export async function createWorkflowWithUI(
  data: WorkflowWithUIRequest
): Promise<number> {
  try {
    const response = await apiClient.post(`/workflows-with-ui`, data);
    return response.data;
  } catch (error) {
    console.error("Error creating device:", error);
    throw error;
  }
}
