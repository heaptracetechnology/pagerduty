FROM golang

RUN go get github.com/PagerDuty/go-pagerduty

WORKDIR /go/src/github.com/microserice/pagerduty

ADD . /go/src/github.com/microserice/pagerduty

RUN go install github.com/microserice/pagerduty

ENTRYPOINT pagerduty

EXPOSE 8000

