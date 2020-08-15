### Informações sobre o exercício "Kubernetes e HPA" (14/08/2020)
<br/>
Exercício realizado conforme solicitado

#### Criando o container Busybox para fazer a carga em go-hpa
``` kubectl run -it loader --image=busybox /bin/sh ```

#### Executando o comando de carga em go-hpa
``` while true; do wget -q -O- http://go-hpa-service.default.svc.cluster.local; done; ```
