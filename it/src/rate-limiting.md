# Rate Limiting
```go
// Il [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// è un'importante tecnica per controllare l'uso delle
// risorse e mantenere un'alta qualità di servizio. Go
// supporta elegantemente il rate limiting tramite le
// goroutine, i channel e i [ticker](ticker).

package main

import "time"
import "fmt"

func main() {

    // Come prima cosa vedremo un semplice esempio di rate
    // limiting. Supponiamo di dover limitare il
    // numero di richieste (per esempio in un server
    // HTTP). Utilizzeremo un channel per
    // mimare l'entrata delle richieste.
    richieste := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        richieste <- i
    }
    close(richieste)

    // Questo channel `limitatore` riceverà un valore ogni
    // 200 millisecondi. Fungerà da regolatore nel
    // nostro schema di rate limiting.
    limitatore := time.Tick(time.Millisecond * 200)

    // Bloccando l'esecuzione su una ricezione da
    // `limitatore` prima di servire ogni richiesta, ci
    // limitiamo a fare al massimo una richiesta ogni
    // 200 millisecondi.
    for ric := range richieste {
        <-limitatore
        fmt.Println("richiesta", ric, time.Now())
    }

    // Potremmo aver bisogno di sostenere brevi
    // "raffiche" di richieste, mantenendo comunque
    // il rate limiting. Possiamo fare ciò mettendo
    // bufferizzando il nostro channel. Questo channel
    // `limitatoreARaffica` ci permetterà di sostenere
    // raffiche fino ad un massimo di 3 richieste.
    limitatoreARaffica := make(chan time.Time, 3)

    // Riempiamo il channel per dire che già dall'inizio
    // possiamo prendere in carico 3 richieste.
    for i := 0; i < 3; i++ {
        limitatoreARaffica <- time.Now()
    }

    // Ogni 200 millisecondi proveremo ad inviare un
    // nuovo valore al `limitatoreARaffica`, fino ad
    // arrivare al suo limite di 3 richieste.
    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            limitatoreARaffica <- t
        }
    }()

    // Ora simuleremo 5 richieste al nostro limitatore.
    // Le prime tre potranno godere della capacità del
    // limitatore di eseguire richieste a raffica.
    richiesteARaffica := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        richiesteARaffica <- i
    }
    close(richiesteARaffica)
    for ric := range richiesteARaffica {
        <-limitatoreARaffica
        fmt.Println("richiesta", ric, time.Now())
    }
}
```
```sh
# Eseguendo il nostro programma vedremo il gruppo
# di richieste che verranno gestite ognuna ogni 200
# millisecondi.
$ go run rate-limiting.go
richiesta 1 2016-04-17 19:05:37.736132953 +0200 CEST
richiesta 2 2016-04-17 19:05:37.936138961 +0200 CEST
richiesta 3 2016-04-17 19:05:38.136209943 +0200 CEST
richiesta 4 2016-04-17 19:05:38.336145582 +0200 CEST
richiesta 5 2016-04-17 19:05:38.536120745 +0200 CEST

# Invece, per il nostro secondo gruppo di richieste
# vedremo che le prime tre vengono eseguite
# all'istante, mentre le altre 2 a distanza di 200
# millisecondi l'una dall'altra.
richiesta 1 2016-04-17 19:05:38.536251527 +0200 CEST
richiesta 2 2016-04-17 19:05:38.536266185 +0200 CEST
richiesta 3 2016-04-17 19:05:38.536277665 +0200 CEST
richiesta 4 2016-04-17 19:05:38.736385724 +0200 CEST
richiesta 5 2016-04-17 19:05:38.936385957 +0200 CEST
```
