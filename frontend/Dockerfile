FROM node:lts-alpine as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
RUN yarn
RUN yarn global add @vue/cli@3.7.0
RUN yarn add node-sass
COPY ./ ./
RUN yarn build

FROM nginx:stable-alpine

RUN rm -rf /usr/share/nginx/html/*
COPY --from=build /app/dist/ /usr/share/nginx/html/
