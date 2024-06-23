# pkgo

`pkgo` is a simple CLI tool to help you find, install and manage Go packages.

## Motivation

I was tired of having to search for the full package name when I wanted to do `go get`.
So, I decided to build a simple tool to automate this process.

## Usage

You need to have Go installed and GOPATH in your PATH.

You can install it by running:

`go install github.com/estevesnp/pkgo@latest`

You have the following commands available:

- `pkgo search <package>`: Search for and list all packages that match the given package name.
- `pkgo get <package>`: Get the package with the given name.
- `pkgo install <package>`: Install the package with the given name.

All commands have a -l flag to define the number of results to be shown, for example `pkgo search gin -l 10`.

Additionally, the `pkgo get` command has a -u flag to update the package, for example `pkgo get gin -u`.

## TODO

- [ ] Add tests
- [ ] Add -v flag to define versions on get and install commands
