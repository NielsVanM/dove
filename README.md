# Dove
Easily run multiple development processes at the same time!


# Features
- Easy customizable TOML config!
- Watch outputu of all processes!
- Restart processes or KILL them!


# Installation
## Go Install
```sh
go install github.com/nielsvanm/dove

dove -v
```

## Prebuilt releases
Check the [Releases Tab](https://github.com/NielsVanM/dove/releases) for the latests pre-built release.


## From Source
```sh
# Clone project
git clone github.com/nielsvanm/dove
go get

# Build Binary
go build cmd/dove/main.go -o dove


# Install in bin
chmod +x dove
mv dove /usr/bin/
```

# Usage

## Initialize 
```text
NAME:
   dove init

USAGE:
   dove init [command options] [arguments...][arguments...]

DESCRIPTION:
   Initializes the local directory with a dove config

OPTIONS:
   --force, -f  Forces the init process even though a dove config is already pressent (default: false)
   --help, -h   show help
```

Results in a bare minimum `.dovecfg` in the current directory.



