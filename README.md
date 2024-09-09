# super simple docker-short-cli ( 1 hours to write )

![cat](./assets/image.png)


## installing 

```bash
make install
```

or

```bash
chmod +x scripts/install.sh
scripts/install.sh
```
---

### Required min 2 argumets after docker-short

#### Usage:

```bash
dcshort <run|build> <image-name> [<container-name>] [<port>]
```
---

#### Build example:

```bash
dcshort build image:namev1 # docker build . -t image:namev1 -f Dockerfile
```

#### Run exmaple 

```bash
dcshort run image:namev1 test-con 9090:9090 # docker run --rm -d --name test-con -p 9090:9090 image:namev1
```

--- 
#### Command List

- run | command who replace docker run commadn with --rm and optional --name and -p flags

- build | command who replace docker build 

--- 
#### alias

```bash
make alias #command who added alias ds="dcshort" in ur zshrc file if this alias doesnt exist 
```