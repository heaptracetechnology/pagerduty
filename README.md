# Pagerduty as a microservice
An OMG service to create, get and get list of incidents and getservice, listservices on PagerDuty

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)

This microservice's goal is to create, get and get list of incidents on PagerDuty

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```

### CLI

##### Create Incident
```sh
$ omg exec createincident -a from=<*FROM_EMAIL*> -a incident=<*CREATE_INCIDENT_OBJECT*> -e access_token=<*ACCESS_TOKEN*>
```
##### Get Incident
```sh
$ omg exec getincident -a id=<*INCIDENT_ID*> -e access_token=<*ACCESS_TOKEN*>
```
##### Get List of Incidents
```sh
$ omg exec listincidents -a <*ARGUMENTS*> -e access_token=<*ACCESS_TOKEN*>
```
##### Get list of Incidentnotes
```sh
$ omg exec listincidentnotes -a id=<*INCIDENT_ID*> -e access_token=<*ACCESS_TOKEN*>
```
##### Get Service
```sh
$ omg exec getservice -a id=<*SERVICE_ID*> -a serviceoptions=<*object*> -e access_token=<*ACCESS_TOKEN*>
```
##### Get List of Service
```sh
$ omg exec listservices -a serviceoptions=<*object*> -e access_token=<*ACCESS_TOKEN*>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build --rm -f "Dockerfile" -t pagerduty:latest .
```
### RUN
```
docker run -p 5000:5000 pagerduty:latest
```

### Usage
PagerDuty is a platform for agile incident management, not a monitoring system. Think of PagerDuty as an add-on to trigger the right actions to all the data from your existing monitoring tools.

### Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
