# markdown-it-icons

Plugin for markdown-it, supports emoji icons and font-awesome icons.


## Installation

```
yarn add markdown-it-icons
```


## Usage

### for node.js

```js
import markdownIt from 'markdown-it'
import markdownItIcons from 'markdown-it-icons'
const mdi = markdownIt()
mdi.use(markdownItIcons, 'emoji')
mdi.use(markdownItIcons, 'font-awesome')
mdi.render('I :heart: you') // <p>I <i class="e1a-heart"></i> you</p>
mdi.render('A :fa-car: runs') // <p>A <i class="fa fa-car"></i> runs</p>
```

### for browser

You also need to import the css:

```js
import 'markdown-it-icons/dist/index.css'
```

Or you can add the css to the web page directly.


## Development

### Build

```
yarn build:watch
```

### Test

```
yarn test
```

### Distribution

```
yarn release && npm publish
```


## Todo
