import { Box, Typography } from "@mui/material";
import { Handle, Position, Node, NodeProps } from "@xyflow/react";
import React from "react";
import { AddNodeFunction } from "../workflow";

interface BeginWorkflowNodeData {
  [key: string]: unknown;
  addNode: AddNodeFunction;
}

type BeginWorkflowNodePropsType = Node<BeginWorkflowNodeData>;

type BeginWorkflowNodeProps = NodeProps<BeginWorkflowNodePropsType>;

export const BeginWorkflowNode = ({ id, data }: BeginWorkflowNodeProps) => {
  const { addNode } = data;

  return (
    <Box
      sx={{
        width: "100%",
        height: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Box>
        <Typography variant="h6">Begin Workflow</Typography>
      </Box>
      <Handle
        position={Position.Right}
        type="source"
        onClick={(event) => {
          event.stopPropagation();
          addNode(id);
        }}
      />
    </Box>
  );
};
