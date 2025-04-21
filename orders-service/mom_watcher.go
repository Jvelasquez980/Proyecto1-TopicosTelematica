package main

import (
    "context"
    "encoding/json"
    "log"
    "os"
    "time"

    pb "orders-service/proto/orders"
    "google.golang.org/grpc"
)

type failedOrder struct {
    UserID        string       `json:"user_id"`
    Items         []*pb.Item   `json:"items"`
    PaymentMethod string       `json:"payment_method"`
    Reason        string       `json:"reason"`
}

func readFailedOrders(path string) ([]failedOrder, error) {
	var orders []failedOrder

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []failedOrder{}, nil // No hay archivo aún
		}
		return nil, err
	}

	if len(data) == 0 {
		return []failedOrder{}, nil
	}

	err = json.Unmarshal(data, &orders)
	return orders, err
}

func overwritePendingOrders(path string, orders []failedOrder) error {
	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}


func main() {
    os.MkdirAll("logs", os.ModePerm)
    logFile, err := os.OpenFile("logs/mom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("❌ No se pudo abrir mom.log: %v", err)
    }
    defer logFile.Close()
    log.SetOutput(logFile)
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)

    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("❌ No se pudo conectar al orders-service: %v", err)
    }
    defer conn.Close()

    client := pb.NewOrderServiceClient(conn)

    log.Println("👀 MOM Watcher iniciado. Monitoreando órdenes pendientes...")

    for {
        time.Sleep(10 * time.Second)

        orders, err := readFailedOrders("logs/pending_orders.json")
        if err != nil {
            log.Printf("❌ Error leyendo órdenes pendientes: %v", err)
            continue
        }

        if len(orders) == 0 {
            log.Println("📭 No hay órdenes pendientes.")
            continue
        }

        var remaining []failedOrder
        for _, ord := range orders {
            log.Printf("🔁 Reintentando orden de %s (%d ítems)", ord.UserID, len(ord.Items))

            req := &pb.OrderRequest{
                UserId:        ord.UserID,
                Items:         ord.Items,
                PaymentMethod: ord.PaymentMethod,
            }

            ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
            defer cancel()

            resp, err := client.CreateOrder(ctx, req)
            if err != nil {
                log.Printf("❌ Error al reenviar orden: %v", err)
                remaining = append(remaining, ord)
                continue
            }

            if resp.Status == "created" {
                log.Printf("✅ Orden procesada correctamente: %s", resp.OrderId)
            } else {
                log.Printf("⚠️ Orden no pudo completarse: estado = %s", resp.Status)
                remaining = append(remaining, ord)
            }
        }

        if err := overwritePendingOrders("logs/pending_orders.json", remaining); err != nil {
            log.Printf("❌ Error al actualizar archivo pending_orders.json: %v", err)
        }
    }
}
