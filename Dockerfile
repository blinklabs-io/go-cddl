FROM golang:1.19 AS build

WORKDIR /code
COPY . .
RUN make build

FROM cgr.dev/chainguard/glibc-dynamic AS cddl
COPY --from=build /code/cddl /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/cddl"]
