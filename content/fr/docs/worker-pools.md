---
weight: 3
---
# Worker Pools
```go
// Dans cet exemple, nous allons regarder comment implémenter un worker poole en utilisant des goroutines et des canaus.

package main

import "fmt"
import "time"

// Voici notre worker, dont nous allons lancer plusieurs instances concurrentes. Ces workers recevront du travail dans le canal `jobs` et enverront le résultat correspondant dans `results`. Nous ferons une pause d'une seconde par tâche pour simuler une opération complexe.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    // Afin d'utiliser notre pool de workers, nous devons leur envoyer du travail et collecter leurs résultats. On crée deux canaux pour celà.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // On démarre 3 workers, initiallement bloqués car ils n'ont pas de travaux pour le moment.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Ici on envoit 9 travaux, puis ferme ce canal pour indiquer que c'est tout le travail que nous avons.
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Finalement, nous collectons les résultats du travail.
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```
```sh
# Notre programme montre, au cours de son exécution, les
# 9 jobs qui s'exécutent, traités par les différents
# workers. Le programme prend seulement environ 3
# secondes à s'exécuter, bien qu'il traite environ 9
# secondes de travail car il y a 3 workers qui
# travaillent de manière concurrente.
$ time go run worker-pools.go 
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 2 processing job 5
worker 3 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 3 processing job 9

real	0m3.149s
```
