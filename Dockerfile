#--------------------------------------
FROM node:13.12.0 AS frontend

RUN yarn global add @vue/cli@4.3.1

COPY ./frontend /app
WORKDIR /app
RUN ls -l
COPY ./frontend/package.json /app

RUN yarn install
RUN yarn build

#--------------------------------------
FROM golang:1.14.2 AS backend

WORKDIR /go/src/app
COPY ./backend ./backend
COPY ./go.* ./

RUN ls -l
ARG KUBE_CONFIG
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.KUBECONFIG=$KUBE_CONFIG" backend/kubedash.go

#--------------------------------------
FROM alpine:3.11.5
RUN adduser --disabled-password --no-create-home kubedash
RUN apk add --no-cache bash
COPY --from=backend go/src/app/kubedash /
COPY --from=frontend app/dist /frontend/dist

COPY ./backend/tmp_deployments/kubeconfig.yaml ./kubeconfig.yaml

# COPY --from=lachlanevenson/k8s-kubectl:v1.10.3 /usr/local/bin/kubectl /usr/local/bin/kubectl
# RUN chmod a+rx /usr/local/bin/kubectl

RUN apk add --update --no-cache curl ca-certificates && \
    curl -L https://dl.k8s.io/v1.10.3/bin/linux/amd64/kubectl --output kubectl && \
    mv kubectl /usr/local/bin/kubectl && \
    chmod +x /usr/local/bin/kubectl

RUN apk add --update --no-cache curl ca-certificates && \
    curl -L https://github.com/FairwindsOps/pluto/releases/download/v3.4.0/pluto_3.4.0_linux_386.tar.gz | tar xvz && \
    mv pluto /usr/bin/pluto && \
    chmod +x /usr/bin/pluto && \
    rm -f /var/cache/apk/*

USER kubedash
# EXPOSE 8008
# EXPOSE 8080

ENTRYPOINT ["/kubedash"]