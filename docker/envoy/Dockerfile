FROM envoyproxy/envoy-dev:latest

RUN apt-get update && apt-get install iproute2 -y

COPY ./docker/envoy/entry.sh .

RUN chmod +x /entry.sh

ENTRYPOINT [ "/entry.sh" ]

CMD ["envoy", "-c", "envoy.yaml"]