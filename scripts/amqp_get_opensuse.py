#!/usr/bin/env python
import pika
import sys

connection = pika.BlockingConnection(pika.URLParameters("amqps://opensuse:opensuse@rabbit.opensuse.org"))
channel = connection.channel()

channel.exchange_declare(exchange='pubsub', exchange_type='topic', passive=True, durable=True)

result = channel.queue_declare("", exclusive=True)
queue_name = result.method.queue
print("Queue:", queue_name)
channel.queue_bind(exchange='pubsub', queue=queue_name, routing_key='#')

print(' [*] Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    print(" [x] %r:%r" % (method.routing_key, body))

channel.basic_consume(queue_name, callback, auto_ack=True)
channel.start_consuming()
