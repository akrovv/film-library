FROM golang:alpine AS builder

RUN mkdir /library
WORKDIR /library

COPY . .
COPY .env .

RUN go mod download && go build cmd/filmlibrary/main.go

FROM alpine

RUN mkdir docs

COPY --from=builder /library/docs ./docs
COPY --from=builder /library/rbac_model.conf .
COPY --from=builder /library/rbac_policy.csv .
COPY --from=builder /library/.env .
COPY --from=builder /library/main .

EXPOSE 8080

CMD [ "./main" ]