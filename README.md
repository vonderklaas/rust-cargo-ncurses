### Order of Development

Project Architecture:

- Set Up Development Environment
- Clean Layered Architecture

Scaffolding API Server

- Set Up HTTP Server and API: net/http, go-chi
- Add Hot Reloading: air
- Environment Variables: direnv

Databases:

- Repository Pattern
- Running PostgreSQL container on Docker
- Configuring the DB Connection Pool
- Persisting Data
- SQL Migrations

Posts CRUD

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

The User Feed
