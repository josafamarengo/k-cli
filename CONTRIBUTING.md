# Contributing

Any contributions you make are **greatly appreciated**. 

Thank you to everyone who contributed and supported this project.

## Development Setup

To get a local copy up and running follow these simple steps.

### Prerequisites
- [Go - 1.23 or later](https://golang.org/)
- [Git](https://git-scm.com/)

## Development Workflow

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