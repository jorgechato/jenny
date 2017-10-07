# Jenny
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/jenny)](https://goreportcard.com/report/github.com/jorgechato/jenny) [![Build Status](https://travis-ci.org/jorgechato/jenny.svg?branch=master)](https://travis-ci.org/jorgechato/jenny)
## Download and run
- Download the latest version from [here](https://github.com/jorgechato/jenny/releases/download/alpha-bones/jenny)
- Move it to your ***/usr/bin*** or add the folder where it lives into your **$PATH**.
- Use the command **jenny** to start using it

```bash
$ jenny
```
## Use
### Profile
Command | Action | Status
--- | --- | ---
profile uri \<JenkinsURI\> | Location of the Jenkins server | ready
profile user \<userName\> | Username credential | ready
profile pwd \<password\> | Password credential | ready
profile project \<projectName\> | Unique id of the Job/Pipeline | ready
profile show | Show the current profile configuration | ready
profile show --uncover | Uncover the password. | ready
profile logout | Logout | ready
profile logout --force | Clean user and password in .jenny.yml. | ready
profile login | Login | ready
profile save | Save current configuration in .jenny.yml file. | ready
profile save --global | Create .jenny.yml in $HOME directory. | ready
profile save --force | Store user and password in the .jenny.yml file. | ready
profile clear | Remove .jenny.yml file. | ready
profile cancel | Discard configuration | ready

### Jenkins
Command | Action | Status
--- | --- | ---
open | Opens the UI dashboard of this project in the browser. | ready
status \<jobName\> \<executionNumber\> | Status of given build id or latest build. | ready
status \<jobName\> --last | Get the last execution. | ready
logs \<jobname\> \<executionNumber\> | Print the logs for a build. | ready
logs \<jobName\> --last | Get the last execution. | ready
stop \<jobname\> \<executionNumber\> | Stop a build execution. | ready
stop \<jobName\> --last | Get the last execution. | ready
build \<jobName\> | Trigger parametrized build. | in progress (ready without params)
describe \<jobName\> | Describe build history of project. | ready

### Shortcuts
KeyBinding          | Description
--------------------|---------------------------------------------------------
<kbd>Ctrl + A</kbd> | Go to the beginning of the line (Home)
<kbd>Ctrl + E</kbd> | Go to the End of the line (End)
<kbd>Ctrl + P</kbd> | Previous command (Up arrow)
<kbd>Ctrl + N</kbd> | Next command (Down arrow)
<kbd>Ctrl + F</kbd> | Forward one character
<kbd>Ctrl + B</kbd> | Backward one character
<kbd>Ctrl + D</kbd> | Delete character under the cursor
<kbd>Ctrl + H</kbd> | Delete character before the cursor (Backspace)
<kbd>Ctrl + W</kbd> | Cut the Word before the cursor to the clipboard.
<kbd>Ctrl + K</kbd> | Cut the Line after the cursor to the clipboard.
<kbd>Ctrl + U</kbd> | Cut/delete the Line before the cursor to the clipboard.

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
