### Folder Structure

`/bin`: compiled binaries (e.g., `air`, build artifacts) <br />
`/cmd/api`: main HTTP API server (entrypoint and route handlers) <br />
`/cmd/migrate`: Database migration SQL files and seed logic <br />
`/docs`: project documentation (designs, API specs, notes) <br />
`/internal/db`: db connection utilities and seeding helpers <br />
`/internal/env`: environment variable configuration and helpers <br />
`/internal/store`: storage layer: database access logic for posts, users, etc. <br />
`/scripts`: supporting SQL and shell scripts <br />

### Order of Development

1. Project Architecture:

-   Set Up Development Environment
-   Clean Layered Architecture

2. Scaffolding API Server:

-   Set Up HTTP Server and API: net/http, go-chi
-   Add Hot Reloading: air
-   Environment Variables: direnv

3. Databases:

-   Repository Pattern
-   Running PostgreSQL container on Docker
-   Configuring the DB Connection Pool
-   Persisting Data
-   SQL Migrations

4. Posts CRUD:

-   Marshalling JSON
-   Creating a Post
-   Getting a Post by ID
-   Internal Errors Package
-   HTTP Payload Validation: go-validator
-   DB Relationships
-   SQL Joins
-   Adding Comments to Posts
-   Updating Posts: PATCH
-   Deleting Posts
-   Standardising JSON Responses
-   Optimistic Concurrency Control
-   Managing SQL Query Timeouts
-   Database Seeding

5. The User Feed:

-   Creating User Profile
-   Get User By ID
-   Add Followers Table
-   SQL Indexes
-   User Feed Algorithm

6. Filtering, Sorting, and Pagination:

-   Pagination
-   Sorting
-   Fuzzy Search

7. Swagger

-   API Docs
-   Documented Routes

8. Logging

-   Adding Structured Logging

10. User Creation

-   User Registration Flow
-   SQL Transactions
-   User Activation
