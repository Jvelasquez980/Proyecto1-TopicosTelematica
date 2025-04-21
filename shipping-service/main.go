package main

import (
    "context"
    "log"
    "net"

    pb "shipping-service/proto"
    "google.golang.org/grpc"
    "github.com/google/uuid"
)

type server struct {
    pb.UnimplementedShippingServiceServer
}

func (s *server) CreateShipment(ctx context.Context, req *pb.ShippingRequest) (*pb.ShippingResponse, error) {
    log.Printf("📦 Preparando envío para el usuario %s con %d ítems", req.UserId, len(req.Items))
    return &pb.ShippingResponse{
        ShippingId: uuid.New().String(),
        Status:     "enviado",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50054")
    if err != nil {
        log.Fatalf("❌ Error al escuchar: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterShippingServiceServer(grpcServer, &server{})

    log.Println("🚀 ShippingService escuchando en el puerto 50054...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("❌ Error al servir: %v", err)
    }
}
