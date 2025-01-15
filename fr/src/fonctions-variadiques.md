# Fonctions variadiques
```go
// Les [_fonctions variadiques_](https://fr.wikipedia.org/wiki/Fonction_variadique)
// peuvent être appelées avec n'importe quel nombre
// d'arguments. Par exemple `fmt.Println` est une
// fonction variadique courante.

package main

import "fmt"

// Voici une fonction qui va prendre un nombre arbitraire
// d'`ints` comme arguments.
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // Les fonctions variadiques peuvent être appelées de
    // la manière habituelle, avec des arguments
    // individuels.
    sum(1, 2)
    sum(1, 2, 3)

    // Si vous avez déjà plusieurs arguments dans une
    // slice, vous pouvez les appliquer à la fonction
    // variadique en utilisant `func(slice...)`
    // comme ceci.
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```
```sh
$ go run variadic-functions.go 
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10

# Un autre aspect clé des fonctions en Go, c'est la 
# possibilité d'en faire des closures, ce que nous allons
# regarder ensuite.
```
