FROM node:lts-alpine AS builder

WORKDIR /app

COPY ["package*.json", "yarn.lock", "vite.config.js" , "./"]

RUN yarn install

COPY ./src ./src
COPY ./index.html ./index.html

RUN yarn build

FROM nginx:latest

COPY ./nginx /etc/nginx
COPY ./ dest

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx","-g","daemon off;"]