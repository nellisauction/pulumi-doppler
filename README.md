# Doppler Resource Provider

The Doppler Resource Provider lets you manage [Doppler](https://doppler.com) secrets and configuration resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @nellisauction/pulumi-doppler
```

or `yarn`:

```bash
yarn add @nellisauction/pulumi-doppler
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_doppler
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/nellisauction/pulumi-doppler/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Doppler
```

## Configuration

The following configuration points are available for the `doppler` provider:

- `doppler:token` (environment: `DOPPLER_TOKEN`) - a Doppler token (service account, service token, or personal token)
- `doppler:host` (environment: `DOPPLER_API_HOST`) - the Doppler API host (default: `https://api.doppler.com`)
- `doppler:verifyTls` (environment: `DOPPLER_VERIFY_TLS`) - whether to verify TLS (default: `true`)

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/doppler/api-docs/).
