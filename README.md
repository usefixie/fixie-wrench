![Fixie](https://usefixie.com/img/logo.svg)

# fixie-wrench

fixie-wrench is a command line utility that makes it easy to proxy any TCP connection through [Fixie Socks](https://usefixie.com/documentation/socks), even if your language or client library does not natively support SOCKSv5 proxies. By connecting through Fixie Socks, your application with have a stable set of outbound IP addresses, making it possible to address a remote service that performs IP address whitelisting from Heroku or other platforms that provide ephemeral instances.

fixie-wrench does port-forwarding, similar to SSH port forwarding, so your remote database, FTP server, or other service will appear to be running locally from the perspective of your application code.

To use fixie-wrench, you must have an active Fixie Socks account. If you do not have a Fixie Socks account, you can sign up on the [Heroku Marketplace](https://addons.heroku.com/provider/addons/fixie-socks).

## Usage
Assuming your `FIXIE_SOCKS_HOST` enironment variable is set, forwarding a connection is as simple as:

```
fixie-wrench $localPort:$host:$remotePort
```

For example, to forward a Postgres database running on port 5432 of my-database.example.com to port 1234 on your application host: `fixie-wrench 1234:my-database.example.com:5432`. With fixie-wrench running, your application can now connect to the remote database by connecting to `localhost:1234`.

## Installation

### Recommended: Installer
The easiest way to install fixie-wrench is to use the installer. In your project directory, run:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/usefixie/fixie-wrench/HEAD/install.sh)"
```

### Alternative: Download pre-compiled binaries

You can download the latest release from the [releases page](https://github.com/usefixie/fixie-wrench/releases/) and place the files in your project directory. You should download at least `fixie-wrench-linux-amd64` (for execution on Heroku), and you may also want to download the binary for your local machine. If you run on more than one platform, we suggest also downloading the `fixie-wrench` shell script, which will automatically select the correct platform-specific binary at runtime.

## Advanced use

### Optional command line flags:

- **-v**: Verbose mode. In verbose mode, fixie-wrench will print logs for each request to STDERR
- **--fixieSocksHost**: If set, fixie-wrench will use this connection string instead of the `FIXIE_SOCKS_HOST` environment variable

### Forwarding multiple ports:

fixie-wrench accepts multiple positional arguments, each specifying a forwarding command in the form of `$localPort:$host:$remotePort`. A single instance of fixie-wrench can proxy to many remote hosts on different local ports.

### Platform-specific binaries and the cross-platform launcher

If you install fixie-wrench from the releases page, `./bin/fixie-wrench` is a bash script which determines the OS and architecture of your machine and loads the correct pre-built fixie-wrench binary. Out of the box, the launcher supports 64 bit x86 chips and 64 bit ARM chips. It supports MacOS and Linux, making it possible to use the same command locally on an Apple Silicon Macbook and an x86 Linux server like Heroku.

You can also call out to a specific prebuilt binary directly (eg. `fixie-wrench-linux-arm64`), but you should not need to in normal operation.

### An advanced example

```
fixie-wrench -v 1234:my-first-database.example.com:5432 1235:my-second-database.example.com:5432
```

This means "using the FIXIE_SOCKS_HOST environment variable, make my-first-database.example.com:5432 available to my application on localhost:1234, and make my-second-database.example.com:5432 available to my application on localhost:1235, and print verbose logs."

## Building for other platforms
Fixie provides prebuilt binaries for common platforms and chip architectures, but if you want to run fixie-wrench on a more esoteric system (eg. FreeBSD running on a 32-bit Intel chip), you can do so by cloning the repository and running `make build` after installing Go 1.17. This will produce a binary (`bin/fixie-wrench`) compiled for your OS and architecture.

## Example app
To see fixie-wrench in action, check out the [example app](https://github.com/usefixie/fixie-wrench-example-app).

You can also deploy the example app with a click:

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/usefixie/fixie-wrench-example-app)

## License

fixie-wrench was created by Fixie Technologies LLC and is released under the MIT License. You are free to use, modify, and fork fixie-wrench without restriction.

While this tool is designed for use with Fixie Socks, and is only tested with Fixie Socks, it should work with any SOCKSv5 proxy server.
