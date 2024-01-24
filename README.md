<p align="center">
  <a><img src="./images/bitcrook.png" width=180 height="180"></a>
  <h1 align="center">Bitcrook</h1>
  <p align="center">
    <a href="https://pkg.go.dev/github.com/ax-i-om/bitcrook"><img src="https://pkg.go.dev/badge/github.com/ax-i-om/tempest.svg" alt="Documentation"></a>
    <a href="https://goreportcard.com/report/github.com/ax-i-om/bitcrook"><img src="https://goreportcard.com/badge/github.com/ax-i-om/bitcrook" alt="Go Report Card"></a>
    <a><img src="https://img.shields.io/badge/tests-8&#47;8-green.svg" alt="Tests"></a>
    <a><img src="https://img.shields.io/badge/version-2.3.0-blue.svg" alt="v2.3.0"></a><br>
    <a href="https://app.deepsource.com/gh/ax-i-om/bitcrook/" target="_blank"><img alt="DeepSource" title="DeepSource" src="https://app.deepsource.com/gh/ax-i-om/bitcrook.svg/?label=active+issues&show_trend=true"/></a><br>
    Centralize and expedite OSINT investigations<br>
  <a href="https://github.com/users/ax-i-om/projects/1">View the roadmap</a><br>
</a>
  </p><br>
</p>

## Table of Contents

- [Information](#information)
  - [About](#about)
  - [Disclaimer](#disclaimer)
  - [Changelog](#changelog)
  - [Preview](#preview)
  - [Attributions](#attributions)
- [Usage](#usage)
  - [Installation](#installation)
  - [CLI](#cli)
  - [GUI](#gui)
  - [Docker](#docker)
  - [Authentication](#authentication)

## Information

### About

Bitcrook is an open-source intelligence apparatus that aims to centralize all of the tools necessary to carry out an investigation. Although investigations will still require human interaction to connect the dots, the interface can be tailored to an individual’s needs to expedite the process of due diligence. Some packages do require an authentication key and others do not. See the [Package Types](#package-types) tables for more information. Bitcrook can be integrated within your application OR it can be used directly from the [terminal/web browser.](#preview)

### Disclaimer

It is the end user's responsibility to obey all applicable local, state, and federal laws. Developers assume no liability and are not responsible for any misuse or damage caused by this program. By using Bitcrook, you agree to the previous statements.

### Changelog

Bitcrook v2.0.0 has been released as of October 11, 2023, only one day after v1.0.0 due to some significant changes to project structure. Here are some highlights of the recent changes in the v1.0.0 and v2.0.0 releases: 

 - Switch web framework from Fiber to Go (enabled Bitcrook to be more lightweight)
 - Complete CLI and GUI design overhaul
 - Proper docker implementation
 - Binary release to support `go install`
 - Switched secret configuration from .json to .env, see [Authentication](#authentication)
 - Enabled API key support in web app
 - Fixed CORS error when accessing web app over LAN
 - Username lookup command now only displays valid results
 - Added HaveIBeenPwned breach check to email and domain lookup commands
 - Added account existence check for emails (Wordpress, Gravatar, Twitter, and Pinterest so far)
 - Added phone and email search to web app
 - Reimplemented Case.Law search (record)
 - Packages now grouped based on field type rather than authentication type
 - Fixed tests and some documentation
 - Bug fixes and general optimizations

### Preview

<a><img src="./images/cliprev.png"></a>
<a><img src="./images/guiprev.png"></a>


### Attributions

 - [Social Preview](./images/card.jpg) created with [Canva.](https://www.canva.com/)
 - CLI output formatting inspired by [Echo.](https://echo.labstack.com)

## Usage

### Installation

 - Fetch the repository via 'git clone': `git clone https://github.com/ax-i-om/bitcrook.git`

 *or*

 - Install via 'go install': `go install github.com/ax-i-om/bitcrook@latest`

### CLI 

1. Navigate to the root directory of Bitcrook via `cd`
2. In your preferred terminal, enter and run: `go run main.go`
3. After running the aforementioned command, all dependencies will be installed and usage help will be printed to the console.

### GUI

1. Navigate to the root directory of Bitcrook via `cd`
2. In your preferred terminal, enter and run: `go run main.go server`
3. In your preferred web browser, navigate to `http://127.0.0.1:6174`

### Docker

You can also host the Bitcrook webapp via docker:

1. Navigate to the root directory of Bitcrook via `cd`
2. Build the Docker image: `docker build -t bitcrook .`
3. Run Bitcrook via Docker: `docker run --restart always --env-file .env --name bitcrook -d -p 6174:6174 bitcrook`
4. Access the interface via `http://127.0.0.1:6174`

### Authentication

Some packages require an API key. Bitcrook will run without specifying any API keys; however, if you would like to implement these functionalites, create a file named `.env` and format it like so:

``` conf
BITCROOK_MLSA=UNSPECIFIED
BITCROOK_HIBP=UNSPECIFIED
BITCROOK_IPTL=UNSPECIFIED
```

Replace the instances of `UNSPECIFIED` with the corresponding API key.

From top to bottom, these environment variables should correspond with the API keys for the services Melissa, HaveIBeenPwned, and IP2Location. Insure that the environnment variables are passed/exported after making any changes. For Docker, append `--env-file .env` to the arguments. In Linux, run the command: `export $(xargs <.env)`. Results may vary depending on your operating system/shell.