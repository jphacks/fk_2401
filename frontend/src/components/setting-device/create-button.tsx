import {
  IconButton,
  Box,
  Button,
  Typography,
  Modal,
  TextField,
  InputAdornment,
  FormControlLabel,
  Switch,
  Collapse,
  Slider,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Divider,
} from "@mui/material";
import { useState, useEffect } from "react";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import { SelectChangeEvent } from "@mui/material";
import { ClimateDataResponse } from "../../types/api";
import { getClimateDatas } from "../../mocks/setting_device_api";

export function CreateDeviceButton() {
  const [modalOpen, setModalOpen] = useState(false);
  const [climateData, setClimateData] = useState("");
  const [setPointChecked, setSetPointChecked] = useState(false);
  const [timerChecked, setTimerChecked] = useState(false);
  const [fetchedClimateDatas, setFetchedClimateDatas] = useState<
    ClimateDataResponse[]
  >([]);

  useEffect(() => {
    const climateDataRes: ClimateDataResponse[] = getClimateDatas();

    setFetchedClimateDatas(climateDataRes);
  }, []);

  const handleClimateDataChange = (event: SelectChangeEvent) => {
    setClimateData(event.target.value as string);
  };
  const handleSetPointChange = () => {
    setSetPointChecked((prev) => !prev);
  };
  const handleTimerChange = () => {
    setTimerChecked((prev) => !prev);
  };

  const handleOpen = () => setModalOpen(true);
  const handleClose = () => setModalOpen(false);

  return (
    <>
      <IconButton size="large" color="primary" onClick={handleOpen}>
        <AddCircleOutlineIcon fontSize="large" />
      </IconButton>
      <Modal
        open={modalOpen}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={modalStyle}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            デバイス追加
          </Typography>
          <Divider />
          <Typography sx={{ my: 2 }}>
            <TextField label="デバイス名" size="small" />
          </Typography>
          <Typography>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label" size="small">
                気象データ
              </InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={climateData}
                label="気象データ"
                onChange={handleClimateDataChange}
                size="small"
              >
                {fetchedClimateDatas.map((data) => (
                  <MenuItem value={data.climateData}>
                    {data.climateData}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Typography>
          <Typography sx={{ my: 2 }}>
            <FormControlLabel
              control={
                <Switch
                  checked={setPointChecked}
                  onChange={handleSetPointChange}
                />
              }
              label={climateData ? `設定${climateData}` : "設定気象データ"}
              disabled={climateData == ""}
            />
            <Collapse in={setPointChecked}>
              <TextField
                type="number"
                size="small"
                label={`設定${climateData}`}
                inputProps={{ step: "0.1" }}
                slotProps={{
                  input: {
                    endAdornment: (
                      <InputAdornment position="end">{}</InputAdornment>
                    ),
                  },
                }}
              />
            </Collapse>
          </Typography>
          <Typography sx={{ my: 2 }}>
            <FormControlLabel
              control={
                <Switch checked={timerChecked} onChange={handleTimerChange} />
              }
              label="タイマー"
              disabled={climateData == ""}
            />
            <Collapse in={timerChecked}>
              <Box
                sx={{ px: 3, py: 4, display: "flex", justifyContent: "center" }}
              >
                <TimerSlider duration={1} />
              </Box>
            </Collapse>
          </Typography>
          <Divider />
          <Box sx={{ mt: 1, display: "flex" }}>
            <Button size="small">保存</Button>
            <Button size="small">キャンセル</Button>
          </Box>
        </Box>
      </Modal>
    </>
  );
}

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

interface TimerSliderProps {
  duration?: number;
}

function TimerSlider({ duration }: TimerSliderProps) {
  const marks = [
    { value: 1, label: "1時間" },
    { value: 12, label: "12時間" },
  ];

  function valuetext(value: number) {
    return `${value}時間`;
  }

  return (
    <Box sx={{ width: 300 }}>
      <Slider
        defaultValue={duration !== undefined ? duration : 1}
        aria-label="Timer duration"
        getAriaValueText={valuetext}
        min={1}
        max={12}
        marks={marks}
        valueLabelDisplay="on"
      />
    </Box>
  );
}
