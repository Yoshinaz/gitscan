# github repository scan application

This is an application for detecting the private_key / public_key in github repository

`run-tests.sh` runs a simplistic test and generates the API
documentation below.


## Pre required

    make sure that the docker application is started


## Run the tests

    ./run-tests.sh

# REST API

The REST API to the repository scan app is described below.

## Scan repository

### Request

`Post /v1/repo/scan`

    curl -k -d {\"name\":\"test\",\"url\":\"https://github.com/Yoshinaz/test_secret\",\"rules_set\":\"any\",\"all_commit\":\"false\"} http://localhost:8080/v1/repo/scan
### Response

    {
        "status": "QUEUED"
    }

## View Scan result

### Request

`GET /thing/id`

    curl -k -d {\"name\":\"test\",\"url\":\"https://github.com/Yoshinaz/test_secret\",\"rules_set\":\"any\",\"all_commit\":\"false\"} http://localhost:8080/v1/repo/view

### Response

    {
        "Name": "test_repo",
        "URL": "https://github.com/Yoshinaz/test_secret",
        "Status": "SUCCESS",
        "Description": "",
        "EnqueuedAt": "2023-02-26T21:27:01+07:00",
        "StartedAt": "2023-02-26T21:28:23+07:00",
        "FinishedAt": "2023-02-26T21:28:23+07:00",
        "CreatedAt": "2023-02-26T21:27:01+07:00",
        "Findings": [
            {
                "Type": "sast",
                "RuleId": "G401",
                "Location": [
                    {
                        "Path": "test/test2.go",
                        "Positions": {
                            "Begin": [
                                {
                                    "Line": "5"
                                }
                            ]
                        }
                    },
                    {
                        "Path": "test/test.go",
                        "Positions": {
                            "Begin": [
                                {
                                    "Line": "5"
                                },
                                {
                                    "Line": "7"
                                }
                            ]
                        }
                    },
                    {
                        "Path": "main.go",
                        "Positions": {
                            "Begin": [
                                {
                                    "Line": "5"
                                }
                            ]
                        }
                    }
                ],
                "Commit": "d675ba1d98ed0ef7c7d08a1c79ce390cd2c95464",
                "Metadata": {
                    "Description": "private / public key detected",
                    "Severity": "HIGH"
                }
            }
        ]
    }
