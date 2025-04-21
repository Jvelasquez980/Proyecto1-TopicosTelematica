package main

import (
    "context"
    "encoding/json"
    "log"
    "net"
    "os"

    pb "mom-service/proto"
    orderpb "orders-service/proto/orders"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
    pb.UnimplementedMomServiceServer
}

type failedOrder struct {
    UserID        string           `json:"user_id"`
    Items         []*orderpb.Item  `json:"items"`
    PaymentMethod string           `json:"payment_method"`
    Reason        string           `json:"reason"`
}

func readPendingOrders(path string) ([]failedOrder, error) {
    var orders []failedOrder
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    if len(data) == 0 {
        return []failedOrder{}, nil
    }
    err = json.Unmarshal(data, &orders)
    return orders, err
}

func writePendingOrders(path string, orders []failedOrder) error {
    data, err := json.MarshalIndent(orders, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(path, data, 0644)
}

func (s *server) ProcessAllPending(ctx context.Context, _ *emptypb.Empty) (*pb.MomSummary, error) {
    log.Println("🛠️ Procesando órdenes pendientes...")
    path := "../orders-service/logs/pending_orders.json"

    orders, err := readPendingOrders(path)
    if err != nil {
        log.Printf("❌ Error leyendo órdenes pendientes: %v", err)
        return &pb.MomSummary{}, nil
    }

    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Printf("❌ No se pudo conectar al orders-service: %v", err)
        return &pb.MomSummary{}, nil
    }
    defer conn.Close()
    client := orderpb.NewOrderServiceClient(conn)

    var remaining []failedOrder
    var reprocessed, failed int32

    for _, ord := range orders {
        req := &orderpb.OrderRequest{
            UserId:        ord.UserID,
            Items:         ord.Items,
            PaymentMethod: ord.PaymentMethod,
        }

        resp, err := client.CreateOrder(context.Background(), req)
        if err != nil || resp.Status != "created" {
            log.Printf("❌ Orden %s falló: %v", ord.UserID, err)
            remaining = append(remaining, ord)
            failed++
        } else {
            log.Printf("✅ Orden procesada: %s → %s", ord.UserID, resp.OrderId)
            reprocessed++
        }
    }

    _ = writePendingOrders(path, remaining)

    return &pb.MomSummary{
        Reprocessed: reprocessed,
        Failed:      failed,
    }, nil
}

func main() {
    os.MkdirAll("logs", os.ModePerm)
    f, _ := os.OpenFile("logs/mom-service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    log.SetOutput(f)
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)

    lis, err := net.Listen("tcp", ":50055")
    if err != nil {
        log.Fatalf("❌ No se pudo escuchar en el puerto 50055: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterMomServiceServer(grpcServer, &server{})

    log.Println("🚀 MomService escuchando en el puerto 50055...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("❌ Fallo al servir: %v", err)
    }
}
