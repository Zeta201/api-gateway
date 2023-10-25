https://loft.sh/blog/tutorial-how-ebpf-improves-observability-within-kubernetes/

docker build -t dinitha/sniffer .

docker run --privileged --rm -ti -v /usr/src:/usr/src -v /lib/modules:/lib/modules -v /linux-kernel:/linux-kernel -v /sys:/sys  http-parser:1.0 
docker exec -it kind-control-plane bash
kind load docker-image http-parser:1.0 --name kind