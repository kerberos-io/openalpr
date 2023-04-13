# OpenALPR with Kerberos Vault

This repository was originally created for illustrating the Golang binding delivered by the [Open Source License Plate Recognition project](https://github.com/openalpr/openalpr). Currently there are several open issues ([#831](https://github.com/openalpr/openalpr/issues/831), [#661](https://github.com/openalpr/openalpr/issues/661)) on the GitHub project asking for support, but unfortunately the project looks deprecated.

Within this repository we will be showing how you can build and containerize the OpenALPR Golang binding, and to integrate it with the [Kerberos Vault](https://github.com/kerberos-io/vault) solution. The goal of this project is that you get a good idea of how you can use OpenALPR (or other algorithms such as [YOLOv3, YOLOv7](https://github.com/kerberos-io/hub-objecttracker), etc.) in a production-ready video management deployment.

## Manage OpenALPR dependencies through Docker

As state in the original repostory it is defined in the `Dockerfile` that [several dependencies are required](https://github.com/openalpr/openalpr/blob/master/Dockerfile#L10-L13), and this is exactly where most people are experiencing issues trying to setup their local environment.

To overcome this issues, and avoid people forcing to install dependencies such as:

- `libleptonica-dev`
- `liblog4cplus-dev`
- `libopencv-dev`
- `libtesseract-dev`

We will build and run our project in a Docker container, that will allow us to run it on whatever host system we want, and avoids us installing another load of dependencies which you'll probably never use again.

The [`Dockerfile`](https://github.com/kerberos-io/openalpr/blob/main/Dockerfile) in this project illustrates how to pull a pre-build OpenALPR base image, install the Go toolchain, import the source code from our repository, build it and provide the `main.go` as the `ENTRYPOINT` of the Docker container.

## Base image

The [`Dockerfile`](https://github.com/kerberos-io/openalpr/blob/main/Dockerfile) in this project inherits from the `kerberos/openalpr-base:latest` base image. This image is a build on top of a customised [OpenALPR Dockerfile](https://github.com/kerberos-io/openalpr-base/blob/master/Dockerfile), which you can find [in our forked GitHub repository](https://github.com/kerberos-io/openalpr-base).

## How to run?

By building the `Dockerfile`, your Go code will be copied into the Docker image, and build to an executable.

    docker build -t kerberos/openalpr .

The output of this command is a `Docker` container with an `ENTRYPOINT` to your Go project (`main.go`). By executing `docker run` you can execute your source code.

    docker run -it kerberos/openalpr

You would expect to see some results, as it will try to recognise `car.png` and `car2.png` in this repository.

    openalpr % docker run -it kerberos/openalpr

    2.3.0
    {"version":2,"data_type":"alpr_results","epoch_time":1681366099708,"img_width":1060,"img_height":908,"processing_time_ms":58.521626,"regions_of_interest":[{"x":0,"y":0,"width":1060,"height":908}],"results":[{"plate":"JE123","confidence":80.169815,"matches_template":0,"plate_index":0,"region":"","region_confidence":0,"processing_time_ms":10.548375,"requested_topn":20,"coordinates":[{"x":458,"y":602},{"x":635,"y":601},{"x":631,"y":637},{"x":457,"y":639}],"candidates":[{"plate":"JE123","confidence":80.169815,"matches_template":0},{"plate":"IJE123","confidence":76.161102,"matches_template":0},{"plate":"1JE123","confidence":82.874680,"matches_template":0},{"plate":"1JVE123","confide...
