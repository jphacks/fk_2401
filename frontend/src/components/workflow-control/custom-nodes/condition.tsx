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
import RuleIcon from "@mui/icons-material/Rule";
import { Node, Handle, Position, NodeProps } from "@xyflow/react";
import { useEffect, useState } from "react";
import { ClimateDataResponse } from "@/types/api";
import { AddNodeFunction, UpdateNodeFunction } from "../workflow-editor";
import { Condition } from "@/types/workflow";

export interface ConditionNodeData {
  [key: string]: unknown;
  climateDataList: ClimateDataResponse[];
  addNode: AddNodeFunction;
  updateNode: UpdateNodeFunction;
  condition: Condition;
}

type ConditionNodePropsType = Node<ConditionNodeData>;

type ConditionNodeProps = NodeProps<ConditionNodePropsType>;

export const ConditionNode = ({ id, data }: ConditionNodeProps) => {
  const { climateDataList, addNode, updateNode } = data as ConditionNodeData;
  const [selectedClimateData, setSelectedClimateData] = useState<
    ClimateDataResponse | undefined
  >(undefined);
  const [cmpOpe, setCmpOpe] = useState<string>("");
  const [condition, setCondition] = useState<Condition>({
    climate_data_id: 0,
    comp_ope_id: 0,
    set_point: 0,
  });

  useEffect(() => {
    updateNode(id, { ...data, condition: condition });
  }, [id, updateNode, condition]);

  const handleClimateDataChange = (event: SelectChangeEvent) => {
    const climateDataID = parseInt(event.target.value);
    const climateDataRec = climateDataList.find(
      (data) => data.id === climateDataID
    );
    condition.climate_data_id = climateDataID;

    setSelectedClimateData(climateDataRec);
    setCondition(condition);
  };

  const handleCmpOpeChange = (event: SelectChangeEvent) => {
    const compOpeID: number = parseInt(event.target.value);
    condition.comp_ope_id = compOpeID;

    setCmpOpe(event.target.value as string);
    setCondition(condition);
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
            backgroundColor: "#F57C00",
            padding: "4px 8px 4px 8px",
          }}
        >
          <RuleIcon />
          <Typography variant="h6">If</Typography>
        </Box>
        <Divider />
        <Box sx={{ padding: 2, display: "flex", justifyContent: "center" }}>
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
              value={selectedClimateData ? String(selectedClimateData.id) : ""}
              label="気象データ"
              onChange={handleClimateDataChange}
              size="small"
              inputProps={{ className: "nodrag nopan nowheel" }}
            >
              {climateDataList.map((data) => (
                <MenuItem key={data.id} value={data.id}>
                  {data.climate_data}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <Select
            value={cmpOpe}
            size="small"
            onChange={handleCmpOpeChange}
            sx={{ flex: 1, marginX: 1 }}
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
                    {selectedClimateData?.unit}
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
        style={{ width: 12, height: 12 }}
        onClick={(event) => {
          event.stopPropagation();
          addNode(id);
        }}
      />
    </Box>
  );
};
