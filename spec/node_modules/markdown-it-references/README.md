# `markdown-it-references`

> Ordered reference injection for [markdown-it](https://github.com/markdown-it/markdown-it).

<div>
  <p align="center">
    <img src="https://raw.githubusercontent.com/studyathome-internationally/markdown-it-plugins/master/packages/markdown-it-references/coverage/badge-branches.svg">
    <img src="https://raw.githubusercontent.com/studyathome-internationally/markdown-it-plugins/master/packages/markdown-it-references/coverage/badge-functions.svg">
    <img src="https://raw.githubusercontent.com/studyathome-internationally/markdown-it-plugins/master/packages/markdown-it-references/coverage/badge-lines.svg">
    <img src="https://raw.githubusercontent.com/studyathome-internationally/markdown-it-plugins/master/packages/markdown-it-references/coverage/badge-statements.svg">
    <a href="https://raw.githubusercontent.com/studyathome-internationally/markdown-it-plugins/master/packages/markdown-it-references/LICENSE" target="_blank">
      <img src="https://badgen.net/github/license/studyathome-internationally/markdown-it-plugins">
    </a>
  </p>
</div>

## Example

```md
# References

![Stormtroopocat](https://octodex.github.com/images/stormtroopocat.jpg "The Stormtroopocat")

<<the-stormtroopocat>> shows an example.
```

```html
<h1>References</h1>
<p>
  <figure id="the-stormtroopocat">
    <img src="https://octodex.github.com/images/stormtroopocat.jpg" alt="Stormtroopocat" title="The Stormtroopocat" />
    <figcaption>
      <a href="#the-stormtroopocat" class="anchor">§</a><a href="#the-stormtroopocat" class="label">Figure 1</a>: The
      Stormtroopocat
    </figcaption>
  </figure>
</p>
<p><a href="#the-stormtroopocat" class="figure-reference">Figure 1</a> shows an example.</p>
<h2 id="list-of-figures" class="list">List of Figures</h2>
<ol class="list">
  <li class="item"><a href="#the-stormtroopocat" class="label">Figure 1</a>: The Stormtroopocat</li>
</ol>
```

## Usage

Works with the following packages in conjunction:

- [markdown-it-figure-references](https://www.npmjs.com/package/markdown-it-figure-references)
- [markdown-it-table-references](https://www.npmjs.com/package/markdown-it-table-references)
- [markdown-it-attribution-references](https://www.npmjs.com/package/markdown-it-attribution-references).

```js
// Figures
const md = require("markdown-it")()
  .use(require("markdown-it-figure-references"), { ns: "figures" })
  .use(require("markdown-it-references"), opts);

// Tables
const md = require("markdown-it")()
  .use(require("markdown-it-table-references"), { ns: "tables" })
  .use(require("markdown-it-references"), opts);

// Attributions
const md = require("markdown-it")()
  .use(require("markdown-it-attribution-references"), { ns: "attributions" })
  .use(require("markdown-it-references"), opts);

// All
const md = require("markdown-it")()
  .use(require("markdown-it-figure-references"), { ns: "figures" })
  .use(require("markdown-it-table-references"), { ns: "tables" })
  .use(require("markdown-it-attribution-references"), { ns: "attributions" })
  .use(require("markdown-it-references"), opts);
```

See a [demo as JSFiddle](https://jsfiddle.net/r62t17av/).

<style>
table { width: 100%;} td:first-child {width: 15%;} td:last-child {width: 45%;}
</style>

The `opts` object can contain:

| Name     | Description                    | Default                        |
| -------- | ------------------------------ | ------------------------------ |
| `labels` | Array of label configurations. | `[ { see below }, { }, .. ] ]` |

<br/>

An object in the `labels` array must contain:

| Name          | Description                   | Example              |
| ------------- | ----------------------------- | -------------------- |
| `ns`          | Namespace.                    | `"figures"`          |
| `text`        | Reference label text.         | `"Figure #"`         |
| `placeholder` | Reference number placeholder. | `"#"`                |
| `class`       | Reference label class         | `"figure-reference"` |

<br/>

By default, `markdown-it-references` contains similar configuration for `figures`, `tables`, and `attributions`, in that order.

**NOTE**  
Label order resolves naming conflicts.
However, the same id shouldn't occur in the same document more than once.

## License

[GPL-3.0](https://github.com/studyathome-internationally/vuepress-plugins/blob/master/LICENSE) &copy; [StudyATHome Internationally](https://github.com/studyathome-internationally/)
