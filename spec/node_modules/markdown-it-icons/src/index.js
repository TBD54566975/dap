import 'emojione/extras/css/emojione-awesome.css'
import './font-awesome.css'

import { rendererRule, coreRuler } from 'markdown-it-regex'
import emoji from 'emojione/emoji.json'
import faIconChars from 'font-awesome-icon-chars'

const iconsPlugin = (md, name) => {
  let options = null
  switch (name) {
    case 'emoji':
      let emojis = []
      Object.keys(emoji).forEach((key) => {
        emojis = emojis.concat(emoji[key].shortname.slice(1, -1)).concat(emoji[key].shortname_alternates.map((item) => item.slice(1, -1)))
      })
      const emojiRegex = new RegExp(`(:(?:${emojis.join('|').replace(/\+/g, '\\+')}):)`)
      options = {
        name,
        regex: emojiRegex,
        replace: (match) => `<i class="e1a-${match.slice(1, -1)}"></i>`
      }
      break
    case 'font-awesome':
      let faIcons = []
      faIconChars.forEach((char) => {
        faIcons = faIcons.concat(char.id).concat(char.aliases || [])
      })
      const faRegex = new RegExp(`(:fa-(?:${faIcons.join('|')}):)`)
      options = {
        name,
        regex: faRegex,
        replace: (match) => `<i class="fa ${match.slice(1, -1)}"></i>`
      }
      break
    default:
      return
  }
  md.renderer.rules[options.name] = (tokens, idx) => {
    return rendererRule(tokens, idx, options)
  }
  md.core.ruler.push(options.name, (state) => {
    coreRuler(state, options)
  })
}

export default iconsPlugin
