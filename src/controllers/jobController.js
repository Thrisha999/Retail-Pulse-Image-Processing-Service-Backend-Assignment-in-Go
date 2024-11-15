import { v4 as uuidv4 } from 'uuid';
import { JobService } from '../services/jobService.js';
import { ImageProcessor } from '../services/imageProcessor.js';

const jobService = new JobService();
const imageProcessor = new ImageProcessor();

export const submitJob = async (req, res) => {
  const { count, visits } = req.body;

  if (!count || !visits || count !== visits.length) {
    return res.status(400).json({ error: "Invalid request parameters" });
  }

  const jobId = uuidv4();
  jobService.createJob(jobId, { count, visits });

  // Process job asynchronously
  imageProcessor.processJob(jobId, visits, jobService);

  res.status(201).json({ job_id: jobId });
};

export const getJobStatus = (req, res) => {
  const { jobid } = req.query;

  if (!jobid) {
    return res.status(400).json({});
  }

  const job = jobService.getJob(jobid);
  if (!job) {
    return res.status(400).json({});
  }

  const response = {
    status: job.status,
    job_id: jobid
  };

  if (job.status === 'failed' && job.errors) {
    response.error = job.errors;
  }

  res.json(response);
};