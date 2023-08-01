Proof of work 
=================

## Summary

A simple web server that can send POW tasks and validate them. 
A simple web client, that calling the server every 5 seconds for a task and trying to resolve it. 

## Start service
  ### Docker
  In order to srart the service, a `Docker` should be installed. 

 How to install `Docker` 
  ``` 
  https://docs.docker.com/get-docker/
```
  If `Docker` is installed, please, run command,s from service root directory
  ```
  docker-compose -f docker/docker-compose.yaml up
``` 

  This command will run 2 docker containers on a local computer. 
  - container with `Server` (service that simulates vehicles movement)
  - container with `Worker` (expose port :8080)

  While both are working, can read chain using `GET http://localhost:8080/chain` endpoint (port is set to :8080 by default config)

