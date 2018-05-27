### terraform-provider-weather
> a custom terraform provider

Compile 
```bash
export GOOS="linux" GOARCH="amd64" 
go build -o terraform-provider-weather
```

Run 
```
terraform init
terraform plan
terraform show
terraform apply --auto-approve
terraform destroy
```