# cbox

<div align="center"><img src="https://raw.githubusercontent.com/dplabs/cbox/master/docs/cbox.png" width="250px"/></div>

[![Go Report Card](https://goreportcard.com/badge/github.com/dplabs/cbox)](https://goreportcard.com/report/github.com/dplabs/cbox)
[![Travis](https://travis-ci.org/dplabs/cbox.svg?branch=master)](https://travis-ci.org/dplabs/cbox)
[![Sonar](https://sonarcloud.io/api/project_badges/measure?project=dplabs_cbox&metric=alert_status)](https://sonarcloud.io/dashboard?id=dplabs_cbox)

**cbox** helps you organizing and finding those useful commands you find time to time and would really love to store for future occasions.

## Introduction

**cbox** is an easy way to store (and search for) those useful commands that you find from time to time and tend to forget. It's also the perfect way to store a set of commands most frequently used in your job and that would be awesome to share with your colleagues, to make everyone's life easier.

### Features

- Store and search commands or snippets in a single place, and just a few key strokes away.
- Assign an small title and description to your commands, as well as keeping a reference to the website where you find it (who knows what else could you need in the future)
- Assign tags to your commands so they are easier to find later on.
- Organize your commands in spaces, i.e. based on project, technology, environment... whatever fits you best.
- Share (and backup) those spaces with the community using **cbox cloud**.
- Hundreds of useful tagged and categorized commands already available to the community in the **cbox cloud**

### Demo / Screenshots

TODO

## Getting started

### Installation

In order to start using **cbox** you have the following options available:

#### Option 1: Homebrew (TODO)

    brew install cbox

#### Option 2: Snap (TODO)

    snap install cbox

#### Option 3: install a pre-compiled release

Go to https://github.com/dplabs/cbox/releases and download latest precompiled release.

We create precompiled packages for all major platforms (RPM & DEB packages also available).

#### Option 3: install with `golang`'s cli

    go get github.com/dplabs/cbox

#### Option 4: re-install from source

If you want to build it yourself from sources, just use:

    make build

or

    make install

if you also want to install **cbox** in the default `golang` bin path.

### Quickstart guide

After installing **cbox** by any of the procedures described before, you are ready to go (a *default* space is created for you the first time you use cbox in a new environment):

    cbox command add

will let you add a new command to your default space

    cbox list 

will list all the content of your **default** space

    cbox search CRITERIA

will list all commands containing criteria as part of the command's code, title or description.

If you want to have a more in depth walkthrough of what **cbox** offers, please check our [tutorial](https://github.com/dplabs/cbox/wiki/Tutorial)

### More info

- [Spaces](https://github.com/dplabs/cbox/wiki/Spaces)
- [Selectors](https://github.com/dplabs/cbox/wiki/Selectors)
- [Organizations](https://github.com/dplabs/cbox/wiki/Organizations)
- [Settings](https://github.com/dplabs/cbox/wiki/Settings)
- [Cloud](https://github.com/dplabs/cbox/wiki/Cloud)
- [Shell automcompletion](https://github.com/dplabs/cbox/wiki/Shell-autocompletion)

## Contributing

Do you have any great idea and would like to make it happen in **cbox**? Great! Just for the repo, hack your solution and create a PR.

We're looking forward to see what can you achieve with **cbox**

For more tips on how to properly set **cbox** to use a *test* cloud, please refer to the [Contributing guidelines](https://github.com/dplabs/cbox/wiki/Contributing)

## About

https://cbox.dplabs.io

Daniel Pecos Martinez - https://danielpecos.com

### License

Copyright Daniel Pecos 2019
