import {
  Box,
  Button,
  Divider,
  Modal,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";

const modalStyle = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  boxShadow: 24,
  p: 4,
};

export function SaveWorkflowButton() {
  const [modalOpen, setModalOpen] = useState(false);

  const handleOpen = () => setModalOpen(true);
  const handleClose = () => setModalOpen(false);

  return (
    <div>
      <Button variant="contained" size="small" onClick={handleOpen}>
        保存する
      </Button>
      <Modal
        open={modalOpen}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={modalStyle}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            ワークフローを保存する
          </Typography>
          <Divider />
          <Box sx={{ my: 2 }}>
            <TextField type="text" size="small" label="Name" />
          </Box>
          <Divider />
          <Box sx={{ mt: 1, display: "flex", gap: 1 }}>
            <Button variant="contained" size="small" onClick={handleClose}>
              保存
            </Button>
            <Button variant="outlined" size="small" onClick={handleClose}>
              キャンセル
            </Button>
          </Box>
        </Box>
      </Modal>
    </div>
  );
}
