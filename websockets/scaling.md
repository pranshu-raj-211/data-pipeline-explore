# Scaling websockets in Go

Since websockets are stateful protocols, having some overhead, there are some challenges in scaling them, especially in a horizontal manner.

## Issues with scaling websockets (vertically)

1. Memory overhead to maintain state info on server (connection object, buffers)
2. Overhead associated with TCP file descriptors on Linux servers (each connection has a file), number of this is limited (refer to GopherCon 19 talk by Eran Yanay, Going Infinite)
3. High bandwidth usage with scale
4. Size and processing of headers and data


## Issues with scaling horizontally

1. Maintaining consistent state information
2. Routing connections to the correct websocket server by load balancer
3. 

A general idea used to scale stateful systems is to add an inter-communication layer between these components to ensure resource consistency. Adding one more communication layer on top often creates challenges in terms of handling high throughput of messages, ensuring consistent execution of concurrent events, and fault tolerance in case of service unavailability.

