# Go Language Essentials
Here are the basic essentials things needed for getting started with Go Language.

### How to install Go
- If you are using WSL2, follow this [Medium](https://medium.com/@benzbraunstein/how-to-install-and-setup-golang-development-under-wsl-2-4b8ca7720374) post (_if you also are using zsh, paste the lines into .zshrc_);
- If you are using MAC, Windows or Linux, go to the [GoLang](https://golang.org/doc/install) website.

### Getting started
#### Initializing a module:
Every new application need a module manager (something like a *package.json*). So run the following command to create the ```go.mod``` file:
```bash
go mod init {module_name}
```
Usually ```{module_name}``` will be where your application is located, something like ```github.com/Guilospanck/GoLanguageEssentials```.
This command is similar to *yarn init*

---
#### Import modules from the internet:
You can go to the [Go dev](https://pkg.go.dev/) website and search for the package you need (something similar to *yarn*, *npm*, *pypi*). Search for
the module you want and then import it in you ```.go``` file.
```go
import "module/name"
```
After doing this, you need to run a command to actually install that module (similar to *npm install* or *yarn install*):
```bash
go mod tidy
```
---
#### To run your .go file:
```bash
go run file.go
```

#### Testing in Go:
- You have to create a new file with the name of the file you are testing, appending ```_test``` at the end. Example:
```
greetings.go  => greetings_test.go
```
To run your tests, go to the directory wanted and run:
```bash
go test
or
go test -v  # to verbose
```

---
#### Some observations about Golang:
- The first line when you're writing a code in ```Go``` is ```package {packageName}```. It must exists only a package per directory. If you wanna have a different package name for a file, you must change its directory. Every Golang program must have a file in which there is the main package with the ```main``` function.

- Functions that begin with an upper-case letter are *exported* by default, hence it can be used in other packages that import our package.

- The ```:=``` (*short assignment statement*) can be used in place of a ```var``` declaration with implicit type. Outside a function, every statement begins with a keyword (var, func, and so on), and so the ```:=``` construct is not available.

- Each source file can define its own ```init``` function to set up whatever state is required (actually, each file can have multiple ```init``` functions).
```init``` is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.
