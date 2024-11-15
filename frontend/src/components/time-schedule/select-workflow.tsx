import { WorkflowResponse } from "@/types/api";
import { FormControl, InputLabel, Select, MenuItem } from "@mui/material";
import { useState } from "react";
import { SelectChangeEvent } from "@mui/material";

interface WorkflowSelectProps {
  initialWorkflow: WorkflowResponse | null;
  workflows: WorkflowResponse[];
  index: number;
  onSelectChange: (index: number, updatedData: WorkflowResponse) => void;
}

export const WorkflowSelect = (props: WorkflowSelectProps) => {
  const { initialWorkflow, workflows, index, onSelectChange } = props;

  const [options, setOption] = useState<WorkflowResponse[]>(workflows);
  const [selectedWorkflow, setSelectedWorkflow] =
    useState<WorkflowSettingResponse | null>(initialWorkflow);

  const handleWorkflowChange = (event: SelectChangeEvent) => {
    const workflowID = parseInt(event.target.value);
    const workflowRec = options.find((data) => data.id === workflowID);

    if (workflowRec) {
      setSelectedWorkflow(workflowRec);
    }
  };

  return (
    <FormControl fullWidth>
      <InputLabel id="workflow-select-label">ワークフロー</InputLabel>
      <Select
        labelId="workflow-select-label"
        id="workflow-select"
        value={String(selectedWorkflow?.id)}
        label="ワークフロー"
        size="small"
        onChange={handleWorkflowChange}
      >
        {options.map((data) => (
          <MenuItem key={data.id} value={data.id}>
            {data.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};
