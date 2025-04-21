package main

import (
    "context"
    "log"
    "math/rand"
    "net"
    "time"

    pb "payment-service/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedPaymentServiceServer
}

func (s *server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
    log.Printf("💳 Procesando pago de %.2f por el usuario %s con método %s", req.Amount, req.UserId, req.Method)

    rand.Seed(time.Now().UnixNano())
    aprobado := rand.Intn(100) >= 20 // 80% de probabilidad de aprobación

    if aprobado {
        log.Println("✅ Pago aprobado")
    } else {
        log.Println("❌ Pago rechazado")
    }

    return &pb.PaymentResponse{Approved: aprobado}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("❌ Error al escuchar: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterPaymentServiceServer(grpcServer, &server{})

    log.Println("🚀 PaymentService escuchando en el puerto 50053...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("❌ Error al servir: %v", err)
    }
}
