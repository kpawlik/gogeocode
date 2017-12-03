# gogeocode

## Instalation

```go get github.com/kpawlik/gogeocode```

```go install github.com/kpawlik/gogeocode```


## Run
```
gogeocode -apiKey <YOUR_API_KEY> -in <CSV_PATH> -out <CSV_PATH>
gogeocode -id <CLIENT_ID> -signature <CLIENT_SIGNATURE> -in <CSV_PATH> -out <CSV_PATH>
```
Example:
```
gogeocode -apiKey asdSDF234FGSDFGdsfg -in /tmp/address_in.csv -out /tmp/address_in.csv
gogeocode -id client_id -signature signature_value -in /tmp/address_in.csv -out /tmp/address_in.csv
```


Input data format:  
```ID,ADDRESS```

Output data format:  
```ID,ADDRESS,LAT,LNG```
