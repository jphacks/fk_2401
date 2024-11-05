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
import { AddNodeFunction } from "../workflow";
import { useState } from "react";

interface SelectDeviceNodeData {
  [key: string]: unknown;
  addNode: AddNodeFunction;
}

type SelectDeviceNodePropsType = Node<SelectDeviceNodeData>;

type SelectDeviceNodeProps = NodeProps<SelectDeviceNodePropsType>;

export const SelectDeviceNode = ({ id, data }: SelectDeviceNodeProps) => {
  const { addNode } = data;
  const [selectedDevice, setSelectedDevice] = useState<number>(0);
  const handleSelectedDeviceChange = (event: SelectChangeEvent) => {
    const selectedDeviceID: number = parseInt(event.target.value, 10);
    setSelectedDevice(selectedDeviceID);
  };

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
                value={selectedDevice.toString()}
                size="small"
                onChange={handleSelectedDeviceChange}
                label="デバイス"
                inputProps={{ className: "nodrag nopan nowheel" }}
              >
                <MenuItem value={1}>{"加温器1"}</MenuItem>
                <MenuItem value={2}>{"窓開閉装置1"}</MenuItem>
                <MenuItem value={3}>{"加温器2"}</MenuItem>
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
