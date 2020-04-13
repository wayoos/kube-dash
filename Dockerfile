#--------------------------------------
FROM node:13.12.0 AS frontend

RUN yarn global add @vue/cli@4.3.1

COPY ./frontend /
WORKDIR /frontend
RUN yarn install
RUN yarn build

#--------------------------------------
FROM golang:1.14.2 AS backend

WORKDIR /go/src/app
COPY ./backend .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' kubedash.go

#--------------------------------------
FROM alpine:3.11.5
RUN adduser --disabled-password --no-create-home kubedash
COPY --from=backend go/src/app/kubedash /

USER kubedash
EXPOSE 8000

ENTRYPOINT ["/kubedash"]