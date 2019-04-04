# # Multistage docker build

# FROM golang:alpine AS build-env
# RUN mkdir -p /go/src/github.com/angadsharma1016/socialcops
# WORKDIR /go/src/github.com/angadsharma1016/socialcops
# ADD . .
# RUN go build -o socialcops

# FROM alpine
# WORKDIR /app
# COPY --from=build-env /go/src/github.com/angadsharma1016/socialcops /app/
# EXPOSE 3000
# ENTRYPOINT ./socialcops

# # docker build -t angadsharma1016/socialcops .
# # docker run --rm angadsharma1016/socialcops

FROM golang

RUN mkdir -p /go/src/github.com/angadsharma1016/socialcops

ADD . /go/src/github.com/angadsharma1016/socialcops

WORKDIR /go/src/github.com/angadsharma1016/socialcops

RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

EXPOSE 3000

ENTRYPOINT watcher -run github.com/angadsharma1016/socialcops/ -watch github.com/angadsharma1016/socialcops