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
import { ClimateDataResponse, CreateDeviceRequest } from "@/types/api";
import { getClimateDatas } from "@/mocks/setting-device-api";
// import { getClimateDatas } from "@/features/api/climate-data/get-climate-data";
import { createDevice } from "@/features/api/device/create-device";

interface CreateDeviceButtonProps {
  houseID: number;
}

export function CreateDeviceButton(props: CreateDeviceButtonProps) {
  const { houseID } = props;

  const [modalOpen, setModalOpen] = useState<boolean>(false);
  const [setPointChecked, setSetPointChecked] = useState<boolean>(false);
  const [timerChecked, setTimerChecked] = useState<boolean>(false);
  const [fetchedClimateDatas, setFetchedClimateDatas] = useState<
    ClimateDataResponse[]
  >([]);
  const [selectedClimateData, setSelectedClimateData] = useState<string>("");

  // States for device creation form inputs
  const [deviceNameInput, setDeviceNameInput] = useState<string>("");
  const [climateDataInput, setClimateDataInput] = useState<number>(0);
  const [setPointInput, setSetPointInput] = useState<number>(0);
  const [durationInput, setDurationInput] = useState<number>(0);

  const formReset = () => {
    setSetPointChecked(false);
    setTimerChecked(false);
    setDeviceNameInput("");
    setSelectedClimateData("");
    setClimateDataInput(0);
    setSetPointInput(0);
    setDurationInput(0);
  };

  useEffect(() => {
    const fetchClimateDatas = async () => {
      const climateDataRes: ClimateDataResponse[] = await getClimateDatas();
      setFetchedClimateDatas(climateDataRes);
    };

    fetchClimateDatas();
  }, []);

  // Handlers for device creation form inputs
  const handleDeviceNameChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const deviceName: string = event.target.value as string;
    setDeviceNameInput(deviceName);
  };
  const handleClimateDataChange = (event: SelectChangeEvent) => {
    const climateData = event.target.value;

    const climateDataRec = fetchedClimateDatas.find(
      (data) => data.climate_data === climateData
    );

    setSelectedClimateData(climateData);

    if (climateDataRec) {
      setClimateDataInput(climateDataRec.id);
    }
  };
  const handleSetPointSwitchChange = () => {
    setSetPointChecked((prev) => !prev);
  };
  const handleSetPointChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const setPoint: number = parseFloat(event.target.value);
    setSetPointInput(setPoint);
  };
  const handleTimerSwitchChange = () => {
    setTimerChecked((prev) => !prev);
  };
  const handleDurationChange = (
    _event: Event,
    newDurationInput: number | number[]
  ) => {
    setDurationInput(newDurationInput as number);
  };

  const handleModalOpen = () => setModalOpen(true);
  const handleModalClose = () => setModalOpen(false);

  const handleSendForm = async () => {
    const req: CreateDeviceRequest = {
      device_name: deviceNameInput,
      climate_data_id: climateDataInput,
      set_point: setPointInput,
      duration: durationInput,
    };

    console.log(req);

    formReset();

    try {
      await createDevice(houseID, req);
      handleModalClose();
    } catch (error) {
      console.error("Error creating device:", error);
    }

    handleModalClose();
  };

  const handleCancel = () => {
    formReset();
    handleModalClose();
  };

  return (
    <>
      <IconButton size="large" color="primary" onClick={handleModalOpen}>
        <AddCircleOutlineIcon fontSize="large" />
      </IconButton>
      <Modal
        open={modalOpen}
        onClose={handleModalClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={modalStyle}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            デバイス追加
          </Typography>
          <Divider />
          <Box sx={{ my: 2 }}>
            <TextField
              label="デバイス名"
              size="small"
              value={deviceNameInput}
              onChange={handleDeviceNameChange}
            />
          </Box>
          <Box>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label" size="small">
                気象データ
              </InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={selectedClimateData}
                label="気象データ"
                onChange={handleClimateDataChange}
                size="small"
              >
                {fetchedClimateDatas.map((data) => (
                  <MenuItem key={data.id} value={data.climate_data}>
                    {data.climate_data}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Box>
          <Box sx={{ my: 2 }}>
            <FormControlLabel
              control={
                <Switch
                  checked={setPointChecked}
                  onChange={handleSetPointSwitchChange}
                />
              }
              label={
                selectedClimateData
                  ? `設定${selectedClimateData}`
                  : "設定気象データ"
              }
              disabled={selectedClimateData == ""}
            />
            <Collapse in={setPointChecked}>
              <TextField
                type="number"
                size="small"
                label={`設定${selectedClimateData}`}
                inputProps={{ step: "0.1" }}
                slotProps={{
                  input: {
                    endAdornment: (
                      <InputAdornment position="end">{}</InputAdornment>
                    ),
                  },
                }}
                value={setPointInput}
                onChange={handleSetPointChange}
              />
            </Collapse>
          </Box>
          <Box sx={{ my: 2 }}>
            <FormControlLabel
              control={
                <Switch
                  checked={timerChecked}
                  onChange={handleTimerSwitchChange}
                />
              }
              label="タイマー"
              disabled={selectedClimateData == ""}
            />
            <Collapse in={timerChecked}>
              <Box
                sx={{ px: 3, py: 4, display: "flex", justifyContent: "center" }}
              >
                <TimerSlider
                  durationInput={durationInput}
                  handleDurationChange={handleDurationChange}
                />
              </Box>
            </Collapse>
          </Box>
          <Divider />
          <Box sx={{ mt: 1, display: "flex" }}>
            <Button size="small" onClick={handleSendForm}>
              保存
            </Button>
            <Button size="small" onClick={handleCancel}>
              キャンセル
            </Button>
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
  durationInput?: number;
  handleDurationChange: (event: Event, value: number | number[]) => void;
}

function TimerSlider({
  durationInput,
  handleDurationChange,
}: TimerSliderProps) {
  const marks = [
    { value: 1, label: "1時間" },
    { value: 12, label: "12時間" },
  ];

  const valuetext = (value: number) => {
    return `${value}時間`;
  };

  return (
    <Box sx={{ width: 300 }}>
      <Slider
        defaultValue={durationInput !== undefined ? durationInput : 1}
        aria-label="Timer duration"
        getAriaValueText={valuetext}
        min={1}
        max={12}
        marks={marks}
        valueLabelDisplay="on"
        value={durationInput}
        onChange={handleDurationChange}
      />
    </Box>
  );
}
