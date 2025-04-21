import grpc
import orders_pb2
import orders_pb2_grpc

def send_order(user_id, items, payment_method):
    channel = grpc.insecure_channel('localhost:50051')  # IP del orders-service
    stub = orders_pb2_grpc.OrderServiceStub(channel)

    item_msgs = [orders_pb2.Item(product_id=i["product_id"], quantity=i["quantity"]) for i in items]
    request = orders_pb2.OrderRequest(
        user_id=user_id,
        items=item_msgs,
        payment_method=payment_method
    )

    response = stub.CreateOrder(request)
    return response
