FROM golang:1.22 AS builder

COPY ${PWD} /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/appbin *.go

FROM scratch
LABEL MAINTAINER Author Diniyar

COPY --from=builder /app /app

WORKDIR /app

EXPOSE 8000

CMD ["./appbin"]
