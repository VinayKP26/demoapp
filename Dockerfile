FROM golang:alpine AS builder
WORKDIR /src/
COPY . . 
RUN GOOS=linux CGO_ENABLED=0 go build -o /src/target/app .

FROM alpine:3.16
WORKDIR /web/ 
COPY --from=builder /src/target/app /web/ 
CMD [ "/web/app" ]