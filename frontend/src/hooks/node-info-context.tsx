import { Workflow } from "@/types/node";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type nodeInfoState = Workflow | null;

type nodeInfoContextType = [
  nodeInfoState,
  Dispatch<SetStateAction<nodeInfoState>>
];

const NodeInfoContext = createContext<nodeInfoContextType>([null, () => {}]);

export const NodeInfoProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [workflowInfo, setWorkflowInfo] = useState<nodeInfoState>(null);

  return (
    <NodeInfoContext.Provider value={[workflowInfo, setWorkflowInfo]}>
      {children}
    </NodeInfoContext.Provider>
  );
};

export const useNodeInfo = () => {
  return useContext(NodeInfoContext);
};

export default NodeInfoContext;
