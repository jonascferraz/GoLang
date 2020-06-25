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


 # Declarando uma Variável em go

 **Código EX:**

```
    package main

    import "fmt"

    func main() {
        var name string

        name = "Jonas"

        fmt.Print(name)
    }
```

O Compliador precisa saber o tipo da variável que sempre fica alocado ao lado direito da variável declarada

Podemos também declarar as variáveis como feito em javascript:

**Exemplo 1** 

    var name = Jonas (o propio sistema já entende que é uma string)

----------

**Exemplo 2**

    name := Jonas

 - A Nível de Package não podemos declarar desta forma se não da erro ao compilar

  Utilizar da forma do Exemplo 1 , ou declarar a variável fora da **func main**

-------

**Tipos de Variáveis**

 - String no name (Devido a ser caracter)

 - bool (built-in) (0 ou 1, true(1) or false(0))

 - int (Número inteiro) Ex: (1 , 2 , 5 ,22 , -15)

 - float32 (Números com ponto) Ex: (1.75 , 1.87)