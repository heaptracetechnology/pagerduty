# Pagerduty
An OMG service to create, get and get list of incidents on PagerDuty

## Installation
Docker build and run

### Usage
PagerDuty is a platform for agile incident management, not a monitoring system. Think of PagerDuty as an add-on to trigger the right actions to all the data from your existing monitoring tools.

### OMG cli

##### Create Incident
* omg exec createincident -a from=<*FROM_EMAIL*> -a incident=<*CREATE_INCIDENT_OBJECT*> -e access_token=<*ACCESS_TOKEN*>
##### Get Incident
* omg exec getincident -a id=<*INCIDENT_ID*> -e access_token=<*ACCESS_TOKEN*>
##### Get List of Incidents
* omg exec listincidents -a <*ARGUMENTS*> -e access_token=<*ACCESS_TOKEN*>

### Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

### License
[MIT](https://choosealicense.com/licenses/mit/)
