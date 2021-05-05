
pipeline {
    environment {
        BRANCH_NAME = "${env.GIT_BRANCH.split("/")[1]}"
        DEPLOY = "${BRANCH_NAME} == 'main' ? 'true' : 'false'"
        IMAGE_REGISTRY = 'suayb/golang-echo-graphql-example'
        IMAGE_VERSION = '1.0.0'
        IMAGE_REGISTRY_CREDENTIAL = 'dockerhubcreds'
    }
    agent none
    stages {
        stage('Docker Build') {
             when {
                environment name: 'DEPLOY', value: 'true'
             }
            steps {
                   sh "docker build -t ${IMAGE_REGISTRY}:${IMAGE_VERSION} ."
            }
        }
        stage('Docker Publish') {
            when {
                environment name: 'DEPLOY', value: 'true'
            }
            steps {
                    withDockerRegistry([credentialsId: "${IMAGE_REGISTRY_CREDENTIAL}", url: ""]) {
                        sh "docker push ${IMAGE_REGISTRY}:${IMAGE_VERSION}"
                    }
            }
        }
    }
}