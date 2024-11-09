import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type DnDState = string | null;

type DnDContextType = [DnDState, Dispatch<SetStateAction<DnDState>>];

const DnDContext = createContext<DnDContextType>([null, () => {}]);

export const DnDProvider = ({ children }: { children: React.ReactNode }) => {
  const [type, setType] = useState<DnDState>(null);

  return (
    <DnDContext.Provider value={[type, setType]}>
      {children}
    </DnDContext.Provider>
  );
};

export const useDnD = () => {
  return useContext(DnDContext);
};

export default DnDContext;
