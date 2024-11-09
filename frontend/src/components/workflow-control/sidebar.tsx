import { useDnD } from "@/hooks/dnd-context";
import { DragEvent } from "react";
import { Box, Divider, Typography } from "@mui/material";
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
    <Box sx={{ height: "100%", width: "300px", backgroundColor: "#ddd" }}>
      <Box sx={{ display: "flex", flexDirection: "column" }}>
        <Box
          sx={{
            mx: 4,
            my: 2,
            padding: 1,
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
          <Typography variant="subtitle2">If</Typography>
          <Divider />
          <Box sx={{ pt: 1, textAlign: "center" }}>条件ノード</Box>
        </Box>
        <Box
          sx={{
            mx: 4,
            my: 2,
            padding: 1,
            border: "1px solid #000",
            borderRadius: "10px",
            backgroundColor: "#FFF",
          }}
          onDragStart={(event) =>
            onDragStart(event, "deviceOperation", {
              operationsList: fetchedOperations,
            })
          }
          draggable
        >
          <Typography variant="subtitle2">Operation</Typography>
          <Divider />
          <Box sx={{ pt: 1, textAlign: "center" }}>オペレーションノード</Box>
        </Box>
      </Box>
    </Box>
  );
};
