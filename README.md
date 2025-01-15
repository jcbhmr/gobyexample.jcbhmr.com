<p align=center>
  <b>You're probably looking for <a href="https://gobyexample.jcbhmr.com/">gobyexample.jcbhmr.com</a></b>
</p>

# Go by Example

🔰 gobyexample.jcbhmr.com website \
🔀 Forked from [mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample)

<p align=center>
  <a href="https://gobyexample.jcbhmr.com/hello-world.html">
    <img src="https://github.com/user-attachments/assets/5845558f-9bc2-4074-a079-421edbbde149">
  </a>
</p>

🌎 Translations all in one place \
📘 Uses mdBook \
🔍 Searchable \
📱 Mobile-friendly \
🔗 Side panel for quick navigation \
💻 Code examples run in the browser!

## Development

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)
![mdBook](https://img.shields.io/static/v1?style=for-the-badge&message=mdBook&color=000000&logo=mdBook&logoColor=FFFFFF&label=)

You'll need:

- [Go](https://golang.org/dl/)
- [mdBook](https://rust-lang.github.io/mdBook/guide/installation.html)

Each language is its own mdBook project. You can start the development server with `mdbook serve <dir>` where `<dir>` is the language directory.

```sh
mdbook serve # English
mdbook serve ./fr/ # French
```

Every so often remember to run [`mdbook test`](https://rust-lang.github.io/mdBook/cli/test.html) but for `go` code blocks:

```sh
./scripts/mdbook-test-go.go # English
./scripts/mdbook-test-go.go ./fr/ # French
```

You can preview the full multilingual site by first building each language (`mdbook build <dir>`) and then serving the `book/` directory with your favorite static file server.

```sh
mdbook build # English
mdbook build ./fr/ # French
# ...
python -m http.server -d ./book/
```

### Why fork?

- **The original is not mobile-friendly.** This is my #1 gripe. This website is mobile friendly. Thanks mdBook!
- The original doesn't let you run code examples like [Rust by Example](https://doc.rust-lang.org/rust-by-example/hello.html). That's a feature that I want. This website has it.
- The original doesn't have a side panel to quickly jump between pages; you have to go back to the index page. This website has that. There's also `<` and `>` buttons.
- The original doesn't have centralized community translations. This website houses all translations in one place.
- The original's build setup is custom & complicated. This website uses mdBook, mdbook-i18n-helpers, and GitHub Pages.

## License

Copyright (c) 2012 Mark McGranaghan \
Copyright (c) 2025 Jacob Hummer \
[Creative Commons Attribution 3.0 Unported License](https://github.com/jcbhmr/gobyexample.jcbhmr.com/blob/main/LICENSE)
