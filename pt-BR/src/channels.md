# Channels
```go
// _Channels_ ou _canais_ são conexões entre goroutines
// concorrentes. Valores podem ser enviados para canais
// de uma goroutine de forma que outra goroutine os receba.

package main

import "fmt"

func main() {

	// Criação de canais pode ser feito com `make(chan val-type)`.
	// Canais são tipados pelo valores que eles recebem.
	messages := make(chan string)

	// Para _Enviar_ um valor para um canal, use a sintaxe
	// `channel <-`. Aqui é enviada, de dentro de uma nova
	// goroutine uma string `"ping"` para o canal `messages`
	// feito acima.
	go func() { messages <- "ping" }()

	// A sintaxe `<-channel` _recebe_ um valor de um canal.
	// Aqui é recebido a string `"ping"` pelo canal message,
	// e após, impresso.
	msg := <-messages
	fmt.Println(msg)
}
```
```sh
# Ao rodar o programa a mensagem `"ping"` é passada com 
# successo de uma goroutine para outra via o canal que
# foi criado.
$ go run channels.go 
ping

# Por padrão, ambos os canais de envio e recebimento 
# ficam bloqueados até que ambos estejam prontos. 
# Esta propriedade permite aguardar ao final do programa 
# pela mensagem sem que seja preciso qualquer outro tipo
# de sincronização.
```
