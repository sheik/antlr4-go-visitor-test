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

*NOTE:* This is currently built using a custom version of ANTLR4, modified to output visitor structure for Golang (github.com/sheik/antlr4).
If you wish to try it out you will need to check out the custom antlr4 and edit the go.mod replace statement to point Go at it. You will
also need to build the custom antlr4 tool in order to run it.

There is a PR for merging this into the official ANTLR4 repository here: https://github.com/antlr/antlr4/pull/3299


