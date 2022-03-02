# terraform-provider-greenhouse

Remember to set GOPROXY (so Go will look at Artifactory) and GOPRIVATE (so it won't look for checksums for internal packages).
```
GOPROXY=https://<user>:<token>@crl.jfrog.io/artifactory/api/go/go
GOPRIVATE=github.com/carnegierobotics/*
```
