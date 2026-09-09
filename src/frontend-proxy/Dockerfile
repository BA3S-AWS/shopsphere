FROM nginx:alpine
RUN apk upgrade --no-cache

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]
