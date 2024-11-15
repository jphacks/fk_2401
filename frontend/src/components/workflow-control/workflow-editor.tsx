import { Box } from "@mui/material";
import {
  Node,
  Edge,
  Connection,
  useNodesState,
  useEdgesState,
  addEdge,
  ReactFlow,
  Background,
  BackgroundVariant,
  useReactFlow,
  ReactFlowProvider,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { useMemo, useEffect, useCallback, DragEvent } from "react";
import {
  SelectDeviceNode,
  SelectDeviceNodeData,
} from "./custom-nodes/select-device";
import { ConditionNode, ConditionNodeData } from "./custom-nodes/condition";
import {
  DeviceOperationNode,
  DeviceOperationNodeData,
} from "./custom-nodes/device-operation";
import { Sidebar } from "./sidebar";
import { DnDProvider, useDnD } from "@/hooks/dnd-context";
import { NodeInfoProvider } from "@/hooks/node-info-context";
import { WorkflowWithUIResponse } from "@/types/api";
import { getWorkflows } from "@/mocks/workflow-api";
import { useWorkflowInfo } from "@/hooks/workflow-info-context";

type CustomNodeData =
  | SelectDeviceNodeData
  | ConditionNodeData
  | DeviceOperationNodeData;

export type AddNodeFunction = (parentNodeId: string) => void;
export type UpdateNodeFunction = (
  id: string,
  updatedData: CustomNodeData
) => void;

const nodeIdMap: Map<string, number> = new Map();
const getId = (type: string) => {
  const currentId = nodeIdMap.get(type) || 1;
  nodeIdMap.set(type, currentId + 1);
  return `${type}_${currentId}`;
};

interface WorkflowEditorProps {
  workflowID: number | null;
}

function WorkflowEditor({ workflowID }: WorkflowEditorProps) {
  const [nodes, setNodes, onNodesChange] = useNodesState<Node>([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>([]);
  const [type] = useDnD();
  const [workflowInfo] = useWorkflowInfo();

  const nodeTypes = useMemo(
    () => ({
      select_device: SelectDeviceNode,
      condition: ConditionNode,
      device_operation: DeviceOperationNode,
    }),
    []
  );

  useEffect(() => {
    if (workflowID) {
      const workflow: WorkflowWithUIResponse = getWorkflows();

      workflow.workflow_ui.nodes.forEach((node) => {
        const currentId = nodeIdMap.get(node.node_type) || 1;
        nodeIdMap.set(node.node_type, currentId + 1);
      });

      const nodes = workflow.workflow_ui.nodes
        .map((node): Node | undefined => {
          if (node.node_type === "select_device") {
            return {
              id: node.workflow_node_id,
              type: node.node_type,
              position: { x: node.position_x, y: node.position_y },
              data: {
                ...(node.data as Record<string, unknown>),
                devicesList: workflowInfo.devices,
                updateNode: updateNodeData,
              },
            };
          } else if (node.node_type === "condition") {
            return {
              id: node.workflow_node_id,
              type: node.node_type,
              position: { x: node.position_x, y: node.position_y },
              data: {
                ...(node.data as Record<string, unknown>),
                climateDataList: workflowInfo.climate_data,
                updateNode: updateNodeData,
              },
            };
          } else if (node.node_type === "device_operation") {
            return {
              id: node.workflow_node_id,
              type: node.node_type,
              position: { x: node.position_x, y: node.position_y },
              data: {
                ...(node.data as Record<string, unknown>),
                operationsList: workflowInfo.operations,
                updateNode: updateNodeData,
              },
            };
          }

          return undefined;
        })
        .filter((node): node is Node => node !== undefined);

      const edges = workflow.workflow_ui.edges.map((edge) => ({
        id: edge.id.toString(),
        source: edge.source_node_id,
        target: edge.target_node_id,
        style: { strokeWidth: 4 },
      }));

      setNodes(nodes);
      setEdges(edges);

      return;
    }

    const initialNode: Node = {
      id: "select_device_1",
      type: "select_device",
      position: { x: 0, y: 300 },
      data: {
        devicesList: workflowInfo.devices,
        updateNode: updateNodeData,
      },
    };

    setNodes([initialNode]);
  }, []);

  const { screenToFlowPosition } = useReactFlow();

  // イベントハンドラー
  const onConnect = useCallback(
    (params: Connection) => {
      const animatedEdge = {
        ...params,
        style: { strokeWidth: 4 },
      };
      setEdges((eds) => addEdge(animatedEdge, eds));
    },
    [setEdges]
  );

  const onDragOver = useCallback((event: DragEvent) => {
    event.preventDefault();
    event.dataTransfer.dropEffect = "move";
  }, []);

  const updateNodeData = useCallback(
    (id: string, updatedData: CustomNodeData) => {
      setNodes((nds) =>
        nds.map((node) =>
          node.id === id
            ? { ...node, data: { ...node.data, ...updatedData } }
            : node
        )
      );
    },
    [setNodes]
  );

  const onDrop = useCallback(
    (event: DragEvent) => {
      event.preventDefault();

      if (!type) {
        return;
      }

      const dataString = event.dataTransfer.getData("application/reactflow");
      const nodeData = dataString ? JSON.parse(dataString) : {};

      const position = screenToFlowPosition({
        x: event.clientX,
        y: event.clientY,
      });
      const newNode = {
        id: getId(type),
        type,
        position,
        data: {
          label: `${type} node`,
          ...nodeData,
          updateNode: updateNodeData,
        },
      };

      setNodes((nds) => nds.concat(newNode));
    },
    [screenToFlowPosition, type]
  );

  // 画面の大きさ調節
  interface Viewport {
    x: number;
    y: number;
    zoom: number;
  }
  const defaultViewport: Viewport = { x: 50, y: 15, zoom: 0 };

  return (
    <>
      <Box sx={{ width: "100%", height: "80vh", backgroundColor: "#eee" }}>
        <Box sx={{ width: "100%", height: "100%", display: "flex" }}>
          <ReactFlow
            nodes={nodes}
            edges={edges}
            nodesDraggable={true} // ノードのドラッグを無効化
            edgesReconnectable={true} // エッジの更新を無効化
            // panOnDrag={false} // 画面全体のドラッグを無効化
            // zoomOnScroll={false} // マウスホイールでのズームを無効化
            zoomOnPinch={false} // ピンチ操作でのズームを無効化
            zoomOnDoubleClick={false} // ダブルクリックでのズームを無効化
            defaultViewport={defaultViewport} // 初期配置と大きさを設定
            nodeTypes={nodeTypes}
            onNodesChange={onNodesChange}
            onEdgesChange={onEdgesChange}
            onConnect={onConnect}
            onDrop={onDrop}
            onDragOver={onDragOver}
          >
            <Background
              color="#000"
              variant={BackgroundVariant.Dots}
            ></Background>
          </ReactFlow>
          <Sidebar />
        </Box>
      </Box>
    </>
  );
}

interface WorkflowWrapperProps {
  workflowID: number;
}

export const WorkflowWrapper = ({ workflowID }: WorkflowWrapperProps) => {
  return (
    <ReactFlowProvider>
      <NodeInfoProvider>
        <DnDProvider>
          <WorkflowEditor workflowID={workflowID} />
        </DnDProvider>
      </NodeInfoProvider>
    </ReactFlowProvider>
  );
};
