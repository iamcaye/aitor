import { initServer } from "./config/express";
import { config } from "./config/env";

initServer(() => {
    console.log(`Server is running on port ${config.PORT}`);
});
