export class JobService {
  constructor() {
    this.jobs = new Map();
  }

  createJob(jobId, jobData) {
    this.jobs.set(jobId, {
      ...jobData,
      status: 'ongoing',
      errors: []
    });
    return jobId;
  }

  getJob(jobId) {
    return this.jobs.get(jobId);
  }

  updateJobStatus(jobId, status, errors = []) {
    const job = this.jobs.get(jobId);
    if (job) {
      job.status = status;
      job.errors = errors;
    }
  }
}