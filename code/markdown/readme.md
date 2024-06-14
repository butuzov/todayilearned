# Markdown

- https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax
- https://daringfireball.net/projects/markdown/syntax#backslash

## Links

- Link [1]
- [foobar][2]
- [Link](http://google.com)
- [example@example.com](mailto:example@example.com)
- [Call Me!](tel:1111111)

[1]: http://google.com
[2]: http://google.com

```markdown
- Link [1]
- [foobar][2]
- [Link](http://google.com)
- [example@example.com](mailto:example@example.com)
- [Call Me!](tel:1111111)

[1]: http://google.com
[2]: http://google.com
```

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

### Diffs

```diff
def example do
-  :ok
+  {:ok, :more_stuff}
end
```
````
```diff
def example do
-  :ok
+  {:ok, :more_stuff}
end
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

## Special Blocks @ Github

> [!NOTE]
> Highlights information that users should take into account, even when skimming.

```markdown
> [!NOTE]
> Highlights information that users should take into account, even when skimming.
```

> [!TIP]
> Optional information to help a user be more successful.

```markdown
> [!TIP]
> Optional information to help a user be more successful.
```

> [!IMPORTANT]
> Crucial information necessary for users to succeed.

```markdown
> [!IMPORTANT]
> Crucial information necessary for users to succeed.
```

> [!WARNING]
> Critical content demanding immediate user attention due to potential risks.

```markdown
> [!WARNING]
> Critical content demanding immediate user attention due to potential risks.
```

> [!CAUTION]
> Negative potential consequences of an action.

```markdown
> [!CAUTION]
> Negative potential consequences of an action.
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
