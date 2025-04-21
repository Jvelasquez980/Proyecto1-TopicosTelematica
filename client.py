import requests
import json
import random
import time

url = "http://localhost:8000/orders"

productos = ["a1", "b2", "c3", "d4", "e5"]

def generar_orden(usuario_id):
    cantidad_items = random.randint(1, 3)
    items = []

    for _ in range(cantidad_items):
        producto = random.choice(productos)
        cantidad = random.randint(1, 5)
        items.append({"product_id": producto, "quantity": cantidad})

    return {
        "user_id": f"user_{usuario_id}",
        "items": items,
        "payment_method": random.choice(["credit_card", "debit_card", "paypal"])
    }

def enviar_orden(orden):
    headers = {"Content-Type": "application/json"}
    try:
        response = requests.post(url, headers=headers, data=json.dumps(orden), timeout=2)
        print(f"✅ Orden enviada: {orden['user_id']} → {response.status_code} → {response.json()}")
    except requests.exceptions.RequestException as e:
        print(f"❌ Error al enviar orden {orden['user_id']}: {e}")

def main():
    num_ordenes = 10
    for i in range(num_ordenes):
        orden = generar_orden(i)
        enviar_orden(orden)
        time.sleep(1)  # Espera de 1 segundo entre órdenes

if __name__ == "__main__":
    main()
