# Common
## Base URL
> http://localhost:5000/api

## version
> /v1

GET /memories                                   # Get all memories
POST /memories                                  # Post a new memory

GET /memories/:memoryId                         # Get specific memory
PUT /memories/:memoryId                         # Post a new memory or replace existing one.

PATCH /memories/:memoryId                       # Update specific memory
DELETE /memories/:memoryId                      # Deletes a specific memory


GET http://localhost:5000/api/v1/memories/:memoryId