## Refresh mini-manual
Negapedia refresh generator and development environment: this package and docker image is responsible of generating [negapedia](http://en.negapedia.org/), a website on social data extracted from [wikipedia](https://en.wikipedia.org).

### Requirements
You will need a machine with internet connection, 16GB of RAM, 300GB of storage and [docker storage base directory properly setted](https://forums.docker.com/t/how-do-i-change-the-docker-image-installation-directory/1169).

### Description of operations flow
This image take in input the nationalization and store the result of the operations in `/data` (in-container folder). All the operation of data fetching are totally automatized and the result is [negapedia website](http://negapedia.org) in the form of a gzipped tarball of gzipped webpages. The operations flow is composed of thee phases:
1. preprocessing of data: CPU intensive, it requires a good internet connection and 16GB of RAM;
2. exporting to csv, CPU intensive.
3. (optional) calculating [TFIDF](https://github.com/negapedia/wikitfidf), CPU and IO intensive.
4. construction of in-container database - IO intensive, requires 300GB of storage, best if SSD.
5. exporting and compressing the static website from quering the database and TFIDF data.

### Refresh options
1. `lang`: [wikipedia nationalization to parse](https://github.com/negapedia/wikiassignment/tree/master/nationalization/internal/languages), default `it`.
2. `url`:  Output base URL, `%s` is the optional placeholder for subdomain, default `http://%s.negapedia.org`.
3. `source`: source of data (`net` or `savepoint`), default `net`.
4. `keep`: keep every savepoint after the execution (`true` or `false`), default `false`.
5. `tfidf`: calculate TFIDF, if `false`, try available precalculated measures (`true` or `false`), default `false`.
6. `test`: Run as test on a fraction of the articles before `savepoint` (`true` or `false`), default `false`.

### Examples
1. `docker run negapedia/negapedia refresh -lang en`: basic usage, run the image on the english nationalization and store the result in the in-containter `/data` folder.
2. `docker run -v /path/2/out/dir:/data negapedia/negapedia --rm refresh -lang en`:
..1. run the image as before.
..2. [mount as a volume](https://docs.docker.com/storage/volumes/) the guest `/data` folder to the host folder `/path/2/out/dir`, the output folder, so that at the end of the operations  `/path/2/out/dir` will contain the result. This folder can be changed to an arbitrary folder of your choice.
..3. remove the image right after the execution.
3. `docker run -v /path/2/out/dir:/data  --rm --init -d negapedia/negapedia refresh -lang en`, **you may want to use this commad** :
..1. run the image as before.
..2. run an init process that will take care of killing eventual zombie processes - just in case.
..3. run the image in detatched mode.
For further explanations please refer to [docker run reference](https://docs.docker.com/engine/reference/run)

### Useful commands
1. `docker pull negapedia/negapedia` Update the image to the last revision.
2. `docker kill --signal=SIGQUIT  $(docker ps -ql)` Quit the last container and log trace dump.
3. `docker kill --signal=SIGUSR1  $(docker ps -ql)` Log the trace dump of the last container without quitting it.
4. `docker logs -f $(docker ps -lq)` Fetch the logs of the last container.
5. `docker system prune -fa --volumes` Remove all unused images and volume without asking for confirmation.