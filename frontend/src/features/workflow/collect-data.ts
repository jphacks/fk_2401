import { ConditionOperation, WorkflowData } from "@/types/workflow";
import { Edge, Node } from "@xyflow/react";
import { SelectDeviceNodeData } from "@/components/workflow-control/custom-nodes/select-device";
import { ConditionNodeData } from "@/components/workflow-control/custom-nodes/condition";
import { DeviceOperationNodeData } from "@/components/workflow-control/custom-nodes/device-operation";

export function collectData(nodes: Node[], edges: Edge[]): WorkflowData | null {
  const selectDeviceConditionEdges: Edge[] = [];
  const conditionOperationEdges: Edge[] = [];

  // デバイスが選択されているか確認する
  const selectedDeviceNode: Node | undefined = nodes.find(
    (node) => node.type === "select_device"
  );
  let selectDeviceNodeData: SelectDeviceNodeData | undefined;
  if (selectedDeviceNode && selectedDeviceNode.data) {
    selectDeviceNodeData = selectedDeviceNode.data as SelectDeviceNodeData;
  }

  if (!selectDeviceNodeData?.device_id) {
    return null;
  }

  // デバイス選択-条件、条件-操作でエッジを分類する
  edges.forEach((edge) => {
    const source: string = edge.source;
    const target: string = edge.target;

    if (source.includes("select_device") && target.includes("condition")) {
      selectDeviceConditionEdges.push(edge);
    } else if (
      source.includes("condition") &&
      target.includes("device_operation")
    ) {
      conditionOperationEdges.push(edge);
    }
  });

  // 条件・操作ノードのつながりを見てデータを作成する
  const conditionOperations: ConditionOperation[] = [];
  conditionOperationEdges.forEach((conOpeEdge) => {
    const conOpeSource: string = conOpeEdge.source;
    const conOpeTarget: string = conOpeEdge.target;

    const sourceNode: Node | undefined = nodes.find(
      (node) => node.id === conOpeSource
    );
    const targetNode: Node | undefined = nodes.find(
      (node) => node.id === conOpeTarget
    );

    let conditionNodeData: ConditionNodeData | undefined;
    if (sourceNode && sourceNode.data) {
      conditionNodeData = sourceNode.data as ConditionNodeData;
    }

    let deviceOperationNodeData: DeviceOperationNodeData | undefined;
    if (targetNode && targetNode.data) {
      deviceOperationNodeData = targetNode.data as DeviceOperationNodeData;
    }

    if (
      !conditionNodeData?.condition ||
      !deviceOperationNodeData?.operationID
    ) {
      return;
    }

    const conditionOperation: ConditionOperation = {
      condition: conditionNodeData.condition,
      operation_id: deviceOperationNodeData.operationID,
    };

    const isConnectedWithSelectDevice: boolean =
      selectDeviceConditionEdges.some(
        (edge) =>
          edge.source.includes("select_device") && edge.target == conOpeSource
      );

    if (isConnectedWithSelectDevice) {
      conditionOperations.push(conditionOperation);
    }
  });

  const workflow: WorkflowData = {
    device_id: selectDeviceNodeData.device_id,
    condition_operations: conditionOperations,
  };

  return workflow;
}
