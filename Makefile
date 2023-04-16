all: run

run:
	docker run --name slidev --rm -it --user node -v ${PWD}/public:/slidev -p 3030:3030 tangramor/slidev:latest

build:
	docker run --name slidev --rm -it --user node -v ${PWD}/public:/slidev --entrypoint="" tangramor/slidev:latest npx slidev build --base=/shellowe-piesci-weza/
	mv dist/* docs/
	rmdir dist
