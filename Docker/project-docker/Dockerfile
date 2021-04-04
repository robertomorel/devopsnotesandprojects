# Imagem
FROM golang:1.14

# Copia todos os arquivos que estão na máquina e cola no container
COPY . .

# Gera o exe do GO
RUN go build main.go

# Responde na porta 8080 caso algum outro container for comunicar
EXPOSE 8080

# Qndo o container rodar, executar esse arquivo main 
ENTRYPOINT [ "./main" ]

# PARA EXECUTAR O DOCKERFILE: docker build -t rmorel-nome-da-imagem .
# Criar uma imagem a partir dos comandos do Dockerfile
# -t rmorel-nome-da-imagem -> nome da imagem
# . -> na pasta atual