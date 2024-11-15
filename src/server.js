import express from 'express';
import { jobRouter } from './routes/jobs.js';

const app = express();
const port = process.env.PORT || 8080;

app.use(express.json());
app.use('/api', jobRouter);

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});