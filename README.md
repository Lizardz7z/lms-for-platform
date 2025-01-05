# LMS Module API

## Endpoints
### GET /courses
Returns a list of all courses.

### GET /courses/{id}
Returns details of a specific course.

### GET /progress/{user_id}
Returns the progress of a specific user.

### POST /progress/{user_id}/{assignment_id}/complete
Marks an assignment as completed for the user.
