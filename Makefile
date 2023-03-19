all: build-image run-container
build-image:
	docker build -t ascii-art .
run-container:
	docker run -d --name=ascii-art-container -p 8080:8080 ascii-art
