import {
  Button,
  Box,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Paper,
} from "@mui/material";
import { TimeTableRow } from "./time-table-row";
import { TimeScheduleResponse, WorkflowResponse } from "@/types/api";
import { useState } from "react";

interface TimeTableProps {
  initialSchedules: TimeScheduleResponse[];
  workflows: WorkflowResponse[];
}

export function TimeTable(props: TimeTableProps) {
  const { initialSchedules, workflows } = props;
  const [timeSchedules, setTimeSchedules] =
    useState<TimeScheduleResponse[]>(initialSchedules);

  const handleRowChange = (
    index: number,
    updatedData: TimeScheduleResponse
  ) => {
    const newSchedules = [...timeSchedules];
    newSchedules[index] = updatedData;
    setTimeSchedules(newSchedules);
  };

  const handleAddRow = () => {
    if (timeSchedules.length < 8) {
      setTimeSchedules([
        ...timeSchedules,
        { start_time: "", end_time: "", workflows: [] },
      ]);
    }
  };

  const handleRemoveRow = () => {
    if (timeSchedules.length > 1) {
      setTimeSchedules(timeSchedules.slice(0, -1));
    }
  };

  return (
    <Box>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell sx={{ width: "20%" }}>開始時間</TableCell>
              <TableCell sx={{ width: "20%" }}>終了時間</TableCell>
              <TableCell sx={{ width: "60%" }}>デバイス設定</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {timeSchedules.length == 0 ? (
              <TimeTableRow
                timeSchedule={null}
                workflows={workflows}
                index={0}
                onRowChange={handleRowChange}
              />
            ) : (
              timeSchedules.map((schedule, index) => (
                <TimeTableRow
                  key={index}
                  timeSchedule={schedule}
                  workflows={workflows}
                  index={index}
                  onRowChange={handleRowChange}
                />
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
      <Box sx={{ mt: 1, display: "flex", gap: 1 }}>
        <Button size="small" variant="contained" onClick={handleAddRow}>
          行を追加
        </Button>
        <Button size="small" variant="outlined" onClick={handleRemoveRow}>
          行を削除
        </Button>
      </Box>
    </Box>
  );
}
