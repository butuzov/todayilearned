<!-- menu: vs-code -->

# Visual Studio Code


## Settings

...


### Local Settings `.vscode/settings.json`

```json
{
    // Ignoring (Hiding Files in the File Tree)
    "files.exclude": {
        "**/_index.md": true
    }
}
```

## Shortcodes

...


## Running `tasks.json`

* `tasks.json` - [Integrate with External Tools via Tasks](https://code.visualstudio.com/Docs/editor/tasks)
* [Variables Reference](https://code.visualstudio.com/docs/editor/variables-reference)


## Debugging

### Python In Container

- https://rednafi.com/python/debug_dockerized_apps_in_vscode/
- https://github.com/microsoft/debugpy
- [Debugging by attaching over a network connection](https://code.visualstudio.com/docs/python/debugging#_debugging-by-attaching-over-a-network-connection)
- [Debug Python within a container](https://code.visualstudio.com/docs/containers/debug-python)

App Structure

- `__init__.py`
- `main.py`
- `requirements.txt`
- `Dockerfile`
- `docker-compose.yml`

#### Code

```yaml
# docker-compose.yml
version: "3.9"

services:
  web:
    build:
      context: .
      dockerfile: ./Dockerfile

    # This overrides the CMD in the Dockerfile
    command: uvicorn main:app --host 0.0.0.0 --port 8000
    ports:
      - 8000:8000
```
```Dockerfile
FROM python:3.12-slim-bookworm
WORKDIR /usr/src/app
COPY . /usr/src/app
RUN pip install --no-cache-dir -r requirements.txt
EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
```
```requirements.txt
starlette
uvicorn
```
```python
from starlette.applications import Starlette
from starlette.responses import JSONResponse
from starlette.routing import Route


async def homepage(request) -> JSONResponse:
    return JSONResponse({"hello": "world"})


app = Starlette(debug=True, routes=[Route("/", homepage)])
```

#### `launch.json`

This instructs the VSCode debugger to attach to a debug server running on localhost through port 5678. The next section will elaborate on how to run the debug server in a container.


- Click on the debugger button and then click on create a launch.json file.
- Select the Python debugger.
- Finally, select the Remote Attach debug config.

```json
{
   "version":"0.2.0",
   "configurations":[
      {
         "name":"Python: Remote Attach",
         "type":"python",
         "request":"attach",
         "connect":{
            "host":"localhost",
            "port":5678
         },
         "pathMappings":[
            {
               "localRoot":"${workspaceFolder}",
               "remoteRoot":"."
            }
         ],
         "justMyCode":true
      }
   ]
}
```
#### `docker-compose.debug.yml`

- `sh -c`: Selects the shell inside the Docker container.
- `pip install debugpy -t /tmp`: Installs the debugpy tool into the /tmp directory of the container.
- `python /tmp/debugpy --wait-for-client --listen 0.0.0.0:5678`: Runs debugpy, sets it to wait for a client connection and listen on all network interfaces at port 5678.
- `-m uvicorn main:app --host 0.0.0.0 --port 8000`: Starts an uvicorn server hosting the application defined in main:app, making it accessible on all network interfaces at port 8000.

```yaml
# docker-compose.debug.yml
version: "3.9"

services:
  web:
    build:
      context: .
      dockerfile: ./Dockerfile
    command:
      - "sh"
      - "-c"
      - |
        pip install debugpy -t /tmp \
        && python /tmp/debugpy --wait-for-client --listen 0.0.0.0:5678 \
        -m uvicorn main:app --host 0.0.0.0 --port 8000
    ports:
      - 8000:8000
      - 5678:5678
```


## Themes

#### Publish Theme Github Action

```yaml
name: Continuous Deployment

on:
  push:
    tags:
      - v*

jobs:
  release:
    name: VSCode Marketplace
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - run: make deps-ci-release
      - run: |
          git config --global user.email "<butuzov@users.noreply.github.com>"
          git config --global user.name "Oleg Butuzov"

      - name: Bump Version (package.json)
        run: |
          export VERSION_RELEASE=${GITHUB_REF/refs\/tags\/v/}
          cat <<< $(jq -r ".version = env.VERSION_RELEASE" ./package.json) > package.json
          git commit -a  -m "chore(release): ${VERSION_RELEASE}"
          git push origin HEAD:refs/heads/main


      - name: Create Package
        run: |
          export VERSION_RELEASE=${GITHUB_REF/refs\/tags\/v/}
          mkdir ./dist
          npm run publish:vsix

      # https://github.com/softprops/action-gh-release
      - name: Github Release
        uses: softprops/action-gh-release@v1
        if: startsWith(github.ref, 'refs/tags/')
        with:
          files: acid.vsix
          draft: true
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: VSC Release
        run: npm run publish:vsix
        env:
          VSCE_PAT: ${{ secrets.VSCE_PAT }}
```

