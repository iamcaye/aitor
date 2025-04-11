import { configDotenv } from "dotenv";
import { z } from "zod";

configDotenv()

const envSchema = z.object({
    PORT: z.string(),
});

export const config = envSchema.parse(process.env)
