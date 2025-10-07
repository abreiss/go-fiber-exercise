FROM golang:1.25 AS builder
#from baseimage; this is the starting point
#for go its gonna be golang:1.25 i think
WORKDIR /src
#caching 
COPY go.mod go.sum  ./
RUN go mod download

#copy everything else
# why is this seperate?
COPY . .

RUN go build -o /src/userapi .

FROM gcr.io/distroless/base-debian12

WORKDIR /app
ENV PORT=80
EXPOSE 80
COPY --from=builder /src/userapi /app/userapi
CMD ["/app/userapi"]