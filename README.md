# SPN2 Go Library and CLI

The SPN2 Go Library provides an easy-to-use interface to interact with the Save Page Now 2 (SPN2) API, which is part of the Internet Archive's Wayback Machine service. This library allows users to capture web pages, retrieve the status of these captures, and check the system and user account status. Additionally, a simple command-line interface (CLI) is included for direct interaction with the SPN2 API.

## SPN2

SPN2 is archive.org's Save Page Now 2 API.

Documentation for the API itself can be found at https://docs.google.com/document/d/1Nsv52MvSjbLb2PCpHlat0gkzw0EvtSgpKHu4mk0MnrA.

## Features

The library supports the following features:
1. Submitting a new URL to be saved (captured).
2. Retrieving the status of a capture using the Job ID.
3. Checking the system status of the SPN2 API.
4. Checking the user's account status.
5. Authentication using the S3 method.

The CLI provides corresponding commands for these features.

## Installation

To use this library and CLI, clone the repository or download the source code into your Go workspace.

## Configuration

To use the CLI, set the following environment variables for authentication:

- `SPN2_ACCESS_KEY`: Your SPN2 access key.
- `SPN2_SECRET_KEY`: Your SPN2 secret key.

These keys are necessary for the API to authenticate your requests.

## CLI Usage

The CLI tool supports the following commands:

- `spn2 save {url}`: Saves the URL and prints the resulting JSON response.
- `spn2 status {job_id}`: Checks the status of the capture with the given Job ID and prints the resulting JSON response.
- `spn2 health`: Prints the system status of the SPN2 API in JSON format.
- `spn2 user`: Prints the user's account status in JSON format.

### Examples

To capture a webpage:

```shell
$ spn2 save http://example.com
```

To check the status of a capture:

```shell
$ spn2 status <job_id>
```

To check the system status:

```shell
$ spn2 health
```

To check the user's account status:

```shell
$ spn2 user
```

## Library Usage

Create a new client and use the methods provided to interact with the SPN2 API:

```go
client := spn2.NewClient("YourAccessKey", "YourSecretKey")
// Use client to interact with the SPN2 API
```

Refer to the library's documentation for detailed information on each method.

## Contributing

Contributions are welcome. Please open an issue or submit a pull request with your changes.

## Support

For support, open an issue in the GitHub repository. Support is on a best-effort basis.
