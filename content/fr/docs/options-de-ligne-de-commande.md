---
weight: 61
---
# Options de ligne de commande
```go
// Les options de [ligne de commande](https://fr.wikipedia.org/wiki/Interface_en_ligne_de_commande#Option)
// sont une manière courante de préciser des options pour des programmes en ligne de commande.
// Par exemple, dans `wc -l`, le `-l` est une option.

package main

// Go fournit un package `flag` qui permet d'analyser des options simples. Nous utiliserons ce package pour implémenter notre programme d'exemple.
import "flag"
import "fmt"

func main() {

    // Des déclarations d'options simples sont disponibles pour les string, les entiers et les booléens. Ici on déclare une option `word` avec comme valeur par défaut `"foo"`, et une courte description. La fonction `flag.String` renvoie un pointeur sur la chaîne (pas une chaîne directement). Nous verrons comment utliser ce pointeur après.
    wordPtr := flag.String("word", "foo", "une chaine")

    // On déclare des options `numb` et `fork`, en utilisant une approche similaire.
    numbPtr := flag.Int("numb", 42, "un int")
    boolPtr := flag.Bool("fork", false, "un booleen")

    // Il est également possible de déclarer l'option en utilisant une variable déjà déclarée ailleurs dans le programme. A noter que l'on doit passer un pointeur à la fonction de déclaration d'option.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Une fois que toutes les options sont déclarées, il faut appeller `flag.Parse()` pour exécuter l'analyse des paramètres de ligne de commande.
    flag.Parse()

    // Ici, on va afficher les options analysées ainsi que les éventuels paramètres restants. A noter qu'il faut déréférencer les pointeurs, par ex. avec `*wordPtr` pour obtenir la véritable valeur des options.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
```
```sh
# Pour expérimenter le programme de test d'options de 
#ligne de commande, le plus simple est de le compiler
# puis d'exécuter le binaire ensuite.
$ go build command-line-flags.go

# Essayer le programme en donnant tout d'abord des 
# valeurs pour toutes les options.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Notez que si vous ne précisez pas certaines options,
# elles prennent automatiquement leurs valeurs par défaut.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Des arguments restants peuvent être fournis après toutes
# les options
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Notez que le package `flag` a besoin que toutes les 
# options apparaissent avant les arguments
# 'positionnels'. Sinon, les options sont traités comme
# des arguments positionnels.
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Utiliseez les options `-h` ou `--help` pour obtenir une
# aide générée automatiquement au sujet des paramètres de
# ligne de commande.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Si vous fournissez une option qui n'était pas précisée 
# au package `flag`, le programme affichera un message
# d'erreur puis le texte d'aide à nouveau.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Ensuite nous allons regarder les variables
# d'environnement, une autre manière courante de
# paramètrer les programmes.
```
