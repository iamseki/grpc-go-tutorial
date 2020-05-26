## :wrench: Setup
  - If you want to generate `chat.pb.go`
  > `sudo apt update && sudo apt install snapd`

  > `sudo snap install protobuf --classic`

  > ` protoc --go_out=plugins=grpc:chat chat.proto`
  
---

### :running: Running the simple
 - to simulate communication between server and client, run them on two different terminals:

`go run server.go`

`go run client.go`

expected output in **client** terminal :

    
     Response from Server: Hello from the Server!
    
expected output in **server** terminal :

    
     Response from Server: Hello from the Server!
    
