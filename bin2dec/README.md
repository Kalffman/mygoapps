# Bin2Dec

A script that convert binary strings, separeted by space " ", in a decimal number and 
validate if this follow the binary number pattern.

Example
---


```bash
$ go run bin2dec.go 00000000
> 00000000 -> 0 #result
```

```bash
$ go run bin2dec.go 11111111
> 11111111 -> 256
```


```bash
$ go run bin2dec.go 00000000 11112111
> 00000000 -> 0
> 11112111 -> is not a binary string
```

That's all Folks :)
---