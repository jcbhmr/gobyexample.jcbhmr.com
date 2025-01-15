(function languageSwitcher() {
    const rightButtons = document.querySelector(".right-buttons");
    const languageLabels = {
        __proto__: null,
        "en": "English",
        "zh": "中文",
        "fr": "Français",
        "it": "Italiano",
        "ja": "日本語",
        "ko": "한국어",
        "ru": "Русский",
        "uk": "Українська",
        "pt-BR": "Português do Brasil",
        "vi": "Tiếng Việt",
        "my": "Bahasa Melayu",
    };
    const altHrefs = {
        __proto__: null,
    }
    const currentLanguage = document.documentElement.lang || "en";
    for (const lang of Object.keys(languageLabels)) {
        const altPrefix = lang === "en" ? "/" : `/${lang}/`;
        const altHref = new URL(location);
        const re = /^\/((?<grandfathered>(en-GB-oed|i-ami|i-bnn|i-default|i-enochian|i-hak|i-klingon|i-lux|i-mingo|i-navajo|i-pwn|i-tao|i-tay|i-tsu|sgn-BE-FR|sgn-BE-NL|sgn-CH-DE)|(art-lojban|cel-gaulish|no-bok|no-nyn|zh-guoyu|zh-hakka|zh-min|zh-min-nan|zh-xiang))|((?<language>([A-Za-z]{2,3}(-(?<extlang>[A-Za-z]{3}(-[A-Za-z]{3}){0,2}))?)|[A-Za-z]{4}|[A-Za-z]{5,8})(-(?<script>[A-Za-z]{4}))?(-(?<region>[A-Za-z]{2}|[0-9]{3}))?(-(?<variant>[A-Za-z0-9]{5,8}|[0-9][A-Za-z0-9]{3}))*(-(?<extension>[0-9A-WY-Za-wy-z](-[A-Za-z0-9]{2,8})+))*(-(?<privateUse>x(-[A-Za-z0-9]{1,8})+))?)|(?<privateUse1>x(-[A-Za-z0-9]{1,8})+))\//;
        if (re.test(altHref.pathname)) {
            altHref.pathname = altHref.pathname.replace(re, altPrefix);
        } else {
            altHref.pathname = altHref.pathname.replace(/^\//, altPrefix);
        }
        altHrefs[lang] = altHref.href;
    }
    for (const [lang, altHref] of Object.entries(altHrefs)) {
        document.head.insertAdjacentHTML("beforeend", `<link rel="alternate" hreflang="${lang}" href="${altHref}" />`);
    }
    document.head.insertAdjacentHTML("beforeend", `<link rel="alternate" hreflang="x-default" href="${altHrefs["en"]}" />`);
    rightButtons.insertAdjacentHTML("afterbegin", `
        <button id="language-toggle" class="icon-button" type="button"
            title="Change language" aria-label="Change language"
            aria-haspopup="true" aria-expanded="false" aria-controls="language-list"
        >
            <i class="fa fa-globe"></i>
        </button>
        <ul id="language-list" class="theme-popup" aria-label="languages" role="menu" style="display: none;">
            ${Object.entries(languageLabels).map(([lang, label]) => {
                return `<li role="none">
                    <button role="menuitem" class="theme ${lang === currentLanguage ? "theme-selected" : ""}">
                        <a id="${lang}" href="${altHrefs[lang]}">${label}</a>
                    </button>
                </li>`;
            }).join("\n")}
        </ul>
        <style>
            #language-list {
                left: auto;
                right: 90px;
            }
            [dir="rtl"] #language-list {
                left: 10px;
                right: auto;
            }
            #language-list a {
                color: inherit;
            }
        </style>
    `);
    requestAnimationFrame(() => {
        let langToggle = document.getElementById("language-toggle");
        let langList = document.getElementById("language-list");
        langToggle.addEventListener("click", (event) => {
            langList.style.display = langList.style.display == "block" ? "none" : "block";
        });

        // When the user clicks a list item, the page jump is performed, just like clicking the internal <a> tag.
        langList.querySelectorAll("li").forEach(function(li) {
            li.addEventListener("click", function(event) {
                event.preventDefault();

                let link = this.querySelector("a");
                if (link && window.location.href !== link.href) {
                    window.location.href = link.href;
                }
            });
        });
    });
})();