FROM centurylink/golang-builder AS builder

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s" -a -installsuffix cgo -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app .

ENTRYPOINT ["/app"]