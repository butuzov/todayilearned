var app = ((debug) => {
  "use strict";

  document.addEventListener("DOMContentLoaded", function () {
    var $this = {
      debug: debug,
      index: null,
      modal: null, // bootraped modal
      searchModal: document.getElementById("searchModal"),
      searchInput: document.getElementById("search-input"),
      _foo: null,

      // Indexing
      isIndexing: false,
      db: [],

      _bar: () => {},

      // empty - is creating empty results lists.
      empty: () => {
        var $board = document.createElement("div");
        $board.classList.add("blankslate");
        var $h2 = document.createElement("h2");
        $h2.classList.add("blankslate-heading");
        $h2.innerHTML = `No results for "${input}"`;
        $board.appendChild($h2);
        return $board;
      },

      highlight: (element, tokens) => {
        const nodeFilter = {
          acceptNode: function (node) {
            return NodeFilter.FILTER_ACCEPT;
          },
        };
        const walker = document.createTreeWalker(
          element,
          NodeFilter.SHOW_TEXT,
          nodeFilter,
          false
        );
        let node = null;

        while ((node = walker.nextNode())) {
          const text = node.textContent.toLowerCase();
          let found = false;
          for (var i = 0; i < tokens.length && !found; i++) {
            const token = tokens[i].toString();
            const startIndex = text.indexOf(token);
            if (startIndex == -1) {
              continue;
            }
            let range = document.createRange();
            range.setStart(node, startIndex);
            range.setEnd(node, startIndex + token.length);
            let mark = document.createElement("mark");
            range.surroundContents(mark);
            found = true;
          }
          walker.nextNode();
        }
      },

      print: (...theArgs) => {
        if ($this.debug) {
          console.log("[Search]", ...theArgs);
        }
      },

      // opens modal window.
      open: (searchTerm = null) => {
        if (!$this.modal) {
          $this.modal = new bootstrap.Modal($this.searchModal, {
            focus: true,
            keyboard: true,
          });
        }
        setTimeout(() => {
          if (searchTerm) {
            $this.searchInput.value = searchTerm;
          }
          $this.searchInput.focus();
        }, 500);
        $this.modal.show();
      },

      activateSearchShortcut: (e) => {
        // reacts to cmd+k or ctrl+k
        if (e.keyCode != 75 || !(e.ctrlKey || e.metaKey)) {
          return;
        }
        $this.open();
        e.preventDefault();
      },

      createEmptyResultElement: (input) => {
        var $blankslate = document.createElement("div");
        $blankslate.classList.add("blankslate");
        var $h2 = document.createElement("h2");
        $h2.classList.add("blankslate-heading");
        $h2.innerHTML = `No results for "<mark>${input}</mark>"`;
        $blankslate.appendChild($h2);
        return $blankslate;
      },

      pipeline_RequestIndex: async () => {
        // Requests index.json
        return await fetch("/index.json", {
          cache: "force-cache",
        }).then((resp) => {
          return resp.json();
        });
      },

      // Persisting decoded json entries
      pipeline_Persist: (items) => {
        if (!$this.db.length) {
          $this.print("persisting", items.length);
          $this.db = items;
        }
        return;
      },

      // Creates lunr index
      pipeline_Index: () => {
        if ($this.index) {
          $this.print("index exists - skipping...");
          return;
        }
        $this.print("index start");
        var builder = new lunr.Builder();
        builder.ref("uri");
        builder.field("title", { boost: 10 });
        builder.field("tags", { boost: 5 });
        builder.field("content");
        $this.db.forEach((doc) => {
          builder.add(doc);
        });
        $this.index = builder.build();
        $this.print("index done");
      },

      //  Actual Search and showing results
      pipeline_Search: () => {
        $this.print("searching...");
        let searchTerm = $this.searchInput.value.trim();

        var results = $this.index.query(function (query) {
          query.term($this.index.pipeline.run(lunr.tokenizer(searchTerm)), {
            usePipeline: false,
            wildcard: lunr.Query.wildcard.TRAILING,
          });
        });
        $this.print("found", results.length);

        // - cleanup first
        var $listContainer = document.querySelector("#searchModal .modal-body");
        if (!results.length) {
          $listContainer.innerHTML =
            $this.createEmptyResultElement(searchTerm).outerHTML;
          return;
        }

        /// - perform search
        var $listGroup = document.createElement("div");
        $listGroup.classList.add("list-group");
        $listGroup.id = "docsearch-list";
        let tabIndex = 120;
        results.slice(0, 30).forEach((found) => {
          let doc = $this.db.filter((k) => k.uri === found.ref)[0];

          // div
          //	nav
          //	a > div.highlighting
          var $div = document.createElement("div");
          $div.tabIndex = tabIndex++;
          $div.classList.add("list-group-item", "list-group-item-action");

          var $nav = document.createElement("nav");
          let href = new URL(doc["uri"]).pathname.split("/");
          $nav.innerHTML = href.slice(1, href.length - 2).join(" / ");
          $div.appendChild($nav);

          // a
          var $item = document.createElement("a");
          $item.href = doc["uri"];
          $item.id = `docsearch-item-${doc["objectID"]}`;

          // a.innerDiv
          let $title = document.createElement("span");
          $title.innerText = doc["title"];
          $this.highlight($title, lunr.tokenizer(searchTerm));
          $item.appendChild($title);

          $div.appendChild($item);

          $listGroup.appendChild($div);
        });
        $listContainer.innerHTML = $listGroup.outerHTML;
        $this.print("done", $listContainer.innerHTML);
      },

      modalNavigationInput: (e) => {
        const left = e.key == "ArrowLeft";
        const right = e.key == "ArrowRight";
        if (!(left || right)) {
          return;
        }

        if (window.getComputedStyle($this.searchModal).display == "none") {
          return;
        }

        $this.searchInput.focus();
      },
      //  EXPERIMENTAL: adds keyboard navigation to results.
      modalNavigationLinksActivate: (e) => {
        if (e.key != "Enter") {
          return;
        }

        if (document.activeElement.classList.contains("list-group-item")) {
          window.location = document.activeElement.querySelector("a").href;
        }
      },
      modalNavigationLinks: (e) => {
        const up = e.key == "ArrowUp";
        const down = e.key == "ArrowDown";
        if (!(up || down)) {
          return;
        }

        if (window.getComputedStyle($this.searchModal).display == "none") {
          return;
        }

        // $this.searchInput.focus();

        var current = false;
        var links = document.querySelectorAll("#searchModal .list-group-item");

        // runs only if input focused or any of selected links are focused.
        let inputActive = document.activeElement == $this.searchInput;
        let linksActive = Array.from(links).some((a, k) => {
          let found = a == document.activeElement;
          if (found) {
            current = k;
          }
          return found;
        });

        $this.print("current current", document.activeElement);
        $this.print("anything active", inputActive, linksActive);
        $this.print("found?", current);

        if (!(inputActive || linksActive)) {
          return;
        }

        if (inputActive && up) {
          current = links.length - 1;
          $this.print("input+up", current);
        } else if (inputActive && down) {
          current = 0;
          $this.print("input+down", current);
        } else if (linksActive && up) {
          current--;
          $this.print("links+up", current);
        } else if (linksActive && down) {
          current++;
          $this.print("links+down", current);
        }

        if (current < 0) {
          current = links.length - 1;
        }
        $this.print("Curernt", current);

        links[current % links.length].focus();
        $this.print(links[current % links.length], document.activeElement);
        $this.print("nope");
      },
      // search pipeline
      searchPipeline: () => {
        let pervious = $this.prev == $this.searchInput.value.trim();
        let empty = $this.searchInput.value.trim() == "";
        if (pervious || empty) {
          return;
        }
        $this.prev = $this.searchInput.value;

        try {
          new Pipeline()
            .pipe($this.pipeline_RequestIndex)
            .pipe($this.pipeline_Persist)
            .pipe($this.pipeline_Index)
            .pipe($this.pipeline_Search)
            .process();
        } catch (error) {
          $this.print("DEBUG", error);
        }
      },

      activate: () => {
        // $this.print("activate");
        $this.searchInput.addEventListener("keyup", $this.searchPipeline);
        window.addEventListener("keydown", $this.activateSearchShortcut); // Open Modal with shortcut
        window.addEventListener("keyup", $this.modalNavigationLinks);
        window.addEventListener("keyup", $this.modalNavigationLinksActivate);
        window.addEventListener("keyup", $this.modalNavigationInput);
      },

      live: (searchTerm) => {
        $this.open(searchTerm);
        window.setTimeout(() => {
          $this.searchInput.dispatchEvent(
            new KeyboardEvent("keyup", {
              code: "Enter",
              key: "Enter",
              charCode: 13,
              keyCode: 13,
              view: window,
              bubbles: true,
            })
          );
        }, 1200);
      },

      load: () => {
        $this.activate();
        // $this.live("y");
      },
    };

    return $this.load();
  });
})(false);
