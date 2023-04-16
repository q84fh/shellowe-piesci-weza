all: run

run:
	docker run --name slidev --rm -it --user node -v ${PWD}:/slidev -p 3030:3030 tangramor/slidev:latest

build:
	docker run --name slidev --rm -it --user node -v ${PWD}:/slidev --entrypoint="" tangramor/slidev:latest npx slidev build --base=/shellowe-piesci-weza/;	rm -rf docs/*; mv dist/* docs/; rmdir dist
