## `Test doubles`

* `Dummy` is only for parameter filling, so we can pass test (`initialize` - `exercise` - `verify`).
* `Fake` working implementation, not suitable for prod, good for tests (`initialize` - `exercise` - `verify`). 
* `Stub` dependecy that hide predefined inputs and behaviour (`initialize` - `exercise` - `verify`). 
* `Mock` is stub, where you need to define your expectations (`initialize` - `set expectations` - `exercise` - `verify`).
* `Spy` interseptor with recorder (`initialize` - `exercise` - `verify`)


#### Reading 

* https://martinfowler.com/articles/mocksArentStubs.html
* https://stackoverflow.com/questions/3459287/whats-the-difference-between-a-mock-stub

##  `Given` – `When` – `Then`


The Given-When-Then formula is a template intended to guide the writing of acceptance tests for a User Story:

- (Given) some context
- (When) some action is carried out
- (Then) a particular set of observable consequences should obtain

An example:

- Given my bank account is in credit, and I made no withdrawals recently,
- When I attempt to withdraw an amount less than my card’s limit,
- Then the withdrawal should complete without errors or warnings


## `!test`

A test is not a unit test if:

 * It talks to the database
 * It communicates across the network
 * It touches the file system
 * It can't run at the same time as any of your other unit tests
 * You have to do special things to your environment to run it.