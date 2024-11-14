export interface Condition {
  climate_data_id: number;
  comp_ope_id: number;
  set_point: number;
}

export interface ConditionOperation {
  condition: Condition;
  operation_id: number;
}

export interface WorkflowData {
  device_id: number;
  condition_operations: ConditionOperation[];
}