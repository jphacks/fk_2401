import { WorkflowData } from "@/types/workflow";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type nodeInfoState = WorkflowData;

type nodeInfoContextType = [
  nodeInfoState,
  Dispatch<SetStateAction<nodeInfoState>>
];

const NodeInfoContext = createContext<nodeInfoContextType>([
  {
    device_id: 0,
    condition_operations: [],
  },
  () => {},
]);

export const NodeInfoProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [nodeInfo, setNodeInfo] = useState<nodeInfoState>({
    device_id: 0,
    condition_operations: [],
  });

  return (
    <NodeInfoContext.Provider value={[nodeInfo, setNodeInfo]}>
      {children}
    </NodeInfoContext.Provider>
  );
};

export const useNodeInfo = () => {
  return useContext(NodeInfoContext);
};

export default NodeInfoContext;
