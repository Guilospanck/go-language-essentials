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

#### Compiling and installing the application
- To generate a binary executable, you can run (in your file directory):
```bash
go build
```
This will generate a binary file with the name of you Module. You can run this binary file as it follows:
```bash
./GoLanguageEssentials    (./{binary_name})
```
But now, to be able to run it from anywhere in your computer, you must install it:
```bash
go install
```
Now, you can run it from anywhere just typing:
```bash
GoLanguageEssentials    ({binary_name})
```

---
#### Some observations about Golang:
- The first line when you're writing a code in ```Go``` is ```package {packageName}```. It must exists only a package per directory. If you wanna have a different package name for a file, you must change its directory. Every Golang program must have a file in which there is the main package with the ```main``` function.

- Functions that begin with an upper-case letter are *exported* by default, hence it can be used in other packages that import our package.

- The ```:=``` (*short assignment statement*) can be used in place of a ```var``` declaration with implicit type. Outside a function, every statement begins with a keyword (var, func, and so on), and so the ```:=``` construct is not available.

- Each source file can define its own ```init``` function to set up whatever state is required (actually, each file can have multiple ```init``` functions).
```init``` is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.

#### Go's basic types:
- <code>bool</code>
- <code>string</code>
- <code>int</code> (all positive and negative values) =>  (int8 16 32 64)
- <code>uint</code> (only positive values) => (uint8 16 32 64 ptr)
- <code>byte</code> (alias for <code>uint8</code>)
- <code>rune</code> (alias for <code>int32</code>. It represents a Unicode code point)
- <code>float32</code> <code>float64</code>
- <code>complex64</code> <code>complex128</code>

For 32-bit systems, ```int``` ```uint``` and ```uintptr``` are 32 bits wide. For 64-bit systems they are 64 bits wide (usually).

##### Zero values:
Variables declared without an explicit initial value are given their *zero value*. The zero value is:
- <code>0</code> for numeric types;
- <code>false</code> for the boolean type, and
- <code>""</code> (the empty string) from strings.

