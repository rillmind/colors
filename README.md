# colors

Paleta [Catppuccin Mocha](https://catppuccin.com/palette) pronta para usar com o
[lipgloss](https://github.com/charmbracelet/lipgloss) em programas de terminal em Go.

## O que faz

O pacote `colors` oferece três coisas para cada cor de destaque da paleta (Rosewater, Flamingo,
Pink, Mauve, Red, Maroon, Peach, Yellow, Green, Teal, Sky, Sapphire, Blue e Lavander):

- uma constante `lipgloss.Color`, por exemplo `colors.MauveColor` (`#cba6f7`);
- um estilo com essa cor de texto, por exemplo `colors.MauveStyle`;
- uma função que imprime a string colorida com quebra de linha, por exemplo `colors.Mauve("texto")`.

Também tem as cores de fundo e de texto `BaseColor`, `TextColor` e `CrustColor` (só as constantes).

## Por que existe

Para não copiar os códigos hexadecimais do Catppuccin em cada projeto. É uma das minhas
bibliotecas pessoais e é usada pelo [new-distro](https://github.com/rillmind/new-distro).

## Instalação

Pré-requisito: Go 1.25.4+ num módulo Go.

```sh
go get github.com/rillmind/colors@latest
```

Exemplo de uso:

```go
package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/rillmind/colors"
)

func main() {
	colors.Mauve("Olá!")                              // imprime em Mauve
	fmt.Println(colors.GreenStyle.Render("tudo certo")) // usa o estilo pronto

	titulo := lipgloss.NewStyle().Bold(true).Foreground(colors.PeachColor)
	fmt.Println(titulo.Render("Catppuccin Mocha"))     // monta o próprio estilo
}
```

## Como testar

Não há testes automatizados. `go vet ./...` no repositório confere o código. Para conferir as
cores, rode o exemplo acima num terminal com suporte a cores (`go run .`): "Olá!" deve aparecer em
lilás, "tudo certo" em verde e "Catppuccin Mocha" em pêssego e negrito. Com a saída redirecionada
para um arquivo ou pipe, o lipgloss não emite cores, só o texto.
