# Serverless Golang REST Endpoint

![Build Status](https://travis-ci.com/nerney/serverless-rest-golang.svg?branch=master)
![Go Report Card](https://goreportcard.com/badge/github.com/nerney/serverless-rest-golang)
[![Test Coverage](https://codecov.io/gh/nerney/serverless-rest-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/nerney/serverless-rest-golang)
![Deploy Status](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiWXNBSisyeXNlQlh4T2RDc1E4VHZkVW9MOVJJa3d0SzFqa1VEdldmZ2hGSHJJbStRVjZhNWRhaGNJQXgxZ2NVT1RkNWhiNlhBWjgzQ2hkQW9QNy84ZXFFPSIsIml2UGFyYW1ldGVyU3BlYyI6Ik91THkyRFBoMS9UaDYvdUwiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)

## Requirements

* Node.js
* Golang

## Scripts

```bash
# installs dependencies
sh scripts/install.sh

# runs tests with coverage
sh scripts/test.sh
# this produces an html coverage report
open coverage.html

# builds the application
sh scripts/build.sh

# deploys the application with serverless
sh scripts/install.sh

```
