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

Have fun.

