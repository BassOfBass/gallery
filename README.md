# gallery
Self-hosted photo gallery

## Usage

### From source

```
go get -u github.com/shivakar/gallery

gallery [DIRECTORIES...]
```

### Using Docker

```
docker pull shivakar/gallery

docker run --detach -p 8080:8080 -v PATH/TO/ALBUMS:/mnt/albums shivakar/gallery gallery /mnt/albums/*
```

## References

* http://photoswipe.com/documentation/getting-started.html
* https://webdesign.tutsplus.com/tutorials/the-perfect-lightbox-using-photoswipe-with-jquery--cms-23587
