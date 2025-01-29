# Tempo
```go
// Go offre un supporto molto esteso per misurare il tempo
// e la durata dei task. Ecco qualche esempio.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // Iniziamo prendendo il tempo attuale.
    now := time.Now()
    p(now)

    // È possibile costruire una struct di tipo `time`
    // passando anno, mese, giorno, etc. Queste struct
    // sono sempre associate ad una `Location`, cioè
    // ad un fuso orario (in questo caso UTC).
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // È possibile estrarre i vari componenti della struct
    // `time` in un modo molto semplice.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // È possibile anche ottenere il giorno della settimana
    // tramite `Weekday`
    p(then.Weekday())

    // Questi metodi confrontano due `time`, controllando
    // quale dei due si riferisce ad un insante precedente
    // all'altro. Oppure il metodo `Equal` permette di confrontare
    // se due `time` sono uguali.
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // Il metodo `Sub` restituisce una `Duration` che rappresenta
    // l'intervallo di tempo fra due `time`.
    diff := now.Sub(then)
    p(diff)

    // È possibile calcolare la lunghezza di questo intervallo
    // in svariate unità di misura.
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // Si può dunque usare `Add` per spostare avanti nel
    // tempo un `time` di una `duration` desiderata. Anteponendo
    // un `-` è possibile spostare indietro un `time` della
    // stessa durata.
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

# Adesso vedremo un concetto molto simile: il tempo legato
# allo Unix epoch.
```
