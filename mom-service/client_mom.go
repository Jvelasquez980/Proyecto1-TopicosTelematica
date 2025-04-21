package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    pb "mom-service/proto"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
    // Crear carpeta de logs si no existe
    os.MkdirAll("logs", os.ModePerm)

    // Abrir archivo de logs
    file, err := os.OpenFile("logs/client_mom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("❌ No se pudo abrir archivo de log: %v", err)
    }
    defer file.Close()

    // Configurar logging
    log.SetOutput(file)
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)

    log.Println("🚀 Iniciando cliente de monitoreo MOM...")

    for {
        conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
        if err != nil {
            log.Printf("❌ No se pudo conectar a mom-service: %v", err)
            time.Sleep(10 * time.Second)
            continue
        }

        client := pb.NewMomServiceClient(conn)
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

        resp, err := client.ProcessAllPending(ctx, &emptypb.Empty{})
        cancel()
        conn.Close()

        if err != nil {
            log.Printf("❌ Error llamando ProcessAllPending: %v", err)
        } else if resp.Reprocessed > 0 || resp.Failed > 0 {
            log.Printf("🔁 Reprocesadas: %d | Fallidas: %d", resp.Reprocessed, resp.Failed)
            fmt.Printf("🌀 Reprocesadas: %d | Fallidas: %d\n", resp.Reprocessed, resp.Failed)
        } else {
            log.Println("📭 No hay órdenes pendientes.")
            fmt.Println("📭 No hay órdenes pendientes.")
        }

        time.Sleep(10 * time.Second)
    }
}
