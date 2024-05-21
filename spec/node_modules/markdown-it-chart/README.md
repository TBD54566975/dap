# markdown-it-chart

Chart.js plugin for markdown-it.


## Installation

```
yarn install markdown-it-chart
```


## Usage

```js
import markdownIt from 'markdown-it'
import markdownItChart from 'markdown-it-chart'
const mdi = markdownIt()
mdi.use(markdownItChart)
mdi.render(`\`\`\`chart
{
  "type": "pie",
  "data": {
    "labels": [
      "Red",
      "Blue",
      "Yellow"
    ],
    "datasets": [
      {
        "data": [
          300,
          50,
          100
        ],
        "backgroundColor": [
          "#FF6384",
          "#36A2EB",
          "#FFCE56"
        ],
        "hoverBackgroundColor": [
          "#FF6384",
          "#36A2EB",
          "#FFCE56"
        ]
      }
    ]
  },
  "options": {}
}
\`\`\``)
```


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
