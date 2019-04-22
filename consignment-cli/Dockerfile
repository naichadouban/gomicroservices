FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
# 将当前目录下的货物信息consignment.json复制到 /app下
ADD consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli

ENTRYPOINT ["./consignment-cli"]
CMD [ "./consignment-cli" ]