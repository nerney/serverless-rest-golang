# Serverless Golang REST Endpoint

![Build Status](https://travis-ci.com/nerney/serverless-rest-golang.svg?branch=master)
[![Test Coverage](https://codecov.io/gh/nerney/serverless-rest-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/nerney/serverless-rest-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/nerney/serverless-rest-golang)](https://goreportcard.com/report/github.com/nerney/serverless-rest-golang)

## Goals

The goal here is to acheive a fully serverless REST endpoint. By "fully serverless", I mean that there are no servers continually running or database resources provisioned.

This is accomplished through a single lambda function that handles crud methods (`GET`,`POST`,`PUT`,`DELETE`).
The serverless environment uses an in-memory cache for holding the data, but ultimately, I would like to have changes immediately synced to s3.

## Requirements

- Node.js
- Golang

## Deploying

![Deploy Status](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiWXNBSisyeXNlQlh4T2RDc1E4VHZkVW9MOVJJa3d0SzFqa1VEdldmZ2hGSHJJbStRVjZhNWRhaGNJQXgxZ2NVT1RkNWhiNlhBWjgzQ2hkQW9QNy84ZXFFPSIsIml2UGFyYW1ldGVyU3BlYyI6Ik91THkyRFBoMS9UaDYvdUwiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)

Deploying this service to an AWS environment can be accomplished in more than one way:

- From your local machine, it will deploy using your locally configured AWS credentials.
- From AWS CodeBuild, using the default `buildspec.yml`

## Scripts

```bash
# installs dependencies
sh scripts/install.sh

# runs tests with coverage
sh scripts/test.sh
open coverage.html # (optional: view the coverage report)

# builds the application
sh scripts/build.sh

# deploy the application with serverless
sh scripts/install.sh

```

#### MORE TO COME...