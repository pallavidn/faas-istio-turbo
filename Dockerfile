# Set the base image

FROM alpine:3.3

# Set the file maintainer

MAINTAINER Pallavi Debnath <pallavi.debnath@turbonomic.com>


ADD _output/faas-istio-turbo.linux /bin/faasistioturbo


ENTRYPOINT ["/bin/faasistioturbo"]
