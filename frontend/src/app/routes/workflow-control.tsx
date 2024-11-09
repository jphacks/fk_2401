import { Box, Button } from "@mui/material";
import { Navigation } from "@/layouts/navigation";
import { Workflow } from "@/components/workflow-control/workflow";

export default function WorkflowControl() {
  return (
    <Navigation>
      <Box sx={{ padding: "16px" }}>
        <Box sx={{ display: "flex" }}>
          <Button>新規作成</Button>
          <Button>保存する</Button>
        </Box>
        <Workflow />
      </Box>
    </Navigation>
  );
}
