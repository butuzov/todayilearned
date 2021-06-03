# Code Reviews

* `CR` - code review
* `CL` - changelist

## Reading 

* https://google.github.io/eng-practices/review/reviewer/
* https://google.github.io/eng-practices/review/developer/
* https://newrelic.com/blog/best-practices/code-review-guidelines
* https://www.michaelagreiler.com/all-posts/

## Automating `CL`

* [ ] tocs/docs (`.gitcommit`)
* [ ] linting (only your `feature` branch, skip the rest)
* [ ] tests (all your changes in tests, then all cahnges)
* [ ] does it build (with tests)?

## Befoer Code Review

* Make Expectations Clear (`when` and `what` is going to happen)
* Don't Demand Changes, Make Suggestions And Explain 
* Be Open For Discussions
* Eliminate Personal Code Authorship
* Reserve Enough Time
* Automate The Boring Stuff (`pull` & `run`)

### What Do Code Reviewers Look For?

- [ ] Design: Is the code well-designed and appropriate for your system?
- [ ] Functionality: Does the code behave as the author likely intended? Is the way the code behaves good for its users?
- [ ] Complexity: Could the code be made simpler? Would another developer be able to easily understand and use this code when they come across it in the future?
- [ ] Tests: Does the code have correct and well-designed automated tests?
- [ ] Naming: Did the developer choose clear names for variables, classes, methods, etc.?
- [ ] Comments: Are the comments clear and useful?
- [ ] Style: Does the code follow our style guides?
- [ ] Documentation: Did the developer also update relevant documentation?


#### TL/DR

* The code is well-designed.
* The functionality is good for the users of the code.
* Any UI changes are sensible and look good.
* Any parallel programming is done safely.
* The code isn’t more complex than it needs to be.
* The developer isn’t implementing things they might need in the future but don’t know they need now.
* Code has appropriate unit tests.
* Tests are well-designed.
* The developer used clear names for everything.
* Comments are clear and useful, and mostly explain why instead of what.
* Code is appropriately documented (generally in g3doc).
* The code conforms to our style guides.

Coding Practices

* [ ] Code is easy to understand, does not repeat itself
* [ ] No redundant code, no spelling errors
* [ ] Descriptive naming, use enums where applicable
* [ ] Functions have been broken up to only do one thing
* [ ] Expected errors thrown are either handled or returned in API response
* [ ] Errors logged to Sentry are actionable
* [ ] Use dependency injection where possible
* [ ] Follow Law of Demeter. A class should know only its direct dependencies