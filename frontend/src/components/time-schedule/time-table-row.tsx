import { TableRow, TableCell, TextField, IconButton } from "@mui/material";
import Grid from "@mui/material/Grid2";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import dayjs from "dayjs";
import { LocalizationProvider, TimePicker } from "@mui/x-date-pickers";
import { WorkflowSelect } from "./select-workflow";
import AddIcon from "@mui/icons-material/Add";
import RemoveIcon from "@mui/icons-material/Remove";
import { WorkflowResponse, TimeScheduleResponse } from "@/types/api";
import { useState } from "react";

interface TimeTableRowProps {
  timeSchedule: TimeScheduleResponse | null;
  workflows: WorkflowResponse[];
  index: number;
  onRowChange: (index: number, updatedData: TimeScheduleResponse) => void;
}

export function TimeTableRow(props: TimeTableRowProps) {
  const { timeSchedule, workflows, index, onRowChange } = props;
  const [selectedWorkflows, setSelectedWorkflows] = useState<
    WorkflowResponse[]
  >(timeSchedule?.workflows || [{ id: 0, name: "" }]);

  const handleStartTimeChange = (value: string) => {
    if (value !== undefined) {
      if (timeSchedule) {
        timeSchedule.start_time = value;

        onRowChange(index, timeSchedule);
      }
    }
  };

  const handleEndTimeChange = (value: string) => {
    if (value !== undefined) {
      if (timeSchedule) {
        timeSchedule.end_time = value;

        onRowChange(index, timeSchedule);
      }
    }
  };

  const handleSelectChange = (index: number, updatedData: WorkflowResponse) => {
    if (timeSchedule) {
      const newWorkflows = [...timeSchedule.workflows];
      newWorkflows[index] = updatedData;
      setSelectedWorkflows(newWorkflows);
    }
  };

  const handleAddWorkflow = () => {
    if (selectedWorkflows.length >= 4) {
      return;
    }
    const newWorkflows = [...selectedWorkflows, { id: 0, name: "" }];
    setSelectedWorkflows(newWorkflows);

    if (timeSchedule) {
      timeSchedule.workflows = newWorkflows;
      onRowChange(index, timeSchedule);
    }
  };

  const handleRemoveWorkflow = () => {
    if (selectedWorkflows.length <= 1) {
      return;
    }

    const newWorkflows: WorkflowResponse[] = selectedWorkflows.slice(0, -1);
    setSelectedWorkflows(newWorkflows);

    if (timeSchedule) {
      timeSchedule.workflows = newWorkflows;
      onRowChange(index, timeSchedule);
    }
  };

  return (
    <TableRow
      sx={{
        "&:hover": {
          backgroundColor: "rgb(245, 245, 245)",
        },
      }}
    >
      <TableCell>
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <TimePicker
            label="開始時間"
            ampm={false}
            value={
              timeSchedule?.start_time
                ? dayjs(timeSchedule.start_time, "HH:mm")
                : null
            }
            onChange={(value) => {
              const startTime = value?.format("HH:mm");
              if (!startTime) {
                return;
              }
              handleStartTimeChange(startTime);
            }}
            slots={{
              textField: (props) => (
                <TextField
                  {...props}
                  size="small"
                  InputLabelProps={{ shrink: true }}
                />
              ),
            }}
            sx={{ marginRight: "8px" }}
          />
        </LocalizationProvider>
      </TableCell>
      <TableCell>
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <TimePicker
            label="終了時間"
            ampm={false}
            value={
              timeSchedule?.end_time
                ? dayjs(timeSchedule.end_time, "HH:mm")
                : null
            }
            onChange={(value) => {
              const endTime = value?.format("HH:mm");
              if (!endTime) {
                return;
              }
              handleEndTimeChange(endTime);
            }}
            slots={{
              textField: (props) => (
                <TextField
                  {...props}
                  size="small"
                  InputLabelProps={{ shrink: true }}
                />
              ),
            }}
            sx={{ marginRight: "8px" }}
          />
        </LocalizationProvider>
      </TableCell>
      <TableCell>
        <Grid container spacing={2}>
          <Grid size={9} container spacing={2}>
            {selectedWorkflows.length == 0 ? (
              <Grid key={0} size={3}>
                <WorkflowSelect
                  initialWorkflow={null}
                  index={0}
                  workflows={workflows}
                  onSelectChange={handleSelectChange}
                />
              </Grid>
            ) : (
              selectedWorkflows.map((workflow, index) => (
                <Grid key={index} size={3}>
                  <WorkflowSelect
                    initialWorkflow={workflow}
                    index={index}
                    workflows={workflows}
                    onSelectChange={handleSelectChange}
                  />
                </Grid>
              ))
            )}
          </Grid>
          <Grid size={2} sx={{ marginLeft: "auto" }}>
            <IconButton
              size="small"
              sx={{
                width: 24,
                height: 24,
                borderRadius: 1,
                color: "#4CAF50",
                border: "2px solid #4CAF50",
                "&:hover": {
                  backgroundColor: "#70CC70",
                  color: "#FFF",
                  border: "2px solid #70CC70",
                },
              }}
              onClick={handleRemoveWorkflow}
            >
              <RemoveIcon />
            </IconButton>
            <IconButton
              size="small"
              sx={{
                width: 24,
                height: 24,
                ml: 1,
                borderRadius: 1,
                color: "#FFF",
                backgroundColor: "#4CAF50",
                "&:hover": {
                  backgroundColor: "#70BB70",
                  color: "#FFF",
                },
              }}
              onClick={handleAddWorkflow}
            >
              <AddIcon />
            </IconButton>
          </Grid>
        </Grid>
      </TableCell>
    </TableRow>
  );
}
