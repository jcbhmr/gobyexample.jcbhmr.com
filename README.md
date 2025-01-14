<p align=center>
  <b>You're probably looking for <a href="https://gobyexample.jcbhmr.com/">gobyexample.jcbhmr.com</a></b>
</p>

# Go by Example

🔰 gobyexample.jcbhmr.com website \
🔀 Forked from [mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample)

<p align=center>
  <img src="https://github.com/user-attachments/assets/5845558f-9bc2-4074-a079-421edbbde149">
</p>

<p align=center>
  <a href="https://gobyexample.jcbhmr.com/">Website</a>
  | <a href="https://crowdin.com/project/gobyexamplejcbhmrcom">Crowdin (translations)</a>
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
- [mdbook-i18n-helpers](https://github.com/google/mdbook-i18n-helpers) (requires [Rust](https://rustup.rs/))
- [gettext](https://gnu.org/software/gettext/) ([Debian](https://packages.debian.org/sid/gettext), [macOS](https://formulae.brew.sh/formula/gettext), [Windows](https://mlocati.github.io/articles/gettext-iconv-windows.html))
- A `.po` editor such as [pofile.net](https://pofile.net/free-po-editor) or [Poedit](https://poedit.net/download)

Launch the dev server:

```sh
./scripts/dev.sh
```

Every so often remember to run [`mdbook test`](https://rust-lang.github.io/mdBook/cli/test.html) but for `go` code blocks:

```sh
./scripts/mdbook-test-go.go
```

### Translations

Create new translation:

```sh
./scripts/create-i18n.sh <lang>
```

> [!TIP]
> Use Google Translate to pre-fill a `po/<lang>.po` file. You'll need [cloud-translate](https://github.com/mgeisler/cloud-translate) and [gcloud](https://cloud.google.com/sdk/docs/install).
>
> ```sh
> cloud-translate <google-cloud-project-id> ./po/<lang>.po 30000
> ```

> [!WARNING]
> `mdbook-gettext` will use the original untranslated text for all entries marked as "fuzzy" (visible as "Needs work" in Poedit). This is especially important when using cloud-translate for initial translation as all entries will be marked as "fuzzy".
>
> If your text isn't translated, double-check that you have removed all "fuzzy" flags from your xx.po file.

Update all translations with content from `src/`:

```sh
./scripts/update-i18n.sh
```

### Why fork?

- **The original is not mobile-friendly.** This is my #1 gripe. This website is mobile friendly. Thanks mdBook!
- The original doesn't let you run code examples like [Rust by Example](https://doc.rust-lang.org/rust-by-example/hello.html). That's a feature that I want. This website has it.
- The original doesn't have a side panel to quickly jump between pages; you have to go back to the index page. This website has that. There's also `<` and `>` buttons.
- ~~The original doesn't have centralized community translations.~~ TODO
- The original's build setup is custom & complicated. This website uses mdBook, mdbook-i18n-helpers, and GitHub Pages.

## License

Copyright (c) 2012 Mark McGranaghan \
Copyright (c) 2025 Jacob Hummer \
[Creative Commons Attribution 3.0 Unported License](https://github.com/jcbhmr/gobyexample.jcbhmr.com/blob/main/LICENSE)
