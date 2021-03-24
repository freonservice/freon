FROM golang:1.16-alpine AS backend

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git

ENV CGO_ENABLED=0

WORKDIR /src

COPY ./api ./api

COPY ./cmd ./cmd

COPY ./internal ./internal

COPY ./migrations ./migrations

COPY ./pkg ./pkg

COPY ./vendor ./vendor

COPY ./go.mod ./go.mod

COPY ./go.sum ./go.sum

RUN go build -o /src/freon ./cmd/freon

FROM scratch AS final

COPY --from=backend /user/group /user/passwd /etc/

COPY --from=backend /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=backend /src/freon /freon

COPY --from=backend /src/migrations /migrations

USER nobody:nobody

ENTRYPOINT ["/freon"]
