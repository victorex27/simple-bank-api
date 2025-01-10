<a id="readme-top"></a>




<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/victorex27/simple-bank-api">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Simple bank api</h3>

  <p align="center">
    A Simple Bank Api: Learning to build with Golang.
    <br />
    <a href="https://github.com/victorex27/simple-bank-api"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/victorex27/simple-bank-api">View Demo</a>
    ·
    <a href="https://github.com/victorex27/simple-bank-api/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/victorex27/simple-bank-api/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
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
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

This project aims to develop a Simple banking application to micmic accounts and transactions in a single bank

Key features include Sign In, Transfers, Withdrawals, Deposits.

### Problem Statement
Using this application as a tool for learning how to create backend applications in golang

### Goals
#### Business Goals
* Secure Authentication: Ensure a robust and seamless authentication process, safeguarding user data at every step.
* Personalization: Collect essential user information during onboarding to tailor services to individual needs.
* Data Protection: Implement strong authentication mechanisms to protect user information and maintain trust.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With


* [![Golang][Golang]][Golang-url]
* [![Docker][Docker]][Docker-url]
* [![Postgres][Postgres]][Postgres-url]
* [Golang Migrate]
* [Sqlc]


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites

You need to have the following installed and setup before installing the application.
* Installation of npm on your local machine
  ```sh
  npm install npm@latest -g
  ```
* Installation of yarn

* Set up a SendGrid account and a SendGrid Api key, Take note of the api key and the sender parameters.
* Set up a Google cloud storage bucket. Take note of the project id and bucket name.
* Set up a Google service account and download the credentials. This service account must have the relevant rights for the bucket
* Set up an OAuth 2.0 Client ID in google cloud. Copy the Client Id and Client Secret. Google OAuth will be used as a sign in option on the application.
* Set Up Mysql and take note of the credentials.
* Set up a Firebase Project. Look for cloud Messaging inside te project and enable it. Copy and store the sdk detail. Go to the Project Overview page, Select "Service Account" and generate the configuration based on the environment you are going to be using, in our case nodejs



### Installation


go mod tidy to clean up our imports

go get github.com/stretchr/testify to install a package


<!-- ROADMAP -->
## Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [ ] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
    - [ ] French
    - [ ] Spanish

See the [open issues](https://github.com/victorex27/simple-bank-api/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>






<!-- CONTACT -->
## Contact

Amaobi Obikobe - [@amaobiwrites](https://x.com/amaobiwrites) - aobikobe@gmail.com

Project Link: [https://github.com/victorex27/simple-bank-api](https://github.com/victorex27/simple-bank-api)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

I would like to give credit to these amazing resources without which, this project would not have been a success!

* [Img Shields](https://shields.io)
* [Setting up SendGrid Account](https://www.adarsha.dev/blog/nestjs-sendgrid-email-service)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[product-screenshot]: images/screenshot.png
[Golang]: https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png
[Golang-url]: https://go.dev/
[Golang Migrate]: https://github.com/golang-migrate/migrate
[Docker]: ./images/docker-logo-blue.svg
[Docker-url]: https://www.docker.com/
[Postgres]: ./images/postgres.png
[Postgres-url]: https://www.postgresql.org/
[Sqlc]: https://docs.sqlc.dev/en/latest/overview/install.html
