## Build

Run `docker build -t repo_name .`. it would build the docker image, then we could

## Running Apps

```sh
docker container create --name image_id -e PORT=8080 -e INSTANCE_ID="mountain dew" -p 8080:8080 image_id
docker container start image_id
docker container run --name image_id --rm -itd -e PORT=8080 -e INSTANCE_ID="moiuntain dew" -p 8080:8080 image_id
docker container stop image_id
```

to know docker image_id run `docker image ls -a`
