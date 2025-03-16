# How do websockets in Go work?

Websockets are a bi directional persistent protocol that is built on top of HTTP. It basically hijacks the underlying TCP connection to send and receive messages, in a stateful manner. Since its built on top of TCP, it's only usable over HTTP 1.1 and 2.0, which basically works just about everywhere.

