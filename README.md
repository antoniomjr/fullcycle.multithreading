
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
![Screenshot 2024-02-20 at 23 01 45](https://github.com/antoniomjr/fullcycle/assets/53837075/5cbf6a88-f229-4cf6-910e-f83c496175a5)


> [!WARNING]  
> 
>                



