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
