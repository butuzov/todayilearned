# Google Cloud Platform

* https://github.com/dennyzhang/cheatsheet-gcp-A4/
* https://github.com/gregsramblings/google-cloud-4-words

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

## General

```bash
gcloud version
gcloud info
gcloud components list
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


## GKE (Google Kubernetes Engine)

  * [Kubernetes Engine: `kubectl config`](https://medium.com/google-cloud/b6270d2b656c)

```bash
# Display a list of credentialed accounts
gcloud auth list
# Set the active account
gcloud config set account <ACCOUNT>
# Set kubectl context
gcloud container clusters get-credentials <cluster-name>
# Change region
gcloud config set compute/region us-west
# Change zone
gcloud config set compute/zone us-west1-b
# List all container clusters
gcloud container clusters list

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


# Buckets
```bash
# List all buckets and files
gsutil ls
gsutil ls -lh gs://<bucket-name>
# Download file
gsutil cp gs://<bucket-name>/<dir-path>/package-1.1.tgz .
# Upload file
gsutil cp <filename> gs://<bucket-name>/<directory>/
# Cat file
gsutil cat gs://<bucket-name>/<filepath>/
# Delete file
gsutil rm gs://<bucket-name>/<filepath>
# Move file
gsutil mv <src-filepath> gs://<bucket-name>/<directory>/<dest-filepath>
# Copy folder
gsutil cp -r ./conf gs://<bucket-name>/
# Show disk usage
gsutil du -h gs://<bucket-name/<directory>
# Create bucket
gsutil mb gs://<bucket-name>
# Caculate file sha1sum
gsha1sum syslog-migration-10.0.2.tgz
shasum syslog-migration-10.0.2.tgz
# Gsutil help
gsutil help
gsutil help cp
gsutil help options

# Security
# Make all files readable
gsutil -m acl set -R -a public-read gs://<bucket-name>/
# Config auth
gsutil config -a
# Grant bucket access
gsutil iam ch user:denny@gmail.com:objectCreator,objectViewer gs://<bucket-name>
# Remove bucket access
gsutil iam ch -d user:denny@gmail.com:objectCreator,objectViewer gs://<bucket-name>
```
