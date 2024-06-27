# Contribution Guide

There are many ways to be an open source contributor, and we're here to help you on your way!
You may:

* Propose ideas in the [`#dap` channel on Discord][tbd-discord]
* Raise an issue or feature request in our [issue tracker]
* Help another contributor with one of their questions, or a code review
* Suggest improvements to our documentation by supplying a Pull Request
* Evangelize our work together in conferences, podcasts, and social media spaces.

This guide is for you.

## Communication

### Issues

Anyone from the community is welcome (and encouraged!) to raise issues via
[GitHub Issues][issue tracker].

### Discussions

Design discussions and proposals take place in the [`#dap` channel on Discord][tbd-discord].

We advocate an asynchronous, written debate model - so write up your thoughts and invite the
community to join in!

### Continuous Integration

Build and Test cycles are run on every commit to every branch on using [GitHub Actions].

## Development

### 🛠️ Prerequisites

If you want to contribute without any local environment setup, click one of the buttons below to use
StackBlitz or CodeSandbox:

[![Open in StackBlitz](https://developer.stackblitz.com/img/open_in_stackblitz.svg)][StackBlitz]
[![Open with CodeSandbox](https://assets.codesandbox.io/github/button-edit-lime.svg)][CodeSandbox]

If you prefer to develop locally, you'll need to install the following requirements:

| Requirement | Tested Version |
| ----------- | -------------- |
| [Node.js]   | 20.14.0        |
| [PNPM]      | 9.1.4          |

You can install these requirements manually, but we recommend using
[<img src="https://github.com/TBD54566975/daps.dev/assets/134693/cba8aa84-4f4a-4cb6-9cc6-ff7179cf7fb1" alt="Hermit Logo" width="20" height="20" style="vertical-align: middle;"> Hermit][Hermit],
which installs tools for software projects in self-contained, isolated sets to ensure consistent
tooling for all local development and continuous integration. To use Hermit, run the following
command:

```shell
. ./bin/activate-hermit
```

This will set up your local environment with the necessary tools and versions specified for this project.

### <img src="https://github.com/TBD54566975/daps.dev/assets/134693/9fdf53e5-c3a7-4630-8f34-681b2861d594" alt="Starlight Logo" width="20" height="20" style="vertical-align: middle;"> Starlight

This documentation site was built with built with [Starlight].

Starlight looks for `.md` or `.mdx` files in the `src/content/docs/` directory. Each file is exposed
as a route based on its file name.

Images can be added to `src/assets/` and embedded in Markdown with a relative link.

Static assets, like favicons, can be placed in the `public/` directory.

### 🧞 Commands

All commands are run from the root of the project, from a terminal:

| Command                   | Action                                           |
| :------------------------ | :----------------------------------------------- |
| `pnpm install`            | Installs dependencies                            |
| `pnpm dev`                | Starts local dev server at `localhost:4321`      |
| `pnpm build`              | Build your production site to `./dist/`          |
| `pnpm preview`            | Preview your build locally, before deploying     |
| `pnpm astro ...`          | Run CLI commands like `astro add`, `astro check` |
| `pnpm astro -- --help`    | Get help using the Astro CLI                     |

## Contribution

We review contributions to the codebase via GitHub's Pull Request mechanism. We have
the following guidelines to ease your experience and help our leads respond quickly
to your valuable work:

* Start by proposing a change either in Issues (most appropriate for small change requests or bug
  fixes) or in Discussions (most appropriate for design and architecture considerations, proposing a
  new feature, or where you'd like insight and feedback).
* Cultivate consensus around your ideas; the project leads will help you pre-flight how beneficial
  the proposal might be to the project. Developing early buy-in will help others understand what
  you're looking to do, and give you a greater chance of your contributions making it into the
  codebase! No one wants to see work done in an area that's unlikely to be incorporated.
* Fork the repo into your own namespace/remote.
* Work in a dedicated feature branch. Atlassian wrote a great
  [description of this workflow][feature-branch-workflow].
* When you're ready to offer your work to the project:
  * Squash your commits into a single one (or an appropriate small number of commits), and rebase
    atop the upstream `main` branch. This will limit the potential for merge conflicts during
    review, and helps keep the audit trail clean. A good writeup for how this is done is
    [this beginner's guide to squashing commits][squashing-commits-guide]. If you are having
    trouble, feel free to ask a member or the community for help or leave the commits as-is, and
    flag that you'd like rebasing assistance in your PR! We're here to support you.
  * Open a PR in the project to bring in the changes from your feature branch.
  * The maintainers noted in the [`CODEOWNERS`](./CODEOWNERS) file will review your PR and optionally
    open a discussion about its contents before moving forward.
  * Remain responsive to follow-up questions, be open to making requested changes, and...
    You're a contributor!
* And remember to respect everyone in our global development community. Guidelines
  are established in our [Code of Conduct](./CODE_OF_CONDUCT.md).

[tbd-discord]: https://discord.gg/tbd
[issue tracker]: https://github.com/TBD54566975/daps.dev/issues
[GitHub Actions]: https://github.com/TBD54566975/daps.dev/actions
[feature-branch-workflow]: https://www.atlassian.com/git/tutorials/comparing-workflows/feature-branch-workflow
[squashing-commits-guide]: https://medium.com/@slamflipstrom/a-beginners-guide-to-squashing-commits-with-git-rebase-8185cf6e62ec
[Node.js]: https://nodejs.org/
[PNPM]: https://pnpm.io/
[StackBlitz]: https://stackblitz.com/github/TBD54566975/daps.dev/tree/main/docs
[CodeSandbox]: https://codesandbox.io/p/sandbox/github/TBD54566975/daps.dev/tree/main/docs
[Starlight]: https://starlight.astro.build
[Hermit]: https://cashapp.github.io/hermit/usage/get-started/
