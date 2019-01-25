FROM golang

RUN go get github.com/PagerDuty/go-pagerduty
RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/microservice/pagerduty

ADD . /go/src/github.com/microservice/pagerduty

RUN go install github.com/microservice/pagerduty

ENTRYPOINT pagerduty

EXPOSE 5000
