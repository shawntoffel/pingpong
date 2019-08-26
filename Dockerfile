FROM golang:1.12 as build
WORKDIR /go/src/github.com/shawntoffel/pingpong
ADD . .
RUN make build-linux

FROM scratch
COPY --from=build --chown=100:101 /go/src/github.com/shawntoffel/pingpong/bin/pingpong .
ENTRYPOINT ["./pingpong"]