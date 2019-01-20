# Searching Tips

## Search Engines

### Google

 Keyword            | Description           | Example
--------------------|-----------------------|----------------------------
`-word`             | Exclude Word          | `serenity -2019`
`site:domain.zone`  | Search in domain zone | `serenity site:ua`
`*`                 | wildcard for words    | `my * revenge`
`related:domain`    | related domains       | `related:news.ycombinator.com`
`OR`                | or search             | `black OR white`
`..`                | range search          | `tv series 2006..2013`
`time *where*`      | time at some place    | `time Mexico`
`filetype:pdf`      | filetype search       | `stochastic math filetype:pdf`

### DuckDuckGO

Will work in (almost) same way as google. If you need google add `!g` at the beginng of your search term.

### Twitter
https://twitter.com/search-home

 Operator                           | Meaning
------------------------------------|-------------------------------------------------------------
`twitter search`                    | Containing both "twitter" and "search". This is the default `operator.
`"happy hour"`                      | Containing the exact phrase "happy hour".
`love OR hate`                      | Containing either "love" or "hate" (or both).
`beer -root`                        | Containing "beer" but not "root".
`#haiku`                            | Containing the hashtag "haiku".
`from:alexiskold`                   | Sent from person "alexiskold".
`to:techcrunch`                     | Sent to person "techcrunch".
`@mashable`                         | Rreferencing person "mashable".
`"happy hour" near:"san francisco"` | Containing the exact phrase "happy hour" and sent near "san francisco".
`near:NYC within:15mi`              |sent within 15 miles of "NYC".
`superhero since:2010-12-27`        |containing "superhero" and sent since date "2010-12-27" (year-month-day).
`ftw until:2010-12-27`              |containing "ftw" and sent up to date "2010-12-27".
`hilarious filter:links`            |containing "hilarious" and linking to URLs.
`news source:"Twitter Lite"`        |containing "news" and entered via Twitter Lite


### Github

https://github.com/search

Most of tips can be found on [official page](https://help.github.com/articles/searching-code/).

 Keyword        | Description           | Example
----------------|-----------------------|----------------------------
`filename`      | Search by Filename    | `filename:serverless.yml`
`in:file`       | Search in Files       | `in:file func main`
`in:path`       | Search in Path        | `in:path tips.pdf`
`extension:pdf` | Search by Extension   | `in:path extension:pdf tips`


## Pro Tips

 * Bookmarkelets for regular search (recent news in the field)

  ```javascript
  // Twitter Search for Recent Tips in Go Programming
  // Normalized view
  javascript:(function(){
    var m = "01,02,03,04,05,06,07,08,09,10,11,12".split(',');
    f = (d) => {return d.getFullYear()+'-'+m[d.getMonth()]+'-'+d.getDate()};
    var d=prompt("Start Date", f(new Date()));
    var today=new Date(d);
    var weekago=new Date(today.getTime()-(7*24*60*60*1000));
    q=`golang tips since:${f(weekago)} until:${f(today)}`;
    window.location.href=`https://twitter.com/search?l=&q=${encodeURI(q)}&src=typd`;
  })();
  ```
