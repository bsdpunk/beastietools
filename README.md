# beastietools
A shell for all those odds and ends network tools


## Description
beastietools is a swiss army shell of odds and ends of network tools

## Usage
beastietools with no arguements drops you into the shell allowing you to use tab completion:
```bash
beastietools> 
arp    extip  ls     quit   
beastietools> extip 
Successfully discovered 1 services:
  Device: Netgear CG3000DCR
    External IP address: 68.52.132.49
beastietools> quit
```

You may also use arguements if you already know what you want to do or if you want to script:
```bash
dusty@xor:~$ beastietools extip
Successfully discovered 1 services:
  Device: Netgear CG3000DCR
    External IP address: 68.52.132.49
```

## Install
Readline.h is a requirement in unix enviroments, so readline-devel or readline-dev packages.

To install, use `go get`:

```bash
$ go get -d github.com/bsdpunk/beastietools
```
You may need to install other Libraries as well:
```bash
$ go get -d github.com/huin/goupnp
```

## Contribution

1. Fork ([https://github.com/bsdpunk/beastietools/fork](https://github.com/bsdpunk/beastietools/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Create a new Pull Request

## Author

[bsdpunk](https://github.com/bsdpunk)
