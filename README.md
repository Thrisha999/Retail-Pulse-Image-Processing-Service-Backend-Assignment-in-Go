# Image Processing Service

A Node.js service to process thousands of images collected from stores, calculating their perimeters and managing processing jobs.

## Description

This service provides REST APIs to:
1. Submit image processing jobs with store visits and image URLs
2. Check job status and results
3. Process images by calculating perimeters with simulated GPU processing time

## Setup Instructions

1. Install dependencies:
```bash
npm install
```

2. Start the server:
```bash
npm start
```

For development with auto-reload:
```bash
npm run dev
```

## API Documentation

### 1. Submit Job
POST `/api/submit`

Request body:
```json
{
  "count": 2,
  "visits": [
    {
      "store_id": "S00339218",
      "image_url": [
        "https://www.gstatic.com/webp/gallery/2.jpg",
        "https://www.gstatic.com/webp/gallery/3.jpg"
      ],
      "visit_time": "2023-11-08T10:00:00Z"
    }
  ]
}
```

### 2. Get Job Status
GET `/api/status?jobid=<job_id>`

## Technologies Used

- Node.js
- Express.js
- sharp (for image processing)
- node-fetch (for downloading images)
- uuid (for generating job IDs)

## Future Improvements

1. Add database persistence for jobs and results
2. Implement proper error handling and retry mechanisms
3. Add metrics and monitoring
4. Implement rate limiting
5. Add authentication and authorization
6. Add unit tests and integration tests
7. Implement concurrent image processing
8. Add logging system
9. Implement job queue system for better scalability
10. Add input validation and sanitization