/*!
https://github.com/rust-lang/mdBook/blob/618a2fa78b29b2dd848cf6b074c9a1ecffa48d68/src/theme/book.js
Mozilla Public License, version 2.0
TODO: Update this file every so often to keep it in sync with the original.
CHANGES:
- Add a runGoCode() function to run Go code snippets.
- Where ".playground" is mentioned, branch on the ".language-<id>" to Rust or Go.
- Hydrate Go playgrounds as though they were also default playgrounds.
*/

"use strict";

// CHANGE: Hydrate Go playgrounds as though they were also default playgrounds.
(function goDefaultPlaygrounds() {
    const defaultPlaygrounds = [...document.querySelectorAll("pre")].filter(pre => {
        if (pre.querySelector(":scope > code.language-go.noplayground")) {
            return false;
        }
        if (pre.querySelector(":scope > code.language-go.mdbook-runnable")) {
            return true;
        }
        if (pre.querySelector(":scope > code.language-go.ignore")) {
            return false;
        }
        if (pre.querySelector(":scope > code.language-go.no_run")) {
            return false;
        }
        return !!pre.querySelector(":scope > code.language-go");
    })
    // Rust playgrounds are `<pre><pre class="playground"><code class="language-rust editable">`. The normal
    // HTML only gives us one `<pre>`, so we need to wrap it again.
    for (const pre of defaultPlaygrounds) {
        pre.classList.add("playground");
        const newPre = document.createElement("pre");
        pre.replaceWith(newPre);
        newPre.append(pre);
    }
})();

// Fix back button cache problem
window.onunload = function () { };

// Global variable, shared between modules
function playground_text(playground, hidden = true) {
    let code_block = playground.querySelector("code");

    if (window.ace && code_block.classList.contains("editable")) {
        let editor = window.ace.edit(code_block);
        return editor.getValue();
    } else if (hidden) {
        return code_block.textContent;
    } else {
        return code_block.innerText;
    }
}

(function codeSnippets() {
    function fetch_with_timeout(url, options, timeout = 6000) {
        return Promise.race([
            fetch(url, options),
            new Promise((_, reject) => setTimeout(() => reject(new Error('timeout')), timeout))
        ]);
    }

    // CHANGE: Only do this for Rust playgrounds.
    const rustPlaygrounds = [...document.querySelectorAll(".playground")].filter(pre => !!pre.querySelector(":scope > code.language-rust"));
    if (rustPlaygrounds.length > 0) {
        fetch_with_timeout("https://play.rust-lang.org/meta/crates", {
            headers: {
                'Content-Type': "application/json",
            },
            method: 'POST',
            mode: 'cors',
        })
        .then(response => response.json())
        .then(response => {
            // get list of crates available in the rust playground
            let playground_crates = response.crates.map(item => item["id"]);
            rustPlaygrounds.forEach(block => handle_crate_list_update(block, playground_crates));
        });
    }
    // CHANGE: handleUpdateGo() which doesn't require fetching crates.
    const goPlaygrounds = [...document.querySelectorAll(".playground")].filter(pre => !!pre.querySelector(":scope > code.language-go"));
    if (goPlaygrounds.length > 0) {
        // Fake a fetch()-like delay so that this code happens AFTER all the below code runs.
        setTimeout(() => {
            for (const block of goPlaygrounds) {
                handleUpdateGo(block);
            }
        }, 10);
    }

    function handle_crate_list_update(playground_block, playground_crates) {
        // update the play buttons after receiving the response
        update_play_button(playground_block, playground_crates);

        // and install on change listener to dynamically update ACE editors
        if (window.ace) {
            let code_block = playground_block.querySelector("code");
            if (code_block.classList.contains("editable")) {
                let editor = window.ace.edit(code_block);
                editor.addEventListener("change", function (e) {
                    update_play_button(playground_block, playground_crates);
                });
                // add Ctrl-Enter command to execute rust code
                editor.commands.addCommand({
                    name: "run",
                    bindKey: {
                        win: "Ctrl-Enter",
                        mac: "Ctrl-Enter"
                    },
                    exec: _editor => run_rust_code(playground_block)
                });
            }
        }
    }

    // CHANGE: Add a handleUpdateGo() function which doesn't require fetching crates.
    function handleUpdateGo(pre) {
        updatePlayButtonGo(pre);
        if (window.ace) {
            let code_block = pre.querySelector("code");
            if (code_block.classList.contains("editable")) {
                let editor = window.ace.edit(code_block);
                editor.addEventListener("change", function (e) {
                    updatePlayButtonGo(pre);
                });
                editor.commands.addCommand({
                    name: "run",
                    bindKey: {
                        win: "Ctrl-Enter",
                        mac: "Ctrl-Enter"
                    },
                    exec: _editor => run_go_code(pre)
                });
            }
        }
    }

    // updates the visibility of play button based on `no_run` class and
    // used crates vs ones available on https://play.rust-lang.org
    function update_play_button(pre_block, playground_crates) {
        var play_button = pre_block.querySelector(".play-button");

        // skip if code is `no_run`
        if (pre_block.querySelector('code').classList.contains("no_run")) {
            play_button.classList.add("hidden");
            return;
        }

        // get list of `extern crate`'s from snippet
        var txt = playground_text(pre_block);
        var re = /extern\s+crate\s+([a-zA-Z_0-9]+)\s*;/g;
        var snippet_crates = [];
        var item;
        while (item = re.exec(txt)) {
            snippet_crates.push(item[1]);
        }

        // check if all used crates are available on play.rust-lang.org
        var all_available = snippet_crates.every(function (elem) {
            return playground_crates.indexOf(elem) > -1;
        });

        if (all_available) {
            play_button.classList.remove("hidden");
        } else {
            play_button.classList.add("hidden");
        }
    }
    
    // CHANGE: Add updatePlayButtonGo() which doesn't require an available crates list.
    function updatePlayButtonGo(pre) {
        const playButton = pre.querySelector(".play-button");
        if (pre.querySelector("code").classList.contains("no_run")) {
            playButton.classList.add("hidden");
        }
        playButton.classList.remove("hidden");
    }

    function run_rust_code(code_block) {
        var result_block = code_block.querySelector(".result");
        if (!result_block) {
            result_block = document.createElement('code');
            result_block.className = 'result hljs language-bash';

            code_block.append(result_block);
        }

        let text = playground_text(code_block);
        let classes = code_block.querySelector('code').classList;
        let edition = "2015";
        if(classes.contains("edition2018")) {
            edition = "2018";
        } else if(classes.contains("edition2021")) {
            edition = "2021";
        }
        var params = {
            version: "stable",
            optimize: "0",
            code: text,
            edition: edition
        };

        if (text.indexOf("#![feature") !== -1) {
            params.version = "nightly";
        }

        result_block.innerText = "Running...";

        fetch_with_timeout("https://play.rust-lang.org/evaluate.json", {
            headers: {
                'Content-Type': "application/json",
            },
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(params)
        })
        .then(response => response.json())
        .then(response => {
            if (response.result.trim() === '') {
                result_block.innerText = "No output";
                result_block.classList.add("result-no-output");
            } else {
                result_block.innerText = response.result;
                result_block.classList.remove("result-no-output");
            }
        })
        .catch(error => result_block.innerText = "Playground Communication: " + error.message);
    }

    // CHANGE: Add a runGoCode() function to run Go code snippets.
    // FIXME: This uses a CORS proxy. See if there's a way to avoid that.
    function runGoCode(codeBlock) {
        let resultBlock = codeBlock.querySelector(".result");
        if (!resultBlock) {
            resultBlock = document.createElement("code");
            resultBlock.className = "result hljs language-bash";
            codeBlock.append(resultBlock);
        }

        const text = playground_text(codeBlock);
        const params = {
            version: "2",
            body: text,
            withVet: "false",
        };

        resultBlock.innerText = "Running...";

        fetch_with_timeout("https://corsproxy.io/?url=https://go.dev/_/compile?backend=", {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
                "User-Agent": "gobyexample.jcbhmr.com",
            },
            method: "POST",
            mode: "cors",
            body: new URLSearchParams(params),
        })
        .then(response => response.json())
        .then(response => {
            let result = "";
            if (response.Errors) {
                result = "Errors:\n" + response.Errors;
            }
            if (response.VetErrors) {
                if (result) {
                    result += "\n";
                }
                result += "VetErrors:\n" + response.VetErrors;
            }
            if (response.Events) {
                if (result) {
                    result += "\nOutput:\n";
                }
                for (const event of response.Events) {
                    if (event.Kind === "stdout") {
                        result += event.Message;
                    } else if (event.Kind === "stderr") {
                        result += event.Message;
                    } else {
                        console.warn("Unknown event kind: %s", event.Kind, event);
                    }
                }
            }

            if (result.trim() === "") {
                resultBlock.innerText = "No output";
                resultBlock.classList.add("result-no-output");
            } else {
                resultBlock.innerText = result;
                resultBlock.classList.remove("result-no-output");
            }
        })
        .catch(error => resultBlock.innerText = "Playground Communication: " + error.message);
    }

    // Syntax highlighting Configuration
    hljs.configure({
        tabReplace: '    ', // 4 spaces
        languages: [],      // Languages used for auto-detection
    });

    let code_nodes = Array
        .from(document.querySelectorAll('code'))
        // Don't highlight `inline code` blocks in headers.
        .filter(function (node) {return !node.parentElement.classList.contains("header"); });

    // CHANGE: Stash "code.language-<id>" as "<pre class='playground' data-language='<id>'>".
    // This is needed because the language class is removed from the code block in editable blocks.
    for (const code of document.querySelectorAll("pre > code[class^='language-']")) {
        const pre = code.parentElement;
        const match = code.className.match(/language-(\S+)/);
        if (match) {
            pre.dataset.language = match[1];
        }
    }

    if (window.ace) {
        // language-rust class needs to be removed for editable
        // blocks or highlightjs will capture events
        code_nodes
            .filter(function (node) {return node.classList.contains("editable"); })
            .forEach(function (block) {
                block.classList.remove('language-rust');
                // CHANGE: Also remove the language-go class.
                block.classList.remove('language-go');
            });

        code_nodes
            .filter(function (node) {return !node.classList.contains("editable"); })
            .forEach(function (block) { hljs.highlightBlock(block); });
    } else {
        code_nodes.forEach(function (block) { hljs.highlightBlock(block); });
    }

    // Adding the hljs class gives code blocks the color css
    // even if highlighting doesn't apply
    code_nodes.forEach(function (block) { block.classList.add('hljs'); });

    Array.from(document.querySelectorAll("code.hljs")).forEach(function (block) {

        var lines = Array.from(block.querySelectorAll('.boring'));
        // If no lines were hidden, return
        if (!lines.length) { return; }
        block.classList.add("hide-boring");

        var buttons = document.createElement('div');
        buttons.className = 'buttons';
        buttons.innerHTML = "<button class=\"fa fa-eye\" title=\"Show hidden lines\" aria-label=\"Show hidden lines\"></button>";

        // add expand button
        var pre_block = block.parentNode;
        pre_block.insertBefore(buttons, pre_block.firstChild);

        pre_block.querySelector('.buttons').addEventListener('click', function (e) {
            if (e.target.classList.contains('fa-eye')) {
                e.target.classList.remove('fa-eye');
                e.target.classList.add('fa-eye-slash');
                e.target.title = 'Hide lines';
                e.target.setAttribute('aria-label', e.target.title);

                block.classList.remove('hide-boring');
            } else if (e.target.classList.contains('fa-eye-slash')) {
                e.target.classList.remove('fa-eye-slash');
                e.target.classList.add('fa-eye');
                e.target.title = 'Show hidden lines';
                e.target.setAttribute('aria-label', e.target.title);

                block.classList.add('hide-boring');
            }
        });
    });

    if (window.playground_copyable) {
        Array.from(document.querySelectorAll('pre code')).forEach(function (block) {
            var pre_block = block.parentNode;
            if (!pre_block.classList.contains('playground')) {
                var buttons = pre_block.querySelector(".buttons");
                if (!buttons) {
                    buttons = document.createElement('div');
                    buttons.className = 'buttons';
                    pre_block.insertBefore(buttons, pre_block.firstChild);
                }

                var clipButton = document.createElement('button');
                clipButton.className = 'clip-button';
                clipButton.title = 'Copy to clipboard';
                clipButton.setAttribute('aria-label', clipButton.title);
                clipButton.innerHTML = '<i class=\"tooltiptext\"></i>';

                buttons.insertBefore(clipButton, buttons.firstChild);
            }
        });
    }

    // Process playground code blocks
    Array.from(document.querySelectorAll(".playground")).forEach(function (pre_block) {
        // Add play button
        var buttons = pre_block.querySelector(".buttons");
        if (!buttons) {
            buttons = document.createElement('div');
            buttons.className = 'buttons';
            pre_block.insertBefore(buttons, pre_block.firstChild);
        }

        var runCodeButton = document.createElement('button');
        runCodeButton.className = 'fa fa-play play-button';
        runCodeButton.hidden = true;
        runCodeButton.title = 'Run this code';
        runCodeButton.setAttribute('aria-label', runCodeButton.title);

        buttons.insertBefore(runCodeButton, buttons.firstChild);
        runCodeButton.addEventListener('click', function (e) {
            // CHANGE: Switch on "data-language" that was stashed above.
            const language = pre_block.dataset.language;
            if (language === "go") {
                runGoCode(pre_block);
            } else {
                // Fall back to Rust.
                run_rust_code(pre_block);    
            }
        });

        if (window.playground_copyable) {
            var copyCodeClipboardButton = document.createElement('button');
            copyCodeClipboardButton.className = 'clip-button';
            copyCodeClipboardButton.innerHTML = '<i class="tooltiptext"></i>';
            copyCodeClipboardButton.title = 'Copy to clipboard';
            copyCodeClipboardButton.setAttribute('aria-label', copyCodeClipboardButton.title);

            buttons.insertBefore(copyCodeClipboardButton, buttons.firstChild);
        }

        let code_block = pre_block.querySelector("code");
        if (window.ace && code_block.classList.contains("editable")) {
            var undoChangesButton = document.createElement('button');
            undoChangesButton.className = 'fa fa-history reset-button';
            undoChangesButton.title = 'Undo changes';
            undoChangesButton.setAttribute('aria-label', undoChangesButton.title);

            buttons.insertBefore(undoChangesButton, buttons.firstChild);

            undoChangesButton.addEventListener('click', function () {
                let editor = window.ace.edit(code_block);
                editor.setValue(editor.originalCode);
                editor.clearSelection();
            });
        }
    });
})();

(function themes() {
    var html = document.querySelector('html');
    var themeToggleButton = document.getElementById('theme-toggle');
    var themePopup = document.getElementById('theme-list');
    var themeColorMetaTag = document.querySelector('meta[name="theme-color"]');
    var themeIds = [];
    themePopup.querySelectorAll('button.theme').forEach(function (el) {
        themeIds.push(el.id);
    });
    var stylesheets = {
        ayuHighlight: document.querySelector("[href$='ayu-highlight.css']"),
        tomorrowNight: document.querySelector("[href$='tomorrow-night.css']"),
        highlight: document.querySelector("[href$='highlight.css']"),
    };

    function showThemes() {
        themePopup.style.display = 'block';
        themeToggleButton.setAttribute('aria-expanded', true);
        themePopup.querySelector("button#" + get_theme()).focus();
    }

    function updateThemeSelected() {
        themePopup.querySelectorAll('.theme-selected').forEach(function (el) {
            el.classList.remove('theme-selected');
        });
        themePopup.querySelector("button#" + get_theme()).classList.add('theme-selected');
    }

    function hideThemes() {
        themePopup.style.display = 'none';
        themeToggleButton.setAttribute('aria-expanded', false);
        themeToggleButton.focus();
    }

    function get_theme() {
        var theme;
        try { theme = localStorage.getItem('mdbook-theme'); } catch (e) { }
        if (theme === null || theme === undefined || !themeIds.includes(theme)) {
            return default_theme;
        } else {
            return theme;
        }
    }

    function set_theme(theme, store = true) {
        let ace_theme;

        if (theme == 'coal' || theme == 'navy') {
            stylesheets.ayuHighlight.disabled = true;
            stylesheets.tomorrowNight.disabled = false;
            stylesheets.highlight.disabled = true;

            ace_theme = "ace/theme/tomorrow_night";
        } else if (theme == 'ayu') {
            stylesheets.ayuHighlight.disabled = false;
            stylesheets.tomorrowNight.disabled = true;
            stylesheets.highlight.disabled = true;
            ace_theme = "ace/theme/tomorrow_night";
        } else {
            stylesheets.ayuHighlight.disabled = true;
            stylesheets.tomorrowNight.disabled = true;
            stylesheets.highlight.disabled = false;
            ace_theme = "ace/theme/dawn";
        }

        setTimeout(function () {
            themeColorMetaTag.content = getComputedStyle(document.documentElement).backgroundColor;
        }, 1);

        if (window.ace && window.editors) {
            window.editors.forEach(function (editor) {
                editor.setTheme(ace_theme);
            });
        }

        var previousTheme = get_theme();

        if (store) {
            try { localStorage.setItem('mdbook-theme', theme); } catch (e) { }
        }

        html.classList.remove(previousTheme);
        html.classList.add(theme);
        updateThemeSelected();
    }

    // Set theme
    var theme = get_theme();

    set_theme(theme, false);

    themeToggleButton.addEventListener('click', function () {
        if (themePopup.style.display === 'block') {
            hideThemes();
        } else {
            showThemes();
        }
    });

    themePopup.addEventListener('click', function (e) {
        var theme;
        if (e.target.className === "theme") {
            theme = e.target.id;
        } else if (e.target.parentElement.className === "theme") {
            theme = e.target.parentElement.id;
        } else {
            return;
        }
        set_theme(theme);
    });

    themePopup.addEventListener('focusout', function(e) {
        // e.relatedTarget is null in Safari and Firefox on macOS (see workaround below)
        if (!!e.relatedTarget && !themeToggleButton.contains(e.relatedTarget) && !themePopup.contains(e.relatedTarget)) {
            hideThemes();
        }
    });

    // Should not be needed, but it works around an issue on macOS & iOS: https://github.com/rust-lang/mdBook/issues/628
    document.addEventListener('click', function(e) {
        if (themePopup.style.display === 'block' && !themeToggleButton.contains(e.target) && !themePopup.contains(e.target)) {
            hideThemes();
        }
    });

    document.addEventListener('keydown', function (e) {
        if (e.altKey || e.ctrlKey || e.metaKey || e.shiftKey) { return; }
        if (!themePopup.contains(e.target)) { return; }

        switch (e.key) {
            case 'Escape':
                e.preventDefault();
                hideThemes();
                break;
            case 'ArrowUp':
                e.preventDefault();
                var li = document.activeElement.parentElement;
                if (li && li.previousElementSibling) {
                    li.previousElementSibling.querySelector('button').focus();
                }
                break;
            case 'ArrowDown':
                e.preventDefault();
                var li = document.activeElement.parentElement;
                if (li && li.nextElementSibling) {
                    li.nextElementSibling.querySelector('button').focus();
                }
                break;
            case 'Home':
                e.preventDefault();
                themePopup.querySelector('li:first-child button').focus();
                break;
            case 'End':
                e.preventDefault();
                themePopup.querySelector('li:last-child button').focus();
                break;
        }
    });
})();

(function sidebar() {
    var body = document.querySelector("body");
    var sidebar = document.getElementById("sidebar");
    var sidebarLinks = document.querySelectorAll('#sidebar a');
    var sidebarToggleButton = document.getElementById("sidebar-toggle");
    var sidebarResizeHandle = document.getElementById("sidebar-resize-handle");
    var firstContact = null;

    function showSidebar() {
        body.classList.remove('sidebar-hidden')
        body.classList.add('sidebar-visible');
        Array.from(sidebarLinks).forEach(function (link) {
            link.setAttribute('tabIndex', 0);
        });
        sidebarToggleButton.setAttribute('aria-expanded', true);
        sidebar.setAttribute('aria-hidden', false);
        try { localStorage.setItem('mdbook-sidebar', 'visible'); } catch (e) { }
    }

    function hideSidebar() {
        body.classList.remove('sidebar-visible')
        body.classList.add('sidebar-hidden');
        Array.from(sidebarLinks).forEach(function (link) {
            link.setAttribute('tabIndex', -1);
        });
        sidebarToggleButton.setAttribute('aria-expanded', false);
        sidebar.setAttribute('aria-hidden', true);
        try { localStorage.setItem('mdbook-sidebar', 'hidden'); } catch (e) { }
    }

    // Toggle sidebar
    sidebarToggleButton.addEventListener('click', function sidebarToggle() {
        if (body.classList.contains("sidebar-hidden")) {
            var current_width = parseInt(
                document.documentElement.style.getPropertyValue('--sidebar-width'), 10);
            if (current_width < 150) {
                document.documentElement.style.setProperty('--sidebar-width', '150px');
            }
            showSidebar();
        } else if (body.classList.contains("sidebar-visible")) {
            hideSidebar();
        } else {
            if (getComputedStyle(sidebar)['transform'] === 'none') {
                hideSidebar();
            } else {
                showSidebar();
            }
        }
    });

    sidebarResizeHandle.addEventListener('mousedown', initResize, false);

    function initResize(e) {
        window.addEventListener('mousemove', resize, false);
        window.addEventListener('mouseup', stopResize, false);
        body.classList.add('sidebar-resizing');
    }
    function resize(e) {
        var pos = (e.clientX - sidebar.offsetLeft);
        if (pos < 20) {
            hideSidebar();
        } else {
            if (body.classList.contains("sidebar-hidden")) {
                showSidebar();
            }
            pos = Math.min(pos, window.innerWidth - 100);
            document.documentElement.style.setProperty('--sidebar-width', pos + 'px');
        }
    }
    //on mouseup remove windows functions mousemove & mouseup
    function stopResize(e) {
        body.classList.remove('sidebar-resizing');
        window.removeEventListener('mousemove', resize, false);
        window.removeEventListener('mouseup', stopResize, false);
    }

    document.addEventListener('touchstart', function (e) {
        firstContact = {
            x: e.touches[0].clientX,
            time: Date.now()
        };
    }, { passive: true });

    document.addEventListener('touchmove', function (e) {
        if (!firstContact)
            return;

        var curX = e.touches[0].clientX;
        var xDiff = curX - firstContact.x,
            tDiff = Date.now() - firstContact.time;

        if (tDiff < 250 && Math.abs(xDiff) >= 150) {
            if (xDiff >= 0 && firstContact.x < Math.min(document.body.clientWidth * 0.25, 300))
                showSidebar();
            else if (xDiff < 0 && curX < 300)
                hideSidebar();

            firstContact = null;
        }
    }, { passive: true });
})();

(function chapterNavigation() {
    document.addEventListener('keydown', function (e) {
        if (e.altKey || e.ctrlKey || e.metaKey || e.shiftKey) { return; }
        if (window.search && window.search.hasFocus()) { return; }
        var html = document.querySelector('html');

        function next() {
            var nextButton = document.querySelector('.nav-chapters.next');
            if (nextButton) {
                window.location.href = nextButton.href;
            }
        }
        function prev() {
            var previousButton = document.querySelector('.nav-chapters.previous');
            if (previousButton) {
                window.location.href = previousButton.href;
            }
        }
        switch (e.key) {
            case 'ArrowRight':
                e.preventDefault();
                if (html.dir == 'rtl') {
                    prev();
                } else {
                    next();
                }
                break;
            case 'ArrowLeft':
                e.preventDefault();
                if (html.dir == 'rtl') {
                    next();
                } else {
                    prev();
                }
                break;
        }
    });
})();

(function clipboard() {
    var clipButtons = document.querySelectorAll('.clip-button');

    function hideTooltip(elem) {
        elem.firstChild.innerText = "";
        elem.className = 'clip-button';
    }

    function showTooltip(elem, msg) {
        elem.firstChild.innerText = msg;
        elem.className = 'clip-button tooltipped';
    }

    var clipboardSnippets = new ClipboardJS('.clip-button', {
        text: function (trigger) {
            hideTooltip(trigger);
            let playground = trigger.closest("pre");
            return playground_text(playground, false);
        }
    });

    Array.from(clipButtons).forEach(function (clipButton) {
        clipButton.addEventListener('mouseout', function (e) {
            hideTooltip(e.currentTarget);
        });
    });

    clipboardSnippets.on('success', function (e) {
        e.clearSelection();
        showTooltip(e.trigger, "Copied!");
    });

    clipboardSnippets.on('error', function (e) {
        showTooltip(e.trigger, "Clipboard error!");
    });
})();

(function scrollToTop () {
    var menuTitle = document.querySelector('.menu-title');

    menuTitle.addEventListener('click', function () {
        document.scrollingElement.scrollTo({ top: 0, behavior: 'smooth' });
    });
})();

(function controllMenu() {
    var menu = document.getElementById('menu-bar');

    (function controllPosition() {
        var scrollTop = document.scrollingElement.scrollTop;
        var prevScrollTop = scrollTop;
        var minMenuY = -menu.clientHeight - 50;
        // When the script loads, the page can be at any scroll (e.g. if you reforesh it).
        menu.style.top = scrollTop + 'px';
        // Same as parseInt(menu.style.top.slice(0, -2), but faster
        var topCache = menu.style.top.slice(0, -2);
        menu.classList.remove('sticky');
        var stickyCache = false; // Same as menu.classList.contains('sticky'), but faster
        document.addEventListener('scroll', function () {
            scrollTop = Math.max(document.scrollingElement.scrollTop, 0);
            // `null` means that it doesn't need to be updated
            var nextSticky = null;
            var nextTop = null;
            var scrollDown = scrollTop > prevScrollTop;
            var menuPosAbsoluteY = topCache - scrollTop;
            if (scrollDown) {
                nextSticky = false;
                if (menuPosAbsoluteY > 0) {
                    nextTop = prevScrollTop;
                }
            } else {
                if (menuPosAbsoluteY > 0) {
                    nextSticky = true;
                } else if (menuPosAbsoluteY < minMenuY) {
                    nextTop = prevScrollTop + minMenuY;
                }
            }
            if (nextSticky === true && stickyCache === false) {
                menu.classList.add('sticky');
                stickyCache = true;
            } else if (nextSticky === false && stickyCache === true) {
                menu.classList.remove('sticky');
                stickyCache = false;
            }
            if (nextTop !== null) {
                menu.style.top = nextTop + 'px';
                topCache = nextTop;
            }
            prevScrollTop = scrollTop;
        }, { passive: true });
    })();
    (function controllBorder() {
        function updateBorder() {
            if (menu.offsetTop === 0) {
                menu.classList.remove('bordered');
            } else {
                menu.classList.add('bordered');
            }
        }
        updateBorder();
        document.addEventListener('scroll', updateBorder, { passive: true });
    })();
})();


// CHANGE: Put this in book.js so that we don't have to copy it to "./theme/custom.js" for each other translation.
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