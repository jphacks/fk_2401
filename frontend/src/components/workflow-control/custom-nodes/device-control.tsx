import {
  Box,
  Select,
  MenuItem,
  SelectChangeEvent,
  Divider,
  FormControlLabel,
  Switch,
  FormControl,
  InputLabel,
  Typography,
} from "@mui/material";
import { Handle, Position } from "@xyflow/react";
import { useState } from "react";

export function DeviceControlNode() {
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
      }}
    >
      <Handle position={Position.Left} type="target" />
      <Box sx={{ padding: "8px" }}>
        <Typography variant="h6">Control</Typography>
        <Divider />
        <Box sx={{ padding: "8px", display: "flex", justifyContent: "center" }}>
          <Box sx={{ flex: 1 }}>
            <FormControlLabel
              control={
                <Switch inputProps={{ className: "nodrag nopan nowheel" }} />
              }
              label="有効"
            />
          </Box>
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
    </Box>
  );
}
