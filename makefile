# Nome do binário
BINARY_NAME=btcgo.exe

# Diretório de fontes
SRC_DIR=./cmd

# Comando para limpar os arquivos compilados
clean:
	rm -f $(BINARY_NAME)

# Comando para fazer pull do repositório
pull:
	git pull

# Comando para compilar o binário
build:
	go build -o $(BINARY_NAME) $(SRC_DIR)

# Comando para rodar o binário
run: build
	./$(BINARY_NAME) $(PARAMS)
