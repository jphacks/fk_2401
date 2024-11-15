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
  >(timeSchedule?.workflows || []);

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

  return (
    <TableRow>
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
              <Grid size={3}>
                <WorkflowSelect
                  key={0}
                  initialWorkflow={null}
                  index={0}
                  workflows={workflows}
                  onSelectChange={handleSelectChange}
                />
              </Grid>
            ) : (
              selectedWorkflows.map((workflow, index) => (
                <Grid size={3}>
                  <WorkflowSelect
                    key={index}
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
                borderRadius: 2,
                color: "#F44336",
                border: "2px solid #F04134",
                "&:hover": {
                  backgroundColor: "#E57373",
                  color: "#FFF",
                  border: "2px solid #E57373",
                },
              }}
            >
              <RemoveIcon />
            </IconButton>
            <IconButton
              size="small"
              sx={{
                ml: 1,
                borderRadius: 2,
                color: "#FFF",
                backgroundColor: "#4CAF50",
                "&:hover": {
                  backgroundColor: "#70BB70",
                  color: "#FFF",
                },
              }}
            >
              <AddIcon />
            </IconButton>
          </Grid>
        </Grid>
      </TableCell>
    </TableRow>
  );
}
