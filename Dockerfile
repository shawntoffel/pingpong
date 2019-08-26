FROM golang:1.12 as build
WORKDIR /go/src/github.com/shawntoffel/pingpong
ADD . .
RUN make build-linux
RUN echo "nobody:x:100:101:/" > passwd

FROM scratch
COPY --from=build /go/src/github.com/shawntoffel/pingpong/passwd /etc/passwd
COPY --from=build --chown=100:101 /go/src/github.com/shawntoffel/pingpong/bin/pingpong .
USER nobody
ENTRYPOINT ["./pingpong"]