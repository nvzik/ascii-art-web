
# Ascii-art-web

A brief description of what this project does and who it's for


## Objectives

Any additional information goes here

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, ascii-art.

The webpage allows use of the different banners:

- shadow
- standard
- thinkertoy

The main page has:
- text input
- radio buttons
- button, which sends a POST request and outputs the result on the page.
- site more appealing, interactive and intuitive.
- site more user friendly and give more feedback to the user.

## HTTP status code

Your endpoints must return appropriate HTTP status codes.
- OK (200), if everything went without errors.
- Not Found, if nothing is found, for example templates or banners.
- Bad Request, for incorrect requests.
- Internal Server Error, for unhandled errors.


## Usage 

```bash
After cloning the repository, complete the following commands:

    - Run the programm
        go run .
        
      After that command click on the localhost link
      
    - Create a docker image with the docker file:
        
        docker build -t ascii-art-docker .

    - Run image and create container connected to port 8080:
        
        docker run -d --name=ascii-container -p8080:8080 ascii-art-docker

    - To check running containers:
        
        docker ps

    -Finally, connect to the localhost:8080 through your browser and start using the static website.

    For your simplicity you can use the make file =))
```


