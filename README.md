# Go

Go é a linguagem desenvolvida pelo Google para resolver problemas da escala que a gigante da tecnologia tem e, sem dúvida, essa linguagem também vai ajudar a turbinar as suas aplicações.

Ele é uma das poucas linguagens que surgiu depois dos processadores com múltiplos núcleos e isso fez com que a linguagem tenha mecanismo de concorrência no cerne da linguagem, o que a torna ideal para aplicações na qual são exigidos um alto grau de desempenho.

Go é uma linguagem enxuta, moderna, compilada (muito rápida), que tem vários recursos que irão te ajudar nos desafios de desenvolver uma aplicação escalável!

Go é uma linguagem concorrente

Paralelismo - Executar código simultaneamente em processadores físicos diferente
Concorrência - Intercalar (administrar) vários processos ao mesmo tempo e isso pode ocorrer em um único processador físico

Concorrência viabiliza paralelismo

É possível que a concorrência seja melhor que o paralelismo!
Paralelismo exirge muito mias do SO e do Hardware.

Concorrência - É a forma de estruturar seu programa
É a composição de processos (tipicamente funções) que executam de forma independente

## Comandos
Instalar MYSQL DRIVE: ```sh go get -u github.com/go-sql-driver/mysql```
Copilar aplicação para rodar no linux ```sh GOOS=linux GOARCH=amd64 go build```
Passando variavel de ambiente ```sh version=5 go run *.go```

## Comando de builds
- Arquitetura Windows de 32bits: ```sh GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o Programa.exe main.go```
- Arquitetura Windows de 64bits: ```sh GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o Programa.exe main.go```
- Arquitetura Linux de 64bits: ```sh GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o Programa64Bits main.go```

## Especificar saída do log com linux
```sh go run main.go &>> applications.log```
