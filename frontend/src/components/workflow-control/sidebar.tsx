import { useDnD } from "./dnd-context";
import { DragEvent } from "react";
import { Box } from "@mui/material";
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
      <Box>You can drag these nodes to the pane on the right.</Box>
      <Box
        onDragStart={(event) =>
          onDragStart(event, "condition", {
            climateDataList: fetchedClimateDatas,
          })
        }
        draggable
      >
        条件ノード
      </Box>
      <Box onDragStart={(event) => onDragStart(event, "dndnode", {})} draggable>
        ファンクションノード
      </Box>
    </Box>
  );
};
