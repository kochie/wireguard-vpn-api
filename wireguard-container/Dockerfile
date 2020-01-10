FROM alpine:3
RUN apk add -U wireguard-tools
RUN ip link add dev wg0 type wireguard
RUN ip address add dev wg0 192.168.2.1/24
RUN wg setconf wg0 myconfig.conf