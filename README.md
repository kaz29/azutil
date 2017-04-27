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

