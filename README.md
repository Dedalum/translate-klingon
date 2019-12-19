# Translate Klingon (from North-American)

Jexia back-end challenge.

Translate a name written in English to Klingon and find out its species using
http://stapi.co

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
clean                          Remove previous build
help                           Display this help screen
```

## Usage

```
Usage of ./translate-klingon:
  -config-path string
        JSON configuration file path (default "config.json")
  -name string
        Klingon character name in Latin alphabet to translate
```

### Structure

```bash
main.go
app/app.go                      # main package for running the app
config/config.go                # config packages for the APP
http/client.go                  # HTTP client for querying the API
translate/translate.go          # Translate package for converting Klingon names
                                # from Latin alphabet to Klingon alphabet

vendor/                         # vendor packages

config.json.example             # example config file
Makefile                        # Makefile 
```

## Development

The application is divided into two main packages:

### HTTP package

The `http` package handles the HTTP client requests to _STAPI_, the API. It
makes the HTTP requests but also parses the received JSON data.
The client is bound to the API version. We expect to update its version along
the API version, therefore the use of constants for the API ressource paths.

### Translate package

The `translate` package loads the alphabet mapping between the Latin/English
alphabet version of the Klingon language with its hexadecimal number version.
To simplify the conversion process, the Latin alphabet part is put in lower
case to the exception of the two letters 'q' and 'Q' (which in English would
result in the same letter but in Klingon are different). For instance, on the
contrary 'D' can be represented as 'd' because 'd' (lower case) is not a
distinct letter in the Klingon alphabet. This avoids having us to set specific
rules for each letter and we only have to consider the specific case of 'q' and
'Q'.

The Latin version of the Klingon alphabet contains at most 3 Latin letter 
('tlh') so to convert a word, we tokenize it and directly convert it:
1. Loop over the name string:
    1. check the 3 characters starting from the current index in the loop and
    try to match them with the map (the only case where this succeeds is for
    the Klingon letter 'tlh')
    2. if no match, check for 2 characters (matches for 'ch' and gh')
    3. if no match, check for 1 character
    4. if no match, we cannot convert the Klingon name as it contains 
    non-Klingon letters and the program exits
2. For every letter we match, we advance in the name string by the number of
characters the Klingon letter is composed of (e.g 'tlh' is a 3-character 
letter in Klingon while represented with Latin alphabet)

## TODO

### Ideas \& improvements
- logger (see github.com/sirupsen/logrus for instance)
