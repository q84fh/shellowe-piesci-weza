all: run

run:
	docker run --name slidev --rm -it --user node -v ${PWD}:/slidev -p 3030:3030 tangramor/slidev:latest
