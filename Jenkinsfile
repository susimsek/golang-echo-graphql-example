pipeline {
    environment {
        BRANCH_NAME = "${env.GIT_BRANCH.split("/")[1]}"
        DEPLOY = "${BRANCH_NAME == "main" || BRANCH_NAME == "develop" ? "true" : "false"}"
        NAME = "app"
        VERSION = '1.0.0'
        DOMAIN = 'localhost'
        REGISTRY = 'suayb/golang-echo-graphql-example'
        REGISTRY_CREDENTIAL = 'dockerhubcreds'
    }
    agent {
        kubernetes {
            defaultContainer 'jnlp'
            yamlFile 'build.yaml'
        }
    }
    stages {
        stage('Docker Build') {
             when {
                environment name: 'DEPLOY', value: 'true'
             }
            steps {
                container('docker') {
                    sh "docker build -t ${REGISTRY}:${VERSION} ."
                }
            }
        }
        stage('Docker Publish') {
            when {
                environment name: 'DEPLOY', value: 'true'
            }
            steps {
                container('docker') {
                    withDockerRegistry([credentialsId: "${REGISTRY_CREDENTIAL}", url: ""]) {
                        sh "docker push ${REGISTRY}:${VERSION}"
                    }
                }
            }
        }
        stage('Kubernetes Deploy') {
              when {
                    environment name: 'DEPLOY', value: 'true'
              }
              steps {
                  container('helm') {
                          sh "helm upgrade --install --force --set name=${NAME} --set golangapp.image.tag=${VERSION} ${NAME} ./helm-chart/app"
                      }
              }
        }
    }
}