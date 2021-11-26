# Programmer's Brain

## READING CODE BETTER

### Decoding Your Confusion While Coding

### Speed Reading For Code

#### Summary 

- The STM has a capacity of two to six elements.
- To overcome the size limitation, your STM collaborates with your LTM when you remember information.
- When you read new information, your brain tries to divide the information into recognizable parts called chunks.
- When you lack enough knowledge in your LTM, you have to rely on low-level reading of code, like letters and keywords. When doing that, you will quickly run out of space in your STM.
- When your LTM stores enough relevant information, you can remember abstract concepts like “a for-loop in Java” or “selection sort in Python” instead of the code at a lower level, occupying less space in your STM.
- When you read code, it is first stored in the iconic memory. Only a bit of the code is subsequently sent to the STM.
- Remembering code can be used as a tool for (self) diagnosis of your knowledge of coding. Because you can most easily remember what you already know, the parts of code that you remember can reveal the design patterns, programming constructs, and domain concepts you are most familiar with.
- Code can contain characteristics that make it easier to process, such as design patterns, comments, and explicit beacons.

#### Excersize


1. Select Code (max 50 lines)
2. Study code (max 2 minutes)
3. Reproduce code (retype it)
4. Reflect:
    - Which parts did you produce correctly with ease?
    - Are there any parts of the code that you reproduced partly?
    - Are there parts of the code that you missed entirely?
    - Do you understand why you missed the lines that you did?
    - Do the lines of code that you missed contain programming concepts that are anfamiliar to you?
    - Do the lines of code that you missed contain domain concepts that are unfamil- iar to you?
5. Compare with someone else (optional)

### How To Learn Programming Syntax More Quickly 

#### Excersize (remembering new programming concepts)

- What concepts does this new concept make you think of? Write down all the related concepts.
- Then, for each of the related concepts you can think of, answer these questions:
    * Why does the new concept make me think of this concept that I already
know?
    * Does it share syntax?
    * Is it used in a similar context?
    * Is this new concept an alternative to one I already know?
- What other ways do you know to write code to achieve the same goal? Try to cre- ate as many variants of this code snippet as you can.
- Do other programming languages also have this concept? Can you write down examples of other languages that support similar operations? How do they dif- fer from the concept at hand?
- Does this concept fit a certain paradigm, domain, library, or framework?

#### Summary

- It’s important to know quite a bit of syntax by heart because more syntax knowl- edge will ease chunking. Also, looking up syntax can interrupt your work.
- You can use flashcards to practice and remember new syntax, with a prompt on one side and code on the other side.
- It’s important to practice new information regularly to fight memory decay.
- The best kind of practice is retrieval practice, where you try to remember information before looking it up.
- To maximize the amount of knowledge you remember, spread your practice
over time.
- Information in your LTM is stored as a connected network of related facts.
- Active elaboration of new information helps strengthen the network of memories the new memory will connect to, easing retrieval.

### How To Read Complex Code

- Cognitive load represents the limit of what the working memory can process. When you experience too much cognitive load, you cannot properly process code.
- There are two types of cognitive load that are relevant in programming: intrinsic cognitive load is created by the inherent complexity of a piece of code, while extraneous cognitive load is added to code either accidentally (by the way it is presented) or because of gaps in the knowledge of the person reading the code.
- Refactoring is a way to reduce extraneous cognitive load by transforming code to align better with your prior knowledge.
- Creating a __dependency graph__ can help you understand a piece of complex and interconnected code.
- Creating a __state table__ containing the intermediate values of variables can aid in reading code that is heavy on calculations.

## THINKING ABOUT CODE

### Reaching A Deeper Understanding Of Code

- _`Fixed value`_ — A variable whose value does not change after initialization plays the role of a fixed value. This can be a constant value if the programming lan- guage you are using allows for values to be fixed, or it can be a variable that is initialized once and afterward is not changed. Examples of fixed-value variables include mathematical constants like pi, or data read from a file or database.
- _`Stepper`_ — When iterating in a loop, there is always a variable stepping through a list of values. That is the role of the stepper, whose value can be predicted as soon as the succession starts. This can be an integer, like the canonical i iterating in a for- loop, but more complicated steppers are also possible, like size = size / 2 in a binary search, where the size of the array to be searched is cut in half on every iteration.

- _`Flag`_ — A variable used to indicate that something has happened or is the case. Typical examples are `is_set`, `is_available`, or `is_error`. Flags are often `Booleans`, but they can be integers or even strings.

- _`Walker`_ — A walker traverses a data structure, similar to a stepper. The difference lies in the way the data structure is traversed. A stepper always iterates over a list of values that are known beforehand, like in a for-loop in Python: `for i in range(0, n)`. A walker, on the other hand, is a variable that traverses a data structure in a way that is unknown before the loop starts. Depending on the programming language, walkers can be pointers or integer indices. Walkers can traverse lists, for example in binary search, but more often traverse data struc- tures like a stack or a tree. Examples of a walker are a variable that is traversing a linked list to find the position where a new element should be added or a search index in a binary tree.

- _`Most recent holder`_ — A variable that holds the latest value encountered in going through a series of values is a most recent holder. For example, it might store the latest line read from a file (line = file.readline()), or a copy of the array element last referenced by a stepper (element = list[i]).

- _`Most wanted holder`_ — Often when you are iterating over a list of values, you are doing that to search for a certain value. The variable that holds that value, or the best value found so far, is what we call a most wanted holder. Canonical examples of a most wanted holder are a variable that stores a minimum value, a maximum value, or the first value meeting a certain condition.

- _`Gatherer`_ — A gatherer is a variable that collects data and aggregates it into one value. This can be a variable that starts at zero and collects values while iterating through a loop, like this:

  ```python
  sum = 0
  for i in range(list):
      sum += list[i]
  ```
  Its value can, however, also be calculated directly in functional languages or lan- guages that encompass certain functional aspects:`functional_total=sum(list)`. 

- _`Container`_ — A container is any data structure that holds multiple elements that can be added and removed. Examples of containers are lists, arrays, stacks, and trees.

- _`Follower`_ — Some algorithms require you to keep track of a previous or subse- quent value. A variable in this role is called a follower and is always coupled to another variable. Examples of follower variables are a pointer that points to a previous element in a linked list when traversing the list, or the lower index in a binary search.

- _`Organizer`_ — Sometimes a variable has to be transformed in some way for further processing. For example, in some languages, you cannot access individual char- acters in a string without converting the string to a character array first, or you may want to store a sorted version of a given list. These are examples of organiz- ers, which are variables that are only used for rearranging or storing values dif- ferently. Often, they are temporary variables.

- _`Temporary`_ — Temporary variables are variables that are used only briefly and are often called temp or t. These variables may be used to swap data or to store the result of a computation that is used multiple times in a method or a fucntion.

### Getting Better At Solving Programming Problems

### Misconceptions: Bugs In Thinking

## WRITING CODE BETTER

### How To Get Better At Naming Things 

### Avoiding Bad Code And Cognitive Load: Two Frameworks

### Getting Better At Solving Complex Problemsaa

## COLLABORATING ON CODE

### The Act Of Writing Code

### Designing And Improving Larger Systems

### How To Onboard New Developers