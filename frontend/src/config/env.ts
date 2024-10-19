import { z } from "zod";

const createEnv = () => {
  const envSchema = z.object({
    API_URL: z.string().url(),
  });

  const envVars = Object.entries(import.meta.env).reduce<
    Record<string, string>
  >((vars, curr) => {
    const [key, value] = curr;
    if (key.startsWith("VITE_")) {
      vars[key.replace("VITE_", "")] = value;
    }
    return vars;
  }, {});

  const parsedEnv = envSchema.safeParse(envVars);

  if (!parsedEnv.success) {
    throw new Error(
      `Invalid env provided.
The following variables are missing or invalid:
${Object.entries(parsedEnv.error.flatten().fieldErrors)
  .map(([k, v]) => `- ${k}: ${v}`)
  .join("\n")}
`
    );
  }

  return parsedEnv.data;
};

export const env = createEnv();
