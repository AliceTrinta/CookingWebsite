# Cooking website
> Website for save and share recipes.

[![Generic badge](https://img.shields.io/badge/Version-1.0.0-<COLOR>.svg)](https://shields.io/)
[![Build Status][travis-image]][travis-url]
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Release History](#release-history)
* [Contact](#contact)
* [Contributing](#contributing)



<!-- ABOUT THE PROJECT -->
## About The Project
A simple recipe website for create profiles to save and share recipes, also has a fuction to show all recipes stored in the database.

### Built With
List of major frameworks used.
* [GoMacaron](https://go-macaron.com/)
* [Mercurius](https://github.com/novatrixtech/mercurius)
* [Bootstrap](https://getbootstrap.com)



<!-- GETTING STARTED -->
## Getting Started

Follow pass-by-pass to install the project in your machine.

### Prerequisites

Firs: Install MySql database in your machine.

### Installation

Clone repository
```sh
git clone https://github.com/AliceTrinta/cooking-website.git
```

### Running

First step: configure the archive app.ini according to your prefenrences.
```sh
cooking-website
              |
              ---conf
                    |
                    ---app.ini
```

Second step: create a database in MySql called web-example.
```sh
CREATE SCHEMA `web-example`;
```

Third step: run main.go by terminal on folder's project
```sh
go run main.go
```



<!-- USAGE -->
## Usage

Study content to help people understand Go for web development with frameworks.



<!-- RELEASE HISTORY -->
## Release History

* 1.0.0
    * Initial project with HTML pages template



<!-- CONTACT -->
## Contact

Alice Trinta – [@malicetrinta](https://www.instagram.com/malicetrinta/) – maria.trinta@aluno.cefet-rj.br

[https://www.researchgate.net/profile/Maria_Alice_Lima2/publications](https://www.researchgate.net/profile/Maria_Alice_Lima2/publications)
Project Link: [https://github.com/AliceTrinta/cooking-website](https://github.com/AliceTrinta/cooking-website)



<!-- CONTRIBUTING -->
## Contributing

1. Fork it (<https://github.com/AliceTrinta/cooking-website>)
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- Markdown link & img dfn's -->
[npm-image]: https://img.shields.io/npm/v/datadog-metrics.svg?style=flat-square
[npm-url]: https://npmjs.org/package/datadog-metrics
[npm-downloads]: https://img.shields.io/npm/dm/datadog-metrics.svg?style=flat-square
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[wiki]: https://github.com/yourname/yourproject/wiki
