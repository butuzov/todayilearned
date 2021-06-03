# Code Comments

## Different types of code comments

1. Documentation comments - defining propose of file or component. `doc.go`, `phpdoc` for the file, etc.
2. Function comments - comments on public method, most probably going to be genrated into some standalone doc.
3. Logic comments - decribes logic.

## "Best" Practices

This list is result of reading article [The Engineer’s Guide to Writing Meaningful Code Comments](https://dzone.com/articles/the-engineers-guide-to-writing-meaningful-code-com). I am not agree with everything written there on 100%. 

1. Make Use of Code Annotations or Tags (phpdoc, javadoc, pyhon annotations)
2. Write Down Why You Are Doing Something. Not **What** but **Why**
3. Don’t Refer to Other Documents or Comments (unless its jira ticket, or github issue).
4. Write Comments ~~While~~ Before Writing Code - it clears your mind. If you can't explain it, its bad idea. 

## 3. Why?

> Unfortunately, many comments are not particularly helpful. The most common reason is that comments repeat the code: all of the information in the comment can easily be deduced from the code next to the comment.

John Ousterhout’s [A Philosophy of Software Design](https://www.abebooks.com/9781732102200/Philosophy-Software-Design-Ousterhout-John-1732102201/plp).

### Architecture Desicion Log

- I am an idiot, I don't understant this at all. It must be correct.
- Context, When, Why


`adr`

https://adr.github.io/