# How to deploy CoC on AWS with k8s

## Service
+ fsvc.yaml -- LoadBalancer -- expose front-end
+ bwsvc.yaml -- NodePort -- expose back-end-weather
+ bcsvc.yaml -- ClusterIP -- expose back-end-chatbox
+ msvc.yaml -- ClusterIP -- expose mysql

## Database
+ mpv.yaml -- persistent volume
+ mpvc.yaml -- persistent volume claim
+ mconfig.yaml -- ConfigMap
+ msecret.yaml -- Secret -- password
+ mysql.yaml -- database
+ exec m-dep pod
  - execute `mysql -u root -p`
  - pwd =  `12345678`
  - execute `create database weather;`

## Backend & Database Initialize
+ bwdep.yaml -- deployment -- backend-weather
+ exec bw-dep pods
  - execute `python manage.py makemigrations`
  - execute `python manage.py migrate`
+ bcdep.yaml -- deployment -- backend-chatbox

## Jobs
+ warning_cj.yaml -- job -- provide warning
+ weather_cj.yaml -- job -- provide weather
+ predict_cj.yaml -- job -- provide predict

## Frontend
+ fdep.yaml -- deployment -- front-end



~~~mermaid
graph TD
  A[User Connection and Identification] -->|User Accesses Application| B[Frontend Message Sending and Receiving]
  B -->|User Inputs Message| C[Backend Message Processing - JDBC+OpenAIAgent]
  C -->|WebSocket Communication| D[Real-Time Communication]
  D -->|Broadcast Message| B
  C -->|Store Message| E[MySQL Database]
  C -->|Injection Prevention and Security| G[HTML Escaping and Input Validation]
  A -->|Generate UUID| H[Unique User ID]
  E -->|Persistent Storage| I[Persistent Volume]
  H -->|Store Locally| B
  I -->|Ensure Data Persistence| E
  A -->|Browser Local Storage| B

  subgraph Technology Stack
    B
    D
    E
    G
    H
    I
  end

  subgraph Cloud Deployment
    J[Kubernetes Configuration]
    K[AWS Platform]
    J -->|Automated Deployment| K
    J -->|Load Balancing| D
    K -->|High Availability| J
  end

~~~

- #### 实现

  1. **技术栈**
     - **前端**：使用Vue.js构建用户界面，通过WebSocket实现实时通信。
     - **后端**：使用Java编写，通过Spring Boot框架处理WebSocket连接和业务逻辑。
     - **数据库**：使用MySQL存储用户消息和点赞数据，通过JDBC与Java应用程序连接。
  2. **用户ID生成**
     - 使用UUID (Universally Unique Identifier) 来生成唯一的用户标识，确保每个用户的唯一性和隐私性。
  3. **防注入安全**
     - 对用户输入进行转义，防止SQL注入和XSS攻击。例如，在发送消息时，对消息内容进行HTML转义，防止恶意脚本执行。
  4. **WebSocket**
     - WebSocket提供了在客户端和服务器之间进行全双工通信的能力，使聊天室能够实现即时消息传递。
     - 在Java后端，通过Spring Boot WebSocket支持，管理WebSocket连接、消息发送和接收。

  #### 部署

  1. **Kubernetes配置**
     - 使用多个YAML文件配置和管理Kubernetes资源，包括服务、部署和持久化存储卷。
     - 例如，`bcdep.yaml`文件用于部署后端聊天服务，`bcsvc.yaml`文件用于定义ClusterIP服务，确保后端服务只能在集群内部访问。
     - 前端通过`fsvc.yaml`和`fdep.yaml`文件配置，使用LoadBalancer服务将前端应用公开给外部访问。
  2. **数据库管理**
     - 使用`mpv.yaml`和`mpvc.yaml`文件定义持久化存储卷和存储卷声明，确保MySQL数据的持久化存储。
     - 通过ConfigMap和Secret管理数据库配置和凭证，确保安全性和配置的灵活性。





### 部署部分

#### Kubernetes配置

在AWS上的Kubernetes平台上，我们使用多个YAML文件来配置和管理我们的Kubernetes资源，确保我们的服务能够高效且可靠地运行。

1. **服务定义**

   - **前端服务 (`fsvc.yaml`)**：使用LoadBalancer类型的服务，将前端应用暴露给外部访问。LoadBalancer会自动创建一个公共IP地址，用户可以通过这个IP地址访问我们的前端应用。

   ```
   yaml复制代码apiVersion: v1
   kind: Service
   metadata:
     name: front-end-service
   spec:
     type: LoadBalancer
     ports:
       - port: 80
     selector:
       app: front-end
   ```

   - **后端天气服务 (`bwsvc.yaml`)**：使用NodePort类型的服务，将后端天气服务暴露给外部访问。NodePort会在每个Node上打开一个指定的端口，用户可以通过这个端口访问我们的后端服务。

   ```
   yaml复制代码apiVersion: v1
   kind: Service
   metadata:
     name: backend-weather-service
   spec:
     type: NodePort
     ports:
       - port: 8000
         nodePort: 30001
     selector:
       app: backend-weather
   ```

   - **后端聊天服务 (`bcsvc.yaml`)**：使用ClusterIP类型的服务，确保后端聊天服务只能在集群内部访问，保证安全性和内部通信的高效性。

   ```
   yaml复制代码apiVersion: v1
   kind: Service
   metadata:
     name: backend-chat-service
   spec:
     type: ClusterIP
     ports:
       - port: 8080
     selector:
       app: backend-chat
   ```

2. **持久化存储**

   - **持久化存储卷 (`mpv.yaml`)**：定义持久化存储卷，确保MySQL数据能够持久化存储，即使Pod重启或迁移也不会丢失数据。

   ```
   yaml复制代码apiVersion: v1
   kind: PersistentVolume
   metadata:
     name: mysql-pv
   spec:
     capacity:
       storage: 1Gi
     accessModes:
       - ReadWriteOnce
     hostPath:
       path: "/mnt/data"
   ```

   - **持久化存储卷声明 (`mpvc.yaml`)**：定义持久化存储卷声明，绑定持久化存储卷和Pod，使Pod能够使用定义的存储资源。

   ```
   yaml复制代码apiVersion: v1
   kind: PersistentVolumeClaim
   metadata:
     name: mysql-pvc
   spec:
     accessModes:
       - ReadWriteOnce
     resources:
       requests:
         storage: 1Gi
   ```

3. **数据库配置**

   - **ConfigMap (`mconfig.yaml`)**：通过ConfigMap管理数据库配置，例如数据库名称、用户名等，方便集中管理和修改配置。

   ```
   yaml复制代码apiVersion: v1
   kind: ConfigMap
   metadata:
     name: mysql-config
   data:
     database: weather
     user: root
   ```

   - **Secret (`msecret.yaml`)**：通过Secret管理敏感信息，例如数据库密码，确保这些信息的安全性。

   ```
   yaml复制代码apiVersion: v1
   kind: Secret
   metadata:
     name: mysql-secret
   type: Opaque
   data:
     password: MTIzNDU2Nzg= # base64 encoded password
   ```

4. **应用部署**

   - **前端应用部署 (`fdep.yaml`)**：定义前端应用的部署，使用Vue.js和Nginx构建并提供前端服务。

   ```
   yaml复制代码apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: front-end-deployment
   spec:
     replicas: 3
     selector:
       matchLabels:
         app: front-end
     template:
       metadata:
         labels:
           app: front-end
       spec:
         containers:
           - name: front-end
             image: nginx:latest
             ports:
               - containerPort: 80
   ```

   - **后端天气服务部署 (`bwdep.yaml`)**：定义后端天气服务的部署，使用Django框架处理天气数据。

   ```
   yaml复制代码apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: backend-weather-deployment
   spec:
     replicas: 3
     selector:
       matchLabels:
         app: backend-weather
     template:
       metadata:
         labels:
           app: backend-weather
       spec:
         containers:
           - name: backend-weather
             image: django:latest
             ports:
               - containerPort: 8000
   ```

   - **后端聊天服务部署 (`bcdep.yaml`)**：定义后端聊天服务的部署，使用Java和Spring Boot框架处理聊天消息。

   ```
   yaml复制代码apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: backend-chat-deployment
   spec:
     replicas: 3
     selector:
       matchLabels:
         app: backend-chat
     template:
       metadata:
         labels:
           app: backend-chat
       spec:
         containers:
           - name: backend-chat
             image: java:latest
             ports:
               - containerPort: 8080
   ```

5. **数据库部署**

   - **MySQL数据库部署 (`mysql.yaml`)**：定义MySQL数据库的部署，确保数据持久化存储和高效访问。

   ```
   yaml复制代码apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: mysql-deployment
   spec:
     selector:
       matchLabels:
         app: mysql
     strategy:
       type: Recreate
     template:
       metadata:
         labels:
           app: mysql
       spec:
         containers:
           - name: mysql
             image: mysql:5.7
             env:
               - name: MYSQL_ROOT_PASSWORD
                 valueFrom:
                   secretKeyRef:
                     name: mysql-secret
                     key: password
             ports:
               - containerPort: 3306
             volumeMounts:
               - name: mysql-persistent-storage
                 mountPath: /var/lib/mysql
         volumes:
           - name: mysql-persistent-storage
             persistentVolumeClaim:
               claimName: mysql-pvc
   ```

### 云端结合的意义

部署在云端，特别是使用Kubernetes和AWS，有以下几个显著的好处：

1. **弹性扩展**
   - Kubernetes能够根据实际负载情况，自动扩展和缩减资源。在用户量激增时，系统能够自动增加Pod数量，确保性能和响应速度；在用户量减少时，系统能够自动缩减资源，降低成本。
2. **高可用性**
   - Kubernetes通过多副本部署和负载均衡，确保即使某个节点发生故障，其他节点可以接管工作，保证服务的连续性和高可用性。
3. **便捷管理**
   - Kubernetes提供了一套完整的自动化部署和管理工具，使得我们的应用在部署、更新、监控和维护方面更加高效。利用Helm等工具，我们可以轻松管理应用的生命周期。
4. **数据安全**
   - 通过使用ConfigMap和Secret，我们可以确保敏感信息的安全性。持久化存储卷保证了数据不会因为Pod的重启或迁移而丢失。

------

希望这些详细信息对你有帮助，如果你有更多问题或需要进一步的修改，请告诉我。
