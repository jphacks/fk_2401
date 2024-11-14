import { Box, Button, Select, FormControl, InputLabel } from "@mui/material";
import { Navigation } from "@/layouts/navigation";
import { WorkflowWrapper } from "./workflow-editor";
import { WorkflowInfoProvider } from "@/hooks/workflow-info-context";
import { useWorkflowInfo } from "@/hooks/workflow-info-context";
import { useEffect } from "react";
import {
  ClimateDataResponse,
  DeviceResponse,
  OperationResponse,
} from "@/types/api";
import {
  getClimateDatas,
  getDevices,
  getOperations,
} from "@/mocks/workflow_api";

export default function Workflow() {
  const [workflowInfo, setWorkflowInfo] = useWorkflowInfo();

  useEffect(() => {
    const fetchWorkflowData = async () => {
      const devicesRes: DeviceResponse[] = await getDevices();
      const climateDataRes: ClimateDataResponse[] = await getClimateDatas();
      const operationsRes: OperationResponse[] = await getOperations();

      workflowInfo.devices = devicesRes;
      workflowInfo.climate_data = climateDataRes;
      workflowInfo.operations = operationsRes;

      setWorkflowInfo(workflowInfo);
    };

    fetchWorkflowData();
  }, [workflowInfo, setWorkflowInfo]);

  return (
    <WorkflowInfoProvider>
      <Navigation>
        <Box sx={{ padding: "16px" }}>
          <Box sx={{ display: "flex" }}>
            <Box sx={{ width: "150px" }}>
              <FormControl fullWidth>
                <InputLabel id={`workflow-select-label`} size="small">
                  ワークフローを選択
                </InputLabel>
                <Select
                  value={"test"}
                  labelId={"workflow-select-label"}
                  id={"workflow-select"}
                  size="small"
                  // onChange={}
                  label="ワークフロー選択"
                  // disabled={deviceOperations.length === 0}
                >
                  {/* {deviceOperations.map((data) => (
                  <MenuItem key={data.id} value={data.id}>
                  {data.name}
                  </MenuItem>
                  ))} */}
                </Select>
              </FormControl>
            </Box>
            <Button>新規作成</Button>
            <Button>保存する</Button>
          </Box>
          <WorkflowWrapper workflowID={1} />
        </Box>
      </Navigation>
    </WorkflowInfoProvider>
  );
}
