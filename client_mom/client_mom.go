package main

import (
    "context"
    "log"
    "os"
    "time"

    pb "client_mom/proto"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
    os.MkdirAll("logs", os.ModePerm)
    f, err := os.OpenFile("logs/client_mom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    log.SetOutput(f)
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)

    for {
        conn, err := grpc.Dial("mom-service:50055", grpc.WithInsecure())
        if err != nil {
            log.Printf("❌ No se pudo conectar a mom-service: %v", err)
            time.Sleep(10 * time.Second)
            continue
        }

        client := pb.NewMomServiceClient(conn)
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        resp, err := client.ProcessAllPending(ctx, &emptypb.Empty{})
        conn.Close()

        if err != nil {
            log.Printf("❌ Error en ProcessAllPending: %v", err)
        } else if resp.Reprocessed > 0 || resp.Failed > 0 {
            log.Printf("🔁 Reprocesadas: %d | Fallidas: %d", resp.Reprocessed, resp.Failed)
        } else {
            log.Println("📭 No hay órdenes pendientes.")
        }

        time.Sleep(10 * time.Second)
    }
}
