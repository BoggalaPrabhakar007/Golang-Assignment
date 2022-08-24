
# Golang Assignment

This project is about getting the json data and storing into the database and doing CRUD oeration on data.

## Project Architecture
![Architecture](/images/PortServiceArchitecture.JPG)
## Clone project

To clone this project run.

```bash
  PS C:\Users\boggala\go\src> git clone https://github.com/BoggalaPrabhakar007/golang-assignment.git
```


Navigate to project folder.

```bash
 PS C:\Users\boggala\go\src> cd golang-assignment
```
## Build the microservice (project) docker image

run below command.

```bash 
   PS C:\Users\boggala\go\src\golang-assignment> docker build -t goland-assignment-image .
   PS C:\Users\boggala\go\src\golang-assignment> docker images                            
   REPOSITORY                TAG       IMAGE ID       CREATED          SIZE  
   goland-assignment-image   latest    50416563fb36   30 minutes ago   1.15GB
```
 
## Deploy the microservice and mongodb 

run below command.

```bash 
 PS C:\Users\boggala\go\src\golang-assignment> docker-compose up -d
 PS C:\Users\boggala\go\src\golang-assignment> docker container ls
 CONTAINER ID   IMAGE                            COMMAND                  CREATED          STATUS          PORTS                      NAMES
 98e0fc155b64   mongo:4.0                        "docker-entrypoint.s…"   29 seconds ago   Up 22 seconds   0.0.0.0:27017->27017/tcp   golang-assignment_mongodb_1
 887c77b26e85   goland-assignment-image:latest   "/build/golang-assig…"   29 seconds ago   Up 22 seconds   0.0.0.0:8080->8080/tcp     golang-assignment_appservice_1
```
## Postman
open the postman in your system follow the below screenshots.
## InsertData
![InsertData](/images/InsertData.JPG)
## GetData
![GetData](/images/GetData.JPG)
## UpdateData
![UpdateData](/images/UpdateData.JPG)
## After Update
![AfterUpdate](/images/AfterUpdate.JPG)
## GetDataById
![GetDataByID](/images/GetDataByID.JPG)
## DeleteDataById
![DeleteDataById](/images/DeleteData.JPG)
## Swagger 
Swagger APIs for the service.
## InsertData
![InsertData](/images/Swagger_InsertData.JPG)
## GetData
![GetData](/images/Swagger_GetData.JPG)
## UpdateData
![UpdateData](/images/Swagger_UpdateData.JPG)

## GetDataById
![GetDataByID](/images/Swagger_GetDataByID.JPG)
## DeleteDataById
![DeleteDataById](/images/Swagger_DeleteData.JPG)
## Swagger APIs
![Swagger](/images/Swagger.JPG)

## To Stop the microservice and mongodb

run below command.

```bash 
   PS C:\Users\boggala\go\src\golang-assignment> docker-compose down
```

  



