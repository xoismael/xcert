# xcert
A cli tool that takes host:port as stdin and output the certificate subject common name.

# Installation
```
go install github.com/xoismael/xcert@latest
```

# Basic Usage:
```
echo example.com | xcert
```
or
```
cat hosts.txt | xcert
```
