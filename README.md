# Jenny
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/jenny)](https://goreportcard.com/report/github.com/jorgechato/jenny) [![Build Status](https://travis-ci.org/jorgechato/jenny.svg?branch=master)](https://travis-ci.org/jorgechato/jenny)
## Download and run
Download the latest version from https://github.com/jorgechato/jenny/releases/download/alpha-bones/jenny
Move it to your **/usr/bin** or add the folder where it lives into your $PATH.
```bash
$ jenny
```
## Use
### Profile
Command | Action | Status
--- | --- | ---
profile uri [var] | Location of the Jenkins server | ready
profile user [var] | Username credential | ready
profile pwd [var] | Password credential | ready
profile project [var] | Unique id of the Job/Pipeline | ready
profile show | Show the current profile configuration | ready
profile show --uncover | Uncover the password. | ready
profile logout | Logout | ready
profile login | Login | ready
profile save | Save current configuration in .jenny.yml file. | ready
profile save --global | Create .jenny.yml in $HOME directory. | ready
profile clear | Remove .jenny.yml file. | ready
profile cancel | Discard configuration | ready

### Jenkins
Command | Action | Status
--- | --- | ---
open | Opens the UI dashboard of this project in the browser. | ready

## Compile from scratch
### Structure
```bash
.
├── Dockerfile
├── jenny
│   ├── client.go
│   ├── completer.go
│   ├── executor.go
│   ├── filter.go
│   ├── model.go
│   ├── option.go
│   └── yaml.go
├── manage.go
└── README.md
```
### Get
```bash
$ git clone https://github.com/jorgechato/jenny.git
$ cd jenny
```
### Install
```bash
$ go get -v ./... # To install dependencies
$ go install
```
### Run without compile
```bash
$ go run manage.go
```
