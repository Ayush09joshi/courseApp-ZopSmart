
# courseApp-ZopSmart

This is a mini project consist of RESTful APIs (CREATE, READ, UPDATE, DELETE) created using GoFr, with integrated backend using MySQL and unit tests.


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file (configs/.env)

```bash
APP_NAME=test-service
HTTP_PORT=9000

DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root123
DB_NAME=test_db
DB_PORT=3306
DB_DIALECT=mysql
```

## Database Setup

```bash
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
```
```bash
docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE courses (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, price INT NOT NULL, author VARCHAR(255) NOT NULL);"
```

## Run 
```bash
go run main.go
```

## Postman Collection
You can visit the following website or download the zip folder for postman collection of the project.

```bash
https://documenter.getpostman.com/view/24503733/2s9YkkeNYZ
```
[RESTfull API-ZopSmart_miniProject.postman_collection.zip](https://github.com/Ayush09joshi/courseApp-ZopSmart/files/13697598/RESTfull.API-ZopSmart_miniProject.postman_collection.zip)


## API
You can run the APIs at 

```bash
 http://localhost:3000
```

## Routes
GET (/get)
```bash
  http://localhost:3000/get
```
POST (/create)
```bash
  http://localhost:3000/create
```
UPDATE (/update/{id})
```bash
   http://localhost:3000/update/2
```
DELETE (/delete/{id})
```bash
   http://localhost:3000/delete/2
```

## Unit Test
Achived Unit Test Coverage 100%, 90.4% for main_test and store_test respectively.

## Diagrams
Sequence

![SEQUENCE DIAGRAM](https://github.com/Ayush09joshi/courseApp-ZopSmart/assets/105715149/3932d954-159d-4c57-8e80-3738d3efd340)

UML

![UML DIAGRAM](https://github.com/Ayush09joshi/courseApp-ZopSmart/assets/105715149/6d9ce490-0e07-44a8-9b98-98fe1876f9f4)
