# Serverless Golang REST Endpoint

![Build Status](https://travis-ci.com/nerney/serverless-rest-golang.svg?branch=master)
[![Test Coverage](https://codecov.io/gh/nerney/serverless-rest-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/nerney/serverless-rest-golang)
![Go Report Card](https://goreportcard.com/badge/github.com/nerney/serverless-rest-golang)

## Goals

The goal here is to acheive a fully serverless REST endpoint. By "fully serverless", I mean that there are no servers continually running or database resources provisioned.

This is accomplished through a single lambda function that handles crud methods (`GET`,`POST`,`PUT`,`DELETE`).
The serverless environment uses an in-memory cache for holding the data, but ultimately, I would like to have changes immediately synced to s3.

#### Status

Done

- the REST lambda to handle all the CRUD operations.

ToDo

- the persistence layer:

  basically, we need to bootstrap the application to pull in the cache from s3 on init, if there isn't one, just create a new one (should only happen on first init), then sync any changes as they occur. 
  to really do this right we will want some `defer` logic to handle doing a final sync (if any updates have occured while the lambda was hot). to be truly serverless we will need to always make sure when performing a sync that we are not overwriting changes that may have occured in some other concurrent lambda, so we will need to diff those guys and make sure they always include everything.

## Requirements

- Node.js
- Golang

## Deploying

![Deploy Status](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiWXNBSisyeXNlQlh4T2RDc1E4VHZkVW9MOVJJa3d0SzFqa1VEdldmZ2hGSHJJbStRVjZhNWRhaGNJQXgxZ2NVT1RkNWhiNlhBWjgzQ2hkQW9QNy84ZXFFPSIsIml2UGFyYW1ldGVyU3BlYyI6Ik91THkyRFBoMS9UaDYvdUwiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)

Deploying this service to an AWS environment can be accomplished in more than one way:

* From your local machine, it will deploy using your locally configured AWS credentials.
* From AWS CodeBuild, using the default `buildspec.yml`

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