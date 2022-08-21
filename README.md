
# Golang Assignment

This project is about getting the json data and storing into the database and doing CRUD oeration on data.

## Clone project

To clone this project run

```bash
  PS C:\Users\boggala\go\src> git clone https://github.com/BoggalaPrabhakar007/golang-assignment.git
```


Navigate to project folder

```bash
 PS C:\Users\boggala\go\src> cd golang-assignment
```
## Build the microservice (project) docker image

run below command

```bash 
   PS C:\Users\boggala\go\src\golang-assignment> docker build -t goland-assignment-image .
   PS C:\Users\boggala\go\src\golang-assignment> docker images                            
   REPOSITORY                TAG       IMAGE ID       CREATED          SIZE  
   goland-assignment-image   latest    50416563fb36   30 minutes ago   1.15GB
```
 
## Deploy the microservice and mongodb 

run below command

```bash 
 PS C:\Users\boggala\go\src\golang-assignment> docker-compose up -d
 PS C:\Users\boggala\go\src\golang-assignment> docker container ls
 CONTAINER ID   IMAGE                            COMMAND                  CREATED          STATUS          PORTS                      NAMES
 98e0fc155b64   mongo:4.0                        "docker-entrypoint.s…"   29 seconds ago   Up 22 seconds   0.0.0.0:27017->27017/tcp   golang-assignment_mongodb_1
 887c77b26e85   goland-assignment-image:latest   "/build/golang-assig…"   29 seconds ago   Up 22 seconds   0.0.0.0:8080->8080/tcp     golang-assignment_appservice_1
```

## To Stop the microservice and mongodb 

run below command

```bash 
   PS C:\Users\boggala\go\src\golang-assignment> docker-compose down
```




  



