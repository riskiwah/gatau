FROM golang:1.15-alpine as builder
RUN apk --no-cache add build-base git mercurial gcc
RUN mkdir /gatau
ADD . /gatau
WORKDIR /gatau
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gatau .

FROM alpine:3.12.0
RUN apk --no-cache add git
COPY --from=builder /gatau .
COPY ./templates .
COPY .git/ .
ENTRYPOINT [ "./gatau" ]
EXPOSE 8080