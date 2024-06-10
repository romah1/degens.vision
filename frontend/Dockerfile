FROM node:20-slim as build

# RUN apt-get update && apt-get install -y python3 build-essential

WORKDIR /app
COPY . /app

RUN yarn install --frozen-lockfile --ignore-engines && yarn run build

FROM nginx:1.16.0-alpine AS production
COPY --from=build /app/build /usr/share/nginx/html

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
