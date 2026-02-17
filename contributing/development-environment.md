# Development environment

`grafana-foundation-sdk` relies on [`devbox`](https://www.jetify.com/devbox/docs/) to manage all
the tools and programming languages it targets.

A shell including all the required tools is accessible via:

```console
$ devbox shell
```

This shell can be exited like any other shell, with `exit` or `CTRL+D`.

One-off commands can be executed within the devbox shell as well:

```console
$ devbox run go version
```

Packages can be installed using:

```console
$ devbox add go@1.23
```

Available packages can be found on the [NixOS package repository](https://search.nixos.org/packages).

## See also

* [Navigating the repository](./navigating-repository.md)
* [Adding new resources](./adding-resources.md)
