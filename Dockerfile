FROM golang:1.11.5 as builder
CMD ["mkdir", "volteo"]
WORKDIR volteo
ADD go.mod go.mod
ADD go.sum go.sum
ADD src src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main src/main.go

FROM alpine:3.7
#RUN apk update && apk add bash
WORKDIR /root
COPY --from=builder /go/volteo/main .
ENV PORT 4000
CMD ["./main"]