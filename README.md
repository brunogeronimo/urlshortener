# URL Shortener

This project contains a URL shortener written in [Go](https://go.dev).

## Basics

The main goal of it is to be able to have a URL Shortener service without the need of a database running on the 
background. All URLs are parsed from a static JSON file ([example](config.json.example)) and remain in memory until the
app execution is finished.

If the config file is changed, the app has to be restarted _(for now)_. If you run it on Google App Engine on a free 
tier, there is no need for restarting, since App Engine kills all instances if they are not used for a while.

## Running the project

The following env variables have to be defined:

| Name              | Description                                              | Required                            | Default Value |
|-------------------|----------------------------------------------------------|-------------------------------------|---------------|
| `CONFIG_URL`      | Contains the config JSON file with all settings          | yes                                 |               |
| `PORT`            | Port to run the application                              | no                                  | `8080`        |
| `CHECK_SIGNATURE` | Checks if URLs have been signed with a valid private key | no                                  | `false`       |
| `PUBLIC_KEY`      | Public key to verify URL signature                       | only if `CHECK_SIGNATURE` is `true` |               |

## Ways to run the project

### Docker

Just pull the desired image from [Docker Hub](https://hub.docker.com/repository/docker/brunogeronimo/urlshortener).

```sh
docker pull -eCONFIG_URL=<url> -p=8080:8080 brunogeronimo/urlshortener:latest
```
