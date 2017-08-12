# drone-plugin-gcr

## Usage

```
pipeline:
  publish:
    image: no0dles/drone-plugin-gcr
    privileged: true
    repo: project-id/image-name
    secrets: [ GOOGLE_TOKEN ]
```

## Parameters

| Name        | Required | Default    |
| ----------- |:--------:| ----------:|
| repo        | yes      | -          |
| tags        | no       | latest     |
| registry    | no       | gcr.io     |
| dockerfile  | no       | Dockerfile |
| buildpath   | no       | .          |
| cache_from  | no       | -          |

## Full Example

```
pipeline:
  publish:
    image: no0dles/drone-plugin-gcr
    privileged: true
    repo: project-id/image-name
    registry: eu.gcr.io
    dockerfile: Dockerfile.prod
    buildpath: /build
    cache_from: latest
    tags:
      - latest
      - ${DRONE_COMMIT}
    secrets: [ GOOGLE_TOKEN ]
```

## GOOGLE_TOKEN
Docs about how to get it [here](https://cloud.google.com/container-registry/docs/advanced-authentication) and [here](https://support.google.com/cloud/answer/6158849#serviceaccounts)

It should look like this
```
{
  "type": "service_account",
  "project_id": "your-project-id",
  "private_key_id": "abc",
  "private_key": "-----BEGIN PRIVATE KEY-----\nLOREM\n-----END PRIVATE KEY-----\n",
  "client_email": "email@your-project-id.iam.gserviceaccount.com",
  "client_id": "1234",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://accounts.google.com/o/oauth2/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/email%40your-project-id.iam.gserviceaccount.com"
}
```
