# Twelf

[![Build Status](http://img.shields.io/travis/jmalloc/twelf/master.svg)](https://travis-ci.org/jmalloc/twelf)
[![Code Coverage](https://img.shields.io/codecov/c/github/jmalloc/twelf/master.svg)](https://codecov.io/github/jmalloc/twelf)
[![Latest Version](https://img.shields.io/github/tag/jmalloc/twelf.svg?label=semver)](https://semver.org)
[![GoDoc](https://godoc.org/github.com/jmalloc/twelf?status.svg)](https://godoc.org/github.com/jmalloc/twelf/src/rinq)
[![Go Report Card](https://goreportcard.com/badge/github.com/jmalloc/twelf)](https://goreportcard.com/report/github.com/jmalloc/twelf)

Twelf is a *very* simple logging interface for [Twelve-Factor](http://12factor.net/)
applications written in Go.

The provided logger simply prints log output to `STDOUT`, as per the
[Twelve-Factor Application logging recommendations](http://12factor.net/logs).

Additionally, the logger discriminates between application messages and debug
messages, as per [Dave Cheney's post about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).
