FROM golang:1.15-alpine as builder
RUN apk --no-cache add build-base git mercurial gcc
ADD . /gatau
WORKDIR /gatau
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin .

FROM scratch
COPY --from=builder /gatau/bin ./gatau
ENTRYPOINT [ "./gatau" ]
EXPOSE 8080