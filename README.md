# Translate Klingon (to North-American)
Jexia back-end challenge

## Table of contents

- [Tools & CI/CD](#tools-&-ci/cd)
- [Usage](#usage)
- [Development](#development)
- [TODO](#todo)

## Tools & CI/CD
There's a Makefile available that allows you to perform useful tasks like:
```
lint                           Lint the files
test                           Run unittests
dep                            Get the dependencies
build                          Build the binary file
dockerize                      Build and tag as latest and current version
clean                          Remove previous build
help                           Display this help screen
```

## Usage

```bash
Usage of ./translate-klingon:
```

### Structure

```
main.go
app/app.go                      # main package for running the app
config/config.go                # config packages for the APP
http/client.go                  # HTTP client for querying the API
vendor/                         # vendor packages

config.json.example             # example config file
Makefile						# Makefile 
```

## Development


## TODO

