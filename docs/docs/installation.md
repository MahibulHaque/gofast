---
sidebar_position: 2
---

# Installation


GoFast provides a convenient CLI tool to effortlessly set up your Go projects. Follow the steps below to install the tool on your system.

### Binary Installation

To install the GoFast CLI tool as a binary, run the following command:

```bash
go install github.com/mahibulhaque/gofast@latest
```

This command installs the Gofast binary, automatically binding it to your `$GOPATH`. 


> **Note** - If you are using Zsh, you will need to add it manually to your `.zshrc` file. 
> 
> After running the installation command, you need to update your PATH environment variable. To do this, you need to find out the correct `GOPATH` for your system. You can do this by running the following command: Check your `GOPATH`.
> ```bash
> go env GOPATH
> ```
>
> Then, add the following line to your `~/.zshrc` file:
>
> `GOPATH=$HOME/go PATH=$PATH:/usr/local/go/bin:$GOPATH/bin`
>
> Save the changes to your `~/.zshrc` file by running the following command:
> 
> ```bash
> source ~/.zshrc
> ```


### Building and Installing from Source

If you prefer to build and install GoFast directly from the source code, you can follow these steps:

Clone the repository from Github:
```bash
git clone https://github.com/mahibulhaque/gofast
```

Build the Gofast binary:
```bash
go build
```

Install in your $PATH to make it accessible system-wide:
```bash
go install
```

Verify the installation by running:
```bash
gofast version
```


This should display the version information of the installed Gofast CLI.

Now you have successfully built and installed GoFast from the source code.
