import Axios from "axios";
import { env } from "@/config/env";

export const apiClient = Axios.create({
  baseURL: env.API_URL,
});
