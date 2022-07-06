SCALER_TAG = 0.9.0
SAFESCALE_TAG = v22.06.0
CA_TAG = 1.22.3

build-scaler-image:
	cd scaler && go build -o ./build/rs-infra-scaler .
	wget https://github.com/CS-SI/SafeScale/releases/download/${SAFESCALE_TAG}/safescale-${SAFESCALE_TAG}-darwin-amd64.tar.gz -O scaler/build/sf.tar.gz
	cd scaler/build && tar -xzf sf.tar.gz safescale
	rm scaler/build/sf.tar.gz
	sudo docker build -t artifactory.coprs.esa-copernicus.eu/rs-docker/rs-infra-scaler:${SCALER_TAG} -f scaler/Dockerfile .

build-safescaled-image:
	wget https://github.com/CS-SI/SafeScale/releases/download/${SAFESCALE_TAG}/safescale-${SAFESCALE_TAG}-darwin-amd64.tar.gz -O scaler/build/sf.tar.gz
	cd scaler/build && tar -xzf sf.tar.gz safescaled
	rm scaler/build/sf.tar.gz
	sudo docker build -t artifactory.coprs.esa-copernicus.eu/rs-docker/rs-infra-scaler:${SAFESCALE_TAG} -f scaler/safescaled.Dockerfile .

build-cluster-autoscaler-image:
	wget https://github.com/kubernetes/autoscaler/archive/refs/tags/cluster-autoscaler-${CA_TAG}.tar.gz -O scaler/build/ca.tar.gz
	cd scaler/build && tar -xzf ca.tar.gz autoscaler-cluster-autoscaler-${CA_TAG}/cluster-autoscaler
	cd scaler/build/autoscaler-cluster-autoscaler-${CA_TAG}/cluster-autoscaler && make build-arch-amd64
	cd scaler/build/autoscaler-cluster-autoscaler-${CA_TAG}/cluster-autoscaler && sudo docker build -t artifactory.coprs.esa-copernicus.eu/rs-docker/cluster-autoscaler:${CA_TAG} -f Dockerfile.amd64 .

push-scaler-image:
	sudo docker push artifactory.coprs.esa-copernicus.eu/rs-docker/rs-infra-scaler:${SCALER_TAG}

push-safescaled-image:
	sudo docker push artifactory.coprs.esa-copernicus.eu/rs-docker/safescaled:${SAFESCALE_TAG}

push-scaler-image:
	sudo docker push artifactory.coprs.esa-copernicus.eu/rs-docker/cluster-autoscaler:${CA_TAG}
