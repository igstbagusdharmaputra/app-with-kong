FROM golang:1.16-alpine AS build-env
RUN apk --no-cache add git gcc
COPY . /src
RUN cd /src && go build -o goapp

FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
CMD ["./goapp"]