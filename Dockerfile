# Menggunakan image Golang sebagai base image
FROM golang:1.21-alpine AS builder

# Set working directory dalam container
WORKDIR /go/src/app

# Copy semua file dari direktori lokal ke dalam container
COPY . .

# Unduh dependencies Go
RUN go mod download

# Build aplikasi Golang
#RUN go build -o app
RUN go build -C ./application
RUN go build -C ./application/queue/server
RUN go build -C ./application/queue/monitoring

# Buat stage untuk image runtime
FROM golang:1.21-alpine

# Set working directory untuk runtime
WORKDIR /go/src/app

# Copy file yang sudah dibuild dari stage sebelumnya
COPY --from=builder /go/src/app .

RUN chmod +x /go/src/app

#CMD ["go run ./application/main.go"]
# Jalankan aplikasi