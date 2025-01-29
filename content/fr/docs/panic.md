---
weight: 42
---
# Panic
```go
// Un `panic` signifie typiquement que quelque chose
// s'est mal passé et de manière inattendue. On l'utilise
// principalement pour faire échouer rapidement des
// erreurs qui ne devraient normalement pas arriver, ou
// que nous ne pouvons pas traiter.

package main

import "os"

func main() {
    // Nous allons utilise panic dans les exemples de ce
    // site pour traiter les cas imprévus. Ce
    // programme-ci est le seul du site qui va paniquer
    // durant son déroulement normal.
    panic("a problem")

    // Une utilisation courante de `panic`, c'est
    // d'abandonner si une fonction renvoie une valeur
    // d'erreur que nous ne savons (ou ne voulons) pas
    // traiter. Voici un exemple de panique si on a une
    // erreur inattendue lors de la création d'un nouveau
    // fichier.
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
```
```sh
# Exécuter ce programme va amener à une panique, qui 
# affiche un message d'erreur et une trace d'exécution, 
# et qui le programme en renvoyant une valeur de status
# non nulle.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# A noter que contrairement à certains langages qui 
# utilisent des exceptions pour la gestion de nombreuses
# erreurs, en Go il est idiomatique d'utiliser des
# valeurs de retour qui indique le status d'erreur dès 
# que possible.
```
