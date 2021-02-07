# build
FROM golang as Builder
WORKDIR /go/src
COPY * .
RUN make

# run
FROM alpine
COPY --from=builder /go/src/signalchannel .
ENTRYPOINT ["./signalchannel"]
