# Because I Can't Trust My Memory (v3) [![build](https://github.com/butuzov/todayilearned/actions/workflows/publish-pages.yaml/badge.svg)](https://butuzov.github.io/)


> This is some kind of a memory dump in proccess of transition. Author intends to use it as quick help from time to time. You also can do that.

## `hugo` vs git style `readme`

This version of til notes will also supports web publication, but in order to keep writing in same style, we going to transform `readme.md` files into `hugo`'s `_index.md` branch bundles.  In order to do so, keep in ming few rules of writing content:

- In order to generate [front matter](https://gohugo.io/content-management/front-matter/), add html comment to the top of the file:

<table>
<thead><tr><th><code>github</code>'s <code>readme.md</code></th><th><code>hugo</code>'s markdown</th></tr></thead>
<tbody>
<tr><td>

```markdown
<!-- test: true -->
<!-- link: false -->
```

</td><td>

```markdown
---
test: true
link: false
---
```

</td></tr>
</tbody></table>

## Titles

- __Menu__. By default `menu` link name correlated with first `# `, in order to make it custom, use `menu: $name` front matter (e.g. `<!-- menu: ssh -->`).
- SEO Titles is `seotitle` (other wise `menu`/`title`/`.content>h1`) will be used.
- With `title: name` first h1 is ignored.

## Redirection
