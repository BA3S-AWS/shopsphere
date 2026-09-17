FROM nginxinc/nginx-unprivileged:alpine

COPY --chown=101:101 index.html /usr/share/nginx/html/index.html
COPY --chown=101:101 app.js /usr/share/nginx/html/app.js
COPY --chown=101:101 nginx.conf /etc/nginx/conf.d/default.conf

USER 101:101

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]
