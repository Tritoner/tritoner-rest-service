# Tritoner Rest Service Application


## Getting Started
To get started install go on your computer. 
Instructions to do so can be found here: https://golang.org/dl/


## Running the Application
Run `go run *.go` 
This will serve the application on port `8080`


## Dynamic Reloading while working on the application
- install fresh: `go get github.com/pilu/fresh`
- run `fresh` in the project 
- now the project will be rebuilt whenever you save to reflect the latest changes. 
If you get a `command fresh not found` error, make sure your PATH variable includes the `bin` directory. Ex: If your GOPATH in `.bashrc` is `$HOME/go` then add `$HOME/go/bin` to the $PATH

NOTES: If there is a compilation error, it will still serve the old application so keep that in mind



