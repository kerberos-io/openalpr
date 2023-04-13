FROM kerberos/openalpr-base:latest

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

ENV LDFLAGS="-Wl,--copy-dt-needed-entries"
# Build OpenALPR
RUN cd $GOPATH/src/github.com/kerberos-io/openalpr && \
    go build -o /usr/local/bin/openalpr

# Entrypoint
ENTRYPOINT ["/usr/local/bin/openalpr"]
