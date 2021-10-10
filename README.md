This is an example grammar from the definitive ANTLR4 reference.

The main files of interest are main.go and parser/expr_visitor.go.

You can see the implementation of Visit functions in parser/expr_base_visitor.g

    jeff@ubuntu-vm:~/Projects/antlr4-go-visitor-test$ go build
    jeff@ubuntu-vm:~/Projects/antlr4-go-visitor-test$ ./antlr4-go-visitor-test 
    calc> 2 + 3
    5
    calc> a = (2 + 5) * (3 + 323)
    calc> a
    2282
    calc> a * 4
    9128
    calc>  


