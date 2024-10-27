import { Box } from "@mui/material";
import {
  Node,
  Edge,
  ReactFlow,
  Background,
  BackgroundVariant,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { useMemo, useState, useEffect, useCallback, useRef } from "react";
import { ConditionNode } from "./custom-nodes/condition";
import { DeviceControlNode } from "./custom-nodes/device-control";
import { BeginWorkflowNode } from "./custom-nodes/begin-workflow";
import { ClimateDataResponse } from "@/types/api";
import { getClimateDatas } from "@/mocks/setting_device_api";

export type AddNodeFunction = (parentNodeId: string) => void;

export function Workflow() {
  const [conditionNodeCount, setConditionNodeCount] = useState<number>(0);
  const [deviceControlNodeCount, setDeviceControlNodeCount] =
    useState<number>(0);
  const maxConditionNode = 8;
  const maxDeviceControlNode = 8;
  const [fetchedClimateDatas, setFetchedClimateDatas] = useState<
    ClimateDataResponse[]
  >([]);
  const [nodes, setNodes] = useState<Node[]>([]);
  const [edges, setEdges] = useState<Edge[]>([]);

  const nodesRef = useRef<Node[]>(nodes);
  useEffect(() => {
    nodesRef.current = nodes;
  }, [nodes]);

  const fetchedClimateDatasRef =
    useRef<ClimateDataResponse[]>(fetchedClimateDatas);
  useEffect(() => {
    fetchedClimateDatasRef.current = fetchedClimateDatas;
  }, [fetchedClimateDatas]);

  const onAddDeviceControlNode: AddNodeFunction = (parentNodeId: string) => {
    const parentNode: Node | undefined = nodesRef.current.find(
      (node) => node.id == parentNodeId
    );
    if (!parentNode) {
      return false;
    }

    setDeviceControlNodeCount((prevCount) => {
      if (prevCount < maxDeviceControlNode) {
        const conditionNodeCountUpdate: number = prevCount + 1;

        setNodes((nodes) => {
          const newNodeId = `${nodes.length + 1}`;
          const x = parentNode.position.x + 600;
          const y = parentNode.position.y;

          const newNode = {
            id: newNodeId,
            type: "deviceControl",
            position: { x, y },
            data: {
              label: `Node ${nodes.length + 1}`,
            },
            parentNodeId: parentNodeId,
            style: {
              width: "350px",
              height: "100px",
              backgroundColor: "#fff",
              border: "1px solid #ccc",
              borderRadius: "10px",
            },
          };

          setEdges((edges) => {
            const newEdge: Edge = {
              id: `e${parentNodeId}-${newNodeId}`,
              source: parentNodeId,
              target: newNodeId,
              type: "default",
              animated: true,
            };

            return [...edges, newEdge];
          });

          return [...nodes, newNode];
        });

        return conditionNodeCountUpdate;
      }
      return prevCount;
    });
  };

  const onAddConditionNode: AddNodeFunction = useCallback(
    (parentNodeId: string) => {
      const parentNode: Node | undefined = nodesRef.current.find(
        (node) => node.id == parentNodeId
      );
      if (!parentNode) {
        return false;
      }

      setConditionNodeCount((prevCount) => {
        if (prevCount < maxConditionNode) {
          const conditionNodeCountUpdate: number = prevCount + 1;

          console.log(fetchedClimateDatasRef.current);

          setNodes((nodes) => {
            const newNodeId = `${nodes.length + 1}`;
            const x = parentNode.position.x + 400;
            let y = parentNode.position.y;
            if (prevCount % 2 == 1) {
              y = y + 150 * (conditionNodeCountUpdate / 2);
            } else {
              y = y - 150 * (conditionNodeCountUpdate / 2);
            }

            const newNode = {
              id: newNodeId,
              type: "condition",
              position: { x, y },
              data: {
                label: `Node ${nodes.length + 1}`,
                climateDataList: fetchedClimateDatasRef.current,
                addNode: onAddDeviceControlNode,
              },
              parentNodeId: parentNodeId,
              style: {
                width: "400px",
                height: "100px",
                backgroundColor: "#fff",
                border: "1px solid #ccc",
                borderRadius: "10px",
              },
              connectable: false,
            };

            setEdges((edges) => {
              const newEdge: Edge = {
                id: `e${parentNodeId}-${newNodeId}`,
                source: parentNodeId,
                target: newNodeId,
                type: "default",
                animated: true,
              };

              return [...edges, newEdge];
            });

            return [...nodes, newNode];
          });

          return conditionNodeCountUpdate;
        }
        return prevCount;
      });
    },
    []
  );

  const nodeTypes = useMemo(
    () => ({
      beginWorkflow: BeginWorkflowNode,
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
      type: "beginWorkflow",
      position: { x: 0, y: 300 },
      data: { label: "Begin Workflow", addNode: onAddConditionNode },
      style: {
        width: "200px",
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

  console.log(nodes);

  interface Viewport {
    x: number;
    y: number;
    zoom: number;
  }
  const defaultViewport: Viewport = { x: 50, y: 15, zoom: 0.8 };

  return (
    <>
      <Box sx={{ width: "100%", height: "100vh" }}>
        <ReactFlow
          nodes={nodes}
          edges={edges}
          nodesDraggable={false} // ノードのドラッグを無効化
          edgesReconnectable={false} // エッジの更新を無効化
          // panOnDrag={false} // 画面全体のドラッグを無効化
          // zoomOnScroll={false} // マウスホイールでのズームを無効化
          zoomOnPinch={false} // ピンチ操作でのズームを無効化
          zoomOnDoubleClick={false} // ダブルクリックでのズームを無効化
          defaultViewport={defaultViewport} // 初期配置と大きさを設定
          nodeTypes={nodeTypes}
        >
          <Background
            color="#000"
            variant={BackgroundVariant.Dots}
          ></Background>
        </ReactFlow>
      </Box>
    </>
  );
}
