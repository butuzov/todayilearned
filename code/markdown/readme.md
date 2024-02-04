# Markdown

- https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax
- https://daringfireball.net/projects/markdown/syntax#backslash

## Links

- Link [1]
- [Link](http://google.com)

[1]: http://google.com

## Language Highlight

Github use [Linguist](https://github.com/github-linguist/linguist) to perform language detection and to select [third-party grammars](https://github.com/github-linguist/linguist/blob/master/vendor/README.md) for syntax highlighting. You can find out which keywords are valid in [the languages YAML file](https://github.com/github-linguist/linguist/blob/master/lib/linguist/languages.yml).

```ruby
require 'redcarpet'
markdown = Redcarpet.new("Hello World!")
puts markdown.to_html
```

As code

````
```ruby
require 'redcarpet'
markdown = Redcarpet.new("Hello World!")
puts markdown.to_html
```
````

## Creating diagrams

Supported: Mermaid, GeoJSON, TopoJSON, and ASCII STL. [Read More](https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/creating-diagrams)


````
```mermaid
graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
```
````
```mermaid
graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
```



## Tables

## Alternative headers

Instead `##` or `#` you can:

````
h2.Header
---------

h1.Header
=========
````
