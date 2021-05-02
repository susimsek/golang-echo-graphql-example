# Golang Echo Graphql Example 
> Golang Echo Graphql Example
>
<img src="https://github.com/susimsek/golang-echo-graphql-example/blob/main/images/golang-echo-graphql-example.png" alt="Golang Graphql Example Using Echo" width="100%" height="100%"/> 

## Prerequisites for Docker-Compose Deployment

* Golang 1.16.x
* Docker 19.03+
* Docker Compose 1.25+

## Installation for Docker-Compose Deployment

```sh
docker-compose up -d 
```

## Prerequisites for Kubernetes Deployment

* Kubernetes 1.12+
* Helm 3.1.0
* PV provisioner support in the underlying infrastructure

## Installation for Kubernetes Deployment

```sh
helm install app helm-chart/app
```

## Installation Using Vagrant for Docker-Compose Deployment

<img src="https://github.com/susimsek/golang-echo-graphql-example/blob/main/images/vagrant-docker-compose-installation.png" alt="Golang Vagrant Docker Compose Installation" width="100%" height="100%"/> 

### Prerequisites for Docker-Compose Deployment

* Vagrant 2.2+
* Virtualbox or Hyperv

```sh
vagrant up
```

```sh
vagrant ssh
```

```sh
cd vagrant/docker-compose-setup
```

```sh
sudo chmod u+x *.sh
```

```sh
./install-prereqs.sh
```

```sh
exit
```

```sh
vagrant ssh
```

```sh
docker-compose up -d
```

You can access the Playground from the following url.

http://192.168.12.11:9000/playground

## Installation Using Vagrant for Kubernetes Deployment

<img src="https://github.com/susimsek/golang-echo-graphql-example/blob/main/images/vagrant-k8s-installation.png" alt="Golang Vagrant Kubernetes Installation" width="100%" height="100%"/> 

### Prerequisites for Kubernetes Deployment

* Vagrant 2.2+
* Virtualbox or Hyperv

```sh
vagrant up
```

```sh
vagrant ssh
```

```sh
cd vagrant/kubernetes-setup
```

```sh
sudo chmod u+x *.sh
```

```sh
./install-prereqs.sh
```

```sh
exit
```

```sh
vagrant ssh
```

```sh
helm install app helm-chart/app
```

You can access the Playground from the following url.

http://localhost:9000/playground

## Used Technologies

* Golang 1.16.3
* Echo
* Query & Mutation & Subscription Example
* Gqlgen
* Gqlparser
* Go Uuid
* Air

## Golang Playground

> You can access the Golang Playground from the following url.

http://app.info/playground

<img src="https://github.com/susimsek/golang-echo-graphql-example/blob/main/images/go-playground.png" alt="Golang Playground" width="100%" height="100%"/> 

## Development

Please note that it requires Go 1.16+ since I use `go mod` to manage dependencies.

```bash

# 1. install dependencies
go mod download

# 2. For less typing, you could add alias air='~/.air' to your .bashrc or .zshrc.
alias air='~/.air'

# 3. run it
air
```


