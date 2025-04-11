import { Request, Response } from 'express';
import { decodeBrotli } from '../../utils/compression';

const makeAudition = (req: Request, res: Response) => {
    console.log(req.body);
    const body = req.body;

    res.send("OK!");
}

export default {
    makeAudition
}
