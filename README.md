# FCM Cleaner

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=alert_status)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=security_rating)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=sqale_index)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=ncloc)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=code_smells)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=omar-bizreh_fcmCleaner&metric=bugs)](https://sonarcloud.io/dashboard?id=omar-bizreh_fcmCleaner)

Simple service to connect to database, read available FCM Tokens and perform a dry run to check which tokens are no longer valid and deletes them

## Requirements

- [golang](https://go.dev/) `v1.17.6`

## Dependencies

- [Go SQL Driver](github.com/go-sql-driver/mysql) `v1.5.0`
- [Go Dot Env](github.com/joho/godotenv) `v1.3.0`

## Setting up

- This service uses plain SQL query to fetch & delete tokens. You need to edit those queries to adapt them to your table. Queries are part of `fcmTokensRepository`

- Check .env.demo file for env values the service uses to complete its work. Use this file for running service locally. Make sure to uncomment `main.go` -> line 17 & 18

- `buildService.sh`: Builds the service and wrap it in a docker container and push to local registry

## Available Endpoints

- **clean_tokens**: Initiates cleaning tokens
