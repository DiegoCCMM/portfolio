# What to do?
The repo contains a skeleton of 4 Spring Boot applications, plus a Docker Compose configuration which spins up the following working environment.

<img width="365" alt="image" src="https://user-images.githubusercontent.com/15728394/199699196-3bf20be2-cc51-4718-8cc2-454c8397c9d4.png">

- _Public Service_: main entry point to system, and the only accesible one to the public.
- _Priority Sale Service_: contains the priority sale logic. 
- _Adi-Club Service_: provides the information of our adiClub members. e-mail, points, registration date...
- _Email Service_: sends a confirmation mail to the lucky ones. 

## 1. Business Requirements

From a functional perspective you will need to implement two functionalities:
- Subscription functionality. 
- Email delivery functionality

#### _Example:_
Considering the following user data:

<img width="250" alt="image" src="https://user-images.githubusercontent.com/15728394/210075353-69cfef77-9de6-474a-af12-ff1c7cd70ccd.png">

This is a timeline example of the behaviour:

![timeline](https://user-images.githubusercontent.com/15728394/210075366-60020f11-6d3d-4f7e-a67d-708604228699.png)

The logic keeps going on and on.

### Business requirements diagrams
For better understanding of the business flow, check the sequence diagrams below.
Please, do not make strict assumptions about the execution order or time of the events, the communication between microservices, or how the data must be saved internally. Those decisions are your own responsibility, and you can make any decision as long as you cover all functional requirements.


![Subscription functionality](https://user-images.githubusercontent.com/5638405/204595709-a1802e67-0ded-4139-9eae-f988a9da954d.png)

![Email delivery functionality](https://user-images.githubusercontent.com/5638405/204595722-8efda463-fa95-4e15-b873-d838bb9beaf9.png)


## 2. Non Functional Requirements:
- **Security** for non-public services
- Proper **exception handling** and REST responses.​
- Unit **testing** for meaningful code (business logic / services).​
- Ability to **scale horizontally** based on incoming load
- **Design for failure**, ability to recover from exception conditions


### Important considerations about the implementation:
- Subscription functionality must be done by exposing a REST API. 
- No need to implement a frontend to create the subscriptions: a command or postman collection will be enough. 
- No need to send a real email: a log line from email service will be enough. 
- adiClub is a service simulating an external web service which owns loyalty program data.
  - It can be assumed that all users are adiClub members.
  - adiClub service will return random data representing the number of points in the club and the date when that user was first registered in the loyalty program.
  - No need to modify adiClub service, you can just query the endpoint provided in the template by using a REST client.

We encourage you to take a look at our architectural principles. And of course, you have total freedom to propose or/and implement the improvements you want! Changes on the architecture, Introducing new services and/or containers, Reactive APIs, Distributed logging.... **Your creativity is more than welcome!**


# Appendix: Building, running and delivering

### 1. What we expect from your delivery:
- The system should be started with all the microservices using docker-compose. 
- The system should reproduce the behaviour described above. It should be able to handle multiple subscriptions to a sale and visualize every minute the email delivery sent to the winner.

### 2. What would you need?
The code requires the following tools:
- Maven 3.8.4
- JDK11
- A Docker container engine: docker desktop, podman, rancher or any other


### 3. How to build & run the project


| **Purpose**                                | **Command**                                          |
|--------------------------------------------|-------------------------------------------------------|
| Build jar files                            | ```mvn clean install```                                    |
| Build docker images & start the containers | ```docker-compose up -d```                                  |
| Cleanup project (docker compose)           | ```docker-compose down --rmi```                             |
| Cleanup project (podman)                   | ```docker-compose down```  <br><br> and delete the images manually |


### 4. Useful commands

| **Purpose** | **Command** | **Expected result** |
|-------------|--------------|---------------------|
| Test dummy endpoint in public-service from your local machine| ```curl localhost:8080/dummy -s``` |should return http status 200 and the message: <br>*“Hello, this is a dummy response from public service”* |
| Test adiclub-service endpoint from priority-sale-service container|```docker exec adidas-be-challenge-prioritysaleservice bash -c "curl -s adidas-be-challenge-adiclubservice:8080/adiclub?emailAddress=test1@gmail.com"```| should return a json response similar to: ```{"email":"test1@gmail.com",```<br>```"points":511,```<br>```"registrationDate":"2022-04-17T08:12:41.467026Z"```<br>```}```|
| Access Open API documentation|Access via browser to: <br> ```http://localhost:8080/swagger-ui/index.html#/```|doc. should be displayed|
| Watch public-service logs| ```docker logs -f adidas-be-challenge-publicservice``` | logs will be displayed|



