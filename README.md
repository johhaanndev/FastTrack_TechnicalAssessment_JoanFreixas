# Technical Assessment - Quiz API
## Contents
- [Context](#context)
- [Implementation](#implementation)
- [Use](#use)

## Context
### Instructions 
The task is to build a super simple quiz with a few questions and a few alternatives for each question. Each with one correct answer. 
### Preferred Stack:
Backend - Golang.
Database - Just in-memory, so no database.
### Preferred Components:
REST API or gRPC.
CLI that talks with the API, preferably using https://github.com/spf13/cobra ;( as CLI framework ).
### User stories/Use cases: 
User should be able to get questions with a number of answers
User should be able to select just one answer per question.
User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
User should see how well they compared to others who have taken the quiz, eg. "You were better than 60% of all quizzers".

## Implementation
### Structure
- main.go: main file on top level of directories. Calls the handler functions and set the listener port to localhost:8080.
- src: directory of all the source code.
- models: object definitions.
- handlers: endpoint functions, process the input and generates the response.
- validations: validates the inputs, in this case the POST /answer request body.
- csvServices: functions to handler read and write csv for local database

### Database

*According to requirements, it shouldn't have any database, just in-memory.*
*A local database with csv with very low level of detail was created in order to make testing easier.*

A csv was created with two columns: id and score.
- Id: unique identifier for each player. Generated randomly when hitting POST /answers.
- Score: score calculated in % fo the correct answers.

The csv is only local, which means the data is not persistent among multiple machines.
The local database will be created the first time an user hit POST /answers with a valid request.
After the first request, a new entry with ID and the score will be added to the file.

*The repository contains an example of how the database looks when created*
*You can remove it before trying the API or you can keep it and start from this point.*


# Use
## CLI
- GET /questions: curl -i -X GET http://localhost:8080/questions
- POST /answers: curl -i -X POST http://localhost:8080/answers -H "Content-Type: application/json" -d "[ \"b\", \"c\", \"a\", \"b\", \"c\" ]"
## Postman
You can test on your postman by downloading the collection under [quiz-app/postman](https://github.com/johhaanndev/FastTrack_TechnicalAssessment_JoanFreixas/tree/master/quiz-app/postman), it contains the two endpoints ready to be run after executing the api.
- GET /questions:
  ![image](https://github.com/user-attachments/assets/4bba77f2-6db2-4f91-9d5a-14b5cc11ddcb)
- POST /answers:
  ![image](https://github.com/user-attachments/assets/77e02a26-712b-4e1b-8b1d-fe3e4d56d6e1)

## Results
- GET /questions: 200, 404, 500
![image](https://github.com/user-attachments/assets/af89dda5-22fc-4173-a5b4-6a9f6d706032)
- POST /answers: 200, 404, 400, 500
![image](https://github.com/user-attachments/assets/c88b8ade-e3ba-4174-9b5c-cd7761563478)
