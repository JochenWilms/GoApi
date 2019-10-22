# GoApi
APi for testing Go language. To start the project, first install the required dependencies.

##Package
Hi, this folder can be used for the project dependencies.
Set your GOPATH and GOBIN to this folder to import the go dependencies in this folder.
```
export GOPATH=$PWD/packages
mkdir $GOPATH/bin
export GOBIN=$GOPATH/bin
```
Please don't push all the dependencies Repo's to our gitRepo.
At the moment this seems a bad idea too me.

To install all the dependencies you can simply do:
```
go get ./
```
Now you can run this project localy. 
```
go run main.go
```
##Docker
There is a docker file to create a docker image from this go API.
```
docker build . -t 'IMAGE_NAME'
```

##Docker Compose
The docker Compose file will boot up a configured MongoDb and the created goapi image.
Please configure the name of the image to the previous docker build image.



Have fun.