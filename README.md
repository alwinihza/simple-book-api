# simple-book-api

# Build Local
## Build with Docker
- Build Docker
```sh
docker build -t simple-golang .
```
- Run the docker
```sh
docker run --name livecode-golang-alwin -p <API PORT>:8000 -e DB_HOST=<DB HOST> -e DB_PORT=<DB_PORT> -e DB_NAME=<DB_NAME> -e DB_USER=<DB_USER> -e DB_PASSWORD=DB_PASSWORD -d simple-golang
```
This will expose API to <API PORT> 
## Build with docker compose
- Run docker compose
```sh
DB_HOST=${DB_HOST} DB_PORT=${DB_PORT} DB_NAME=${DB_NAME} DB_USER=${DB_USER} DB_PASSWORD=${DB_PASSWORD} API_PORT=${API_PORT} ${DOCKER_APP} compose up -d
```

# INTEGRATION STEPS
1. Create golang project (you can see in main.go)
2. Create Dockerfile
3. Test Dockerfile Run

![Test Dockerfile 1](./assets/docker-build.png)

![Test Dockerfile 2](./assets/docker-build2.png)

![Docker Images](./assets/docker-images.png)

![Docker Run](./assets/docker-run.png)

4. Create docker-compose.yml
5. Test docker-compose.yml

![Docker compose Up](./assets/docker-compose-up.png)

6. Create Jenkinsfile
7. Create github
8. Configure webhook

![Webhook setting](./assets/webhook%20setting.png)

9. Configure Jenkins' Job with Pipeline Configuration

![COnfigure 1](./assets/configure-1.png)

![COnfigure 2](./assets/configure-2.png)

![COnfigure 3](./assets/configure-3.png)


# Result
## Local Run
### Local Screenshot

![LOcal API](./assets/local-API.png)

![LOcal API List](./assets/local-API-List.png)

![LOcal API Post](./assets/local-API-Post.png)

## Server Run
### Server Screenshot

![Server List](./assets/server%20list.png)

![Server Post](./assets/server%20post.png)

## Notification
### EMail Notif

![Email Notif](./assets/email%20notif.png)

### Slack NOtif

![Slack NOtif](./assets/slack%20notif.png)