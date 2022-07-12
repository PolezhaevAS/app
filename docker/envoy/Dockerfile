FROM envoyproxy/envoy-dev:46cfdb57d552a58a0ad109ef313d004eb4eb7adc

RUN apt-get update && apt-get install iproute2 -y

COPY ./docker/envoy/entry.sh .

RUN chmod +x /entry.sh

ENTRYPOINT [ "/entry.sh" ]

CMD ["envoy", "-c", "envoy.yaml"]