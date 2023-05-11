FROM golang
COPY . /App
WORKDIR /App
CMD go run main.go
EXPOSE 1234