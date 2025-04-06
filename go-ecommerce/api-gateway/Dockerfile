FROM ubuntu AS build

WORKDIR /app

RUN apt-get update && apt-get install -y golang-go

ENV GO111MODULE=off

COPY . .

RUN CGO_ENABLED=0 go build -o /app/app

FROM scratch

COPY --from=build /app/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]

