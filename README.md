### Folder Structure

`/bin`:binary executables
`/cmd/api`: HTTP layer
`/docs`:
`/internal/db`:
`/internal/env`:
`/internal/store`: storage layer

### Order of Development

1. Project Architecture:

- Set Up Development Environment
- Clean Layered Architecture

2. Scaffolding API Server:

- Set Up HTTP Server and API: net/http, go-chi
- Add Hot Reloading: air
- Environment Variables: direnv

3. Databases:

- Repository Pattern
- Running PostgreSQL container on Docker
- Configuring the DB Connection Pool
- Persisting Data
- SQL Migrations

4. Posts CRUD:

- Marshalling JSON
- Creating a Post
- Getting a Post by ID
- Internal Errors Package
- HTTP Payload Validation: go-validator
- DB Relationships
- SQL Joins
- Adding Comments to Posts
- Updating Posts: PATCH
- Deleting Posts
- Standardising JSON Responses
- Optimistic Concurrency Control
- Managing SQL Query Timeouts
- Database Seeding

5. The User Feed:

- Creating User Profile
- Get User By ID
- Add Followers Table
- SQL Indexes
- User Feed Algorithm

6. Filtering, Sorting, and Pagination:

- WIP
