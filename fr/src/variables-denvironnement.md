# Variables d'environnement
```go
// Les [variables d'environnement](https://fr.wikipedia.org/wiki/Variable_d%27environnement) sont un mécanisme universel pour fournir des [informations de configuration à des programmes Unix](https://12factor.net/fr/config). Regardons comment créer, récupérer et lister des variables d'environnement.

package main

import "os"
import "strings"
import "fmt"

func main() {
    // Pour créer une paire clé/valeur, on utilise `os.Setenv`. Pour récupérer une valeur à partir d'une clé, il y a `os.Getenv`. Elle renvoie une chaîne vide si la clé n'est pas présente dans l'environnement.
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    // `os.Environ` permet de lister toutes les paires clé/valeur de variables d'environnement. Elle renvoie une slice de string sous la forme `CLE=valeur`. On peut obtenir les clés et valeurs avec `strings.Split`. Ici, on affiche toutes les clés.
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }
}
```
```sh
# Exécuter le programme montre que nous choisissons une 
# valeur pour FOO, mais que `BAR` est vide.
$ go run environment-variables.go
FOO: 1
BAR: 

# La liste des clés de l'environnement dépendra de
# votre machine.
TERM_PROGRAM
PATH
SHELL
...

# So nous affectons `BAR` dans l'environnement avant de
# lancer le programme, il récupérera sa valeur.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
```
