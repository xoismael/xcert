# xcert
A cli tool that takes host:port as stdin and output the certificate subject common name.

# Installtion
```console
go install github.com/xoismael/xcert@latest
```

# Basic Usage:
```console
echo example.com | xcert
```
or
```console
cat hosts.txt | xcert
```
