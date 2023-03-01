# swarm02
## nginx-golang

### Ref
* https://github.com/docker/awesome-compose/tree/master/nginx-nodejs-redis
### url
* https://spcn29swarm01.xops.ipv9.me/

ถ้าทำตามswarm01 แล้ว สามารถทำขั้นตอน ... ได้เลย
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
3. ติดตั้ง docker engine 
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
### ติดตั้ง Portainer CE Docker swarm
1.  ดาวน์โหลดและติดตั้ง
    ```
    curl -L https://downloads.portainer.io/ce2-17/portainer-agent-stack.yml -o portainer-agent-stack.yml
    ```
    ```
    docker stack deploy -c portainer-agent-stack.yml portainer
    ```


