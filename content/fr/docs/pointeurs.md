---
weight: 18
---
# Pointeurs
```go
// Go supporte les [pointeurs](https://fr.wikipedia.org/wiki/Pointeur_(programmation)),
// ce qui permet de passer des références à des valeurs
// dans les programmes

package main

import "fmt"

// Nous allons voir comment fonctionnent les pointeurs,
// par opposition aux valeurs, dans deux fonctions :
// `zeroval` et `zeroptr`.
// `zeroval` a un paramètre, donc les arguments lui sont
// passés par valeur. `zeroval` aura une copie de  `ival`
// distincte de celle de la fonction appellante.
func zeroval(ival int) {
    ival = 0
}

// `zeroptr` a cette fois-ci un paramètre `*int`,
// c'est-à-dire un pointeur sur un `int`. Le `*iptr` dans
// le corps de la fonction permet de _deréférencer_ le
// pointeur de son adresse mémoire pour obtenir la valeur
// à cette adresse. Assigner une valeur à un pointeur
// déréférence change la valeur à l'adresse référencée.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // La syntaxe `&i` donne l'adresse mémoire de  `i`,
    // c'est à dire un pointeur vers `i`.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // On peut afficher les pointeurs également.
    // Pointers can be printed too.
    fmt.Println("pointer:", &i)
}
```
```sh
# `zeroval` ne change pas le  `i` dans `main`, mais
# `zeroptr` le fait car elle a une référence vers 
# l'adresse mémoire de cette variable.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
```
