pipeline {
    environment {
        BRANCH_NAME = "${env.GIT_BRANCH.split("/")[1]}"
        DEPLOY = "${BRANCH_NAME == "main" || BRANCH_NAME == "develop" ? "true" : "false"}"
        CHART_NAME = "app"
        CHART_NAMESPACE= "app"
        CHART_DIR = "helm-chart/app"
        IMAGE_REGISTRY = 'suayb/golang-echo-graphql-example'
        IMAGE_VERSION = '1.0.4'
        IMAGE_REGISTRY_CREDENTIAL = 'dockerhubcreds'
    }
    agent {
        kubernetes {
            defaultContainer 'jnlp'
            yamlFile 'build.yaml'
        }
    }
    stages {
        stage('Kubernetes Version Control') {
             when {
                environment name: 'DEPLOY', value: 'true'
             }
            steps {
                container('kubectl') {
                    sh "kubectl version"
                }
            }
        }
        stage('Docker Build') {
             when {
                environment name: 'DEPLOY', value: 'true'
             }
            steps {
                container('docker') {
                    sh "docker build -t ${IMAGE_REGISTRY}:${IMAGE_VERSION} ."
                }
            }
        }
        stage('Docker Publish') {
            when {
                environment name: 'DEPLOY', value: 'true'
            }
            steps {
                container('docker') {
                    withDockerRegistry([credentialsId: "${IMAGE_REGISTRY_CREDENTIAL}", url: ""]) {
                        sh "docker push ${IMAGE_REGISTRY}:${IMAGE_VERSION}"
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
                          sh "helm upgrade --install ${CHART_NAME} ${CHART_DIR} --namespace ${CHART_NAMESPACE}"
                      }
              }
        }
    }
}