# Google Cloud Platform

## Configuration

  * [Kubernetes Engine: `kubectl config`](https://medium.com/google-cloud/b6270d2b656c)

```bash
brew cask install google-cloud-sdk
# authenteficate
gcloud auth login
# what config is used atm
cat ~/.config/gcloud/active_config
# what configs we have?
gcloud config configurations list
# create new config
gcloud config configurations create <NAME>
# switch to other type of configuration
gcloud config configurations activate default
# just in case if you want to delete
gcloud config configurations delete <NAME>
```


## Projects

```bash
# list projects
gcloud projects list
# set current project id
gcloud config set project <PROJECT_ID>
# list current project info
gcloud config list project
# delete project
gcloud projects delete <PROJECT_ID>
# changed my mind
gcloud projects undelete <PROJECT_ID>

# formating and filtering
gcloud projects list --format='value(projectId)'
# is
gcloud projects list --format='value(projectId)' --filter='projectId ~ go-app-eng'
# regexes
gcloud projects list --format='value(projectId)' --filter='projectId ~ go-app-eng'

# export as json
gcloud projects list --format='json'
# with jq pipe
gcloud projects list --format='json' | jq '.[] | select (.projectId | test("go") ) | .projectId '
```


## Kubernetes

  * [Kubernetes Engine: `kubectl config`](https://medium.com/google-cloud/b6270d2b656c)

```bash
# default project will be used?
gcloud container clusters create kubecluster
# within project
gcloud container clusters create kubecluster --project=<PROJECT_ID>
# updates context of the kubectl
gcloud container clusters get-credentials kubecluster

# drop cluster
gcloud container clusters delete --project=<PROJECT_ID>

```

### Zones & Regions

```bash
# zone and region
# https://cloud.google.com/compute/docs/regions-zones/

# region europe-west3 (Frankfurt)
# zones a,b,c

gcloud config set compute/zone europe-west3-b
gcloud config set compute/region europe-west3
```

### Google App Engine

 - https://cloud.google.com/appengine/
 - https://github.com/GoogleCloudPlatform/golang-samples/tree/master/appengine
 - https://cloud.google.com/appengine/docs/go/config/appconfig


### Quick Start

`app.go`

```go
package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, wowld!")
}
```


`app.yaml`
```yaml
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app

```



#### Local Dev Running
```bash
# running (localy)
dev_appserver.py --port=9999 app.yml

# sendmail properties
dev_appserver.py --smtp_host=smtp.example.com --smtp_port=25 \
    --smtp_user=USERNAME --smtp_password=PASSWORD [PATH_TO_YOUR_APP]

# checking logs
gcloud app logs tail -s default

# open url in browser
gcloud app browse
```



#### Runtime Detection

Using `SERVER_SOFTWARE` env var.

```bash
if [[ "${SERVER_SOFTWARE}" == "Google App Engine/*" ]]; then
  echo "You are in GCP now"
fi
```


## Cloud Functions

In case if it's a http triggered function, its simple `HandlerFunc` (see example).

```bash
# enable service
gcloud services enable cloudfunctions.googleapis.com

# hello world code.
curl -LO https://github.com/GoogleCloudPlatform/golang-samples/archive/master.zip
unzip master.zip
cd golang-samples-master/functions/codelabs/gopher

# deploy function
gcloud functions deploy HelloWorld --runtime go111 --trigger-http
```


### Services

```
# list enabled services
gcloud services list

# enable service
gcloud services enable cloudfunctions.googleapis.com
```
