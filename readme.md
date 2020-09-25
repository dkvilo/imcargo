## IMGROOT

Simple Image processing/storing service

### Usage

```
[POST]
http://localhost:8080/upload

size=800x800|800x0 ...

type=centered|default

blur=50

accessToken=

```

## Example Response
```json
{
    "success": true,
    "message": "Avatar was uploaded successfully",
    "data": {
        "size": {
            "width": 800,
            "height": 800
        },
        "path": "static/avatar/3bb7a3fe-a6fd-4696-8efd-01a3d05b9fee.webp"
    }
}
```

### Todo:
  - Clean up the codebase
  - Make hmac auth stabile
  - Use .env file for the secrets
  - Multiple bucket support
  - Delete Support
  - RT Image manipulation EP
  - Store Entity in database (SQLte maybe)
  - Write tests
  - Write THE DOCS!
  - Docker support please



**Author**: David Kviloria
