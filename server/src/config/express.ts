import express, { Express } from "express";
import { config } from "./env";
import morgan from "morgan";
import indexRouter from "../modules/index.router";

export const initServer = (callback: (app: Express) => void): void => {
    const app: Express = express();
    app.use(morgan("dev"))

    app.use(express.json());
    app.use(express.urlencoded({ extended: true }));

    app.use('/api', indexRouter);

    app.listen(config.PORT, () => callback(app));
}
