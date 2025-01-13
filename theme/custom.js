"use strict";

// playground_text() is provided by book.js which is loaded before this script.

(function () {
    var default_playgrounds = Array.from(document.querySelectorAll("pre"));
    default_playgrounds = default_playgrounds.filter(pre => {
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
        return pre.querySelector(":scope > code.language-go");
    });
    // Rust playgrounds are `<pre><pre class="playground"><code class="language-rust editable">`. The normal
    // HTML only gives us one `<pre>`, so we need to wrap it again.
    default_playgrounds.forEach(pre => {
        pre.classList.add("playground");
        var new_pre = document.createElement("pre");
        pre.replaceWith(new_pre);
        new_pre.append(pre);
    });
    // Since this code happens after book.js has already hydrated the playgrounds, we need to manually
    // do the work of book.js here again with checks to make sure we don't do it twice.

    // Undo application of a copy button that wouldn't have been added had it already had the .playground class.
    if (window.playground_copyable) {
        default_playgrounds.forEach(function (pre_block) {
            if (pre_block.classList.contains('playground')) {
                var buttons = pre_block.querySelector(".buttons");
                if (!buttons) {
                    buttons = document.createElement('div');
                    buttons.className = 'buttons';
                    pre_block.insertBefore(buttons, pre_block.firstChild);
                }

                var clipButton = pre_block.querySelector('button.clip-button');
                if (clipButton) {
                    clipButton.remove();
                }
            }
        });
    }

    function fetch_with_timeout(url, options, timeout = 6000) {
        return Promise.race([
            fetch(url, options),
            new Promise((_, reject) => setTimeout(() => reject(new Error('timeout')), timeout))
        ]);
    }

    var playgrounds = Array.from(document.querySelectorAll(".playground"));
    playgrounds = playgrounds.filter(playground => playground.querySelector(":scope > code.language-go"));
    if (playgrounds.length > 0) {
        setTimeout(() => {
            playgrounds.forEach(block => handle_update(block));
        }, 1);
    }

    function handle_update(playground_block) {
        update_play_button(playground_block);

        // install on change listener to dynamically update ACE editors
        if (window.ace) {
            let code_block = playground_block.querySelector("code");
            if (code_block.classList.contains("editable")) {
                let editor = window.ace.edit(code_block);
                editor.addEventListener("change", function (e) {
                    update_play_button(playground_block);
                });
                // add Ctrl-Enter command to execute go code
                editor.commands.addCommand({
                    name: "run",
                    bindKey: {
                        win: "Ctrl-Enter",
                        mac: "Ctrl-Enter"
                    },
                    exec: _editor => run_go_code(playground_block)
                });
            }
        }
    }

    function update_play_button(pre_block) {
        var play_button = pre_block.querySelector(".play-button");

        // skip if code is `no_run`
        if (pre_block.querySelector('code').classList.contains("no_run")) {
            play_button.classList.add("hidden");
            return;
        }

        console.log(play_button, pre_block.outerHTML);
        return
        play_button.classList.remove("hidden");
    }

    function run_go_code(code_block) {
        var result_block = code_block.querySelector(".result");
        if (!result_block) {
            result_block = document.createElement('code');
            result_block.className = 'result hljs language-bash';

            code_block.append(result_block);
        }

        let text = playground_text(code_block);
        let classes = code_block.querySelector('code').classList;
        var params = {
            version: "2",
            body: text,
            withVet: "true",
        };

        result_block.innerText = "Running...";

        // The https://go.dev/play/ backend doesn't support CORS? So we use a CORS proxy.
        // https://corsproxy.io/
        fetch_with_timeout("https://corsproxy.io/?url=https://go.dev/_/compile?backend=", {
            headers: {
                'Content-Type': "application/x-www-form-urlencoded; charset=UTF-8",
                'User-Agent': "gobyexample.jcbhmr.com",
            },
            method: 'POST',
            mode: 'cors',
            body: new URLSearchParams(params)
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
            if (result) {
                result += "\nOutput:\n";
            }
            if (response.Events) {
                for (let i = 0; i < response.Events.length; i++) {
                    let event = response.Events[i];
                    if (event.Kind === "stdout") {
                        result += event.Message;
                    } else if (event.Kind === "stderr") {
                        result += event.Message;
                    } else {
                        console.warn("Unknown event kind: " + event.Kind, event);
                    }
                }
            }

            if (result.trim() === '') {
                result_block.innerText = "No output";
                result_block.classList.add("result-no-output");
            } else {
                result_block.innerText = result;
                result_block.classList.remove("result-no-output");
            }
        })
        .catch(error => result_block.innerText = "Playground Communication: " + error.message);
    }

    let code_nodes = Array
        .from(document.querySelectorAll('code'))
        // Don't highlight `inline code` blocks in headers.
        .filter(function (node) {return !node.parentElement.classList.contains("header"); });

    if (window.ace) {
        // language-go class needs to be removed for editable
        // blocks or highlightjs will capture events
        code_nodes
            .filter(function (node) {return node.classList.contains("editable"); })
            .forEach(function (block) { block.classList.remove('language-go'); });
    }

    // Process playground code blocks
    Array.from(document.querySelectorAll(".playground")).forEach(function (pre_block) {
        if (!pre_block.querySelector(":scope > code.language-go")) {
            return;
        }

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
            run_go_code(pre_block);
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