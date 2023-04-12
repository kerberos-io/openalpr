FROM kerberos/openalpr-base:latest
MAINTAINER kerberos.io

############################
# Build Golang

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN ARCH=$(uname -m) && \
    ARCH=$([ "$(uname -m)" = "armv7l" ] && echo "armv6l" || echo $ARCH) && \
    ARCH=$([ "$(uname -m)" = "x86_64" ] && echo "amd64" || echo $ARCH) && \
    ARCH=$([ "$(uname -m)" = "aarch64" ] && echo "arm64" || echo $ARCH) && \
    wget "https://dl.google.com/go/go1.19.linux-$ARCH.tar.gz" && \
    tar -xvf "go1.19.linux-$ARCH.tar.gz" && \
    rm -rf go1.19.linux-$ARCH.tar.gz && \
    mv go /usr/local

# Copy OpenALPR go project
RUN mkdir -p $GOPATH/src/github.com/kerberos-io/openalpr
COPY . $GOPATH/src/github.com/kerberos-io/openalpr

# Workdir
WORKDIR $GOPATH/src/github.com/kerberos-io/openalpr

# Copy OpenALPR config
COPY openalpr.conf /etc/openalpr/openalpr.conf

# Move runtime data to /var/lib/openalpr
RUN mkdir -p /var/lib/openalpr && \
    mv runtime_data /var/lib/openalpr/runtime_data 

# Build OpenALPR
RUN cd $GOPATH/src/github.com/kerberos-io/openalpr && \
    go build -o /usr/local/bin/openalpr

# Entrypoint
ENTRYPOINT ["/usr/local/bin/openalpr"]
