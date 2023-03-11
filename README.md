# swarm02
## nginx-golang

### Ref
* https://github.com/docker/awesome-compose/tree/master/nginx-golang
### url
* https://spcn29swarm02.xops.ipv9.me/
### wakatime project
* https://wakatime.com/@spcn29/projects/nkcbladfhw
---

### ขั้นตอนสำหรับทำงาน    
* nginx-golang

  clone แอพ จาก github

1. compose Up ไฟล์ compose.yaml
2. docker login

    ถ้าเคยเข้าสู่ระบบแล้วสามารถข้ามได้เลย

    ```
    docker login 
    ```
    ใส่ Username , password ของ dockerhub

3. เพิ่ม tag 
    
    โดยจะใช้ tag 0227 คือวันที่ในการ push เดือน 02 วันที่ 27
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
    ![](https://user-images.githubusercontent.com/109591322/224491497-54bc4234-e419-4f3b-9386-74bfd0fd0a88.png)

    เมื่อ push สำเร็จจะขึ้นหน้าตาแบบรูปด้านบน ที่ dockerhub
   
### เพิ่ม stacks บน portainer
1. เปิด https://portainer.ipv9.me
2. กด Stacks > add stacks 
3. ใส่ code ลงในส่วนของ Web editor บน portainer
* อยู่ในไฟล์ compose-revp.yaml
```
version: '3.7'
services:
 
 db: 
    image: postgres
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
  * Version ระบุเป็น version ที่สนับสนุนกับ application ที่ใช้งาน ควรมากกว่า 3 
  * Service ใช้ระบุส่วนที่ใช้งาน ประกอบไปด้วย image, command, volumes, restart, networks, environment, expose และ deploy เป็นต้น
  * volumes ใช้สร้างที่เก็บข้อมูล
  * Networks ใช้งานการตั้งค่าเครือข่ายภายนอกคอนเทนเนอร์
4. Add an environment variables
    * Environment variables > Add an environment variables
    * name APPNAME value spcns29swarm02
5. ทดสอบการใช้งาน 
    * ไปที่ https://spcn29swarm02.xops.ipv9.me/  
    ถ้าสามารถใช้งานได้จะขึ้นหน้าตาดังรูปด้านล่าง
![](https://user-images.githubusercontent.com/109591322/222915175-9c633d94-6c9a-44d4-bb1f-8e621602084d.png)