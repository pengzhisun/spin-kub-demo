FROM crproxy.trafficmanager.net:5000/library/golang

ADD . /go/src/github.com/azure-devops/k8s-demo

RUN go install github.com/azure-devops/k8s-demo

ADD ./content /content

ENTRYPOINT /go/bin/k8s-demo
