<!-- PROJECT LOGO -->
<p align="center">
  <a href="">
    <img src="./assets/gban.png" alt="Logo">
  </a>
  <p align=center>
    <a target="_blank" href="https://goreportcard.com/report/github.com/audioo/goseek" title="report"><img src="https://goreportcard.com/badge/github.com/audioo/goseek"></a>
    <a href="https://pkg.go.dev/github.com/audioo/goseek"><img src="https://pkg.go.dev/badge/github.com/audioo/goseek.svg" alt="Go Reference"></a>
    <a href="https://inventory.raw.pm/"><img src="https://inventory.raw.pm/img/badges/Rawsec-inventoried-FF5050_flat.svg" alt="Rawsec&#39;s CyberSecurity Inventory"></a>
    <a target="_blank" href="#" title="VERSION"><img src="https://img.shields.io/badge/Version-0.6.2-blue.svg"></a>
  </p>
  <p align="center">
    Open-Source Intelligence and Operations Security Framework Written in Golang.
    <br />
    <a href="https://audioo.github.io/goseek"><strong>See more »</strong></a>
    <br />
</p>


<!-- ABOUT THE PROJECT -->
## About The Project

**I have always found Open-Source Intelligence and Operations Security to be intriguing. I started learning Golang in the beginning of February of 2021 and have been looking for new projects to help strengthen my knowledge. I also wanted these projects to keep me engaged over time and prove somewhat useful instead of making just another calculator. I hope to build on GoSeek and expand it's use cases in a way that others can benefit**

Features:

- [x] Username Lookup - Inspired by [Sherlock](https://github.com/sherlock-project/sherlock), includes account deletion sites.
- [x] IP Lookup - Uses [ip-api](https://ip-api.com/)
- [x] License Plate & VIN Lookup - Uses [htmlquery](https://github.com/antchfx/htmlquery)
- [x] Info Cull - Data narrowing using [gofpdf](https://github.com/jung-kurt/gofpdf), inspired by [HINTS](https://github.com/jadekeys/hints)
- [x] Fake Identity Generator - Name (male, female, or gender neutral), Location, Date of Birth, Username, and Password 

<!-- USAGE EXAMPLES -->
## Install and Run

Download via 'go get':
<code>go get -u github.com/audioo/goseek</code>

Run in terminal:
<code>goseek [command] [args]</code>

**OR**

Download via 'git clone':
<code>git clone https://github.com/audioo/goseek.git</code>

Navigate into GoSeek directory:
<code>cd goseek</code>

Run via 'go run': <code>go run main.go [command] [args]</code>
or build: <code>go build main.go</code>

## Running with Docker
1 - Build: <code>docker build -t goseek . </code>

2 - Run w/ Flags: <code>docker run --rm -it goseek [command] [args] . </code>

3 - Example: <code>docker run --rm -it goseek user audioo n .</code>

### Attributions

- Icon made using assets by [FreePik](https://www.freepik.com) from [www.flaticon.com](https://www.flaticon.com/)