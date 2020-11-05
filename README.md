# Cutos - (Curl to Swagger) 

Cutos provides a terminal UI for writing OpenAPI3 and documentation with the easiest trim by copying сurl queries and json responses.
The project is in the early alpha version and is not a candidate for the main developer installation.
The project is a weekend project and was created exclusively to meet the needs of the author.

---

## Installation

cutos is available on Linux, macOS and Windows platforms.

* Binaries for Linux, Windows and Mac are available as tarballs in the [release](https://github.com/cjp2600/cutos/releases) page.
* Use `go get` to install the latest version of the tools. This command will install the cutos with dependencies:
```shell
go get -u github.com/cjp2600/cutos
```                                                           

---

## Create new swagger document
```shell
catos init swagger.json
```

## Append new path from curl
```shell
cutos add curl swagger.json
```
