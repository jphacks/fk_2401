import { Box } from "@mui/material";
import {
  Node,
  Edge,
  ReactFlow,
  Background,
  BackgroundVariant,
  applyEdgeChanges,
  applyNodeChanges,
  NodeChange,
  EdgeChange,
  useReactFlow,
  ReactFlowProvider,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { useMemo, useState, useEffect, useCallback } from "react";
import { ConditionNode } from "./custom-nodes/condition";
import { DeviceControlNode } from "./custom-nodes/device-control";
import { SelectDeviceNode } from "./custom-nodes/select-device";
import { ClimateDataResponse } from "@/types/api";
import { getClimateDatas } from "@/mocks/setting_device_api";
import { Sidebar } from "./sidebar";
import { DnDProvider, useDnD } from "./dnd-context";

export type AddNodeFunction = (parentNodeId: string) => void;

let id = 0;
const getId = () => `dndnode_${id++}`;

function DnDWorkflow() {
  const [fetchedClimateDatas, setFetchedClimateDatas] = useState<
    ClimateDataResponse[]
  >([]);
  const [nodes, setNodes] = useState<Node[]>([]);
  const [edges, setEdges] = useState<Edge[]>([]);
  const [type] = useDnD();

  const onNodesChange = useCallback(
    (changes: NodeChange[]) =>
      setNodes((nds) => applyNodeChanges(changes, nds)),
    [setNodes]
  );
  const onEdgesChange = useCallback(
    (changes: EdgeChange[]) =>
      setEdges((eds) => applyEdgeChanges(changes, eds)),
    [setEdges]
  );

  const nodeTypes = useMemo(
    () => ({
      selectDevice: SelectDeviceNode,
      condition: ConditionNode,
      deviceControl: DeviceControlNode,
    }),
    []
  );

  useEffect(() => {
    const fetchClimateDatas = async () => {
      const climateDataRes: ClimateDataResponse[] = await getClimateDatas();
      setFetchedClimateDatas(climateDataRes);
    };

    const initialNode: Node = {
      id: "1",
      type: "selectDevice",
      position: { x: 0, y: 300 },
      data: { label: "Begin Workflow" },
      style: {
        width: "350px",
        height: "100px",
        backgroundColor: "#fff",
        border: "1px solid #000",
        borderRadius: "10px",
      },
      connectable: false,
    };

    setNodes([initialNode]);

    fetchClimateDatas();
  }, []);

  const { screenToFlowPosition } = useReactFlow();

  const onDragOver = useCallback((event) => {
    event.preventDefault();
    event.dataTransfer.dropEffect = "move";
  }, []);

  const onDrop = useCallback(
    (event) => {
      event.preventDefault();

      // check if the dropped element is valid
      if (!type) {
        return;
      }

      const dataString = event.dataTransfer.getData("application/reactflow");
      const nodeData = dataString ? JSON.parse(dataString) : {};

      // project was renamed to screenToFlowPosition
      // and you don't need to subtract the reactFlowBounds.left/top anymore
      // details: https://reactflow.dev/whats-new/2023-11-10
      const position = screenToFlowPosition({
        x: event.clientX,
        y: event.clientY,
      });
      const newNode = {
        id: getId(),
        type,
        position,
        data: { label: `${type} node`, ...nodeData },
      };

      setNodes((nds) => nds.concat(newNode));
    },
    [screenToFlowPosition, type]
  );

  interface Viewport {
    x: number;
    y: number;
    zoom: number;
  }
  const defaultViewport: Viewport = { x: 50, y: 15, zoom: 0.8 };

  return (
    <>
      <Box sx={{ width: "100%", height: "100vh", backgroundColor: "#eee" }}>
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
    <DnDProvider>
      <DnDWorkflow />
    </DnDProvider>
  </ReactFlowProvider>
);
