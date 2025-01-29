---
weight: 49
---
# Dates
```go
// Go fournit un support complet des fonctions usuelles
// de gestion du temps et des durées. Voici quelques
// exemples.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // Commençons pour récupérer l'heure actuelle
    // We'll start by getting the current time.
    now := time.Now()
    p(now)

    // On peut construire un structure `time` en fournissant année, mois, jour, etc. Les `time`s sont toujours associés à un lieu, c'est à dire  à un fuseau horaire.
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // On peut extraire les différentes valeurs des composantes du temps comme on s'y attend.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // Une méthode `Weekday` est également disponible pour connaitre le jour de la semaine (Monday-Sunday)
    p(then.Weekday())

    // Ces méthodes comparent deux instances de `time`, pour savoir si la première est arrivée respectivement avant, après ou au même moment que la seconde.
    // These methods compare two times, testing if the
    // first occurs before, after, or at the same time
    // as the second, respectively.
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // La méthode `Sub` renvoie une `Duration`, qui représente l'intervalle écoulé entre les dates contenues dans deux isntances de `time`.
    diff := now.Sub(then)
    p(diff)

    // On peut calculer la longueur de la durée dans plusieurs unités.
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // On peut aussi utiliser `Add` pour avancer à une date selon une certaine durée, ou reculer avec `-`.
    p(then.Add(diff))
    p(then.Add(-diff))
}
```
```sh
$ go run time.go
2012-10-31 15:50:13.793654 +0000 UTC
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
25891h15m15.142266763s
25891.25420618521
1.5534752523711128e+06
9.320851514226677e+07
93208515142266763
2012-10-31 15:50:13.793654 +0000 UTC
2006-12-05 01:19:43.509120474 +0000 UTC

# Ensuite nous allons regarder l'idée liée de temps
# relativement au temps Unix.
```
