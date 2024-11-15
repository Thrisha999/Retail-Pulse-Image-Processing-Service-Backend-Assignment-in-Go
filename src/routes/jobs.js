import express from 'express';
import { submitJob, getJobStatus } from '../controllers/jobController.js';

const router = express.Router();

router.post('/submit', submitJob);
router.get('/status', getJobStatus);

export { router as jobRouter };