all: build run clean

build:
	go build

run:
	go run ex02.go --old ..\tests\snapshot1.txt --new ..\tests\snapshot2.txt
	go run ex02.go --old ..\tests\snapshot2.txt --new ..\tests\snapshot1.txt
	go run ex02.go --old ..\tests\recipes.json --new ..\tests\recipes.xml
clean:
	rm -f ex02.exe