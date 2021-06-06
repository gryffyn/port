# port
[![Build Status](https://ci.neveris.one/api/badges/gryffyn/port/status.svg?ref=refs/heads/main)](https://ci.neveris.one/gryffyn/port)
[![Go Report Card](https://goreportcard.com/badge/git.neveris.one/gryffyn/port)](https://goreportcard.com/report/git.neveris.one/gryffyn/port)
Go application for showing port info.

## requirements
Requires netstat.

Archlinux: `net-tools`

## usage
```
§ port -h
Usage: port [--simple] PORT

Positional arguments:
  PORT

Options:
  --simple, -s           Show simplified output (just the protocol, local address, and PID/process)
  --help, -h             display this help and exit

```
