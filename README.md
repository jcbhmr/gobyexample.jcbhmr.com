<p align=center>
  <b>You're probably looking for <a href="https://gobyexample.jcbhmr.com/">gobyexample.jcbhmr.com</a></b>
</p>

# Go by Example

🔰 gobyexample.jcbhmr.com website \
🔀 Forked from [mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample)

<p align=center>
  <img src="https://github.com/user-attachments/assets/5845558f-9bc2-4074-a079-421edbbde149">
</p>

## Development

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)
![mdBook](https://img.shields.io/static/v1?style=for-the-badge&message=mdBook&color=000000&logo=mdBook&logoColor=FFFFFF&label=)

This project uses [mdBook](https://rust-lang.github.io/mdBook/) for the website itself.

Launch the dev server:

```sh
mdbook serve
```

Run the custom script that's like [`mdbook test`](https://rust-lang.github.io/mdBook/cli/test.html) but for `go` code blocks:

```sh
./scripts/mdbook-test-go.go
```

### Translations

Build the site with translations:

```sh
./scripts/mdbook-build-i18n.go
```

> [!TIP]
> Use a static HTTP server to preview a built site.
>
> ```sh
> go run github.com/eliben/static-server@latest ./book/
> ```

Create initial `po/<lang>.po` file:

```sh
# First need to generate a po/messages.pot file.
MDBOOK_OUTPUT='{"xgettext": {}}' mdbook build --dest-dir ./po/

# Now we can generate the po/<lang>.po file from the messages.pot file.
msginit --input ./po/messages.pot --locale <lang> --output ./po/<lang>.po
```

> [!TIP]
> Use Google Translate to pre-fill a `po/<lang>.po` file:
>
> ```sh
> cargo install cloud-translate
> cloud-translate <google-cloud-project-id> po/<lang>.po 30000
> ```

> [!WARNING]
> `mdbook-gettext` will use the original untranslated text for all entries marked as "fuzzy" (visible as "Needs work" in Poedit). This is especially important when using cloud-translate for initial translation as all entries will be marked as "fuzzy".
>
> If your text isn't translated, double-check that you have removed all "fuzzy" flags from your xx.po file.

Update an existing `po/<lang>.po` file:

```sh
# First generate a fresh up-to-date po/messages.pot file.
MDBOOK_OUTPUT='{"xgettext": {}}' mdbook build --dest-dir ./po/

# Now use the po/messages.pot file to update the po/<lang>.po file.
msgmerge --update ./po/<lang>.po ./po/messages.pot
```

Contributor translations of the Go by Example site are available in:

* [Chinese](https://gobyexample-cn.github.io/) by [gobyexample-cn](https://github.com/gobyexample-cn)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Italian](https://gobyexample.it) by the [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Japanese](http://spinute.org/go-by-example) by [spinute](https://github.com/spinute)
* [Korean](https://mingrammer.com/gobyexample/) by [mingrammer](https://github.com/mingrammer)
* [Russian](https://gobyexample.com.ru/) by [badkaktus](https://github.com/badkaktus)
* [Ukrainian](https://butuzov.github.io/gobyexample/) by [butuzov](https://github.com/butuzov/gobyexample)
* [Brazilian Portuguese](https://lcslitx.github.io/GoEmExemplos/) by [lcslitx](https://github.com/LCSLITX)
* [Vietnamese](https://gobyexample.vn/) by [s6k Gopher](https://github.com/s6k-gopher/gobyexample-vn)
* [Burmese](https://setkyar.github.io/gobyexample) by [Set Kyar Wa Lar](https://github.com/setkyar/gobyexample)

### Why fork mmcgrana/gobyexample?

- **It's not mobile-friendly.** This is my #1 gripe. This website is mobile friendly. Thanks mdBook!
- It doesn't let you run code examples like [Rust by Example](https://doc.rust-lang.org/rust-by-example/hello.html). That's a feature that I want. This website has it.
- It doesn't have a side panel to quickly jump between pages; you have to go back to the index page. This website has that. There's also `<` and `>` buttons.
- ~~It doesn't have centralized community translations.~~ TODO
- The build setup is custom & complicated. This website uses mdBook.

## License

Copyright (c) 2012 Mark McGranaghan \
Copyright (c) 2025 Jacob Hummer \
[Creative Commons Attribution 3.0 Unported License](https://github.com/jcbhmr/gobyexample.jcbhmr.com/blob/main/LICENSE).
