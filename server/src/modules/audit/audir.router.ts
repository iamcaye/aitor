import { Router } from "express";
import auditController from "./audit.controller";


const auditRouter = Router();

auditRouter.post('/', auditController.makeAudition);

export default auditRouter;
