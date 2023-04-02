

## Upload file from blog text editor

Request:
```
curl --location 'https://localhost:8080/api/v1/files/post/dkjfkjdbkj12ye872362' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0NzkwNjksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.7O9NHabJqTHowQorUZ2ToSjkkgyGg3RljjLm07RBHf4' \
--form 'file=@"/home/dave/Pictures/20211007_180313_62.jpg"'
```

Response
```
"uploaded"
```

## Get file from blog text editor

Request:
```
curl --location 'https://localhost:8080/api/v1/files/post/anmpthjerfiletest/20211007_180313_62.jpg' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0NzkwNjksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.7O9NHabJqTHowQorUZ2ToSjkkgyGg3RljjLm07RBHf4'
```

Response
```
binary
```

## Delete a file
Request
```
curl --location --request DELETE 'https://localhost:8080/api/v1/files/post/anmpthjerfiletest/20211007_180313_62.jpg' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0NzkwNjksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.7O9NHabJqTHowQorUZ2ToSjkkgyGg3RljjLm07RBHf4'
```

Response
```
{
    "message": "successfully deleted",
    "status": 200
}
```