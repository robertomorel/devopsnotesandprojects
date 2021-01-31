# Imagem
FROM php:7.3.6-fpm-alpine3.9

RUN apk add --no-cache shadow mysql-client nodejs npm

#RUN apk add --no-cache shadow
# Instalação dos pacotes bash e mysql-client
#RUN apk add bash mysql-client
# Instalando o nodejs, npm e yarn
#RUN apk add nodejs npm

# Instalação das extensões do php
RUN docker-php-ext-install pdo pdo_mysql

# Instalar o composer, gerenciador de pacotes do php
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

####################################
# Instalação do dockerrize para garantir que o DB (mysql) esteja pronto para rodar alguns comandos no container de aplicação
# https://github.com/jwilder/dockerize
RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
####################################

# Vai para o diretório /var/www do container
WORKDIR /var/www

# Remove a pasta html
RUN rm -rf /var/www/html

# Copia o que tem na raiz executada para a pasta /var/www
COPY . /var/www

RUN chown -R www-data:www-data /var/www

# Cria um link simbólico da pasta 'public' do Laravel para uma pasta virtual 'html'
RUN ln -s public html

RUN usermod -u 1000 www-data
USER www-data

# Responde na porta 9000
EXPOSE 9000

# Executa arquivo 'php-fpm'
ENTRYPOINT [ "php-fpm" ]