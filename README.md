
# Go Challenge for FullCycle Pós Go Expert

## Overview

This repository contains a Go application developed as a part of the "`Desafios/multitrade`" (Go Challenge) from the Pós Go Expert.

## Challenge Requirements

In this challenge you will have to use what you learn with Multithreading and APIs to seek the fastest result between two different APIs.

The two requests will be made simultaneously to the following APIs:

https://brasilapi.com.br/api/cep/v1/cep

http://viacep.com.br/ws/" + cep + "/json/

The requirements for this challenge are:

- Find an API that delivers the fastest response and discards the slowest response.

- The result of the request should be displayed on the command line with the address data, as well as which API to envy.

- Limit response time to 1 second. Otherwise, the timeout error should be displayed.
  
## Project Structure

- `main.go`: The Go source file that contains the application and function.

## Instructions

To build and run this application, follow these steps:

#### Running the Go application
This will run the application
```bash
go run main.go cep
```
![Screenshot 2024-06-26 at 22 58 03](https://github.com/antoniomjr/fullcycle.multitrade/assets/53837075/9eab7413-4bec-40b6-aa77-8454c8187f4f)
![Screenshot 2024-06-26 at 22 58 31](https://github.com/antoniomjr/fullcycle.multitrade/assets/53837075/c86e14d4-caef-4619-93b3-23b4b84b252f)

> [!WARNING]  
> 
>                



