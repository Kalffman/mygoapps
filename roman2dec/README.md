# roman2dec

The program that validate and convert a roman numeral into decimal number.

Example
---

```bash
$ go run roman2dec.go I
> MDCCLXXVI -> 1
```

```bash
$ go run roman2dec.go MDCCLXXVI
> MDCCLXXVI -> 1776
```

```bash
$ go run roman2dec.go MMMCMXCIX
> MDCCLXXVI -> 3999
```

Validations
---

Empty
```bash
$ go run roman2dec.go
> At least one roman numeral required.
> Eg.: MMXXII VI X
```

Invalid
```bash
$ go run roman2dec.go 2022
> 2022 -> is not valid roman numeral
```


:)
---