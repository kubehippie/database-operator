# Database Operator

[![Build Status](http://github.dronehippie.de/api/badges/kubehippie/database-operator/status.svg)](http://github.dronehippie.de/kubehippie/database-operator)
[![Stories in Ready](https://badge.waffle.io/kubehippie/database-operator.svg?label=ready&title=Ready)](http://waffle.io/kubehippie/database-operator)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d5bdf000fe8149adb2b3733b161325ac)](https://www.codacy.com/app/kubehippie/database-operator?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kubehippie/database-operator&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/kubehippie/database-operator?status.svg)](http://godoc.org/github.com/kubehippie/database-operator)
[![Go Report](http://goreportcard.com/badge/github.com/kubehippie/database-operator)](http://goreportcard.com/report/github.com/kubehippie/database-operator)
[![](https://images.microbadger.com/badges/image/kubehippie/database-operator.svg)](http://microbadger.com/images/kubehippie/database-operator "Get your own image badge on microbadger.com")

TBD


## Install

You can download prebuilt binaries from our [GitHub releases](https://github.com/kubehippie/database-operator/releases), or you can use our Docker images published on [Docker Hub](https://hub.docker.com/r/kubehippie/database-operator/tags/).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8.

```bash
go get -d github.com/kubehippie/database-operator
cd $GOPATH/src/github.com/kubehippie/database-operator

# install retool
make retool

# sync dependencies
make sync

# generate code
make generate

# build binary
make build

./bin/database-operator -h
```


## Security

If you find a security issue please contact thomas@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
