FROM nginx:1.27.5

RUN rm /etc/nginx/conf.d/default.conf
COPY status.conf /etc/nginx/conf.d/status.conf
COPY htpasswd /etc/nginx/.htpasswd
