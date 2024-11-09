import {
  Box,
  Divider,
  Typography,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
} from "@mui/material";
import { Handle, Position, Node, NodeProps } from "@xyflow/react";
import { AddNodeFunction, UpdateNodeFunction } from "../workflow";
import { useCallback, useState } from "react";
import { DeviceResponse } from "@/types/api";
import { useNodeInfo } from "@/hooks/node-info-context";

export interface SelectDeviceNodeData {
  [key: string]: unknown;
  devicesList: DeviceResponse[];
  addNode: AddNodeFunction;
  updateNode: UpdateNodeFunction;
  device_id: number;
}

type SelectDeviceNodePropsType = Node<SelectDeviceNodeData>;

type SelectDeviceNodeProps = NodeProps<SelectDeviceNodePropsType>;

export const SelectDeviceNode = ({ id, data }: SelectDeviceNodeProps) => {
  const { devicesList, addNode, updateNode } = data;
  const [selectedDevice, setSelectedDevice] = useState<string>("");
  const [workflowInfo, setWorkflowInfo] = useNodeInfo();
  const handleSelectedDeviceChange = useCallback(
    (event: SelectChangeEvent) => {
      const selectedDeviceID: number = parseInt(event.target.value, 10);
      workflowInfo.device_id = selectedDeviceID;

      setSelectedDevice(event.target.value);
      setWorkflowInfo(workflowInfo);
      updateNode(id, { ...data, device_id: selectedDeviceID });
    },
    [id, updateNode, data, setSelectedDevice]
  );

  return (
    <Box
      sx={{
        border: "1px solid #000",
        borderRadius: "10px",
        backgroundColor: "#FFF",
        width: "350px",
      }}
    >
      <Box sx={{ padding: "8px" }}>
        <Typography variant="h6">Select Device</Typography>
        <Divider />
        <Box sx={{ padding: "8px", display: "flex", justifyContent: "center" }}>
          <Box sx={{ flex: 2 }}>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">デバイス</InputLabel>
              <Select
                value={selectedDevice}
                size="small"
                onChange={handleSelectedDeviceChange}
                label="デバイス"
                inputProps={{ className: "nodrag nopan nowheel" }}
              >
                {devicesList.map((data) => (
                  <MenuItem key={data.id} value={data.id}>
                    {data.device_name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Box>
        </Box>
      </Box>
      <Handle
        position={Position.Right}
        type="source"
        onClick={(event) => {
          event.stopPropagation();
          addNode(id);
        }}
      />
    </Box>
  );
};
