t1:
	/usr/local/go/bin/go build -o ./main ./task_1 ;
	./main

t2:
	go build -v ./task_2 ;
	./main

t3:
	go build -v ./task_3
	./main

t4:
	go build -v ./task_4 ;
	./main

.DEFAULT_GOAL := .build
.PHONY: t1 t2 t3 t4