# pap

A simplistic **pa**permc hel**p**er.

## Table of contents

- [pap](#pap)
  - [Table of contents](#table-of-contents)
  - [Examples](#examples)
    - [Download the latest papermc jarfile](#download-the-latest-papermc-jarfile)
    - [Sign the EULA](#sign-the-eula)
    - [Generate a script to run the jarfile](#generate-a-script-to-run-the-jarfile)
  - [Why though?](#why-though)
  - [Install](#install)
    - [Build Dependencies](#build-dependencies)
    - [Arch linux](#arch-linux)
    - [Unix](#unix)
      - [Unix - From Releases](#unix---from-releases)
        - [Unix - System wide from releases](#unix---system-wide-from-releases)
        - [Unix - Local from releases](#unix---local-from-releases)
      - [Unix - From Source](#unix---from-source)
        - [Unix - System wide from source](#unix---system-wide-from-source)
        - [Unix - Local from source](#unix---local-from-source)
    - [Windows](#windows)
      - [Windows - From Releases _(recommended)_](#windows---from-releases-recommended)
      - [Windows - From Source](#windows---from-source)
    - [Common issues](#common-issues)
      - [Local installation not found](#local-installation-not-found)
        - [Bash](#bash)
        - [Zsh](#zsh)
        - [Fish](#fish)
  - [Contributing](#contributing)
  - [Dependencies](#dependencies)

## Examples

### Download the latest papermc jarfile

`pap download`

### Sign the EULA

`pap sign`

### Generate a script to run the jarfile

`pap script --jar server.jar`

## Why though?

pap was created to simplify the annoying parts of making a minecraft server, so you can do everything from an easy CLI.

And one of the most annoying parts, is downloading the latest papermc jarfile, which is the root purpose of pap.

pap looks at the papermc api, and gets the latest version automatically, without you needing to navigate to the papermc website and copy a link and paste it into a `wget` command.

## Install

### Build Dependencies

If you are obtaining pap from source, you will need these dependencies:

- [Go](https://go.dev/) 1.18 or later
- [Git](https://git-scm.com/)

### Arch linux

If you wish, pap can be installed from the AUR:

```sh
yay -S pap
```

### Unix

#### Unix - From Releases

You can go to the [latest release](https://github.com/talwat/pap/releases/latest)
and download the fitting binary for your system from there.

pap is available on esentially all architectures and operating systems, so you will rarely need to compile it from source.

Simply mark the downloaded binary as executable and move it.

##### Unix - System wide from releases

```sh
sudo mv pap* /usr/local/bin/pap
sudo chmod +x /usr/local/bin/pap
```

##### Unix - Local from releases

```sh
mv pap* ~/.local/bin/pap
chmod +x ~/.local/bin/pap
```

#### Unix - From Source

First, clone pap:

```sh
git clone https://github.com/talwat/pap
cd pap
```

Switch to the latest tag _(optional)_:

```sh
git tag # get all tags
git checkout <tag>
```

And then build:

```sh
go build .
```

Finally, move it into your binary directory:

##### Unix - System wide from source

```sh
sudo mv pap /usr/local/bin/pap
sudo chmod +x /usr/local/bin/pap
```

##### Unix - Local from source

```sh
mv ~/.local/bin/pap
chmod +x ~/.local/bin/pap
```

### Windows

pap **does** work on windows, but the install steps listed are for unix-like systems.

#### Windows - From Releases _(recommended)_

If you want to download from releases, download the fitting windows exe and [put it into path](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows#:~:text=Go%20to%20%22My%20computer%20%2D%3E,exe%20's%20directory%20into%20path.).

#### Windows - From Source

First, clone pap:

```sh
git clone https://github.com/talwat/pap
cd pap
```

Switch to the latest tag _(optional)_:

```sh
git tag # get all tags
git checkout <tag>
```

And then build:

```sh
go build .
```

Finally, [put it into path](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows#:~:text=Go%20to%20%22My%20computer%20%2D%3E,exe%20's%20directory%20into%20path.).

### Common issues

#### Local installation not found

Usually this is because `~/.local/bin` is not in PATH.

You can add `~/.local/bin` to PATH through your shell:

##### Bash

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
```

##### Zsh

```sh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
```

##### Fish

Look at the [fish docs](https://fishshell.com/docs/current/tutorial.html#path) for more detailed instructions.

```sh
fish_add_path $HOME/.local/bin
```

## Contributing

Anyone is welcome to contribute, and if someone can port pap to various package managers, it would be greatly appreciated.

## Dependencies

- [schollz/progressbar](https://github.com/schollz/progressbar)
- [urfave/cli](https://github.com/urfave/cli)
