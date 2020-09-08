
## Overview
__clio__ is command line interface application for access Github API using golang.

## Requirements
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
## How To Use ?
1. Clone 		=> `clio clone owner/project`
2. Fork 		=> `clio fork owner/project`
3. Search 	=> `clio search topic:go`
4. Docs 		=> `clio docs owner/project`
5. Pull Request => `clio pullrequest -d owner/project:branch-dest -t "Title PR" -m "message PR" -s owner:branch-source"` 
6. Issue => `clio issue owner/project -t "Title Issue" -m "message issue"`
7. Create Repository => `clio repository owner/project -d "description repo" -p false`

# Some Feature
 - [x] Clone
 - [x] Fork
 - [x] Search
 - [x] Checkout
 - [x] Docs
 - [x] Pull Request
 - [x] Issue
 - [x] Repository(Create)
