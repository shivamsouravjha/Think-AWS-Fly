# Think-AWS-Fly

**Tech Used**
* Golang 
* AWS S3 Bucket

**Curls**
* curl --location --request GET 'localhost:3001/api/v1/fetch/buckets' \
--header 'Cookie: Cookie_1=value' \
--form 'image=@"/Users/shivamsouravjha/Downloads/1638782968982.jpeg"'
* curl --location --request POST 'localhost:3001/api/v1/upload/image' \
--header 'Cookie: Cookie_1=value' \
--form 'image=@"/Users/shivamsouravjha/Downloads/1638782968982.jpeg"'