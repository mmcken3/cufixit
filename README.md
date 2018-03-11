# Clemson University Fix It

An application where people in Clemson can go to report any feedback they have about the university. They can see issues with facilities
like a broken light, broken water fountain and then report it to the University Facilities for maintaince through one single application.
They could also use it to give feedback to the univerity for things like making a building more accessible by attaching an image and description
about the issue and submitting it to the accessibility counsil for work. 

This is a project for the CUHackIt Hackathon.

# Getting Started
We have developed a react-native mobile application for submitting the request to the university. Login to the add, upload or take a picture, add a description, select type of issue and description then you are ready to submit. 

# CUFixIt API

This api is accessible through Amazon ECS at http://54.211.84.167:8002/.

HTTP request | Description
------------ | ------------- 
**GET** /    | returns a hello world to check on API connection |
**GET** /v1/getall    | returns a slice of all reuqest in the DB |
**GET** /v1/submit    | submits a request from the json body |

# Getting Started For Development

## Mac
Download the [docker container engine](https://store.docker.com/editions/community/docker-ce-desktop-mac).

If you would like to ensure that your docker engine is working well enough to run the DB, visit this [site](https://docs.docker.com/docker-for-mac/) and go up to step 3. 

Install [GoLang](https://golang.org).
