all:
	# Build C object
	gcc -c -o clib/hello.o clib/hello.c
	# Archive object as static library
	ar cru clib/libhello.a clib/hello.o
	# Build Go excutable with C library
	go build main.go
	./main

clean:
	rm ./clib/*.o ./clib/*.a
	rm ./main

