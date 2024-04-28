<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- [![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url] -->
<!-- [![Stargazers][stars-shield]][stars-url] -->
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">Golang log generator</h3>

  <p align="center">
    Synthetic log generator writen in golang
  <br />

  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

The necessity to build this project comes from the idea of being able to test log aggregation systems (ELK stack, Grafana Loki) by being able to produce log lines at a high rate in a controlled manner. The application accomplishes this by making use of Golang's concurrency engine (Go routines) as well as using an external package (github.com/go-co-op/gocron)
to control intervals at which log linges are flushed down to flat file.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

[![Golang][Golang.dev]][golang-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started
This project contains a Dockerfile so it can be built and executed in a distributed environment if needed. It also contains docker-compose yaml file for ease of local execution.

### Prerequisites

Docker desktop needs to be installed on the system if aiming for local testing. In order to build the container image simply issue the following command.

* docker
  ```sh
  docker build . -t log-generator:0.0.5
  ```

### Installation

_The following steps are needed in order to build the image locally and launch the application. The application can be configured_

1. Clone the repo
   ```sh
   git clone git@gitlab.com:rodolfo-mora/go-carbon-docker.git
   ```
2. Build docker compose image
   ```sh
   docker build . -t log-generator:0.0.5
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage
Application behavior can be controlled by configuring two specific environment variables

* INTERVAL: Integer number (in ms) used to specify how often WORKERS will flush new log lines.
* WORKERS: Workers take care of generating one log line per INTERVAL. The more workers, the more log lines will be flushed down in parallel.

```docker
version: '3'
services:
  log-generator:
    image: log-generator:0.0.5
    environment:
      - WORKERS=1
      - INTERVAL=1000
    network_mode: host
    volumes:
    - ./vector:/etc/vector/
    - ./logs:/var/log/nginx/
    deploy:
      replicas: 1
```
<!-- _For more examples, please refer to the [Documentation](https://example.com)_ -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [ ] Add Changelog
- [ ] Update code to be imported as a github package
- [ ] Add querier that can execute queries in a locust.io style
- [ ] Update prometheus exporter for logger
- [ ] Add prometheus exporter for querier
- [ ] Add "components" document to easily copy & paste sections of the readme

See the [open issues](https://github.com/rodolfo-mora/golang-log-generator/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Rodolfo Mora Gonzalez - rodolfo.mora.gonzalez@gmail.com

Project Link: [https://github.com/your_username/repo_name](https://github.com/your_username/repo_name)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/rodolfo-mora/golang-log-generator/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/rodolfo-mora/golang-log-generator/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/rodolfo-mora-2214a9b7
[golang-url]: https://go.dev/
[Golang.dev]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white