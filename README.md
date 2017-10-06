# Jenny [![Build Status](https://travis-ci.org/jorgechato/jenny.svg?branch=master)](https://travis-ci.org/jorgechato/jenny)
## Run
```bash
$ jenny
```
## Use
Command | Action | Status
--- | --- | ---
profile uri <var> | Location of the Jenkins server | ready
profile user <var> | Username credential | ready
profile pwd <var> | Password credential | ready
profile name <var> | If you have multiple Jenkins profiles, default: Default | ready
profile use <var> | Use different Jenkins credentials | pending
profile show | Show the current profile configuration | ready
profile show -u | Uncover the password. | ready
profile show --uncover | Uncover the password. | ready
profile cancel | Close and discard configuration | ready
profile save | Save and close configuration | ready
profile save -f | Save current configuration in .jenny.yml file. | ready
profile save --force-save | Save current configuration in .jenny.yml file. | ready

## Compile from scratch
### Install
```bash
$ go get # To install dependencies
$ go install
```
### Run without compile
```bash
$ go run manage.go
```
