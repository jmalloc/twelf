# Twelf

Twelf is a *very* simple logging interface for [Twelve-Factor](http://12factor.net/)
applications written in Go.

The provided logger simply prints log output to `STDOUT`, as per the
[Twelve-Factor Application logging recommendations](http://12factor.net/logs).

Additionally, the logger discriminates between application messages and debug
messages, as per [Dave Cheney's post about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).
