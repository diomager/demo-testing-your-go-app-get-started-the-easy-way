# Testing Your Go App: Get Started the Easy Way

## Introduction

This is the demo code corresponding to the post: [Testing your Go App: Get Started the Easy Way](https://rms1000watt.github.io/post/testing-you-go-app-get-started-the-easy-way).

## Contents

- [Download](#download)
- [Configure Git](#configure-git)
- [Commit Code](#commit-code)

## Download

```bash
git clone git@github.com:rms1000watt/demo-testing-your-go-app-get-started-the-easy-way.git

cd demo-testing-your-go-app-get-started-the-easy-way
```

## Configure Git

```bash
# Give execution privileges to the .githook
chmod a+x .githooks/pre-commit

# Copy the githook to the proper githooks directory
cp .githooks/pre-commit .git/hooks/pre-commit
```

## Commit Code

```bash
# Add arbitrary data to the repo so git will notice the change
echo "." >> dot.txt

# Add and commit the code. Watch the tests occur
git add .
git commit -m "testing precommit hook"
```