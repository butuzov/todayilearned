# Bash Arithmetic Expansion

The shell allows arithmetic expressions to be evaluated, as one of the shell expansions or by using the `((` `))` compound command, the `let` builtin, or the `-i` option to the declare builtin.

Evaluation is done in fixed-width integers with no check for overflow, though division by 0 is trapped and flagged as an error. The operators and their precedence, associativity, and values are the same as in the C language. The following list of operators is grouped into levels of equal-precedence operators. The levels are listed in order of decreasing precedence.

```bash
echo 'Pre Increment', $((++i))
echo 'Pre Decrement', $((--i))

echo 'Post Increment', $((i++))
echo 'Post Decrement', $((i--))

echo 'Exponentiation', $((3**3))
echo 'Addition', $((i+4))
echo 'Subtraction', $((i-2))
echo 'Multiplication', $((i*3))
echo 'Division', $((i/3))
echo 'Remainder', $((i%3))

# Comparison Operators - return 1 or 0
echo 'Equality (==)', $((i%3==0))
echo 'Inequality (!=)', $((i%3!=0))

echo 'Comparison - Greater or Equal (>=)', $((i%3 >= 0))
echo 'Comparison - Greater (>)', $((i%3>0))
echo 'Comparison - Less or Equal (<=)', $(( i + 1 <= 10 ))
echo 'Comparison - Less (<)', $(( i + 1 < 0 ))

# Logical Operators
echo 'Logical Negation', $(( ! i ))
echo 'Logical AND', $(( i && 2==2 ))
echo 'Logical OR', $(( 0 || 1==2 ))
echo 'Tenar Comparison (expr ? expr : expr) ', $(( i % 2 == 0 ? 4 : 9 ))
```

# Declare
Allow to use `typed` variables

```bash
declare -i a
a=10
echo $a
a=hello
echo $a
```
