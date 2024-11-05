import {
  Box,
  TextField,
  Select,
  MenuItem,
  InputAdornment,
  SelectChangeEvent,
  Divider,
  FormControl,
  InputLabel,
  Typography,
} from "@mui/material";
import { Node, Handle, Position, NodeProps } from "@xyflow/react";
import { useState } from "react";
import { ClimateDataResponse } from "@/types/api";
import { AddNodeFunction } from "../workflow";

interface ConditionNodeData {
  [key: string]: unknown;
  climateDataList: ClimateDataResponse[];
  addNode: AddNodeFunction;
}

type ConditionNodePropsType = Node<ConditionNodeData>;

type ConditionNodeProps = NodeProps<ConditionNodePropsType>;

export const ConditionNode = ({ id, data }: ConditionNodeProps) => {
  const { climateDataList, addNode } = data as ConditionNodeData;

  const [selectedClimateData, setSelectedClimateData] = useState<string>("");
  const [selectedClimateDataRec, setSelectedClimateDataRec] =
    useState<ClimateDataResponse>();
  const [cmpOpe, setCmpOpe] = useState<string>("");

  const handleClimateDataChange = (event: SelectChangeEvent) => {
    const climateData = event.target.value;
    const climateDataRec = climateDataList.find(
      (data) => data.climate_data === climateData
    );
    setSelectedClimateData(climateData);
    setSelectedClimateDataRec(climateDataRec);
  };

  const handleCmpOpeChange = (event: SelectChangeEvent) => {
    setCmpOpe(event.target.value as string);
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
        <Typography variant="h6">If</Typography>
        <Divider />
        <Box sx={{ padding: "8px", display: "flex", justifyContent: "center" }}>
          <FormControl sx={{ flex: 4 }}>
            <InputLabel
              id={`climate-data-node-select-label-${id}`}
              size="small"
            >
              気象データ
            </InputLabel>
            <Select
              className="nodrag"
              labelId={`climate-data-node-select-label-${id}`}
              id={`climate-data-node-select-${id}`}
              value={selectedClimateData}
              label="気象データ"
              onChange={handleClimateDataChange}
              size="small"
              inputProps={{ className: "nodrag nopan nowheel" }}
            >
              {climateDataList.map((data) => (
                <MenuItem key={data.id} value={data.climate_data}>
                  {data.climate_data}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <Select
            value={cmpOpe}
            size="small"
            onChange={handleCmpOpeChange}
            sx={{ flex: 1, marginX: "8px" }}
            inputProps={{ className: "nodrag nopan nowheel" }}
          >
            <MenuItem value={1}>{"="}</MenuItem>
            <MenuItem value={2}>{">"}</MenuItem>
            <MenuItem value={3}>{"<"}</MenuItem>
            <MenuItem value={4}>{"≥"}</MenuItem>
            <MenuItem value={5}>{"≤"}</MenuItem>
          </Select>
          <TextField
            type="number"
            size="small"
            slotProps={{
              input: {
                endAdornment: (
                  <InputAdornment position="end">
                    {selectedClimateDataRec?.unit}
                  </InputAdornment>
                ),
              },
            }}
            inputProps={{ step: "0.1", className: "nodrag nopan nowheel" }}
            sx={{ flex: 3, marginRight: "8px" }}
          />
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
