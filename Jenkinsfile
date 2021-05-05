
pipeline {
    environment {
        IMAGE_REGISTRY = 'suayb/golang-echo-graphql-example'
        IMAGE_VERSION = '1.0.0'
        IMAGE_REGISTRY_CREDENTIAL = 'dockerhubcreds'
    }
    stages {
        stage('Docker Build') {
            steps {
                   sh "docker build -t ${IMAGE_REGISTRY}:${IMAGE_VERSION} ."
            }
        }
        stage('Docker Publish') {
            steps {
                    withDockerRegistry([credentialsId: "${IMAGE_REGISTRY_CREDENTIAL}", url: ""]) {
                        sh "docker push ${IMAGE_REGISTRY}:${IMAGE_VERSION}"
                    }
            }
        }
    }
}