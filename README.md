# GoKeyHunt-CLI 1.0.0

## Descrição

Este projeto tem como objetivo implementar um algoritmo de força bruta que tentará todas as chaves privadas dentro de um determinado espaço até encontrar a chave privada correspondente à carteira desejada. O grande desafio é a vasta quantidade de chaves possíveis, o que torna esse processo computacionalmente extremamente custoso. Além disso, o projeto visa o aprendizado e a implementação de tecnologias como Golang e goroutines. Para isso, são utilizados sistemas avançados de controle de fluxo que evitam a verificação repetida da mesma chave e permitem alta performance.

### O que é o puzzle Bitcoin

O "Bitcoin Private Key Puzzle" é um desafio de 2015 onde é preciso encontrar chaves privadas de Bitcoin em intervalos específicos para ganhar recompensas crescentes de BTC. O total disponível é 32,896 BTC. O desafio destaca a dificuldade de encontrar chaves devido ao vasto espaço de busca. Algumas chaves foram resolvidas, mas muitas ainda estão por descobrir.

## Índice

- [Requisitos](#requisitos)
- [Instalação](#instalação)
- [Uso](#uso)
- [Funcionalidades](#funcionalidades)
- [Contribuição](#contribuição)
- [Licença](#licença)
- [Contato](#contato)

## Requisitos

- Go 1.22.3
- Git (caso deseje clonar o projeto diretamente)

## Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/LucasKalil-Programador/GoKeyHunt-CLI.git
    ```
2. Navegue até o diretório do projeto:
    ```sh
    cd GoKeyHunt-CLI
    ```
3. Instale as dependências:
    ```sh
    go mod tidy
    ```
4. Compile o projeto
    ```sh
    # Windows
    go build -o GoKeyHunt.exe ./cmd
    # Linux
    go build -o GoKeyHunt ./cmd
    ```

## Uso

O software é baseado em parâmetros de execução, e existem diversas formas de combiná-los, resultando em diferentes comportamentos.

1. Esse comando executa uma varredura em 1 milhão de chaves sequencialmente, começando de um endereço aleatório dentro da carteira 66.
    ```sh
    ./GoKeyHunt.exe -w 66 -bs 1_000_000 -rng
    ```

2. Esse comando executa a operação do item um mil vezes.
    ```sh
    ./GoKeyHunt.exe -w 66 -bs 1_000_000 -bc 1_000 -rng
    ```

3. Também é possível pré-configurar os parâmetros. Para isso, modifique o exemplo em [presets.json](./data/presets.json).
    ```sh
    ./GoKeyHunt.exe -preset wallet-66
    ```

## Funcionalidades

- **Alta flexibilidade**
  - Usando os parâmetros, é possível executar de milhares de formas diferentes.

- **Gerenciamento de colisões**
  - Ao executar em modo aleatório, é importante evitar a repetição, caso o mesmo índice seja gerado mais de uma vez. Por isso, foi implementado um sistema que evita colisões e armazena os índices já checados em um arquivo.

- **Console inteligente**
  - Os dados exibidos ao usuário são apresentados de forma a facilitar o entendimento da execução, incluindo estimativas de tempo, tempo decorrido, progresso geral e específico, entre outros.

## Contribuição

1. Faça um fork do projeto
2. Crie um branch para sua feature (`git checkout -b feature/nome-da-feature`)
3. Commite suas alterações (`git commit -m 'Adicionar feature'`)
4. Faça o push para o branch (`git push origin feature/nome-da-feature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](./LICENSE) para mais detalhes.

## Contato

Lucas - [lucas.prokalil2020@outlook.com](mailto:lucas.prokalil2020@outlook.com)

Repositório no GitHub: [GoKeyHunt-CLI](https://github.com/LucasKalil-Programador/GoKeyHunt-CLI)
