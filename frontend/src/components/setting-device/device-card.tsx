import {
  Card,
  CardActions,
  CardContent,
  Button,
  Box,
  Typography,
  Modal,
  TextField,
  InputAdornment,
  Slider,
  FormControlLabel,
  Switch,
  Collapse,
  Divider,
} from "@mui/material";
import { useState } from "react";

interface DeviceProps {
  name: string;
  setPoint?: number;
  duration?: number;
  climateData: string;
  unit: string;
}

export function DeviceCard(props: DeviceProps) {
  const { name, setPoint, duration, climateData, unit } = props;

  return (
    <Card raised={true} sx={{ minWidth: 275 }}>
      <CardContent>
        <Typography gutterBottom sx={{ color: "text.main" }}>
          {name}
        </Typography>
        <Box>
          <Typography gutterBottom sx={{ mt: 2, color: "text.secondary" }}>
            設定{climateData}
          </Typography>
          <Box sx={{ display: "flex", justifyContent: "center" }}>
            <Typography variant="h4" component="div">
              {setPoint !== undefined ? `${setPoint}` : `--`}
            </Typography>
            <Typography
              sx={{ pb: 1, color: "text.secondary", alignSelf: "flex-end" }}
            >
              {unit}
            </Typography>
          </Box>
        </Box>
        {duration !== undefined && (
          <Box sx={{ mt: 2, display: "flex", alignItems: "center" }}>
            <Typography component="span" sx={{ color: "text.secondary" }}>
              残り動作時間
            </Typography>
            <Typography component="span" sx={{ ml: 3, mr: 1 }}>
              {duration}
            </Typography>
            <Typography component="span" sx={{ color: "text.secondary" }}>
              時間
            </Typography>
          </Box>
        )}
      </CardContent>
      <CardActions>
        <SettingModalButton
          name={name}
          setPoint={setPoint}
          duration={duration}
          climateData={climateData}
          unit={unit}
        />
      </CardActions>
    </Card>
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

function SettingModalButton(props: DeviceProps) {
  const { name, setPoint, duration, climateData, unit } = props;
  const [modalOpen, setModalOpen] = useState(false);
  const [timerChecked, setTimerChecked] = useState(false);

  const handleTimerChange = () => {
    setTimerChecked((prev) => !prev);
  };
  const handleOpen = () => setModalOpen(true);
  const handleClose = () => setModalOpen(false);

  return (
    <div>
      <Button size="small" onClick={handleOpen}>
        設定
      </Button>
      <Modal
        open={modalOpen}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={modalStyle}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            {name}
          </Typography>
          <Divider />
          <Box sx={{ my: 2 }}>
            <TextField
              {...(setPoint !== undefined && { defaultValue: setPoint })}
              type="number"
              size="small"
              inputProps={{ step: "0.1" }}
              label={`設定${climateData}`}
              slotProps={{
                input: {
                  endAdornment: (
                    <InputAdornment position="end">{unit}</InputAdornment>
                  ),
                },
              }}
            />
          </Box>
          <Box sx={{ my: 2 }}>
            <FormControlLabel
              control={
                <Switch checked={timerChecked} onChange={handleTimerChange} />
              }
              label="タイマー"
            />
            <Collapse in={timerChecked}>
              <Box
                sx={{ px: 3, py: 4, display: "flex", justifyContent: "center" }}
              >
                <TimerSlider duration={duration} />
              </Box>
            </Collapse>
          </Box>
          <Divider />
          <Box sx={{ mt: 1, display: "flex" }}>
            <Button size="small">保存</Button>
            <Button size="small">キャンセル</Button>
          </Box>
        </Box>
      </Modal>
    </div>
  );
}

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
