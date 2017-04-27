#!/bin/bash

cd "/home/qemu"

## Create empty raw disk
qemu-img create -f raw image.raw "${IMAGE_SIZE}"

## Format as ext4
case "${MKFS_TYPE:-"ext4"}" in
	"ntfs")
		echo "formatting as ntfs"
		mkfs.ntfs -F ./image.raw
		;;
	"xfs")
		echo "formatting as xfs"
		mkfs.xfs ./image.raw
		;;
	"ext4")
		echo "formatting as ext4"
		mkfs.ext4 ./image.raw
		;;
esac

## Convert raw->vhd
qemu-img convert -f raw -o subformat=fixed,force_size -O vpc image.raw image.vhd
rm image.raw
