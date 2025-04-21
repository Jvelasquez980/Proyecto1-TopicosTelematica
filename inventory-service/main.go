package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "inventory-service/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedInventoryServiceServer
}

func (s *server) CheckInventory(ctx context.Context, req *pb.InventoryCheckRequest) (*pb.InventoryCheckResponse, error) {
    log.Printf("📦 Solicitud recibida para verificar inventario de %d productos", len(req.Items))

    // Aquí puedes implementar lógica real más adelante
    return &pb.InventoryCheckResponse{Available: true}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("❌ Error al escuchar: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterInventoryServiceServer(grpcServer, &server{})

    fmt.Println("🚀 InventoryService escuchando en el puerto 50052...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("❌ Error al servir: %v", err)
    }
}
