package main

import (
    "context"
    "fmt"
    "log"
    "time"

    pb "orders-service/proto/orders"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("❌ No se pudo conectar a orders-service: %v", err)
    }
    defer conn.Close()

    client := pb.NewOrderServiceClient(conn)

    req := &pb.OrderRequest{
        UserId: "user_test_1",
        Items: []*pb.Item{
            {ProductId: "a1", Quantity: 2},
            {ProductId: "c3", Quantity: 1},
        },
        PaymentMethod: "paypal",
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    resp, err := client.CreateOrder(ctx, req)
    if err != nil {
        log.Fatalf("❌ Error al crear la orden: %v", err)
    }

    fmt.Println("✅ Orden enviada exitosamente:")
    fmt.Printf("🆔 OrderID: %s
", resp.OrderId)
    fmt.Printf("📦 Estado: %s
", resp.Status)
}
