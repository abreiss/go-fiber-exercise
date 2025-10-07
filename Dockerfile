FROM golang:1.25 AS builder
#from baseimage; this is the starting point
#eventually maybe switch to golang:1.25-alpine
#temp workspace inc ontainer
WORKDIR /src
#build caching optimization
#copy dependencies files to /src
COPY go.mod go.sum  ./
#read and install dependencies
RUN go mod download
#copy everything else
COPY . .
#build go program at src/userapi, from src dir
RUN go build -o /src/userapi .
#from google container repo, build with lightweight image
FROM gcr.io/distroless/base-debian12
#workdir for build
WORKDIR /app
#use port 80
ENV PORT=80
EXPOSE 80
#copy from builder stage into app
COPY --from=builder /src/userapi /app/userapi
#run app when called on
CMD ["/app/userapi"]