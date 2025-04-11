import { Router } from 'express';
import auditRouter from './audit/audir.router';

const indexRouter = Router();

indexRouter.get('/', (_, res) => {
    res.send('Hello API!');
});

indexRouter.use('/audit', auditRouter);

export default indexRouter;
