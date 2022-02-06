# FCM Cleaner

Simple service to connect to database, read available FCM Tokens and perform a dry run to check which tokens are no longer valid and deletes them

## Requirements

- [golang](https://go.dev/)

## Dependencies

- [Go SQL Driver](github.com/go-sql-driver/mysql) `v1.5.0`
- [Go Dot Env](github.com/joho/godotenv) `v1.3.0`

## Setting up

- This service uses plain SQL query to fetch & delete tokens. You need to edit those queries to adapt them to your table. Queries are part of `fcmTokensRepository`

- Check .env.demo file for env values the service uses to complete its work
