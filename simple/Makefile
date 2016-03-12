calc: y.go
	go build -o calc y.go

y.go: calc.y
	go tool yacc -p Calc calc.y

clean:
	rm -f calc y.go y.output
