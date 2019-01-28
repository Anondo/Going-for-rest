# GOREST

Gorest is an application written in golang which serves a blog related RESTful web service.

## Endpoints

1. ```GET```  ```{{base_url}}/v1/blogs``` ```List all the blogs```
1. ```GET```   ```{{base_url}}/v1/blogs/1``` ```List a single blog```
1. ```POST```  ```{{base_url}}/v1/blogs/``` ```Create a blog```
1. ```PUT``` ```{{base_url}}/v1/blogs/1``` ```Update all the fields of a blog```
1. ```PATCH``` ```{{base_url}}/v1/blogs/1``` ```Update certain fields of a blog```
1. ```DELETE``` ```{{base_url}}/v1/blogs/1``` ```Delete a single blog```


## Running The Application

### Step-1:
Go the terminal &

```
    docker-compose up
```
or
```
  docker-compose up -d
```
The latter will start the containers as a daemon

### Step-2:

Hit the url ```localhost:8500```

![config](https://github.com/Anondo/Going-for-rest/blob/master/screenshots/config.png)

Go to ```Key/Value``` , hit ```Create```, use ```gorest``` as ```key``` & then copy the
contents of the config.yml file from the code base and paste it as the value. Just like the
the given screenshot. Hit ```Save```. This will up the config.

### Step-3:

Next, in the terminal

```
  source run
```
This will both build the app & start the server. For database migrations

```
  ./gorest migration up
```
```
  ./gorest migration down
```
Run these.

Thats it. Good to go. Play with the endpoints given.

## Requirements
1. Docker /Docker Compose

Or if you want to run everything locally

1. Postgresql
1. Adminer
1. Consul
