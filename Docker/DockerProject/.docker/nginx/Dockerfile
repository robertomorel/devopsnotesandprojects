FROM nginx:1.15.0-alpine

# Deleta o arquivo de conf. padrão do nginx
RUN rm /etc/nginx/conf.d/default.conf

# Copia o conteúdo do arquivo ./nginx.conf para /etc/nginx/conf.d
COPY ./nginx.conf /etc/nginx/conf.d