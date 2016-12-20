_Fat Ubuntu images, you make the rockin' world go round..._

# x86-64 and armhf Docker images in a single repository

This repository uses the support for fat manifests added in Docker's
[version 2 schema](https://docs.docker.com/registry/spec/manifest-v2-2/) to
build a single Docker Hub repository that can be pulled and used on both x86-64
and armhf for a couple of useful Docker Hub repositories (namely, `ubuntu` and
`redis`).

(This is very convenient if you're building Dockerised microservices for a
Raspberry Pi but want to be able to run them on your x86-64 development machine
as well for easier development and debugging.)

This uses the extremely nifty
[manifest-tool](https://github.com/estesp/manifest-tool) tool to actually do
the work of merging the different images into the same repository.

## Redis

The images that are bundled together are the current
[redis](https://hub.docker.com/_/redis/) and
[armhf/redis](https://hub.docker.com/r/armhf/redis/) images on Docker Hub.

## Ubuntu

The images that are bundled together are the current
[ubuntu](https://hub.docker.com/_/ubuntu/) and
[armhf/ubuntu](https://hub.docker.com/r/armhf/ubuntu/) images on Docker Hub.
Only tags that exist in both images are pushed: at present, that's `precise`,
`trusty`, `xenial` and `yakkety`.

## How do I use these images?

You can use these like the official repositories.

For instance, on x86-64, you can use `lawngnome/ubuntu` as a drop in
replacement for the normal `ubuntu` image:

```
adam@x86-64:~$ docker run -it --rm lawngnome/ubuntu:xenial 
root@9cbbef46747e:/# uname -a
Linux 9cbbef46747e 4.8.13-1-ARCH #1 SMP PREEMPT Fri Dec 9 07:24:34 CET 2016 x86_64 x86_64 x86_64 GNU/Linux
root@9cbbef46747e:/# cat /etc/issue
Ubuntu 16.04.1 LTS \n \l
```

Similarly, on armhf, you can use `lawngnome/ubuntu` as a drop in replacement
for `armhf/ubuntu`:

```
adam@armhf:~$ docker run -it --rm lawngnome/ubuntu:xenial
root@7d7a6ad94c98:/# uname -a
Linux 7d7a6ad94c98 4.1.18-v7+ #846 SMP Thu Feb 25 14:22:53 GMT 2016 armv7l armv7l armv7l GNU/Linux
root@7d7a6ad94c98:/# cat /etc/issue
Ubuntu 16.04.1 LTS \n \l
```

(Yes, I should probably update the kernel on my Raspberry Pi.)

## How often are these images updated?

I've got a script on a server that should update this daily. Generally
speaking, the actual underlying images do not actually update that often;
please check their respective Docker Hub pages for their last update times.

## How do I build these images?

As long as `build.sh` can find the `manifest` binary created by building
[manifest-tool](https://github.com/estesp/manifest-tool) in the path, it will
build every YAML file under the `manifests` directory and push it. You will, of
course, need write access to the target images to be able to use this.
