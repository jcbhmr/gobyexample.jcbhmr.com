---
weight: 21
---
# Interfaces
```go
// Les _interfaces_ sont des collections nommées de
// signatures de méthodes.

package main

import "fmt"
import "math"

// Voici une interface simple pour une forme géométrique.
type geometry interface {
    area() float64
    perim() float64
}

// Pour notre example, nous allons implémenter cette
// interface sur les types `rect` et `circle`.
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// Pour implémenter une interface en Go, il suffit
// simplement d'en implémenter toutes les méthodes.
// Ici, nous implémentons `geometry` chez `rect`s.
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    // Voici maintenant l'implémentation pour les `circle`.
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// Si une variable a une interface, alors on peut appeler
// les méthodes de l'interface. Voici une fonction
// générique qui utilise ceci pour travailler avec
// n'importe quelle `geometry`.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // Mes structures `circle` et `rect` implémentent
    // tout deux les interfaces de `geometry`, on peut
    // donc utiliser des instances de ces structures en
    // argument de `measure`.
    measure(r)
    measure(c)
}
```
```sh
$ go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793

# Pour en apprendre plus sur les interfaces en Go, vous
# pouvez lire ce [bon article](http://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go).
```
