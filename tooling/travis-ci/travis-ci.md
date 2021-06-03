![](travis-ci.svg)

# Travis_CI

Free for OpenSource CI/CD platform

## Branches

Its possible to limit travis ci only for specific branches:

```yaml
branches:
  only:
    - 'master'             # Master Branch
    - '/^5\.[0-9]+$/'      # Number based 5.0
```

## Environment

  VARIABLE            | Values
----------------------|-------------
 `TRAVIS_BUILD_NUMBER`| number
 `TRAVIS_OS_NAME`     | string (os name)
 `TRAVIS_COMMIT`      | sha hash
 `TRAVIS_BRANCH`      | string (branch)
 `TRAVIS_BUILD_DIR`   | string (directory)
 `TRAVIS_REPO_SLUG`   | string (repo)
 `TRAVIS_PULL_REQUEST`| string ("true" or "false")


## TODO - Build Matrixes

```yaml
matrix:
  include:
  - os: linux
    language: c
    dist: trusty
    sudo: required
  - os: osx
    language: objective-c
    osx_image: xcode8.3
    env:
    - SIMULATOR=true

```

## TODO - Jobs and Stages

## TODO - Continues Deployment

## Notifications

-  https://docs.travis-ci.com/user/notifications/

### Email:

Disable Any Email Messages:

```yaml
notifications:
  email: false
```

Other Possible Settings

```yaml
notifications:
  email:
    recipients:
    - one@example.com 
    on_success: change # default: always
    on_failure: always # default: always
    on_start: change   # default: never
    on_cancel: always  # default: always
    on_error: always   # default: always
```

### Telegram

* bot: `@travisci_build_bot`

```yaml
notifications:
  webhooks: https://fathomless-fjord-24024.herokuapp.com/notify
```

### Slack

* Formating: https://api.slack.com/docs/message-formatting
* Settings: https://docs.travis-ci.com/user/notifications/#configuring-slack-notifications

```yaml
notifications:
  slack:
    rooms:
      - made-ua:XGUkOpb0PsEuqmm0CvHaNfPk#builds
    template:
      - '*%{repository_slug}@%{branch}*<%{compare_url}|#%{commit}> (_%{author}_) - <%{build_url}|Travis CI: Build #%{build_number}>'
      - '*Title:* %{commit_message}'
      - '----------------------------------------------------------------------- '
      - '(`%{build_id}` is `%{result}`) _%{message}_ in `%{elapsed_time}`'

```

## Encryption (PyPi API Keys)

```bash
# first login to github (using token)
> travis login --org --github-token 07cc694b9b3fc636710fa08b6922c42b7abd6ff0
Successfully logged in as butuzov!

> travis encrypt --org pypi-auth-key/pass

Please add the following to your .travis.yml file:

  secure: "0JPQvtGA0ZbQu9CwINGB0L7RgdC90LAsINC/0LDQu9Cw0LvQsCwK0J/RltC0INC90LXRjiDQtNGW0LLQutCwINGB0YLQvtGP0LvQsC4KCtCf0ZbQtCDQvdC10Y4g0LTRltCy0LrQsCDRgdGC0L7Rj9C70LAsCtCg0YPRgdGP0LLRgyDQutC+0YHRgyDRh9C10YHQsNC70LAuCgrigJMg0J7QuSDQutC+0YHQuCwg0LrQvtGB0Lgg0LLQuCDQvNC+0ZcsCtCU0L7QstCz0L4g0YHQu9GD0LbQuNC70Lgg0LLQuCDQvNC10L3Rli4KCtCR0ZbQu9GM0YjQtSDRgdC70YPQttC40YLRjCDQvdC1INCx0YPQtNC10YLQtSwK0J/RltC0INCx0ZbQu9C40Lkg0LLQtdC70YzQvtC9INC/0ZbQtNC10YLQtS4KCtCf0ZbQtCDQsdGW0LvQuNC5INCy0LXQu9GM0L7QvSwg0L/RltC0INGF0YPRgdGC0LrRgywK0JHRltC70YzRiCDQvdC1INC/0ZbQtNC10YLQtSDQt9CwINC00YDRg9C20LrRgy4KCtCf0ZbQtCDQsdGW0LvQuNC5INCy0LXQu9GM0L7QvSDQtyDQutGW0L3RhtGP0LzQuCwK0JHRltC70YzRiCDQvdC1INC/0ZbQtNC10YLQtSDQtyDRhdC70L7Qv9GG0Y/QvNC4LgoK0JPQvtGA0ZbQu9CwINGB0L7RgdC90LAsINGB0LzQtdGA0LXQutCwLArQodC/0L7QtNC+0LHQsNCy0YHRjCDRhdC70L7Qv9C10YbRjCDQt9C00LDQu9C10LrQsC4KCtCh0L/QvtC00L7QsdCw0LLRgdGMINGF0LvQvtC/0LXRhtGMINGC0LAg0Lkg0L3QsNCy0ZbQuiwK0KLQtdC/0LXRgCDQstC20LUg0LLRltC9INC80ZbQuSDRh9C+0LvQvtCy0ZbQui4="
  
```

and `.travis.yml` will look like this...

```yaml
...
jobs:
  include:
    - stage: deploy
      name: "Deploy to PyPi"
      python: "3.7"
      script: skip
      provider: pypi
        user: butuzov
        password:
          secure: "0JPQvtGA0ZbQu9CwINGB0L7RgdC90LAsINC/0LDQu9Cw0LvQsCwK0J/RltC0\
                   INC90LXRjiDQtNGW0LLQutCwINGB0YLQvtGP0LvQsC4KCtCf0ZbQtCDQvdC1\
                   0Y4g0LTRltCy0LrQsCDRgdGC0L7Rj9C70LAsCtCg0YPRgdGP0LLRgyDQutC+\
                   0YHRgyDRh9C10YHQsNC70LAuCgrigJMg0J7QuSDQutC+0YHQuCwg0LrQvtGB\
                   0Lgg0LLQuCDQvNC+0ZcsCtCU0L7QstCz0L4g0YHQu9GD0LbQuNC70Lgg0LLQ\
                   uCDQvNC10L3Rli4KCtCR0ZbQu9GM0YjQtSDRgdC70YPQttC40YLRjCDQvdC1\
                   INCx0YPQtNC10YLQtSwK0J/RltC0INCx0ZbQu9C40Lkg0LLQtdC70YzQvtC9\
                   INC/0ZbQtNC10YLQtS4KCtCf0ZbQtCDQsdGW0LvQuNC5INCy0LXQu9GM0L7Q\
                   vSwg0L/RltC0INGF0YPRgdGC0LrRgywK0JHRltC70YzRiCDQvdC1INC/0ZbQ\
                   tNC10YLQtSDQt9CwINC00YDRg9C20LrRgy4KCtCf0ZbQtCDQsdGW0LvQuNC5\
                   INCy0LXQu9GM0L7QvSDQtyDQutGW0L3RhtGP0LzQuCwK0JHRltC70YzRiCDQ\
                   vdC1INC/0ZbQtNC10YLQtSDQtyDRhdC70L7Qv9GG0Y/QvNC4LgoK0JPQvtGA\
                   0ZbQu9CwINGB0L7RgdC90LAsINGB0LzQtdGA0LXQutCwLArQodC/0L7QtNC+\
                   0LHQsNCy0YHRjCDRhdC70L7Qv9C10YbRjCDQt9C00LDQu9C10LrQsC4KCtCh\
                   0L/QvtC00L7QsdCw0LLRgdGMINGF0LvQvtC/0LXRhtGMINGC0LAg0Lkg0L3Q\
                   sNCy0ZbQuiwK0KLQtdC/0LXRgCDQstC20LUg0LLRltC9INC80ZbQuSDRh9C+\
                   0LvQvtCy0ZbQui4="
      on:
        branch: master
        tags: true
```