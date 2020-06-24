# Instalação e configuração do GoLang

1. Baixar o instalador do Golang de acordo com a plataforma utilizada em seu computador (Windows, linux,osx), site para download: https://golang.org/ 

 - Executar o instalador
 - I Accept
 - Next
 - Next
 - Finish

2. Abrir o cmd e testar para ver se o Go ja configurou as variaveis de ambiente

 - Digite: go version

 caso não tenha funcionado siga os passos abaixo

3. Ajuste a variável PATH

É necessário configurar a variável de ambiente **GO_HOME**, para você poder rodar o compilador de qualquer pasta.

Para isso, abra o painel de controle do windows, clique em sistema, agora em configurações avançadas do sistema, variáveis de ambiente.

Na parte de baixo, em Variáveis do sistema, clique em novo, em nome da variável coloque **GO_HOME** e em Valor da variável coloque a pasta onde o Go foi instalado no seu windows, barra bin. No meu caso: **C:\Go**

Após isso selecionar na parte de cima  a Variavel **Path** - Editar - NOVO = **%GO_HOME%\bin**

Finalizar e testar novamente 


[GoLang](https://golang.org/ )