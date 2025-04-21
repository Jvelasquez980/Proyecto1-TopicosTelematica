# Proyecto1-TopicosTelematica: Jeronimo Velasquez Escobar, Miguel Angel Correa, Julian Osorio Alturo

# 🧾 Sistema Distribuido de Gestión de Órdenes

Este proyecto implementa una arquitectura de microservicios basada en Go y gRPC para gestionar órdenes, pagos, inventario y envíos. Los servicios se comunican entre sí mediante gRPC, y se expone un API REST vía FastAPI para permitir interacción desde clientes externos.

---

## 📐 Arquitectura General

- Microservicios:
  - orders-service
  - inventory-service
  - payment-service
  - shipping-service
- Comunicación interna: gRPC
- Entrada externa: API Gateway con FastAPI (REST)
- Logging local por servicio
- Simulación de fallos y cola de respaldo (MOM)
- Pruebas con scripts Python
- Desplegable en AWS

---

## 🧩 Microservicios

| Servicio          | Lenguaje | Funcionalidad                                   |
|------------------|----------|--------------------------------------------------|
| orders-service   | Go       | Recibe y coordina órdenes, comunica con otros.  |
| inventory-service| Go       | Valida y actualiza stock.                        |
| payment-service  | Go       | Procesa pagos, aprueba o rechaza.               |
| shipping-service | Go       | Genera info de envío y despacho.                |

Cada uno tiene su propio main.go, proto/, logs/, y conexión gRPC.

---

## 🚀 Instrucciones para correr localmente

Requisitos:

- Go ≥ 1.20
- Python ≥ 3.10
- protoc con plugin de Go
- uvicorn (para API Gateway)

Pasos:

1. Clonar el repositorio:

git clone https://github.com/tu-usuario/proyecto-ordenes.git
cd proyecto-ordenes

## Generar los archivos .pb.go para cada .proto:
protoc --go_out=. --go-grpc_out=. proto/orders/orders.proto
protoc --go_out=. --go-grpc_out=. proto/inventory/inventory.proto
# Y así sucesivamente para payment y shipping

## Ejecutar cada microservicio, en cuatro terminales diferentes:
go run main.go  # en orders-service
go run main.go  # en inventory-service
go run main.go  # en payment-service
go run main.go  # en shipping-service

## Ejecutar API Gateway:
cd api-gateway
uvicorn main:app --reload --port 8000

## Probar el sistema:
python client/client_batch.py

---

# 📝 Ejemplo de payload (POST a /orders)

{
  "user_id": "u123",
  "items": [
    {"product_id": "a1", "quantity": 2},
    {"product_id": "b2", "quantity": 1}
  ],
  "payment_method": "credit_card"
}

---


