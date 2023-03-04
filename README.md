# swarm02
## nginx-golang

### Ref
* https://github.com/docker/awesome-compose/tree/master/nginx-golang
### url
* https://spcn29swarm01.xops.ipv9.me/
### wakatime project
* https://wakatime.com/@spcn29/projects/nkcbladfhw

ถ้าทำตาม swarm01 แล้ว สามารถทำขั้นตอน clone แอพ จาก github ได้เลย
### ขั้นตอนการติดตั้งใน VM
1. สร้าง VM โดยมี spec ดังนี้
    * CPU 2 cores
    * Ram 2 GB
    * HDD 32 GB
    * Ubuntu 22.04
2. ตั้งเวลา

    เป็นเวลาในประเทศไทย
    ```
    sudo -i
    ```
    ```
    timedatectl set-timezone Asia/Bangkok
    ```
3. ตรวจสอบ IP address ของ VM เพื่อ Remote ssh
    ```
    ip a
    ```
### เชื่อมต่อ Remote ssh ผ่าน VS Code และติดตั้ง docker engine 
1. ติดตั้ง Docker, wakatime และ ssh remote ผ่าน VS Code
2. เชื่อมต่อ Remote ssh 
    * กดปุ่นสีเขียวด้านซ้ายล่าง
![images](https://user-images.githubusercontent.com/109591322/222915170-eea6290c-3494-4998-a50e-504b6d00b3ca.png)

    กด Connect to Host > Configure SSH Hosts > /Users/p/.ssh/config > ใส่คำสั่งด้านล่าง

    
        Host        // ชื่อ Host ที่ต้องการตั้ง
        HostName    // ip address
        User        // ชื่อ hostname จากเครื่องที่จะเชื่อม
    
    ![images](https://user-images.githubusercontent.com/109591322/222915172-0b26924e-d083-4126-80e5-6cec87b32832.png)
4. ติดตั้ง Docker, wakatime ที่เครื่อง SSH ที่เชื่อมต่อ 

5. ติดตั้ง docker engine 
    ```
    apt update; apt upgrade -y #อัปเดตแพ็คเกจภายในเครื่อง

    apt-get install ca-certificates curl wget gnupg lsb-release -y #ติดตั้งแพ็คเกจ

    mkdir -m 0755 -p /etv/apt/keyrings

    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg #ดาวโหลดไฟล์แพ็คเกจ Docker

    echo \ "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \ $(lsb_release -cs) stable" |  tee /etc/apt/sources.list.d/docker.list > /dev/null

    apt-get update #อัปเดทไฟล์แพ็คเกจเพื่อไว้สำหรับให้ติดตั้ง
    apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y #ติดตั้ง Docker

    reboot
    ```
4. clone มา 3 Node ได้แก่
    * manage
    * work1
    * work2
5. ตั้ง hostname
    ```
    hostnamectl set-hostname "ชื่อHostname"
    ```
6. Reset Machine ID 

    เพื่อไม่ให้ IP ชนกัน
    ```
    cp /dev/null /etc/machine-id
    rm /var/lib/dbus/machine-id
    ln -s /etc/machine-id /var/lib/dbus/machine-id
    init 0
    ```


### เตรียม stack swarm  
1. สร้าง Token 
    เพื่อนำไปใส่ที่ Node ที่ต้องการให้เชื่อมต่อ
    * พิมพ์คำสั่งในเครื่อง manage
    ```
    docker swarm init 
    ```
    * เมื่อได้ Token แล้วนำไปใส่ใน Node ที่ต้องการให้เชื่อมกัน
        * work1
        * work2
    
    ตรวจสอบการเชื่อมต่อของ Node
    ```
    docker node ls
    ```
### deploy portainer for swarm 
1.  ดาวน์โหลดและติดตั้ง
    ```
    curl -L https://downloads.portainer.io/ce2-17/portainer-agent-stack.yml -o portainer-agent-stack.yml
    ```
    ```
    docker stack deploy -c portainer-agent-stack.yml portainer
    ```
### clone แอพ จาก github    
* nginx-golang
1. compose Up ไฟล์ compose.yaml
2. docker login
    ```
    docker login 
    ```
    ใส่ Username , password ของ dockerhub

    ถ้าเคยใส่แล้วสามารถข้ามได้เลย
3. เพิ่ม tag 
    ``` 
    docker images
    ```
    ```
    docker tag nginx-golang-backend :latest phisit11/nginx-golang:0227
    ```
4. push Image to DockerHub
    ```
    docker push phisit11/nginx-golang:0227
    ```
### เพิ่ม stacks บน portainer
1. เปิด https://portainer.ipv9.me
2. กด Stacks > add stacks 
3. ใส่ code ลงในส่วนของ Web editor บน portainer
* อยู่ในไฟล์ compose-revp.yaml
```
version: '3.7'
services:
 
 db: 
    environment:
     POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    networks:
      - default
    volumes:
      - db_data:/var/lib/postgres/data

  backend:
    image: phisit11/nginx-golang:0227
    volumes:
      - static_data:/usr/src/app/static
    networks:
      - webproxy
      - default
    depends_on:
      - db  
    deploy:
      replicas: 1
      labels:
        - traefik.docker.network=webproxy
        - traefik.enable=true
        - traefik.constraint-label=webproxy
        - traefik.http.routers.${APPNAME}-https.entrypoints=websecure
        - traefik.http.routers.${APPNAME}-https.rule=Host("${APPNAME}.xops.ipv9.me")
        - traefik.http.routers.${APPNAME}-https.tls.certresolver=defauit
        - traefik.http.services.${APPNAME}.loadbalancer.server.port=80
      restart_policy:
        condition: any
      update_config:
        delay: 5s

      
volumes:
    static_data:
    db_data:

networks:
    default:
      driver: overlay
    webproxy:
      external: true

```    
4. Add an environment variables
    * Environment variables > Add an environment variables
    * name APPNAME value spcns29warm02
5. ทดสอบการใช้งาน 
    * ไปที่ https://spcn29swarm02.xops.ipv9.me/  
    ถ้าสามารถใช้งานได้จะขึ้นหน้าตาดังรูปด้านล่าง
![](https://user-images.githubusercontent.com/109591322/222915175-9c633d94-6c9a-44d4-bb1f-8e621602084d.png)