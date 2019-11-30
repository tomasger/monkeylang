# monkeylang
**monkeylang** is a repository for an interpreter of a custom programming language as seen in Thorsten Ball's book "Writing an Interpreter in Go".

The code follows the book closely, but I intend to apply cleaner code practices if/when possible and implement other useful features that have been mentioned, but are out of the scope of this book. 

## Usage
Run the program:
```
go run .
```
**monkeylang** currently prints out the token representation of the entered statements. Example:
```
Welcome to monkeylang!
This prompt now accepts your commands:
>> let x = "asd";
{Type:LET Literal:let}
{Type:IDENT Literal:x}
{Type:= Literal:=}
{Type:ILLEGAL Literal:"}
{Type:IDENT Literal:asd}
{Type:ILLEGAL Literal:"}
{Type:; Literal:;}
```
