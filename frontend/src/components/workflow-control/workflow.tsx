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
import { useMemo, useEffect, useCallback, DragEvent, useState } from "react";
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
import { DeviceResponse } from "@/types/api";
import { getDevices } from "@/mocks/workflow_api";

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

function WorkflowEditor() {
  const [nodes, setNodes, onNodesChange] = useNodesState<Node>([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>([]);
  const [type] = useDnD();
  const [fetchedDevices, setFetchedDevices] = useState<DeviceResponse[]>([]);

  const nodeTypes = useMemo(
    () => ({
      select_device: SelectDeviceNode,
      condition: ConditionNode,
      device_operation: DeviceOperationNode,
    }),
    []
  );

  useEffect(() => {
    const fetchDevices = async () => {
      const devicesRes: DeviceResponse[] = await getDevices();
      setFetchedDevices(devicesRes);
    };

    fetchDevices();
  }, []);

  useEffect(() => {
    const initialNode: Node = {
      id: "select_device_1",
      type: "select_device",
      position: { x: 0, y: 300 },
      data: {
        label: "Begin Workflow",
        devicesList: fetchedDevices,
        updateNode: updateNodeData,
      },
      connectable: false,
    };

    setNodes([initialNode]);
  }, [fetchedDevices]);

  const { screenToFlowPosition } = useReactFlow();

  // イベントハンドラー
  const onConnect = useCallback(
    (params: Connection) => {
      const animatedEdge = {
        ...params,
        animated: true,
        style: { strokeWidth: 5 },
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

export const Workflow = () => (
  <ReactFlowProvider>
    <NodeInfoProvider>
      <DnDProvider>
        <WorkflowEditor />
      </DnDProvider>
    </NodeInfoProvider>
  </ReactFlowProvider>
);
