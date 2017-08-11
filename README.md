# drone-plugin-gcr

## Usage

```
pipeline:
  publish:
    image: no0dles/drone-plugin-gcr
    repo: project-id/image-name
    secrets: [ GOOGLE_TOKEN ]
```

## Parameters

| Name       | Required | Default    |
| ---------- |:--------:| ----------:|
| repo       | yes      | -          |
| tags       | no       | latest     |
| registry   | no       | gcr.io     |
| dockerfile | no       | Dockerfile |
| buildpath  | no       | .          |