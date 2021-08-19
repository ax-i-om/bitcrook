<p align="center">
  <a><img src="./images/omni.png" width=180 height="180"></a>
  <h1 align="center">Omniscient - Information Gathering Apparatus</h1>
  <p align="center">
    <a href="https://goreportcard.com/report/github.com/maraudery/omniscient"><img src="https://goreportcard.com/badge/github.com/maraudery/omniscient" alt="Go Report Card"></a>
    <a><img src="https://img.shields.io/badge/tests-7&#47;8-orange.svg" alt="s"></a>
    <a><img src="https://img.shields.io/badge/version-0.4-blue.svg" alt="s"></a>
    <a href="https://pkg.go.dev/github.com/maraudery/omniscient"><img src="https://pkg.go.dev/badge/github.com/maraudery/omniscient.svg" alt="Go Report Card"></a><br>
    <a href="https://www.buymeacoffee.com/maraudery"><img src="https://cdn.buymeacoffee.com/buttons/default-red.png" height="40" width="170"></a>
  </p><br>
</p>

## Table of Contents

- [Information](#information)
  - [About](#about)
  - [Disclaimer](#disclaimer)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Implementation](#implementation)
  - [Testing](#testing)
  - [Attributions](#attributions)
- [Package Types](#package-types) *a-z*
  - [Court Cases](#court-cases)
  - [Discord](#discord)
  - [IP Address](#ip-address)
  - [Multi-Use](#multi-use)
  - [Username](#username)
  - [Vehicle](#vehicle)



## Information

### About

Omniscient, an information gathering apparatus, is a conglomerate of tools including custom algorithms, API wrappers, etc... in order to make the reconnaissance process significantly quicker. Some packages do require an authentication key and others do not. See the [Package Types](#package-types) tables for more information. Omniscient can be integrated within your application OR it can be used directly from the terminal.

### Disclaimer

It is the end user's responsibility to obey all applicable local, state, and federal laws. Developers assume no liability and are not responsible for any misuse or damage caused by this program. By using Omniscient, you agree to the previous statements.

### Installation

 - Fetch the repository via 'git clone': `git clone https://github.com/maraudery/omniscient.git`

### Usage

1. Navigate into the root directory of Omniscient.
2. In your preferred terminal, enter and run: `go run main.go`
3. After running the aforementioned command, all dependencies will be installed and usage help will be printed to the console.
4. When prompted, allow the application to communicate on your network.
5. Navigate to `http://localhost:6174/`

### Implementation

Instructions/Documentation are provided for each and every package, all you have to do is find what you need in the [Package Types](#package-types) section.

### Testing

Omniscient is currently passing all tests; however, I have provided instructions for properly running the tests if you would like to do so.

1. In the root directory of Omniscient, create a file named `keyconfig.json`
2. In `keyconfig.json`, paste the following text:
``` json
{
    "melissaKeyCred": "Paste Melissa Key with Credits Here",
    "hibpKey": "Paste Have I Been Pwned Key Here",
    "dataGovKey": "Paste Data.gov Key Here"
}
```
3. Paste in your API keys. The test will fail without a valid API key.
4. In your preferred terminal, enter and run `go test ./...`

### Attributions

[Social Preview](./images/card.jpg) created by [TheOneTrueDude](https://github.com/FirstTrueDude)

## Package-Types

### Court Cases

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Case Law](https://github.com/maraudery/omniscient/tree/main/pkg/noauth/caselaw)           | Court Case Search                            |  `none`  |

### Discord

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Discord Token Lookup](https://github.com/maraudery/omniscient/tree/main/pkg/noauth/discord)| Discord Token Lookup                        |  `none`  |

### IP Address

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [IPV4 Address Lookup](https://github.com/maraudery/omniscient/tree/main/pkg/noauth/ip)     | IPV4 Address Lookup                          |  `none`  |

### Multi-Use

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Have I Been Pwned](https://github.com/maraudery/omniscient/tree/main/pkg/authpaid/hibp)   | Email and Password Vulnerability - (Breaches)|  `paid`  |
| [Melissa](https://github.com/maraudery/omniscient/tree/main/pkg/authfree/melissa)          | Lookups - Email, Phone Number, IP Address    |  `free`  |


### Username

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Username Lookup](https://github.com/maraudery/omniscient/tree/main/pkg/noauth/userlookup) | Username Lookup - (Comparable to Sherlock)   |  `none`  |

### Vehicle

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [VIN Lookup](https://github.com/maraudery/omniscient/tree/main/pkg/noauth/vin)             | Vehicle Identification Number Lookup         |  `none`  |
