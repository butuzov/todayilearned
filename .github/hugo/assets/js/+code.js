// ---------------------------------------------------------------------------
// Inserts copy button into each code node.
// ---------------------------------------------------------------------------

(() => {
  document.querySelectorAll("pre > code").forEach(function (codeBlock) {
    console.log(codeBlock.parentNode.parentNode.className);
    if (codeBlock.parentNode.parentNode.className == "highlight") {
      return;
    }

    var div = document.createElement("div");
    div.className = "highlight";
    div.innerHTML = codeBlock.parentNode.outerHTML;
    codeBlock.parentNode.parentNode.insertBefore(div, codeBlock.parentNode);
    codeBlock.remove();
  });

  const addCopyButtons = (clipboard) => {
    document.querySelectorAll("pre > code").forEach(function (codeBlock) {
      var div = document.createElement("div");
      div.className = "bd-clipboard";

      var button = document.createElement("button");
      button.className = "copy-code-button";
      button.type = "button";
      button.innerText = "Copy";
      button.addEventListener("click", () => {
        let code = codeBlock.innerText.replaceAll("\n\n", "\n");
        clipboard.writeText(code).then(
          () => {
            button.blur();
            button.innerText = "Copied!";
            setTimeout(() => {
              button.innerText = "Copy";
            }, 2000);
          },
          (error) => {
            console.log(error);
            button.innerText = "Error";
          }
        );
      });

      div.append(button);

      var pre = codeBlock.parentNode;
      pre.parentNode.insertBefore(div, pre);
    });
  };

  // Initiating clipboarding.
  if (navigator && navigator.clipboard) {
    addCopyButtons(navigator.clipboard);
  } else {
    var script = document.createElement("script");
    script.src =
      "https://cdnjs.cloudflare.com/ajax/libs/clipboard-polyfill/2.7.0/clipboard-polyfill.promise.js";
    script.integrity = "sha256-waClS2re9NUbXRsryKoof+F9qc1gjjIhc2eT7ZbIv94=";
    script.crossOrigin = "anonymous";
    script.onload = function () {
      addCopyButtons(clipboard);
    };
    document.body.appendChild(script);
  }
})();
