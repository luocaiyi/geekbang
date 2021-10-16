FROM golang:1.17-alpine AS build
COPY . /go/src/project/
WORKDIR /go/src/project/
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o /bin/amd64

FROM scratch
LABEL images.authors="caiyiluo"
EXPOSE 80
COPY --from=build /bin/amd64 /bin/amd64
ENTRYPOINT ["/bin/amd64"]

