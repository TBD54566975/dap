const references = (md, opts) => {
  opts = loadOptions(opts);

  md.inline.ruler.push("references", references_rule(opts));
  md.renderer.rules.reference = reference_renderer(opts);
};

function references_rule(opts) {
  const references = (state) => {
    const tokens = state.tokens;

    const start = state.pos;
    const max = state.posMax;

    // should be at least 5 chars - "<<x>>"
    if (start + 4 > max) {
      return false;
    }

    if (state.src.charCodeAt(start) !== 0x3c /* < */) {
      return false;
    }
    if (state.src.charCodeAt(start + 1) !== 0x3c /* < */) {
      return false;
    }

    for (pos = start + 2; pos < max - 1; pos++) {
      if (state.src.charCodeAt(pos) === 0x20) {
        return false;
      }
      if (state.src.charCodeAt(pos) === 0x0a /* \n */) {
        return false;
      }
      if (state.src.charCodeAt(pos) === 0x3e /* > */ && state.src.charCodeAt(pos + 1) === 0x3e) {
        break;
      }
    }

    // no empty references
    if (pos === start + 2) {
      return false;
    }

    if (state.pending) {
      state.pushPending();
    }

    const targetId = state.src.slice(start + 2, pos);
    token = new state.Token("reference", "", 0);
    token.content = targetId;
    token.meta = { targetId };
    tokens.push(token);

    state.pos = pos + 2;
    state.posMax = max;
    return true;
  };
  return references;
}

function reference_renderer(opts) {
  return (tokens, idx, options, env /* , self */) => {
    const token = tokens[idx];
    const id = token.meta.targetId;

    if (id && Array.isArray(opts.labels)) {
      for (const labelCfg of opts.labels) {
        const { ns, text, placeholder, class: className, renderer: customRenderer } = labelCfg;
        if (ns && env[ns] && env[ns].refs && env[ns].refs[id]) {
          const { index } = env[ns].refs[id];
          const renderer = customRenderer ? customRenderer : references.defaults.renderer;
          return renderer(id, className, text, placeholder, index);
        }
      }
    }

    return `&lt;&lt;${token.meta.targetId}&gt;&gt;`;
  };
}

function loadOptions(options) {
  return options
    ? {
        labels: Array.isArray(options.labels) ? options.labels : references.defaults.labels,
      }
    : references.defaults;
}

references.defaults = {
  labels: [
    {
      ns: "figures",
      text: "Figure #",
      placeholder: "#",
      class: "figure-reference",
    },
    {
      ns: "tables",
      text: "Table #",
      placeholder: "#",
      class: "table-reference",
    },
    {
      ns: "attributions",
      text: "Attribution #",
      placeholder: "#",
      class: "attribution-reference",
    },
  ],
  renderer: (id, className, text, placeholder, index) =>
    `<a ${id ? `href="#${id}"` : ""} ${className ? `class="${className}"` : ""}>${text.replace(
      placeholder,
      index
    )}</a>`,
};

module.exports = references;
