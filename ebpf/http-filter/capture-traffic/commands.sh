docker run --privileged --rm -ti -v /usr/src:/usr/src -v /lib/modules:/lib/modules -v /linux-kernel:/linux-kernel -v /sys:/sys  sniffer:1.0 

docker exec -it kind-control-plane bash
kind load docker-image sniffer:1.0 --name kind