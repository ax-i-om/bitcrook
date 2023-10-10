<p align="center">
  <a><img src="./images/bitcrook.png" width=180 height="180"></a>
  <h1 align="center">Bitcrook</h1>
  <p align="center">
    <a href="https://goreportcard.com/report/github.com/ax-i-om/bitcrook"><img src="https://goreportcard.com/badge/github.com/ax-i-om/bitcrook" alt="Go Report Card"></a>
    <a><img src="https://img.shields.io/badge/tests-nil&#47;nil-orange.svg" alt="Tests"></a>
    <a><img src="https://img.shields.io/badge/version-1.0.0-blue.svg" alt="v1.0.0"></a><br>
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
- [Package Types](#package-types) *a-z*
  - [Court Cases](#court-cases)
  - [Discord](#discord)
  - [IP Address](#ip-address)
  - [Multi-Use](#multi-use)
  - [Tax ID](#tax-id)
  - [Username](#username)
  - [Vehicle](#vehicle)

## Information

### About

Bitcrook is an open-source intelligence apparatus that aims to centralize all of the tools necessary to carry out an investigation. Although investigations will still require human interaction to connect the dots, the interface can be tailored to an individual’s needs to expedite the process of due diligence. Some packages do require an authentication key and others do not. See the [Package Types](#package-types) tables for more information. Bitcrook can be integrated within your application OR it can be used directly from the [terminal/web browser.](#preview)

### Disclaimer

It is the end user's responsibility to obey all applicable local, state, and federal laws. Developers assume no liability and are not responsible for any misuse or damage caused by this program. By using Bitcrook, you agree to the previous statements.

### Changelog

Bitcrook v1.0.0 has been released as of October 10, 2023! Here are some of the highlights:

 - Switch web framework from Fiber to Go (enabled Bitcrook to be more lightweight)
 - Completely revamped the CLI output scheme
 - Switched secret configuration from .json to .env
 - Enabled API key support in web app
 - Fixed CORS error when accessing web app over LAN
 - Proper docker implementation
 - Binary release to support `go install`

Although this release hasn't changed much of the foundations of Bitcrook, it does illuminate
a path of where it will be taken next.

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
3. Run Bitcrook via Docker: `docker run --env-file .env -d -p 6174:6174 bitcrook`
4. Access the interface via `http://127.0.0.1:6174`

### Authentication

Some packages require an API key. Bitcrook will run without specifying any API keys; however, if you would like to implement these functionalites, create a file named `.env` in the Bitcrook root directory and format it like so:

``` conf
BITCROOK_MLSA=UNSPECIFIED # Melissa API key
BITCROOK_HIBP=UNSPECIFIED # HaveIBeenPwned API key
BITCROOK_IPTL=UNSPECIFIED # IP2Location API key
```

Replace the instances of `UNSPECIFIED` with the corresponding API key.

## Package-Types

### Court Cases

| Package                                                                                    | Description                                  |   Auth   | Location | Status |
| :-----------------: | -------------------------------------------- | :------: | -------- | :----: |
| CaseLaw             | Court Case Search                            |  `none`  | US | Functioning | 

### Discord

| Package                                                                                    | Description                                  |   Auth   | Location | Status |
| :----------------------------: | -------------------------------------------- | :------: | -------- | :----: |
| Discord Token Lookup           | Returns information regarding the passed token.                            |  `none`  | Global | Functioning | 

### IP Address

| Package                                                                                    | Description                                  |   Auth   | Location | Status |
| :---------------------: | -------------------------------------------- | :------: | -------- | :----: |
| IPV4 Address Lookup     | IPV4 Address Lookup                          |  `none`  | Global | Functioning |

### Multi-Use

| Package                                                                                    | Description                                  |   Auth   | Location | Status |
| :----------------------------: | -------------------------------------------- | :------: | -------- | :----: |
| Have I Been Pwned              | Email and Password Vulnerability - (Breaches)|  `paid`  | Global | Functioning |
| Melissa                        | Lookups - Email, Phone Number, IP Address    |  `free`  | US | Functioning |
| IP2LOCATION                    | Whois Lookup, IP Lookup              |  `free`  | - | Functioning | 

### Tax ID

| Package                                                                                    | Description                                  |   Auth   | Location | Status |
| :------------: | -------------------------------------------- | :------: | -------- | :----: |
| Tax ID Lookup  | Returns public information regarding a Russian INN.   |  `none`  | Russia | Functioning |

### Username

| Package                                                                                    | Description                                  |   Auth   | Location | Status |
| :-------------: | -------------------------------------------- | :------: | -------- | :----: |
| Username Lookup | Username Lookup - (Comparable to Sherlock)   |  `none`  | Global | Functioning |

### Vehicle

| Package                                                                                    | Description                                  |   Auth   | Location | Status |
| :--------------: | -------------------------------------------- | :------: | -------- | :----: |
| VIN Lookup       | Vehicle Identification Number Lookup         |  `none`  | - | Functioning | 
