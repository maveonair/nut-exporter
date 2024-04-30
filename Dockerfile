# ------------- BUILD --------------- #
FROM golang:1.22 as build

RUN mkdir -p /src/build
WORKDIR /src/build

COPY . .

RUN make build

# -------------- RUN ---------------- #
FROM scratch

COPY --from=build /src/build/dist/nut-exporter ./

EXPOSE 9055

CMD [ "./nut-exporter"]
