import fetch from 'node-fetch';
import sharp from 'sharp';

export class ImageProcessor {
  async processJob(jobId, visits, jobService) {
    const errors = [];

    try {
      for (const visit of visits) {
        for (const imageUrl of visit.image_url) {
          try {
            await this.processImage(imageUrl);
          } catch (error) {
            errors.push({
              store_id: visit.store_id,
              error: error.message
            });
            jobService.updateJobStatus(jobId, 'failed', errors);
            return;
          }
        }
      }

      jobService.updateJobStatus(jobId, 'completed');
    } catch (error) {
      jobService.updateJobStatus(jobId, 'failed', [{
        error: 'Internal server error'
      }]);
    }
  }

  async processImage(url) {
    // Download image
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error('Failed to download image');
    }

    const buffer = await response.buffer();
    
    // Get image dimensions using sharp
    const metadata = await sharp(buffer).metadata();
    const perimeter = 2 * (metadata.width + metadata.height);

    // Simulate GPU processing time
    await new Promise(resolve => 
      setTimeout(resolve, 100 + Math.random() * 300)
    );

    return perimeter;
  }
}