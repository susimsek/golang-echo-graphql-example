
pipeline {
    agent any
    environment {
        IMAGE_REGISTRY = 'suayb/golang-echo-graphql-example'
        IMAGE_VERSION = '1.0.0'
        IMAGE_REGISTRY_CREDENTIAL = 'dockerhubcreds'
    }
    stages {
        stage('Build') {
           agent {
             docker {
               image 'golang:1.16-buster'
               args '-v /go/pkg/mod:/go/pkg/mod'
             }
           }
           steps {
             sh """
                ln -s `pwd` /go/src/app
                cd /go/src/app
                go mod download
                go build -v -o server
                 """
           }
        }
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