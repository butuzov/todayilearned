# `github`

## Tooling

- https://github.com/sneas/sensus (chrome only)
- https://github.com/refined-github/refined-github

## Code Reviews

1. [Embeding code snippets](https://github.blog/2017-08-15-introducing-embedded-code-snippets/)
3. butuzov/deadlinks#76 to issue
4. @butuzov user itself
5. suggestions

      ```suggestion
      code goes here.
      ```

## Search

 Keyword        | Description           | Example
----------------|-----------------------|----------------------------
`path:pdf`      | In Path               | `path:docker-compose.yaml kafka`



### Searching for alive people

Original Code posted by Denis Dinkevich in telegram sourcing comunity.

[github.bookmarkelet.js](github.bookmarkelet.js)

### Legacy

https://github.com/search

Most of tips can be found on [official page](https://help.github.com/articles/searching-code/), API for this still in business - `https://api.github.com/search/code?q=`

 Keyword        | Description           | Example
----------------|-----------------------|----------------------------
`filename`      | Search by Filename    | `filename:serverless.yml`
`in:file`       | Search in Files       | `in:file func main`
`in:path`       | Search in Path        | `in:path tips.pdf`
`extension:pdf` | Search by Extension   | `in:path extension:pdf tips`

```python
#!/usr/bin/env python

import requests

url = f"https://api.github.com/search/code?q={q}"
headers = {
  'Authorization': f'Token {GH}'
}
requests.request("GET", url, headers=headers)
```


## Examples

### Usage of `GITHUB_TOKEN`

```shell
git pull https://${{ secrets.GIT_TOKEN }}:x-oauth-basic@github.com/susomejias/portfolio.git master

git init
git remote add origin https://${{ secrets.GIT_TOKEN }}:x-oauth-basic@github.com/$user/$repo.git
git pull
...
git add -A
git commit --allow-empty -m "${DEST_MESSAGE}"
git push -q origin master
```

