#this is the alpine image with the go tools added
FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=/Users/ckumar/gocode/src/github.com/gok8s/app-code

#add the source code:
ADD . $SRC_DIR

#build it
RUN cd $SRC_DIR; go build -o httpExample; cp httpExample /app/

ENTRYPOINT ["./httpExample"]