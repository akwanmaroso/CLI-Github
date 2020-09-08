# Overview
__clio__ is command line interface application for access Github API using golang.

# How to use ?
1. create clio.yaml
2. Add this code
	```
	---
	location: /home/akwanmaroso/Documents // path to your file
	token: mysecrettoken // token from github
	```
3. Run :
	```
	go install ./...
	clio --help
	```

# Some Feature
 - [x] Clone
 - [x] Fork
 - [x] Search
 - [x] Checkout
 - [x] Docs
 - [x] Remote 
 - [x] Pull Request
 - [x] Issue
 - [x] Repository(Create)
