### Real2Centavos

Application that converts a float number, in Brazilian Reais (R$),
to number of coins, Reais cents, needed to fill the value

#### Example

```bash
$ go run real2centavos.go 1.49

> 2 coins -> 50 centavos
> 1 coins -> 25 centavos
> 2 coins -> 10 centavos
> 0 coins -> 5 centavos
> 4 coins -> 1 centavos
```

```bash
$ go run real2centavos.go 2.98

> 5 coins -> 50 centavos
> 1 coins -> 25 centavos
> 2 coins -> 10 centavos
> 0 coins -> 5 centavos
> 3 coins -> 1 centavos
```

:)