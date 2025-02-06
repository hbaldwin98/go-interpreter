# Notes and such
## Lexical Analysis
Lexing, a Lexer, Lexical Analysis all refer to the same thing. We take an input and generate tokens based on that input. These tokens
represent valid or invalid inputs. We can then use these tokens to parse out an AST.

### Challenges

 - Add greater support for more common program language features
    - Floats (**DONE**) 
    - Tuples (this may be a parsing problem?)
    - Postfix (increment) operations.
    - more...

## Parsing

Two main strategies
 - Top-down
 - Buttom-up

We are using a **recursive descent parser**, otherwise known as a **top down operator precendance** parser.
Also called a **Pratt parser**, after it's creater Vaughan Pratt. The goal of parsing is to generate some kind of tree,
an **Abstract Syntrax Tree** (**AST**).

Parsing is essentially a solved problem in computer science and, using a **Context-Free Grammer** (**Context-Free Grammar**), could generate one using
a tool such as [Bison](https://www.gnu.org/software/bison/), however, doing it ourselves like everything is an invaluable experience. Also, building our own
probably would be better if we wanted to add more features to our language.

*Precendence* is done through top down operator precendence (or Pratt parsing). 
*Right Binding Power* UNDERSTAND
*Left Binding Power* UNDERSTAND

### Challenges

 - Write a formal proof of this language

## Statements, Expressions, Identifiers

Statements don't produce values: `let x = 5` doesn't produce a value in of itself.
Expressions produce values: `5 + 5` or even just `5`
Identifiers represent values: `let <identifier> = <expression>`

