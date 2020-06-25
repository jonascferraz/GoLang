# Primeiro Código 

No Primeiro código o programa imprimirá na tela do **CMD** a mensagem 'Olá Go!"
De o nome que quiser ao arquivo, vou chamar de primeiro.go

Abaixo código fonte completo:

**primeiro.go**

```
    package main
    import "fmt"
    func main() {
        fmt.Println("Olá Go!")
    }
```
-------
Para executar o programa use: **go run primeiro.go**

 Isso dentro da Pasta a qual você criou o arquivo primeiro.go no meu caso: C:\My Projects\Go

Para criar o programa em binário use: **go build primeiro.go**

 - Isso cria o binário **primeiro.exe**

 Podemos então executar o binário construído diretamente

 - primeiro.exe

 --------------------------


 **Declarando uma Variável em go**

 Código EX:

    ```
    package main

    import "fmt"

    func main() {
        var name string

        name = "Jonas"

        fmt.Print(name)
    }

    ```
o Compliador precisa saber o tipo da variável

 - String no name (Devido a ser caracter)