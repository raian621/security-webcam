# muli-stage build (2 stages)
# - build_site: compile and bundle frontend code and assets
# - build_server: copy build_site artifacts and compile server
#   binary

# compile and bundle frontend code and assets:
FROM node:alpine as build_site
WORKDIR /src
ADD ./client .
RUN npm i
ENV VITE_HLS_SERVER_URL="{{.HlsServerUrl}}"
RUN npm run build

# copy frontend package from build_site container and 
# build HLS server binary:
FROM golang:1.21-alpine
WORKDIR /src
RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg 
ADD ./hls-server .
COPY --from=build_site /src/dist /var/www/site
RUN go build -o /bin/hls-server && rm -rf ./*
ADD ./scripts/start-docker-server .
CMD ["./start-docker-server"]