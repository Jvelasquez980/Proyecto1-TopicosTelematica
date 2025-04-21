package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "orders-service/proto"

    "google.golang.org/grpc"
    "github.com/google/uuid"
)

type server struct {
    pb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    log.Printf("🧾 [Order] Recibida orden de %s con %d ítems.\n", req.UserId, len(req.Items))

    return &pb.OrderResponse{
        OrderId: uuid.New().String(),
        Status:  "created",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Fallo al escuchar: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterOrderServiceServer(grpcServer, &server{})

    fmt.Println("🚀 OrderService escuchando en el puerto 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Fallo al servir: %v", err)
    }
}
