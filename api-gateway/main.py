from fastapi import FastAPI
from pydantic import BaseModel
from order_client import send_order

app = FastAPI()

class Item(BaseModel):
    product_id: str
    quantity: int

class OrderRequest(BaseModel):
    user_id: str
    items: list[Item]
    payment_method: str

@app.post("/orders")
def create_order(order: OrderRequest):
    response = send_order(
        user_id=order.user_id,
        items=[item.dict() for item in order.items],
        payment_method=order.payment_method
    )
    return {
        "order_id": response.order_id,
        "status": response.status
    }
