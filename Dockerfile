FROM golang:1.17 as build
WORKDIR /go/src/github.com/shawntoffel/pingpong
ADD . .
RUN make build-linux

FROM scratch
COPY --from=build --chown=20000:20001 /go/src/github.com/shawntoffel/pingpong/bin/pingpong .
ENTRYPOINT ["./pingpong"]