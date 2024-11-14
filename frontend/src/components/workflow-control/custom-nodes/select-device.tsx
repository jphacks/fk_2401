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
import DevicesIcon from "@mui/icons-material/Devices";
import { Handle, Position, Node, NodeProps } from "@xyflow/react";
import { AddNodeFunction, UpdateNodeFunction } from "../workflow-editor";
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
  const [nodeInfo, setNodeInfo] = useNodeInfo();
  const handleSelectedDeviceChange = useCallback(
    (event: SelectChangeEvent) => {
      const selectedDeviceID: number = parseInt(event.target.value, 10);
      nodeInfo.device_id = selectedDeviceID;

      setSelectedDevice(event.target.value);
      setNodeInfo(nodeInfo);
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
      <Box>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gap: 1,
            borderRadius: "10px 10px 0 0",
            color: "#FFF",
            backgroundColor: "#42A5F5",
            padding: "4px 8px 4px 8px",
          }}
        >
          <DevicesIcon />
          <Typography variant="h6">Select Device</Typography>
        </Box>
        <Divider />
        <Box sx={{ padding: 2, display: "flex", justifyContent: "center" }}>
          <Box sx={{ flex: 2 }}>
            <FormControl fullWidth>
              <InputLabel
                id={`select-device-node-select-label-${id}`}
                size="small"
              >
                デバイス
              </InputLabel>
              <Select
                value={selectedDevice}
                size="small"
                labelId={`select-device-node-select-label-${id}`}
                id={`select-device-node-select-${id}`}
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
        style={{ width: 12, height: 12 }}
        onClick={(event) => {
          event.stopPropagation();
          addNode(id);
        }}
      />
    </Box>
  );
};
