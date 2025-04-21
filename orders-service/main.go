package main

import (
    "context"
    "encoding/json"
    "log"
    "net"
    "os"

    orderpb "orders-service/proto/orders"
    invpb "orders-service/proto/inventory"
    paymentpb "orders-service/proto/payment"
    shippingpb "orders-service/proto/shipping"

    "google.golang.org/grpc"
    "github.com/google/uuid"
)

type server struct {
    orderpb.UnimplementedOrderServiceServer
    inventoryClient invpb.InventoryServiceClient
    paymentClient   paymentpb.PaymentServiceClient
    shippingClient  shippingpb.ShippingServiceClient
}

func guardarOrdenFallida(req *orderpb.OrderRequest, motivo string) {
	type failedOrder struct {
		UserID        string              `json:"user_id"`
		Items         []*orderpb.Item     `json:"items"`
		PaymentMethod string              `json:"payment_method"`
		Reason        string              `json:"reason"`
	}

	// Cargar órdenes previas si existen
	var pendientes []failedOrder
	filePath := "logs/pending_orders.json"

	data, err := os.ReadFile(filePath)
	if err == nil && len(data) > 0 {
		if err := json.Unmarshal(data, &pendientes); err != nil {
			log.Printf("⚠️ No se pudo parsear el archivo de órdenes pendientes, lo reiniciamos: %v", err)
			pendientes = []failedOrder{}
		}
	}

	// Nueva orden fallida
	nueva := failedOrder{
		UserID:        req.UserId,
		Items:         req.Items,
		PaymentMethod: req.PaymentMethod,
		Reason:        motivo,
	}

	pendientes = append(pendientes, nueva)

	// Serializar y guardar como arreglo válido
	out, err := json.MarshalIndent(pendientes, "", "  ")
	if err != nil {
		log.Printf("❌ Error al serializar orden fallida: %v", err)
		return
	}

	if err := os.WriteFile(filePath, out, 0644); err != nil {
		log.Printf("❌ Error al guardar orden fallida: %v", err)
	} else {
		log.Printf("📝 Orden fallida guardada en pending_orders.json: usuario=%s, motivo=%s", req.UserId, motivo)
	}
}


func (s *server) CreateOrder(ctx context.Context, req *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
    log.Printf("🧾 [Order] Recibida orden de %s con %d ítems.", req.UserId, len(req.Items))

    requestJson, _ := json.MarshalIndent(req, "", "  ")
    log.Println("📦 Contenido completo de la orden:")
    log.Println(string(requestJson))

    log.Println("📦 Verificando inventario...")

    invReq := &invpb.InventoryCheckRequest{}
    for _, item := range req.Items {
        invReq.Items = append(invReq.Items, &invpb.Item{
            ProductId: item.ProductId,
            Quantity:  item.Quantity,
        })
    }

    invResp, err := s.inventoryClient.CheckInventory(ctx, invReq)
    if err != nil {
        log.Printf("❌ Error al llamar a inventory-service: %v", err)
        guardarOrdenFallida(req, "inventory_service_down")
        return &orderpb.OrderResponse{
            OrderId: "PENDIENTE",
            Status:  "en cola MOM (inventory)",
        }, nil
    }

    if !invResp.Available {
        log.Println("❌ Inventario insuficiente. Cancelando orden.")
        guardarOrdenFallida(req, "inventory_unavailable")
        return &orderpb.OrderResponse{
            OrderId: "RECHAZADA",
            Status:  "sin inventario",
        }, nil
    }

    log.Println("✅ Inventario confirmado.")
    log.Println("💳 Procesando pago...")

    amount := float32(len(req.Items)) * 10.0
    payReq := &paymentpb.PaymentRequest{
        UserId: req.UserId,
        Amount: amount,
        Method: req.PaymentMethod,
    }

    payResp, err := s.paymentClient.ProcessPayment(ctx, payReq)
    if err != nil {
        log.Printf("❌ Error al llamar a payment-service: %v", err)
        guardarOrdenFallida(req, "payment_service_down")
        return &orderpb.OrderResponse{
            OrderId: "PENDIENTE",
            Status:  "en cola MOM (payment)",
        }, nil
    }

    if !payResp.Approved {
        log.Println("❌ Pago rechazado. Cancelando orden.")
        guardarOrdenFallida(req, "payment_rejected")
        return &orderpb.OrderResponse{
            OrderId: "RECHAZADA",
            Status:  "pago rechazado",
        }, nil
    }

    log.Println("✅ Pago aprobado.")
    log.Println("🚚 Enviando orden al servicio de shipping...")

    shipReq := &shippingpb.ShippingRequest{
        UserId: req.UserId,
    }
    for _, item := range req.Items {
        shipReq.Items = append(shipReq.Items, &shippingpb.Item{
            ProductId: item.ProductId,
            Quantity:  item.Quantity,
        })
    }

    shipResp, err := s.shippingClient.CreateShipment(ctx, shipReq)
    if err != nil {
        log.Printf("❌ Error al llamar a shipping-service: %v", err)
        guardarOrdenFallida(req, "shipping_service_down")
        return &orderpb.OrderResponse{
            OrderId: "PENDIENTE",
            Status:  "en cola MOM (shipping)",
        }, nil
    }

    log.Printf("✅ Envío creado con ID: %s", shipResp.ShippingId)
    return &orderpb.OrderResponse{
        OrderId: uuid.New().String(),
        Status:  "created",
    }, nil
}

func main() {
    os.MkdirAll("logs", os.ModePerm)
    file, err := os.OpenFile("logs/orders.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("❌ No se pudo abrir archivo de log: %v", err)
    }
    defer file.Close()
    log.SetOutput(file)
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)

    invConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("❌ No se pudo conectar a inventory-service: %v", err)
    }
    defer invConn.Close()
    inventoryClient := invpb.NewInventoryServiceClient(invConn)

    payConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("❌ No se pudo conectar a payment-service: %v", err)
    }
    defer payConn.Close()
    paymentClient := paymentpb.NewPaymentServiceClient(payConn)

    shipConn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("❌ No se pudo conectar a shipping-service: %v", err)
    }
    defer shipConn.Close()
    shippingClient := shippingpb.NewShippingServiceClient(shipConn)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("❌ Fallo al escuchar: %v", err)
    }

    grpcServer := grpc.NewServer()
    orderpb.RegisterOrderServiceServer(grpcServer, &server{
        inventoryClient: inventoryClient,
        paymentClient:   paymentClient,
        shippingClient:  shippingClient,
    })

    log.Println("🚀 OrderService escuchando en el puerto 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("❌ Fallo al servir: %v", err)
    }
}
