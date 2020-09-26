## IMCARGO

Simple Image processing/storing service

## Deploying And Running
```
  docker-compose up --build
```

### API Documentation
### [POST] /upload
| Param   |      Description      |  Required  | Default Value | Type |
|----------|:-------------:|------:|------:|------:| 
| accessToken | Token to authorize | Yes | Available in Terminal | URL Param |
| size | Croped Image size  |  No  | 128x128 | URL Param |
| type | Crop Origin Point | No | default (centered) | URL Param |
| blur | Image blur value | No | 0 | URL Param |
| image | Form Data | Yes |  | Form-Data |
| Content-Type | Content type | Yes | multipart/form-data | Header |



### [GET] /static/avatar/{name}.jpg

## Example Request
```bash
http://localhost:8080/upload?size=800x800&type=centered&blur=0&accessToken=38fc3efe471ac435ed97a8668f53b8ef0ece2c721cee46652c8d810d6efda009
```

```json
{
    "success": true,
    "message": "Avatar was uploaded successfully",
    "data": {
        "size": {
            "width": 800,
            "height": 800
        },
        "path": "static/avatar/3bb7a3fe-a6fd-4696-8efd-01a3d05b9fee.jpg"
    }
}
```

### Todo:
  - [] Clean up the codebase
  - [] Make hmac auth stabile
  - [] Use .env file for the secrets
  - [] Multiple bucket support
  - [] Delete Support
  - [] RT Image manipulation EP
  - [] Store Entity in database (SQLte maybe)
  - [] Write tests
  - [X] Write THE DOCS!
  - [x] Docker support please



**Author**: David Kviloria
