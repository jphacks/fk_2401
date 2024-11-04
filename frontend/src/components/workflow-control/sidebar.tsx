import { useDnD } from "./dnd-context";
import { DragEvent } from "react";
import { Box, Divider, Typography } from "@mui/material";
import { useState, useEffect } from "react";
import { ClimateDataResponse } from "@/types/api";
import { getClimateDatas } from "@/mocks/setting_device_api";

export const Sidebar = () => {
  const [, setType] = useDnD();
  const [fetchedClimateDatas, setFetchedClimateDatas] = useState<
    ClimateDataResponse[]
  >([]);

  useEffect(() => {
    const fetchClimateDatas = async () => {
      const climateDataRes: ClimateDataResponse[] = await getClimateDatas();
      setFetchedClimateDatas(climateDataRes);
    };

    fetchClimateDatas();
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
      <Box>You can drag these nodes to the pane on the left.</Box>
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
          onDragStart={(event) => onDragStart(event, "dndnode", {})}
          draggable
        >
          <Typography variant="subtitle2">Function</Typography>
          <Divider />
          <Box sx={{ pt: 1, textAlign: "center" }}>ファンクションノード</Box>
        </Box>
      </Box>
    </Box>
  );
};
