# SHA256 Hashes
```go
// [_SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2) são
// frequentemente utilizadas para computar identidades em binário
// ou pequenos textos. Por exemplo, certificados TLS/SSL usam SHA256
// para computar a assinatura do certificado. Aqui estã como computar
// hashes SHA256 em Go.

package main

// Go implementa várias funções hash em seus respectivos
// pacotes `crypto/*`.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Aqui é iniciado uma nova hash.
	h := sha256.New()

	// `Write` espera bytes. É possível utilizar `[]byte(s)`
	// em uma string `s` para converte-la para bytes.
	h.Write([]byte(s))

	// Aqui o resultado da hash é retornado como slice de bytes.
	// O argumento passado para `Sum` pode ser usado para acrescentar
	// algo para um hash em slice de byte já existente.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
```
```sh
# Ao executar o código, é exibida a hash em formato hex
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# Também é possível computar outras hashes usando padrão 
# similar visto anteriormente. Por exemplo, para computar
# hashes SHA512 é preciso importar `crypto/sha512` e
# utilizar `sha512.New()`.

# Note que se você precisa de hashes criptograficamente 
# seguras é bom pesquisar um pouco sobre
# [hash strength](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
```
