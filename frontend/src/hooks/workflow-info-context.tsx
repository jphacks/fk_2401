import {
  ClimateDataResponse,
  DeviceResponse,
  OperationResponse,
} from "@/types/api";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type workflowInfoState = {
  devices: DeviceResponse[];
  climate_data: ClimateDataResponse[];
  operations: OperationResponse[];
};

type workflowInfoContextType = [
  workflowInfoState,
  Dispatch<SetStateAction<workflowInfoState>>
];

const WorkflowInfoContext = createContext<workflowInfoContextType>([
  {
    devices: [],
    climate_data: [],
    operations: [],
  },
  () => {},
]);

export const WorkflowInfoProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [workflowInfo, setWorkflowInfo] = useState<workflowInfoState>({
    devices: [],
    climate_data: [],
    operations: [],
  });

  return (
    <WorkflowInfoContext.Provider value={[workflowInfo, setWorkflowInfo]}>
      {children}
    </WorkflowInfoContext.Provider>
  );
};

export const useWorkflowInfo = () => {
  return useContext(WorkflowInfoContext);
};

export default WorkflowInfoContext;
