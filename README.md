# Cutos - (Curl to Swagger) 

Cutos provides a terminal UI for writing OpenAPI3 and documentation with the easy way by copying сurl queries and json responses.
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
If a file with documentation is not created, it should be created using the `init` command. 
```shell
catos init swagger.json
```
> The utility will ask for basic data to create a documentation file

## Append new path from curl
If the file has already been created or was already created, you can add a new path using the following command:
```shell
cutos add curl swagger.json
```
> The utility will open a new listener of your clipboard to create a patch.
> You can easily use programs that allow you to export to the curl format. Such as [Insomnia](https://insomnia.rest/download/), [Postman](https://www.postman.com/) or Browser

## Configuration
By default the utility looks for the configuration file in this path default is **$HOME/.cutos.yaml**
But you can always rewrite the path to the configuration file explicitly in the following parameter **--config=**
```shell
cutos add curl swagger.json --config=config.yml
```
## Example configuration:
```YAML
#
# Basic application settings
#
application_settings:

  # List of headers params to be skipped
  skipped_headers:
    - accept
    - authority
    - user-agent
    - content-type
    - origin
    - sec-fetch-site
    - sec-fetch-mode
    - sec-fetch-dest
    - referer
```
 
