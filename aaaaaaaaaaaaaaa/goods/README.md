# 用户中心微服务

#### 介绍
包含用户中心相关api

#### 目录结构

1.  api 用于注册grpc api 参数验证相关业务需要再api目录实现
2.  service 用于实现我的业务
3.  model 用于处理数据库 
4.  utils 用与于工具类编写

#### 生成加密证书步骤
1. 定义证书配置文件： ``openssl.cnf``

````
    [ req ]
    default_bits       = 2048
    prompt             = no
    default_md         = sha256
    distinguished_name = req_distinguished_name
    req_extensions     = req_ext
    
    [ req_distinguished_name ]
    C  = US
    ST = YourState
    L  = YourCity
    O  = YourOrganization
    CN = www.2108a.com
    
    [ req_ext ]
    subjectAltName = @alt_names
    
    [ alt_names ]
    DNS.1 = www.2108a.com
    DNS.2 = 2108a.com  # 添加其他需要的域名或IP
````

2. 生成CA私钥：
    ``openssl genpkey -algorithm RSA -out ca.key``
3. 生成CA证书：``openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 -out ca.pem -config openssl.cnf
   ``
4. 生成服务器私钥：``openssl genpkey -algorithm RSA -out server_key.pem``
5. 生成服务器证书请求（CSR）：``openssl req -new -key server_key.pem -out server.csr -config openssl.cnf``
6. 使用CA签发服务器证书：``openssl x509 -req -in server.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out server_cert.pem -days 365 -sha256 -extfile openssl.cnf -extensions req_ext
   ``
