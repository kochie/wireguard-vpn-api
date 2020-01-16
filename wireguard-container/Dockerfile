FROM alpine
RUN apk add -U wireguard-tools
RUN ip link add dev wg0 type wireguard
RUN ip address add dev wg0 20.10.0.1/24
RUN wg setconf wg0 wg0.conf
RUN iptables -A FORWARD -i wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE