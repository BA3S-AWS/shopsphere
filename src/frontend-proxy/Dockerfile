FROM nginxinc/nginx-unprivileged:alpine
USER root
RUN apk upgrade --no-cache

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080

USER 101:101

CMD ["nginx", "-g", "daemon off;"]
