# What is this?

This is a live chat application.

# Prerequisites

### Install build tools

On Linux
```sh
sudo apt-get install build-essential
```

On Darwin (MacOS)

```sh
xcode-select --install
```

### Install  `frizzante`

```sh
go install github.com/razshare/frizzante@latest
```

>[!TIP]
>Remember to add Go binaries to your path.
>
> ```sh
> export GOPATH=$HOME/go
> export PATH=$PATH:$GOPATH/bin
> ```

# Get Started

Configure project

```sh
make configure
```

Start development mode with

```sh
make dev
```

# Build

Build for production with

```sh
make build
```

This will create a standalone `.gen/bin/app` binary file.