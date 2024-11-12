export interface Condition {
  climate_data_id: number;
  comp_ope_id: number;
  set_point: number;
  operations_id: number;
}

export interface Workflow {
  device_id: number;
  condition_operations: Condition[];
}
