# Command-Line Flag
```go
// I _Command-line flag_ sono uno dei modi prioritari
// per specificare delle opzioni in una qualsiasi
// [interfaccia a riga di comando](https://it.wikipedia.org/wiki/Interfaccia_a_riga_di_comando)
// Ad esempio in `wc -l`, il `-l` è un flag.

package main

// Go offre il pacchetto `flag` per supportare il parsing
// dei flag da linea di comando. Utilizzeremo questo pacchetto
// per implementare il nostro programma di esempio.
import "flag"
import "fmt"

func main() {

    // Il pacchetto offre semplici funzioni per parsare
    // stringhe, interi e booleani. Qui dichiariamo un
    // flag di tipo string chiamato `word` che ha come
    // valore di default `"foo"` ed una breve descrizione.
    // Questa funzione `flag.String` restituisce il puntatore
    // ad una stringa (non un valore di tipo stringa).
    // Vedremo dopo come utilizzare questo puntatore.
    wordPtr := flag.String("word", "foo", "a string")

    // Qui dichiariamo due flag, `numb` e `fork`,
    // utilizzando un approccio simile a `word`.
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // È inoltre possibile dichiarare un flag che vada ad
    // utilizzare una variabile precedentemente dichiarata
    // nel programma con `var`. Nota che è necessario passare
    // il puntatore a tale variabile.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Una volta che tutti i flag sono stati dichiarati è
    // possibile invocare `flag.Parse()` per eseguire il
    // parsing dei flag dalla linea di comando.
    flag.Parse()

    // Qui stampiamo i valori dei flag. Otterremo i valori
    // di default se i flag non sono presenti sulla linea
    // di comando. Nota che è necessario dereferenziare il
    // puntatore (vedi `*wordPtr`) per ottenere il valore
    // del flag.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
```
```sh
# Per fare qualche prova con la linea di comando è meglio
# prima compilare il programma ed eseguire direttamente
# il compilato.
$ go build command-line-flags.go

# Proviamo ad eseguire passando tutti i flag
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Nota che se non si settano i flag, vengono caricati
# i valori di default
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# I parametri posizionali possono essere passati in
# fondo dopo tutti i flag
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Nota che il pacchetto `flag` richiede che tutti i flag
# siano possizionati necessarimente prima dei parametri.
# Eventuali flag che si trovano in coda vengono trattati
# come parametri e non riconosciuti correttamente
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]

# Usa `-h` o `--help` per ottenere i suggerimenti 
# automatici sull'utilizzo dei flag (che vengono 
# generati a partire dalle definizioni che mettiamo
# nel programma).
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Se si utilizza un flag che non è stato dichiarato,
# il programma mostrerà un messaggio di errore e stamperà
# di nuovo il messaggio di informazione.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Vedremo adesso le variabili d'ambiente. Un altro modo
# comune per rendere parametrici i programmi.
```
