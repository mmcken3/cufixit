# Clemson University Fix It

This is a project for the CUHackIt Hackathon.

An application where people in Clemson can go to report any feedback they have about the university. They can see issues with facilities
like a broken light, broken water fountain and then report it to the University Facilities for maintaince through one single application.
They could also use it to give feedback to the univerity for things like making a building more accessible by attaching an image and description
about the issue and submitting it to the accessibility counsil for work. 

Currently the standard for submitting maintenance request to the University is long and tedious. The form is no easy to find online
and there is a long set of directions on how to use it. We hope that CUFixIt can centralize request into one mobile application
that is simple and easy to use. 

# Getting Started
We have developed a react-native mobile application for submitting the request to the university. Login to the add, upload or take a picture, add a description, select type of issue and description then you are ready to submit. Once someone has submitted a request the type of issue, building
where the issue is, image they chose, and their user name will be submitted to the database. The application will then determine based off of 
the building and type of issue who to contact and will email the contact for that issue type with all of the information.

# CUFixIt API

This api is accessible through Amazon ECS at http://[IP Address]:8002/. This was hosted on Amazon ECS for the CUHackIT but has since been taken down. You can always test running it locally using Docker.

HTTP request | Description
------------ | ------------- 
**GET** /    | returns a hello world to check on API connection |
**GET** /v1/getall    | returns a slice of all reuqest in the DB |
**POST** /v1/submit    | submits a request from the json body |
**GET** /v1/get/type/{type}   | returns a slice of all the request of type {type} |
**GET** /v1/get/building/{building}   | returns a slice of all the request of building {building} |
**GET** /v1/get/user/{user}   | returns a slice of all the request from user {user}|

# CUFixIt Database
We have created a database using Postgresql that is being hosted using Amazon RDS. This database can be accessed from your local terminal or application,
but it connected to our Amazon ECS instance so that the API can interact with the database.

# CUFixIt Image Store
We have created an image store using Amazon S3 in order to store all of the images coming from our users request. When they submit a valid form the 
image the selected or took is uploaded to S3 and the url or the location is stored into the DB.

# Getting Started For Development

## API Development
Start the docker containers by running `docker-compose build` and then `docker-compose up`.

## React Development
Start the react-native app by running `react-native run-ios`.

## Tools Used
Golang 1.9
Docker Container Image
Postgresql
Amazon Web Services
    RDS
    ECS
    EC2
    S3 
React Native
Node.js
