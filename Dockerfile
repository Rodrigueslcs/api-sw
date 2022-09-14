FROM golang:latest as api
COPY . /api-sw
WORKDIR /api-sw
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api-sw main.go


FROM alpine
EXPOSE 8081
COPY --from=api /api-sw /api-sw
WORKDIR /api-sw
ENTRYPOINT [ "./api-sw", "api" ] 