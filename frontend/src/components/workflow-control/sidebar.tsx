import { useDnD } from "@/hooks/dnd-context";
import { DragEvent } from "react";
import { Box, Divider, Typography } from "@mui/material";
import PlayCircleIcon from "@mui/icons-material/PlayCircle";
import RuleIcon from "@mui/icons-material/Rule";
import { useState, useEffect } from "react";
import { ClimateDataResponse, OperationResponse } from "@/types/api";
import { getClimateDatas, getOperations } from "@/mocks/workflow_api";

export const Sidebar = () => {
  const [, setType] = useDnD();
  const [fetchedClimateDatas, setFetchedClimateDatas] = useState<
    ClimateDataResponse[]
  >([]);
  const [fetchedOperations, setFetchedOperations] = useState<
    OperationResponse[]
  >([]);

  useEffect(() => {
    const fetchClimateDatas = async () => {
      const climateDataRes: ClimateDataResponse[] = await getClimateDatas();
      setFetchedClimateDatas(climateDataRes);
    };

    const fetchOperations = async () => {
      const operationsRes: OperationResponse[] = await getOperations();
      setFetchedOperations(operationsRes);
    };

    fetchClimateDatas();
    fetchOperations();
  }, []);

  const onDragStart = (
    event: DragEvent,
    nodeType: string,
    nodeData: object
  ) => {
    setType(nodeType);

    event.dataTransfer.setData(
      "application/reactflow",
      JSON.stringify(nodeData)
    );
    event.dataTransfer.effectAllowed = "move";
  };

  return (
    <Box sx={{ height: "100%", width: "300px", backgroundColor: "#E0E0E0" }}>
      <Box sx={{ display: "flex", flexDirection: "column" }}>
        <Box
          sx={{
            mx: 4,
            my: 2,
            border: "1px solid #000",
            borderRadius: "10px",
            backgroundColor: "#FFF",
          }}
          onDragStart={(event) =>
            onDragStart(event, "condition", {
              climateDataList: fetchedClimateDatas,
            })
          }
          draggable
        >
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
            <Typography variant="body1">If</Typography>
          </Box>
          <Divider />
          <Box sx={{ padding: 1, textAlign: "center" }}>条件</Box>
        </Box>
        <Box
          sx={{
            mx: 4,
            my: 2,
            border: "1px solid #000",
            borderRadius: "10px",
            backgroundColor: "#FFF",
          }}
          onDragStart={(event) =>
            onDragStart(event, "device_operation", {
              operationsList: fetchedOperations,
            })
          }
          draggable
        >
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
            <Typography variant="body1">Operation</Typography>
          </Box>
          <Divider />
          <Box sx={{ padding: 1, textAlign: "center" }}>操作</Box>
        </Box>
      </Box>
    </Box>
  );
};
