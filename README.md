# azutil
---

## Create vhd image


```
$ cd qemu
```

### Cusomize image format and size 

Customize environment `MKFS_TYPE` and `IMAGE_SIZE`.

```
$ vim docker-compose.yml
```

docker-compose.yml

``` 
qemu:
  build: .
  environment:
    MKFS_TYPE: ext4
    IMAGE_SIZE: 10G
  volumes:
    - .:/home/qemu
```

### Run Docker container.

```
$ docker-compose run qemu
$ ls -alh image.vhd
-rw-r--r--  1 user  staff   1.0G  4 27 12:58 image.vhd
```
## Create Storage Account and container

### Create a Service Principle yaml file

Create a yaml file for Storing Service Principle for accessing Azure.

Save the yaml file as `azureConfig.yml` for example.

```
AZURE_SUBSCRIPTION_ID: {YOUR_SUBSCRIPTION}
AZURE_TENANT_ID: {YOUR_TENANT_ID}
AZURE_CLIENT_ID: {YOUR_CLIENT_ID}
AZURE_CLIENT_SECRET: {YOUR_CLIENT_SECRET}
```

Then execute the azutil command.

```
azutil create container -r {Resource Group} -l {Location} -s {Storage Account Name} -c {Container Name} -f azureConfig
```

also, you can see the help using `azutil create container --help`

## Upload VHD to the container

```
azutil upload vhd -r {Resource Group} -s {Storage Account Name} -c {Container Name} -f azureConfig -v {VHD file name} 
```

also, you can see the help using `aztuil upload vhd --help`