function setCookie(key, val) {
  expDays = 7;
  let date = new Date();
  date.setTime(date.getTime() + expDays * 24 * 60 * 60 * 1000);

  const exp = date.toUTCString();

  const cookie =
    escape(key) +
    "=" +
    escape(val) +
    ";SameSite=Lax;expires=" +
    exp +
    ";path=/";

  document.cookie = cookie;
  console.log(cookie);
  console.log(
    "New cookie with key: " + key + " value: " + val + " expiration: " + exp
  );
}

function getCookie(name) {
  let key = name + "=";
  let cookies = document.cookie.split(";");
  for (let i = 0; i < cookies.length; i++) {
    let cookie = cookies[i];
    while (cookie.charAt(0) === " ") {
      cookie = cookie.substring(1, cookie.length);
    }
    if (cookie.indexOf(key) === 0) {
      return cookie.substring(key.length, cookie.length);
    }
  }
  return null;
}

(() => {
  document.addEventListener("DOMContentLoaded", function () {
    var $this = {
      // el
      body: document.querySelector("body"),
      // indicators
      showTypographyGrid: getCookie("gt") == "true",
      showColoredGrid: getCookie("ct") == "true",
      // methods...
      grid_toggle: function () {
        bodyClasses = $this.body.classList;
        if (this.showTypographyGrid) {
          bodyClasses.add("typography-grid");
          return;
        }
        bodyClasses.remove("typography-grid");
      },

      colors_toggle: function () {
        bodyClasses = $this.body.classList;

        console.log(getCookie("ct"));

        if (this.showColoredGrid) {
          bodyClasses.add("colored");
          return;
        }
        bodyClasses.remove("colored");
      },
      grid: (e) => {
        if (!(e.keyCode == 219 && (e.ctrlKey || e.metaKey))) {
          return;
        }
        e.preventDefault();

        console.log("toggle typography");
        setCookie("gt", !$this.showTypographyGrid);
        $this.showTypographyGrid = getCookie("gt") == "true";
        $this.grid_toggle();
      },
      colors: (e) => {
        if (!(e.keyCode == 221 && (e.ctrlKey || e.metaKey))) {
          return;
        }

        console.log("toggle colors");
        e.preventDefault();
        setCookie("ct", !$this.showColoredGrid);
        $this.showColoredGrid = getCookie("ct") == "true";
        $this.colors_toggle();

        return;
      },
      init: () => {
        $this.colors_toggle();
        document.addEventListener("keydown", $this.colors);
        $this.grid_toggle();
        document.addEventListener("keydown", $this.grid);
      },
    };
    return $this.init();
  });
})();
