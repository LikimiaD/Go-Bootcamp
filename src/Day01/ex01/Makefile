all: build run clean

build:
	go build

run:
	go run ex01.go --old ..\tests\recipes.xml --new ..\tests\recipes.json
	go run ex01.go --old ..\tests\recipes.xml --new ..\tests\recipes.xml
	go run ex01.go --old ..\tests\recipes.json --new ..\tests\recipes.xml
clean:
	rm -f ex01.exe