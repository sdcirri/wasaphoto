![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white) ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E) ![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D) ![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
# WASAPhoto
![wasa-screenshot](https://github.com/user-attachments/assets/f2860f88-6ccc-47b3-bc59-328d06eb585c)
Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can
upload your photos directly from your PC, and they will be visible to everyone following you.

This project was realized for the [Web And Software Architecture course](http://gamificationlab.uniroma1.it/en/wasa/) at Sapienza, based on the [Fantastic Coffee (decaffeinated)](https://github.com/sapienzaapps/fantastic-coffee-decaffeinated)
 template.
# How to run
### Running in debug mode
#### Backend
```shell
    $ go run ./cmd/webapi &
```
#### Frontend
```shell
    $ ./open-npm.sh
    $ npm run dev
```
### Running in docker
#### docker compose
```shell
    $ docker compose up -d
```
#### Backend only
```shell
    $ docker build -t wasaphoto-backend:latest -f Dockerfile.backend .
    $ docker run -t --rm -p 3000:3000 -v ./demo/db:/app/db wasaphoto-backend:latest
```
#### Frontend only
```shell
    $ docker build -t wasaphoto-frontend:latest -f Dockerfile.frontend .
    $ docker run -t --rm -p <PORT>:80 wasaphoto-frontend:latest
```
Your instance of WASAPhoto will be reachable through http://localhost:PORT/
By default, the docker-compose will expose port 8080

