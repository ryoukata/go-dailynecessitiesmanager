FROM scratch

COPY go-dailynecessitiesmanager .

ENTRYPOINT ["./go-dailynecessitiesmanager"]

