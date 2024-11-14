import {
  Box,
  Select,
  MenuItem,
  SelectChangeEvent,
  Divider,
  FormControl,
  InputLabel,
  Typography,
} from "@mui/material";
import PlayCircleIcon from "@mui/icons-material/PlayCircle";
import { Handle, Position, Node, NodeProps } from "@xyflow/react";
import { useEffect, useState } from "react";
import { UpdateNodeFunction } from "../workflow-editor";
import { OperationResponse } from "@/types/api";
import { useNodeInfo } from "@/hooks/node-info-context";

export interface DeviceOperationNodeData {
  [key: string]: unknown;
  operationsList: OperationResponse[];
  updateNode: UpdateNodeFunction;
  operationID: number;
}

type DeviceOperationNodePropsType = Node<DeviceOperationNodeData>;

type DeviceOperationNodeProps = NodeProps<DeviceOperationNodePropsType>;

export function DeviceOperationNode({ id, data }: DeviceOperationNodeProps) {
  const { operationsList, updateNode } = data;
  const [selectedOperation, setSelectedOperation] = useState<string>("");
  const handleSelectedDeviceChange = (event: SelectChangeEvent) => {
    const selectedOperationID: number = parseInt(event.target.value, 10);
    setSelectedOperation(event.target.value);
    updateNode(id, { ...data, operationID: selectedOperationID });
  };
  const [deviceOperations, setDeviceOperations] = useState<OperationResponse[]>(
    []
  );
  const [nodeInfo] = useNodeInfo();

  useEffect(() => {
    const selectBoxOperations: OperationResponse[] = operationsList.filter(
      (data) => data.device_id === nodeInfo?.device_id
    );

    setDeviceOperations(selectBoxOperations);

    // const climateDataRec = climateDataList.find(
    //   (data) => data.id === climateDataID
    // );
  }, [operationsList, nodeInfo.device_id]);

  return (
    <Box
      sx={{
        border: "1px solid #000",
        borderRadius: "10px",
        backgroundColor: "#FFF",
        width: "350px",
      }}
    >
      <Handle
        position={Position.Left}
        type="target"
        style={{ width: 12, height: 12 }}
      />
      <Box>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gap: 1,
            borderRadius: "10px 10px 0 0",
            color: "#FFF",
            backgroundColor: "#66BB6A",
            padding: "4px 8px 4px 8px",
          }}
        >
          <PlayCircleIcon />
          <Typography variant="h6">Operation</Typography>
        </Box>
        <Divider />
        <Box sx={{ padding: 2, display: "flex", justifyContent: "center" }}>
          <Box sx={{ flex: 2 }}>
            <FormControl fullWidth>
              <InputLabel
                id={`device-operation-node-select-label-${id}`}
                size="small"
              >
                操作
              </InputLabel>
              <Select
                value={selectedOperation}
                labelId={`device-operation-node-select-label-${id}`}
                id={`device-operation-node-select-${id}`}
                size="small"
                onChange={handleSelectedDeviceChange}
                label="操作"
                inputProps={{ className: "nodrag nopan nowheel" }}
                disabled={deviceOperations.length === 0}
              >
                {deviceOperations.map((data) => (
                  <MenuItem key={data.id} value={data.id}>
                    {data.name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
