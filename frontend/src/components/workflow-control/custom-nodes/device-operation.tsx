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
import { Handle, Position } from "@xyflow/react";
import { useState } from "react";

export function DeviceOperationNode() {
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
      <Handle position={Position.Left} type="target" />
      <Box sx={{ padding: "8px" }}>
        <Typography variant="h6">Operation</Typography>
        <Divider />
        <Box sx={{ padding: "8px", display: "flex", justifyContent: "center" }}>
          <Box sx={{ flex: 2 }}>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">操作</InputLabel>
              <Select
                value={selectedDevice.toString()}
                size="small"
                onChange={handleSelectedDeviceChange}
                label="操作"
                inputProps={{ className: "nodrag nopan nowheel" }}
              >
                <MenuItem value={1}>{"送風"}</MenuItem>
                <MenuItem value={2}>{"加温"}</MenuItem>
              </Select>
            </FormControl>
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
