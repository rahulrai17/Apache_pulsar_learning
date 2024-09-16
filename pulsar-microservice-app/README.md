## Application workflow and Informations

Here’s is the explaination about this Application , what all microservices it has how can you run this to see the work flow and some more details.

### Prerequisites

- Go installed on your system.
- Apache Pulsar either running locally (manually installed or via Docker) or using a managed cloud Pulsar instance.
- Basic understanding of Go and microservice architecture.

### Steps

#### 1. **Create the Microservice Structure**

1. Each microservice will be responsible for one part of the sentence:

- **Service 1**: "Hello"
- **Service 2**: "World"
- **Service 3**: "Whatsup"

2. **Use Pulsar for Communication**: Each service will publish its word to a Pulsar topic and then forward the result to the next service.

3. **Final Service**: Collects all words and prints them on the web page.

4. **Web Service**: Handles HTTP requests and prints the final message.

Let’s create a directory structure:

```
pulsar-microservice-app/
  ├── hello-service/
        -main.go
  ├── world-service/
        -main.go
  ├── whatsup-service/
        -main.go
  ├── web-service/
        -main.go
  -go.mod
  -go.sum

```

#### 2. **Install Pulsar Client for Go**

You’ll need to install the Apache Pulsar Go client to allow your services to publish and consume messages via Pulsar.

Run the following command to install Pulsar Go client:

```bash
go get -u github.com/apache/pulsar-client-go/pulsar
```

### Installation of Pulsar Client

1. **Installation Location**:

   - You should install the Pulsar Go client in the root of your project (`pulsar-microservice-app`), not inside each individual service directory. This is because Go modules handle dependencies at the root level of your project. When you run `go get`, it updates your `go.mod` and `go.sum` files in the root directory of your project.

2. **Setting Up Go Modules**:

   - Ensure you are using Go modules for dependency management. In the root directory (`pulsar-microservice-app`), initialize the Go module if you haven't already:

     ```bash
     go mod init pulsar-microservice-app
     ```

   - Then, install the Pulsar Go client:

     ```bash
     go get -u github.com/apache/pulsar-client-go/pulsar
     ```

   - This command will update the `go.mod` file to include the Pulsar client as a dependency.

3. **Accessing the Pulsar Client**:

   - In each service directory (`hello-service`, `world-service`, `whatsup-service`), you can import and use the Pulsar client as needed. Make sure each service has its own `main.go` and Go module setup if needed. For example, within `hello-service`, you might have:

     ```go
     package main

     import (
         "github.com/apache/pulsar-client-go/pulsar"
         // other imports
     )
     ```

#### 3. **Set up Apache Pulsar (Standalone Mode)**

Start Pulsar on your system:

**Option 1: Using Docker**

```bash
docker run -it -p 6650:6650 -p 8080:8080 --name pulsar-standalone apachepulsar/pulsar:latest bin/pulsar standalone
```

**Option 2: Local Pulsar Setup**

If you installed Apache Pulsar locally without Docker, run it as follows:

```bash
cd apache-pulsar-2.9.1
bin/pulsar standalone
```

#### 4. **Test the Complete Flow**

1. Start all three services in different terminals:

   ```bash
   cd hello-service && go run main.go
   cd world-service && go run main.go
   cd whatsup-service && go run main.go
   cd web-service && go run main.go
   ```

2. Once the flow completes, visit [http://localhost:8080](http://localhost:8080) in your browser, and you should see:

```
Final Message: Hello World Whatsup
```

Here's an easy explanation of the flow of this microservice application, how the code works, and how Apache Pulsar is involved in the process:

### Overview:

You are building a system where three separate services each handle one part of the sentence "Hello World Whatsup." These services communicate through **Apache Pulsar**, which is a messaging platform, and pass the parts of the sentence between them. The final result, "Hello World Whatsup," is printed on a webpage by another service.

### Flow of the Application:

1. **Start of the System**:

   - The application starts by setting up **Apache Pulsar** on your local machine or docker engine or aws etc. Pulsar acts like a post office that handles messages (in this case, words of the sentence) between different services.
   - Each of the microservices you built (Hello Service, World Service, Whatsup Service, and Web Service) works independently but communicates with each other through **topics** in Pulsar.

2. **Hello Service**:

   - **What it does**: This is the first microservice. It sends the message "Hello" to a Pulsar **topic** called `"hello-topic"`.
   - **How it works**: When the Hello Service runs, it connects to Pulsar, publishes the word "Hello" to `"hello-topic"`, and then stops.

3. **World Service**:

   - **What it does**: This service is listening (subscribing) to the `"hello-topic"` where "Hello" was sent. Once it receives "Hello," it adds "World" to it, creating "Hello World." Then, it sends this new message to the next topic, `"world-topic"`.
   - **How it works**: World Service is waiting for messages in `"hello-topic"`. Once "Hello" arrives, it appends "World" and sends the new message to `"world-topic"`.

4. **Whatsup Service**:

   - **What it does**: This service is subscribed to `"world-topic"`. Once it receives the message "Hello World," it adds "Whatsup," making the full sentence "Hello World Whatsup," and sends this to `"final-topic"`.
   - **How it works**: Whatsup Service listens to `"world-topic"`, processes the message, appends "Whatsup," and sends the final sentence to `"final-topic"`.

5. **Web Service**:
   - **What it does**: This service is subscribed to `"final-topic"`. It waits for the final message "Hello World Whatsup," and when it receives it, it displays this message on a webpage.
   - **How it works**: Web Service listens to `"final-topic"` and once the complete message arrives, it prints it on a web page. When you visit `http://localhost:8080`, you will see the final sentence.

### How Apache Pulsar Works in This:

- **Pulsar's Role**: Apache Pulsar acts as the communication platform between the microservices. It's like a middleman that receives messages from one service and delivers them to another.
  - Pulsar uses **topics** to organize this communication. Each microservice sends or listens to specific topics.
  - In this example, three topics are used: `"hello-topic"`, `"world-topic"`, and `"final-topic"`.
- **Behind the Scenes**:
  - When a microservice publishes a message (like "Hello" or "World"), Pulsar stores it temporarily in the corresponding topic.
  - Other microservices, which are subscribed to these topics, receive the message from Pulsar, process it, and then publish the next message to the following topic.
  - This flow allows the services to work together asynchronously without directly interacting with each other.

### Summary of the Flow:

1. **Hello Service** sends "Hello" to `"hello-topic"`.
2. **World Service** listens to `"hello-topic"`, receives "Hello," adds "World," and sends "Hello World" to `"world-topic"`.
3. **Whatsup Service** listens to `"world-topic"`, receives "Hello World," adds "Whatsup," and sends "Hello World Whatsup" to `"final-topic"`.
4. **Web Service** listens to `"final-topic"` and displays the final message "Hello World Whatsup" on the webpage.

This setup showcases how microservices can work together using Pulsar, which ensures that the services communicate by passing messages through topics. Each service only focuses on its own task, and Pulsar manages the communication in the background.
