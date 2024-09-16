Here’s a detailed installation guide for **Apache Pulsar** on different platforms, including Windows, Linux, and AWS. The guide will also show you how to use Pulsar in each environment once installed.

---

## 1. **Installation on Windows**

Apache Pulsar is a Java-based service, and it can be run on Windows by installing Java and downloading Pulsar’s standalone version.

### Prerequisites

- **Java**: Apache Pulsar requires Java 8 or 11. You need to install it before setting up Pulsar.

#### Step-by-Step Guide:

### Step 1: Install Java

1. Download and install Java Development Kit (JDK 8 or 11) from the [Oracle JDK site](https://www.oracle.com/java/technologies/javase-jdk11-downloads.html) or [OpenJDK](https://adoptopenjdk.net/).
2. Add Java to your PATH environment variables:
   - Right-click on **This PC** > **Properties** > **Advanced System Settings** > **Environment Variables**.
   - In the **System variables** section, find the `Path` variable, edit it, and add the path to your `java\bin` directory.
3. Verify Java installation by running:
   ```bash
   java -version
   ```

### Step 2: Download Pulsar

1. Go to the [Pulsar Download Page](https://pulsar.apache.org/download).
2. Download the latest **binary release** for your version (choose a `.tar.gz` or `.zip` file).
3. Extract the file to your desired location.

### Step 3: Run Pulsar in Standalone Mode

1. Open **Command Prompt** and navigate to the directory where Pulsar is extracted:
   ```bash
   cd apache-pulsar-x.y.z
   ```
2. Run Pulsar in standalone mode:
   ```bash
   bin\pulsar standalone
   ```
   Pulsar will start up, and you can access its admin interface at `http://localhost:8080`.

### Step 4: Using Pulsar

- **Publish a message** to a topic:
  ```bash
  bin\pulsar-client produce my-topic --messages "Hello Pulsar"
  ```
- **Consume messages** from a topic:
  ```bash
  bin\pulsar-client consume my-topic -s "sub-1"
  ```

---

## 2. **Installation on Linux**

You can install Apache Pulsar on Linux (Ubuntu/Debian) either manually or via Docker.

### Prerequisites

- **Java**: Pulsar requires Java 8 or 11 on Linux as well.

#### Step-by-Step Guide:

### Step 1: Install Java

1. Install Java via `apt`:
   ```bash
   sudo apt update
   sudo apt install openjdk-11-jdk
   ```
2. Verify Java installation:
   ```bash
   java -version
   ```

### Step 2: Download Pulsar

1. Download the latest **Pulsar release**:
   ```bash
   wget https://downloads.apache.org/pulsar/pulsar-2.10.0/apache-pulsar-2.10.0-bin.tar.gz
   ```
2. Extract the tarball:
   ```bash
   tar xvfz apache-pulsar-2.10.0-bin.tar.gz
   cd apache-pulsar-2.10.0
   ```

### Step 3: Run Pulsar in Standalone Mode

1. Start Pulsar in standalone mode:
   ```bash
   bin/pulsar standalone
   ```
   You can now access the Pulsar admin interface at `http://localhost:8080`.

### Step 4: Using Pulsar

- **Publish a message** to a topic:
  ```bash
  bin/pulsar-client produce my-topic --messages "Hello Pulsar"
  ```
- **Consume messages** from a topic:
  ```bash
  bin/pulsar-client consume my-topic -s "sub-1"
  ```

---

## 3. **Running Pulsar on AWS**

You can either deploy Pulsar manually on an EC2 instance or use managed services like **StreamNative Cloud**, which offers Pulsar as a service. Here's how to set it up on an EC2 instance manually.

### Prerequisites

- **AWS account** with an EC2 instance set up (Amazon Linux or Ubuntu).
- **SSH access** to your EC2 instance.
- **Java** installed on the EC2 instance.

### Step-by-Step Guide:

### Step 1: Set Up EC2 Instance

1. Launch an EC2 instance from the AWS Management Console (Amazon Linux or Ubuntu).
2. Set up the necessary security group rules to allow traffic on the required ports (8080 for Pulsar’s web interface, 6650 for client connections).
3. SSH into your EC2 instance.

### Step 2: Install Java

1. Install Java on your EC2 instance:
   ```bash
   sudo yum update -y   # For Amazon Linux
   sudo yum install java-11-openjdk -y
   ```
   Or on Ubuntu:
   ```bash
   sudo apt update
   sudo apt install openjdk-11-jdk -y
   ```

### Step 3: Download and Install Pulsar

1. Download the Pulsar binaries on the EC2 instance:
   ```bash
   wget https://downloads.apache.org/pulsar/pulsar-2.10.0/apache-pulsar-2.10.0-bin.tar.gz
   ```
2. Extract the file:
   ```bash
   tar xvfz apache-pulsar-2.10.0-bin.tar.gz
   cd apache-pulsar-2.10.0
   ```

### Step 4: Start Pulsar in Standalone Mode

1. Run Pulsar:

   ```bash
   bin/pulsar standalone
   ```

2. Access the web interface at:
   ```
   http://<your-ec2-public-ip>:8080
   ```

### Step 5: Using Pulsar

- **Publish a message** to a topic:
  ```bash
  bin/pulsar-client produce my-topic --messages "Hello Pulsar"
  ```
- **Consume messages** from a topic:
  ```bash
  bin/pulsar-client consume my-topic -s "sub-1"
  ```

---

## 4. **Using Managed Apache Pulsar on Cloud (StreamNative Cloud)**

If you prefer not to manage Pulsar installations and infrastructure yourself, you can use **StreamNative Cloud** for a managed Pulsar service.

### Step-by-Step Guide:

### Step 1: Sign Up for StreamNative Cloud

1. Go to [StreamNative Cloud](https://streamnative.io/) and sign up.
2. Create a new Pulsar cluster from the dashboard.

### Step 2: Get Cluster Credentials

1. After creating a cluster, download the connection credentials (certificates and configurations).
2. Set up your Go or Java Pulsar client with these credentials.

### Step 3: Using Pulsar from Your Application

With the cluster running, use the **Pulsar service URL** and credentials to connect from your application.

For example, in Go:

```go
client, err := pulsar.NewClient(pulsar.ClientOptions{
    URL: "pulsar+ssl://<your-streamnative-cluster-url>:6651",
    Authentication: pulsar.NewAuthenticationTLS("client-cert.pem", "client-key.pem"),
})
```

---

## Summary of Commands and URLs:

### 1. **Running Pulsar in Standalone Mode**

- **Windows**:
  ```bash
  bin\pulsar standalone
  ```
- **Linux/Mac**:
  ```bash
  bin/pulsar standalone
  ```

### 2. **Publishing a Message**:

```bash
bin/pulsar-client produce my-topic --messages "Hello Pulsar"
```

### 3. **Consuming a Message**:

```bash
bin/pulsar-client consume my-topic -s "sub-1"
```

### Web Interface:

- Pulsar Admin Dashboard: `http://localhost:8080` or `http://<EC2_PUBLIC_IP>:8080` (for EC2 instances).

---

This installation guide should help you set up and run Apache Pulsar in various environments. Let me know if you need any additional details or troubleshooting tips!
