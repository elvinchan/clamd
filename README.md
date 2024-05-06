# clamd

Golang Clamd Client

## Description

clamd is a Golang library and cmdline tool that implements the
Clamd client protocol used by ClamAV.

Forked from https://github.com/baruwa-enterprise/clamd

## Requirements

* Golang 1.16.x or higher

## Getting started

### Clamd client

The clamd client can be installed as follows

```console
$ go get github.com/elvinchan/clamd/cmd/clamdscan
```

Or by cloning the repo and then running

```console
$ make build
$ ./bin/clamdscan
```

### Clamd library

To install the library

```console
go get github.com/elvinchan/clamd
```

You can then import it in your code

```golang
import "github.com/elvinchan/clamd"
```

### Testing

``make test``

## License

MPL-2.0
