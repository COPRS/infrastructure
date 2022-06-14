build:
	go build -o rs-infra-scaler
	sudo docker build -t artifactory.coprs.esa-copernicus.eu/rs-docker/rs-infra-scaler:0.8.0-rc2 .
	
push:
	sudo docker push artifactory.coprs.esa-copernicus.eu/rs-docker/rs-infra-scaler:0.8.0-rc2
