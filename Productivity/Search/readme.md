# Search Tips

## Search Engines
1. [Google](google.md)
1. [DuckDuckGo](google.md#duckduckgo)
1. [Twitter](#twitter)
1. [Github](../GitHub)

## [ProPips](pro-tips)

----

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
`near:NYC within:15mi`              | Sent within 15 miles of "NYC".
`superhero since:2010-12-27`        | Containing "superhero" and sent since date "2010-12-27" (year-month-day).
`ftw until:2010-12-27`              | Containing "ftw" and sent up to date "2010-12-27".
`hilarious filter:links`            | Containing "hilarious" and linking to URLs.
`news source:"Twitter Lite"`        | Containing "news" and entered via Twitter Lite



## Pro Tips

 * Bookmarkelet for regular twitter search on term for a week period

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

  * [ifttt: Save tweets featuring specific content to a spreadsheet](https://ifttt.com/applets/P45PCZKW-save-tweets-featuring-specific-content-to-a-spreadsheet)
