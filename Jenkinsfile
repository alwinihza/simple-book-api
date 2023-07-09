pipeline {
    agent any
    environment {
        GIT_URL = 'https://github.com/alwinihza/simple-book-api.git'
        BRANCH = 'main'
        IMAGE = 'book-api'
        CONTAINER = 'my-book-api-app'
        DOCKER_APP = 'docker'
        DB_HOST = 'db'
        DB_USER = 'postgres'
        DB_NAME = 'book_management_system'
        DB_PASSWORD = 'password'
        DB_PORT = '5432'
        API_PORT = '8000'
    }
    stages {
        stage("Cleaning up") {
            steps {
                echo 'Cleaning up'
                sh "${DOCKER_APP} rm -f ${CONTAINER} || true"
                sh "docker compose down"
            }
        }

        stage("Clone") {
            steps {
                echo 'Clone'
                git branch: "${BRANCH}", url: "${GIT_URL}"
            }
        }

        stage("Build and Run") {
            steps {
                echo 'Build and Run'
                sh "DB_HOST=${DB_HOST} DB_PORT=${DB_PORT} DB_NAME=${DB_NAME} DB_USER=${DB_USER} DB_PASSWORD=${DB_PASSWORD} API_PORT=${API_PORT} ${DOCKER_APP} compose up -d"
            }
        }
    }
    post {
        always {
            emailext to: "alwinihza@gmail.com",
            subject: "Livecode Pipeline Notification",
            body: "Pipeline ${env.JOB_NAME} running with ${env.GIT_COMMIT}",
            attachLog: true
            slackSend botUser: true, 
            channel: 'general', 
            color: '#00ff00', 
            message: "Pipeline ${env.JOB_NAME} running with ${env.GIT_COMMIT}", 
            tokenCredentialId: 'Slack OAUTH'
        }
        success {
            echo 'This will run only if successful'
        }
        failure {
            echo 'This will run only if failed'
        }
    }
}