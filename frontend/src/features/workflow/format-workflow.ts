import {
  EdgeRequest,
  NodeRequest,
  WorkflowRequest,
  WorkflowUIRequest,
  WorkflowWithUIRequest,
} from "@/types/api";
import { Edge, Node } from "@xyflow/react";

export function formatWorkflowRequest(
  nodes: Node[],
  edges: Edge[],
  workflowName: string
): WorkflowWithUIRequest {
  const nodesReq: NodeRequest[] = [];
  nodes.forEach((node) => {
    const nodeReq: NodeRequest = {
      workflow_node_id: node.id,
      node_type: node.type || "",
      data: node.data,
      position_x: node.position.x,
      position_y: node.position.y,
    };

    nodesReq.push(nodeReq);
  });

  const edgesReq: EdgeRequest[] = [];
  edges.forEach((edge) => {
    const edgeReq: EdgeRequest = {
      source_node_id: edge.source,
      target_node_id: edge.target,
    };

    edgesReq.push(edgeReq);
  });

  const workflowReq: WorkflowRequest = workflowName;
  const workflowUIReq: WorkflowUIRequest = {
    nodes: nodesReq,
    edges: edgesReq,
  };

  const workflowWithUIReq: WorkflowWithUIRequest = {
    workflow: workflowReq,
    workflow_ui: workflowUIReq,
  };

  return workflowWithUIReq;
}
