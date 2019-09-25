Linux Terraform Provider
==================

[![Maintainability](https://api.codeclimate.com/v1/badges/2692b09c7f370d4d7143/maintainability)](https://codeclimate.com/github/sam-myers/terraform-provider-linux/maintainability)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/aa4e5171703a4db6bc18ef2c9fc530f1)](https://www.codacy.com/manual/sam-myers/terraform-provider-linux?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=sam-myers/terraform-provider-linux&amp;utm_campaign=Badge_Grade)

Usage
---------------------

```
provider "linux" {}
```

See the documentation for examples

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/sam-myers/terraform-provider-linux`

```sh
$ mkdir -p $GOPATH/src/github.com/sam-myers; cd $GOPATH/src/github.com/sam-myers
$ git clone git@github.com:sam-myers/terraform-provider-linux
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/sam-myers/terraform-provider-linux
$ make build
```

Documentation
----------------------

Build the documentation locally

```sh
make website
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-template
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
