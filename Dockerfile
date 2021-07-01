FROM golang:1.16-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 go build -o main

FROM alpine
COPY --from=build /src/main .
EXPOSE 8080
CMD [ "./main" ]