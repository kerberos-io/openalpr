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
