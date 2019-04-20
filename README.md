# spy-api

Server for monitoring of mobile phones' moves.

## How to Run Service

1. Make sure `docker` with `docker-compose` is installed.
2. Build service with `docker-compose build`
3. Execute `docker-compose up` to start service on address `http://<host>:8080/ourell`.

## How To Run Tests

Unit tests are inside docker container and run automatically during container build.
Manual curl tests can be found inside `test` folder.

## API Description

Server has online documentation that is accessible on `http://<host>:8080/ourell/docs`.
Documentation in the Open API format can be found in the file `./api/spec.yaml`.

### POST /bbinput

Accepts GPS coordinates from the mobile and saves them to the database

#### Parameters

All parameters located in **body**.

Content-Type: **application/json**.

Parameters:

| Name        | Required | Type   | Description|
| ----------- | -------- | ------ | ---------- |
| number      | Yes      | string | Phone      |
| ip          | No       | string | IP address |
| imei        | Yes      | string | IMEI       |
| timestamp   | No       | string | EET timestamp in "YYYY/MM/DD-hh:mm:ss" format |
| coordinates | Yes      | object | GPS coordinates of the phone's location       |

**coordinates**:

| Name      | Required | Type   | Description          |
| --------- | -------- | ------ | -------------------- |
| longitude | Yes      | number | Longitude in degrees |
| latitude  | Yes      | number | Latitude in degrees  |

#### Body example

```
{
    "number": "+380991926482",
    "ip": "35.25.21.123",
    "imei": "502507345219189",
    "timestamp": "2019/03/22-15:50:20",
    "coordinates": {
        "longitude": 22.1832284135991,
        "latitude": 60.4538416572538
    }
}
```

#### Responses

| Code | Description           | Schema          |
| ---- | --------------------- | ----------------|
| 200  | OK. Coordinates saved |                 |
| 400  | Invalid arguments     | [Error](#error) |
| 500  | General server error  | [Error](#error) |

**Error**:

| Name    | Type   | Required |
| ------- | ------ | -------- |
| message | string | Yes      |

### POST /bbs

Shows how much time two phones are located in the same room

#### Parameters

All parameters located in **body**.

Content-Type: **application/json**.

Parameters:

| Name        | Required | Type   | Description|
| ----------- | -------- | ------ | ---------- |
| number1     | Yes      | string | Phone of user 1      |
| number2     | Yes      | string | Phone of user 2      |
| from        | Yes      | string | EET timestamp in "YYYY/MM/DD-hh:mm:ss" format |
| to          | Yes      | string | EET timestamp in "YYYY/MM/DD-hh:mm:ss" format |
| minDistance | Yes      | numeric | Distance in meters between users      |

#### Responses

| Code | Description           | Schema          |
| ---- | --------------------- | ----------------|
| 200  | OK.                   | [BbsSchema](#bbsSchema) |
| 400  | Invalid arguments     | [Error](#error) |
| 500  | General server error  | [Error](#error) |

**BbsSchema**:

| Name    | Type   | Required |
| ------- | ------ | -------- |
| percentage | number | Yes      |

### POST /bbfastDrive

Returns phone numbers of fast drivers

#### Parameters

All parameters located in **body**.

Content-Type: **application/json**.

Paramters:

See online documentation.

#### Responses

| Code | Description           | Schema          |
| ---- | --------------------- | ----------------|
| 200  | OK.                   | [BbsSchema](#bbsSchema) |
| 400  | Invalid arguments     | [Error](#error) |
| 500  | General server error  | [Error](#error) |
