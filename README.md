# Fixie Wrench

Fixie Wrench is a command line utility that makes it easy to proxy any TCP connection through [Fixie Socks](https://usefixie.com/documentation/socks), even if your language or wrenchent library does not natively support SOCKSv5 proxies. By connecting through Fixie Socks, your application with have a stable set of outbound IP addresses, making it possible to address a remote service that performs IP address whitelisting from Heroku or other platforms that provide ephemeral instances.

Fixie Wrench does port-forwarding, similar to SSH port forwarding, so your remote database, FTP server, or other service will appear to be running locally from the perspective of your application code.

To use Fixie Wrench, you must have an active Fixie Socks account. If you do not have a Fixie Socks account, you can sign up on the [Heroku Marketplace](https://addons.heroku.com/provider/addons/fixie-socks).

## Usage
Assuming your `FIXIE_SOCKS_HOST` enironment variable is set, forwarding a connection is as simple as:

```
fixie-wrench $localPort:$host:$remotePort
```

For example, to forward a Postgres database running on port 5432 of my-database.example.com to port 1234 on your application host: `fixie-wrench 1234:my-database.example.com:5432`. With Fixie Wrench running, your application can now connect to the remote database by connecting to `localhost:1234`.

## Installation
Download the latest release from the [releases page](https://github.com/usefixie/fixie-wrench/releases/) and place the files in your project directory. You can now execute Fixie Wrench locally by running `./fixie-wrench` in the terminal or, in production, by adding your `fixie-wrench` command to your procfile.

## Advanced use

### Optional command line flags:

- **-v**: Verbose mode. In verbose mode, Fixie Wrench will print logs for each request to STDERR
- **--socksConnectionString**: If set, Fixie Wrench will use this connection string instead of the `FIXIE_SOCKS_HOST` environment variable

### Forwarding multiple ports:

Fixie Wrench accepts multiple positional arguments, each specifying a forwarding command in the form of `$localPort:$host:$remotePort`. A single instance of Fixie Wrench can proxy to many remote hosts on different local ports.

### Platform-specific binaries and the cross-platform launcher

If you install Fixie Wrench from the releases page, `fixie-wrench` is a bash script which determines the OS and architecture of your machine and loads the correct pre-built Fixie Wrench binary. Out of the box, the launcher supports 64 bit x86 chips and 64 bit ARM chips. It supports MacOS and Linux, making it possible to use the same command locally on an Apple Silicon Macbook and an x86 Linux server like Heroku.

You can also call out to a specific prebuilt binary directly (eg. `fixie-wrench-linux-arm64`), but you should not need to in normal operation.

### An advanced example

```
fixie-wrench -v 1234:my-first-database.example.com:5432 1235:my-second-database.example.com:5432
```

This means "using the FIXIE_SOCKS_HOST environment variable, make my-first-database.example.com:5432 available to my application on localhost:1234, and make my-second-database.example.com:5432 available to my application on localhost:1235, and print verbose logs."

## Building for other platforms
Fixie provides prebuilt binaries for common platforms and chip architectures, but if you want to run Fixie Wrench on a more esoteric system (eg. FreeBSD running on a 32-bit Intel chip), you can do so by cloning the repository and running `make build` after installing Go 1.17. This will produce a binary (`bin/fixie-wrench`) compiled for your OS and architecture.

## License

Fixie Wrench was created by Fixie Technologies LLC and is released under the MIT License. You are free to use, modify, and fork Fixie Wrench without restriction.

While Fixie Wrench is designed for use with Fixie Socks, and is only tested with Fixie Socks, it should work with any SOCKSv5 proxy server.

We believe the MIT license is a reason to choose Fixie Wrench. If you are going to run code on your server, you should be able to read the code and build from source. The MIT license does not encumber customers who are building closed-source projects as the GPL, AGPL, and other copy-left licenses may.