import { Box, Tabs, Tab } from "@mui/material";
import { Navigation } from "@/layouts/navigation";
import { Workflow } from "@/components/workflow-control/workflow";
import { HouseResponse } from "@/types/api";
import { getHouses } from "@/mocks/setting_device_api";
import { useState, useEffect } from "react";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

export default function WorkflowControlTabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`setting-device-tabpanel-${index}`}
      aria-labelledby={`setting-device-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `setting-device-tabpanel-${index}`,
    "aria-controls": `setting-device-tabpanel-${index}`,
  };
}

export function WorkflowControl() {
  const [houses, setHouses] = useState<HouseResponse[]>([]);
  const [selectedTab, setSelectedTab] = useState(0);

  useEffect(() => {
    const fetchHouseAndDevices = async () => {
      const housesRes: HouseResponse[] = await getHouses();

      setHouses(housesRes);
    };

    fetchHouseAndDevices();
  }, []);

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setSelectedTab(newValue);
  };

  return (
    <Navigation>
      <Box sx={{ padding: "16px" }}>
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
          <WorkflowControlTabPanel
            key={index}
            index={index}
            value={selectedTab}
          >
            <Workflow />
          </WorkflowControlTabPanel>
        ))}
      </Box>
    </Navigation>
  );
}
