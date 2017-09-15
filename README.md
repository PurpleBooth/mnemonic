# mnemonic

[![Go Report Card](https://goreportcard.com/badge/github.com/PurpleBooth/mnemonic)][3]
[![codebeat badge](https://codebeat.co/badges/f29f1414-8688-431d-b67d-3e607af1a357)][4]

Late night hack session for a toy to generate mnemonic's.

## Installing

```bash
go get -u github.com/golang/dep/cmd/dep
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
dep ensure
(cd mnemonic && go test)
(cd cmd/mnemonic/ && go install)
```

## Running

Uses [wordnet dictionary files][1]. Once installed run

```bash
$ mnemonic /tmp/dict "ROYGBIV"
resistless ocellated turkey yaw gracefully. befouled interconnection victimize.
```

## Docker

Alternatively you can run the docker container

```bash
$ docker build -t mnemonic:latest . && docker run -it --rm mnemonic:latest "ROYGBIV"
rocky open yield gainfully. bonzer incapability vote in.
```

## Links

* [WordNet Dictionary][1]
* [Go Docs][2]

[1]: http://wordnet.princeton.edu/wordnet/download/current-version/
[2]: https://godoc.org/github.com/PurpleBooth/mnemonic/mnemonic
[3]: https://goreportcard.com/report/github.com/PurpleBooth/mnemonic
[4]: https://codebeat.co/projects/github-com-purplebooth-mnemonic-master