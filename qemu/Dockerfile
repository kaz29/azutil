FROM ubuntu:16.04

RUN set -xe \
	&& apt-get update && apt-get install -y qemu-utils --no-install-recommends && rm -rf /var/lib/apt/lists/*

#COPY make_vhd.sh /home/qemu/

#WORKDIR /home/qemu
#CMD ["sh", "make_vhd.sh"]
