## Quiz Game

First excercise of Gophercises is to build a Quiz Game. The idea is build a command line quiz to ask simple math questions and output the final score. Read more about the offical exercise [here](). I also wrote a blog post on my solution [here]()

- [Quiz Game](#quiz-game)
  - [Installation](#installation)
  - [Prerequisite](#prerequisite)
  - [Running It](#running-it)

### Installation

Steps mentioned below are for Ubuntu 18.04

- Follow the steps [here](https://golang.org/doc/install)
- sudo nano ~/.profile.d
- Add two lines 
  - export GOPATH=$HOME
  - export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
- Ctrl + S
- Ctrl + X


### Prerequisite

Completed the following
- [Tour of Go](https://tour.golang.org/welcome/4)

### Running It

- cd quiz-game
- go run main.go
- Additionally you could do
  - go run main.go --help  // to see the command line flags available