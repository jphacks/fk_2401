import { Box, Tab, Tabs } from "@mui/material";
import { TimeTable } from "./time-table";
import { Navigation } from "@/layouts/navigation";
import { useState, useEffect } from "react";
import {
  HouseResponse,
  TimeScheduleResponse,
  WorkflowResponse,
} from "@/types/api";
import { getHouses } from "@/mocks/setting-device-api";
import {
  getTimeSchedules,
  getWorkflows,
} from "@/mocks/setting-time-schedule-api";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function a11yProps(index: number) {
  return {
    id: `setting-device-tabpanel-${index}`,
    "aria-controls": `setting-device-tabpanel-${index}`,
  };
}

function TimeScheduleTabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <Box
      role="tabpanel"
      hidden={value !== index}
      id={`time-schedule-tabpanel-${index}`}
      aria-labelledby={`time-schedule-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </Box>
  );
}

export function TimeSchedule() {
  const [houses, setHouses] = useState<HouseResponse[]>([]);
  const [selectedTab, setSelectedTab] = useState<number>(0);
  const [timeScheduleMap, setTimeScheduleMap] = useState<
    Map<number, TimeScheduleResponse[]>
  >(new Map());
  const [selectBoxWorkflows, setSelectBoxWorkflows] = useState<
    WorkflowResponse[]
  >([]);

  useEffect(() => {
    const fetchInitData = async () => {
      const housesRes: HouseResponse[] = await getHouses();
      const timeScheduleMap: Map<number, TimeScheduleResponse[]> = new Map();
      for (const house of housesRes) {
        const timeScheduleRes: TimeScheduleResponse[] = await getTimeSchedules(
          house.id
        );
        timeScheduleMap.set(house.id, timeScheduleRes);
      }
      const workflowsRes: WorkflowResponse[] = await getWorkflows();

      setHouses(housesRes);
      setTimeScheduleMap(timeScheduleMap);
      setSelectBoxWorkflows(workflowsRes);
    };

    fetchInitData();
  }, []);

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setSelectedTab(newValue);
  };

  return (
    <Navigation>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs
          value={selectedTab}
          onChange={handleTabChange}
          aria-label="basic tabs example"
        >
          {houses.map((house, index) => (
            <Tab key={index} label={house.name} {...a11yProps(index)} />
          ))}
        </Tabs>
      </Box>
      {houses.map((house, index) => (
        <TimeScheduleTabPanel key={index} index={index} value={selectedTab}>
          <TimeTable
            initialSchedules={timeScheduleMap.get(house.id) || []}
            workflows={selectBoxWorkflows}
          />
        </TimeScheduleTabPanel>
      ))}
    </Navigation>
  );
}
