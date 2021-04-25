# build
FROM golang as Builder
WORKDIR /go/src
ADD . .
RUN make

# run
FROM alpine
COPY --from=builder /go/src/signalingchannel .
ENTRYPOINT ["./signalingchannel"]
