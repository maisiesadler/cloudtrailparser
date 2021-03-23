FROM golang:1.15-alpine as build
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go get -d ./...
RUN go install ./...

COPY . .
RUN go build 

FROM alpine:3.11.3
WORKDIR /usr/local/bin
COPY --from=build /build/cloudtrailparser .
RUN ls

ENTRYPOINT ["cloudtrailparser"]
