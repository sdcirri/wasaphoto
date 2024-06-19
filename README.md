# WASAPhoto
Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can
upload your photos directly from your PC, and they will be visible to everyone following you.

# How to run
### Running in debug mode
#### Backend
```shell
    $ go run ./cmd/webapi &
```
#### Frontend
```shell
    $ cd webui
    $ npm run dev &
```
### Running in docker
#### docker compose
```shell
    $ docker compose up -d
```
#### Backend only
```shell
    $ docker build -t wasaphoto-webapi:latest -f Dockerfile.backend .
    $ docker run -t --rm -p 3000:3000 -v ./demo/db:/srv/wasaphoto wasaphoto-webapi:latest
```
#### Frontend only
```shell
    $ docker build -t wasaphoto-webui:latest -f Dockerfile.frontend .
    $ docker run -t --rm -p <PORT>:80 wasaphoto-webui:latest
```
Your instance of WASAPhoto will be reachable through http://localhost:PORT/
By default, the docker-compose will expose port 8080

