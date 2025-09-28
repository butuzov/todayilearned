// ---------------------------------------------------------------------------
// Domain in link
// Customizes links on content area so it follow pattern domain + text/url.part
// for external domains.
//
// domain as prefix + suffix is text
// ---------------------------------------------------------------------------

var domains = (function () {
	document.addEventListener("DOMContentLoaded", function () {
		// console.log("loaded");
		var $this = {
			localLinks: ".doc-content a",
			githubDomain: "github.com",

			// -- GitHub
			github: (el, url, text) => {
				// console.log(el);

				let parts = url.pathname.split("/");
				let external = url.hostname;
				let hasUser =
					parts.length > 2 && parts[1] != "topics" && parts[2].length > 0;
				let textIsLink = false;

				let linkIsDuplicated = false;
				let textIsLinkGithub = false;
				try {
					let text_as_url = new URL(text);
					if (text_as_url.host == url.host) {
						textIsLinkGithub = true;
						linkIsDuplicated = url.toString() == text_as_url.toString();
					}
					textIsLink = true;
				} catch (error) { }

				console.log({
					hasUser: hasUser,
					parts: parts,
					text: text,
					textIsLink: textIsLink,
					textIsLinkGithub: textIsLinkGithub,
				});

				//  el.innerHTML = `${text} <code class="domain">@ ${external}</code>`;

				// case 1: text is link
				if (textIsLink && !textIsLinkGithub) {
					text = text.split("//")[1];
					el.innerHTML = `${text} <code class="domain">@ ${external}</code>`;
					return;
				}

				// case 2: text is not
				if (!textIsLink) {
					el.innerHTML = `${text} <code class="domain">@ ${external}</code>`;
					return;
				}

				// case 3: has github user, but both text and links are same
				// caseId = 4;
				text = url.pathname;

				if (hasUser && linkIsDuplicated) {
					//   caseId = 3;
					text = text.replace("/pull/", "#");
					text = text.replace("/issue/", "#");
					if (text.includes("commit")) {
						text = text.replace("/commit/", "@");
						text = text.replace(parts[4], parts[4].slice(0, 8));
					}
				}

				text = text.replace(/\/$/gm, "");
				text = text.replace(/^\//gm, "");
				if (text == "") {
					text = "home";
				}
				el.innerHTML = `${text} <code class="domain">@ ${external}</code>`;
			},

			default: (el, url, text) => {
				// if test is domain and same as link
				// then remove domain and trailing slash
				try {
					let text_as_url = new URL(text);
					if (text_as_url.host == url.host) {
						text = url.pathname;
						text = text.replace(/\/$/gm, "");
						text = text.replace(/^\//gm, "");
					} else {
						text = text.split("//")[1];
					}

					console.log(text)
					if (text == "") {
						el.innerHTML = `<code class="domain">${url.host}</code>`;
					} else {
						el.innerHTML = `${text} <code class="domain">@ ${url.host}</code>`;
					}

				} catch (error) {
					el.innerHTML = `${text} <code class="domain">@ ${url.host}</code>`;
					return;
				}
			},

			load: () => {
				document.querySelectorAll($this.localLinks).forEach((el) => {
					el.setAttribute("target", "_black");
					el.setAttribute("rel", "nofollow");

					if (el.classList.contains("github")) {
						return;
					}

					let url = new URL(el.href);
					let text = el.innerHTML;

					if (
						url.host == "butuzov.github.io" ||
						url.host.startsWith("localhost")
					) {
						return;
					}

					if (url.host == $this.githubDomain) {
						return $this.github(el, url, text);
					}

					return $this.default(el, url, text);
				});
			},
		};

		return $this.load();
	});
})();
