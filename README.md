# blog-api
simple blog api with authentication built with golang

## Prerequisites

- You need to have Go installed on your computer. The version used for the tutorial is  **1.13**.

- Postgres database running either locally or remote.

## Get started

- Clone this repository to your filesystem.

- `cd` into the project directory, install dependencies and start project, with command "go run main.go.

## Documentation
[find here](https://documenter.getpostman.com/view/4823089/T17Gennu?version=latest)

## Endpoints

- POST Register /v1/register

- POST Login /v1/login

- GET Get all posts /v1/post

- GET Get post detail /v1/post/detail

- POST create post (authenticated user only) /v1/post

- PUT edit post (authenticated user only) /v1/post

- DELETE delete post (authenticated user only) /v1/post 
