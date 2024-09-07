<a id="readme-top"></a>

[![LinkedIn][linkedin-shield]][linkedin-url]
[![Issues][issues-shield]][issues-url]
[![Contributors][contributors-shield]][contributors-url]
[![Stars][stars-shield]][stars-url]
[![Forks][forks-shield]][forks-url]
[![License][license-shield]][license-url]


  <br />
<div align="center">
  <a href="https://github.com/josafamarengo/k-cli">
    <img src="logo.png" width="80" height="80" alt="k-cli"/>
  </a>
  <br>

  <h3 align="center">K CLI</h3>

  <p align="center">
    <strong>K</strong> is a versatile command-line interface (CLI) tool designed to simplify the management and installation of essential development utilities.
     <br />
    <a href="https://github.com/github_username/repo_name"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <!--<a href="https://linkedin-k-cli.netlify.app/">View Demo</a>
    ·-->
    <a href="https://github.com/josafamarengo/k-cli/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/josafamarengo/k-cli/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>

</div>

## 🔍 Table of Contents

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
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>

<a id="about-the-project"></a>
## ℹ️ About The Project

![k-cli Home Screen Shot](./src/assets/img/screenshots/home.png)

I built this project to streamline the setup of various developer tools, enabling a more efficient workflow. It also serves as a skills assessment tool, allowing you to test your setup before using it in larger projects.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<a id="getting-started"></a>
## 🚀 Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites
- [Go](https://golang.org/)
- [Git](https://git-scm.com/)


### Installation

1. Clone the repo

```bash
git clone https://github.com/josafamarengo/k-cli.git
```

2. Go to project folder

```bash
cd k-cli
```

3. Build the project

```bash
go build
```

4. Run the application

```bash
./k
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<a id="roadmap"></a>
## 🗺️ Roadmap

See the [open issues](https://github.com/josafamarengo/k-cli/issues) for a list of proposed features (and known issues).

<a id="contributing"></a>
## 👥 Contributing

Any contributions you make are **greatly appreciated**. 

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

### How to Contribute

1. Fork this repository to your own GitHub account.
2. Follow the steps on [Getting Started Section](#getting-started).
3. Create a new branch based on the `develop` branch. Use the following naming conventions:
    - For new features:
      ```bash
        git checkout -b feature/your-feature
      ```
    - For bug fixes:
      ```bash
        git checkout -b bugfix/bug
      ```
    - For documentation, improvements, refactoring, and optimizations:
      ```bash
        git checkout -b requirement/your-requirement
      ```
    _Note: **Make sure to always [keep your develop branch up-to-date](#steps-to-keep-your-fork-updated) before creating a new branch!**_
4. Make your changes.
5. Please ensure all tests pass before submitting a pull request. 
    - Run the following command to execute the tests:
      ```bash
        go test ./...
      ```
    _Note: **If any test fails, make sure to resolve the issues before proceeding.**_
6. Commit your changes with a clear and concise commit message.
7. Keep your code up-to-date following [this steps](#steps-to-keep-your-fork-updated).
8. Push your changes to your forked repository.
9. After pushing your changes to your fork, open a pull request with the following details:
    - A clear and descriptive title (e.g., "Fix bug in authentication flow").
    - A summary of the changes you made.
    - Any relevant issue numbers (e.g., "Closes #42").

Please make sure to write clear commit messages and to follow our coding conventions. We appreciate your contributions and will review them as soon as possible!

### Steps to Keep Your Fork Updated
If you've forked this repository and want to keep your fork updated with the latest changes from the original repository, follow these steps:

1. First, make sure you have a reference to the original repository. You only need to do this once:
  ```bash
    git remote add upstream https://github.com/josafamarengo/k-cli.git
  ```
2. To get the latest changes from the original repository (not your fork), you need to fetch them:
  ```bash
    git fetch upstream
  ```
3. Check out your local develop branch and pull in the changes from the develop branch of the original repository:
  ```bash
    git checkout develop
    git pull upstream develop
  ```
4. Now that your local develop branch is up to date, switch to your feature branch and rebase it on top of develop:
  ```bash
    git checkout feature/your-feature
    git rebase develop
  ```
5. After successfully rebasing, push your updated feature branch to your fork:
  ```bash
    git push origin feature/your-feature --force
  ```

### Issue Tracker

If you encounter any bugs or have feature requests, please open an issue on our [Issue Tracker][issues-url]. Be sure to include a clear and concise description of the issue, any steps needed to reproduce the problem, and any relevant code snippets.

### Top contributors:

<a href="https://github.com/josafamarengo/k-cli/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=josafamarengo/k-cli" alt="contrib.rocks image" />
</a>


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<a name="license"></a>
## 📜 License

Distributed under the GNU General Public License v3.0. See `LICENSE` for more information.

<a name="contact"></a>
## 📧 Contact

[![Linkedin][linkedin-shield]][linkedin-url]
[![Site][site-shield]][site-url]

<a name="acknowledgments"></a>
## 🙏 Acknowledgments
Thank you to everyone who contributed and supported this project.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<div align="center">
  <sub>Built with ❤︎ by <a href="https://josafa.com.br">Josafá Marengo</a>
</div>

<!-- REPO LINK -->
[repo-url]: https://github.com/josafamarengo/k-cli
[issues-url]: https://github.com/josafamarengo/k-cli/issues

[contributors-shield]: https://img.shields.io/github/contributors/josafamarengo/k-cli.svg?style=flat
[contributors-url]: https://github.com/josafamarengo/k-cli/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/josafamarengo/k-cli.svg?style=flat
[forks-url]: https://github.com/josafamarengo/k-cli/network/members
[stars-shield]: https://img.shields.io/github/stars/josafamarengo/k-cli.svg?style=flat
[stars-url]: https://github.com/josafamarengo/k-cli/stargazers
[issues-shield]: https://img.shields.io/github/issues/josafamarengo/k-cli.svg?style=flat
[issues-url]: https://github.com/josafamarengo/k-cli/issues
[license-shield]: https://img.shields.io/badge/License-GPL%20v3-blue.svg
[license-url]: https://github.com/josafamarengo/k-cli/blob/main/LICENSE.md

<!-- SOCIAL LINKS -->
[linkedin-shield]: https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white
[linkedin-url]: https://linkedin.com/in/josafamarengo

[site-shield]: https://img.shields.io/badge/website-000000?style=flat&logo=Google-chrome&logoColor=white
[site-url]: https://josafa.com.br
