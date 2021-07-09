FROM golang:1.16-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 go build ./cmd/stockupdate.go

FROM alpine
COPY --from=build /src/stockupdate .
EXPOSE 8080
CMD [ "./stockupdate" ]