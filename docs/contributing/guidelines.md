---
sidebar_position: 1
---
# Guidelines
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [How to get involved?](#how-to-get-involved)
- [Submitting PRs](#submitting-prs)
  - [Choosing something to work on](#choosing-something-to-work-on)
  - [Developing rancher-turtles](#developing-rancher-turtles)
  - [Asking for help](#asking-for-help)
  - [PR submission guidelines](#pr-submission-guidelines)
    - [Commit message formatting](#commit-message-formatting)
- [Opening Issues](#opening-issues)
- [How the Maintainers process contributions](#how-the-maintainers-process-contributions)
  - [Prioritizing issues](#prioritizing-issues)
  - [Reviewing PRs](#reviewing-prs)
- [ADRs (Architectural Decision Records)](#adrs-architectural-decision-records)
  - [Process](#process)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

Thank you for taking the time to contribute to Rancher CAPI extension [projects](https://github.com/rancher?q=turtles&type=all&language=&sort=) :heart:

Improvements to all areas of the project; from code, to documentation;
from bug reports to feature design and UI enhancement are gratefully welcome.
This guide should cover all aspects of how to interact with the project
and how to get involved in development as smoothly as possible.

Reading docs is often tedious, so let's put the important contributing rule
right at the top: **Always be kind!**

Looking forward to seeing your contributions in the repo! :sparkles:

## How to get involved?

We'd love to accept your patches in pretty much all areas of projects development!

If you’re a new to the project and want to help, but don’t know where to start, here is a non-exhaustive list of ways you can help out:

1. Submit a [Pull Request](#submitting-prs) :rocket:

    Beyond fixing bugs and submitting new features, there are other things you can submit
    which, while less flashy, will be deeply appreciated by all who interact with the codebase:

      - Extending test coverage!
      - Refactoring!
      - Reviewing and updating [documentation][user-docs]!
      - Adding a new UI functionality!

   (See also [Choosing something to work on](#choosing-something-to-work-on) below.)
1. Open an [issue](#opening-issues) :interrobang:

    We have 2 forms of issues: bug reports and feature requests. If you are not sure which category you need, just make the best guess and provide as much information as possible.

1. Interested in helping to improve:

- Rancher CAPI extension backend? Chime in on [`bugs`](https://github.com/rancher/turtles/issues?q=is%3Aopen+is%3Aissue+label%3Akind%2Fbug+) or
    [`help wanted` issues](https://github.com/rancher/turtles/labels/help-wanted).
    If you are seeking to take on a bigger challenge or a more experienced contributor, check out [`feature requests`](https://github.com/rancher/turtles/issues?q=is%3Aopen+is%3Aissue+label%3Akind%2Ffeature).
- extension UI? Take a look at [`open`](https://github.com/rancher/capi-ui-extension) or
    [`help wanted` issues](https://github.com/rancher/capi-ui-extension/labels/help-wanted).
- maybe, user-docs? :nerd_face: Then, jump straight into [`open` issues](https://github.com/rancher/turtles-docs/issues) in the docs repository.

## Opening Issues

These guides aim to help you write issues in a way which will ensure that they are processed
as quickly as possible.

_See below for [how issues are prioritized](#prioritizing-issues)_.

**General rules**:

1. Before opening anything, take a good look through existing issues.

1. More is more: give as much information as it is humanly possible to give.
  Highly detailed issues are more likely to be picked up because they can be prioritized and
  scheduled for work faster. They are also more accessible
  to the community, meaning that you may not have to wait for the core team to get to it.

1. Please do not open an issue with a description that is **just** a link to another issue,
  a link to a slack conversation, a quote from either one of those, or anything else
  equally opaque. This raises the bar for entry and makes it hard for the community
  to get involved. Take the time to write a proper description and summarise key points.

1. Take care with formatting. Ensure the [markdown is tidy](https://docs.github.com/en/free-pro-team@latest/github/writing-on-github/getting-started-with-writing-and-formatting-on-github),
  use [code blocks](https://docs.github.com/en/free-pro-team@latest/github/writing-on-github/creating-and-highlighting-code-blocks) etc etc.
  The faster something can be read, the faster it can be dealt with.

1. Keep it civil. Yes, it is annoying when things don't work, but it is way more fun helping out
  someone who is not... the worst. Remember that conversing via text exacerbates
  everyone's negativity bias, so throw in some emoji when in doubt :+1: :smiley: :rocket: :tada:.

## Submitting PRs

### Choosing something to work on

If you are here to ask for help or request some new behaviour, this
is the section for you. We have curated a set of issues for anyone who simply
wants to build up their open-source cred :muscle:.

- Issues labelled [`good first issues`](https://github.com/search?q=org%3Agithub%2Francher+repo%3Arancher%2Fturtles+repo%3Arancher%2Fcapi-ui-extension+repo%3Arancher%2Fturtles-docs+is%3Aopen+label%3A%22good+first+issue%22+&type=issues&ref=advsearch)
  should be accessible to folks new to the repos, as well as to open source in general.

  These issues should present a low/non-existent barrier to entry with a thorough description,
  easy-to-follow reproduction (if relevant) and enough context for anyone to pick up.
  The objective should be clear, possibly with a suggested solution or some pseudocode.
  If anything similar has been done, that work should be linked.

  If you have come across an issue tagged with `good first issue` which you think you would
  like to claim but isn't 100% clear, please ask for more info! When people write issues
  there is a _lot_ of assumed knowledge which is very often taken for granted. This is
  something we could all get better at, so don't be shy in asking for what you need
  to do great work :smile:.

  See more on [asking for help](#asking-for-help) below!

- [`help wanted` issues](https://github.com/search?q=org%3Agithub%2Francher+repo%3Arancher%2Fturtles+repo%3Arancher%2Fcapi-ui-extension+repo%3Arancher%2Fturtles-docs+is%3Aopen+label%3A%22help+wanted%22+&type=issues&ref=advsearch)
  are for those a little more familiar with the code base, but should still be accessible enough
  to newcomers.

- All other issues labelled `kind/<x>` or `area/<x>` are also up for grabs, but
  are likely to require a fair amount of context.

### Developing rancher-turtles

Check out the dedicated [notes](development.md) on getting started with development.

### Asking for help

If you need help at any stage of your work, please don't hesitate to ask!

- To get more detail on the issue you have chosen, it is a good idea to start by asking
  whoever created it to provide more information.

- If you are struggling with something while working on your PR, or aren't quite
  sure of your approach, you can open a [draft](https://github.blog/2019-02-14-introducing-draft-pull-requests/)
  (prefix the title with `WIP:`) and explain what you are thinking.

### PR submission guidelines

1. Fork the desired repo, develop and test your code changes.
1. Push your changes to the branch on your fork and submit a pull request to the original repository
against the `main` branch.

```bash
git push <remote-name> <feature-name>
```

1. Submit a pull request.
    1. All code PR must be labeled with one of
        - ⚠️ (`:warning:`, major or breaking changes)
        - ✨ (`:sparkles:`, feature additions)
        - 🐛 (`:bug:`, patch and bugfixes)
        - 📖 (`:book:`, documentation or proposals)
        - 🌱 (`:seedling:`, minor or other)

Where possible, please squash your commits to ensure a tidy and descriptive history.

If your PR is still a work in progress, please open a [Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/)
and prefix your title with the word `WIP`. When your PR is ready for review, you
can change the title and remove the Draft setting.

We recommend that you regularly rebase from `main` of the original repo to keep your
branch up to date.

In general, we will merge a PR once a maintainer has reviewed and approved it.
Trivial changes (e.g., corrections to spelling) may get waved through.
For substantial changes, more people may become involved, and you might get asked to resubmit the PR or divide the changes into more than one PR.

#### Commit message formatting

_For more on how to write great commit messages, and why you should, check out
[this excellent blog post](https://chris.beams.io/posts/git-commit/)._

We follow a rough convention for commit messages that is designed to answer three
questions: what changed, why was the change made, and how did you make it.

The subject line should feature the _what_ and
the body of the commit should describe the _why_ and _how_.
If you encountered any weirdness along the way, this is a good place
to note what you discovered and how you solved it.

The format can be described more formally as follows:

```text
<short title for what changed>
<BLANK LINE>
<why this change was made and what changed>
<BLANK LINE>
<any interesting details>
<footer>
```

The first line is the subject and should be no longer than 70 characters, the
second line is always blank, and other lines should be wrapped at a max of 80 characters.
This allows the message to be easier to read on GitHub as well as in various git tools.

There is a template recommend for use [here](https://gist.github.com/yitsushi/656e68c7db141743e81b7dcd23362f1a).

## How the Maintainers process contributions

### Prioritizing issues

The core team regularly processes incoming issues. There may be some delay over holiday periods.

Every issue will be assigned a `priority/<x>` label. The levels of priorities are:

- [`critical-urgent`](https://github.com/rancher/turtles/labels/priority%2Fcritical-urgent): These are time sensitive issues which should be picked up immediately.
  If an issue labelled `critical` is not assigned or being actively worked on,
  someone is expected to drop what they're doing immediately to work on it.
  This usually means the core team, but community members are welcome to claim
  issues at any priority level if they get there first. _However, given the pressing
  timeframe, should a non-core contributor request to be assigned to a `critical` issue,
  they will be paired with a core team-member to manage the tracking, communication and release of any fix
  as well as to assume responsibility of all progess._

- [`important-soon`](https://github.com/rancher/turtles/labels/priority%2Fimportant-soon): Must be assigned as soon as capacity becomes available.
  Ideally something should be delivered in time for the next release.

- [`important-longterm`](https://github.com/rancher/turtles/labels/priority%2Fimportant-longterm): Important over the long term, but may not be currently
  staffed and/or may require multiple releases to complete.

- [`backlog`](https://github.com/rancher/turtles/labels/priority%2Fbacklog): There appears to be general agreement that this would be good to have,
  but we may not have anyone available to work on it right now or in the immediate future.
  PRs are still very welcome, although it might take a while to get them reviewed if
  reviewers are fully occupied with higher priority issues, for example immediately before a release.

These priority categories have been inspired by [the Kubernetes contributing guide](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md).

Other labels include:

- [`adr-required`](https://github.com/rancher/turtles/labels/adr-required):
  Indicates that the issue or PR contains a decision that needs to be documented in a [ADR](#adrs-architectural-decision-records) _before_
  it can be worked on.

- [`needs-investigation`](https://github.com/rancher/turtles/labels/needs-investigation):  There is currently insufficient information to either categorize properly,
  or to understand and implement a solution. This could be because the issue opener did
  not provide enough relevant information, or because more in-depth research is required
  before work can begin.

### Reviewing PRs

The core team aims to clear the PR queue as quickly as possible. Community members
should also feel free to keep an eye on things and provide their own thoughts and expertise.

High-value and/or high priority contributions will be processed as quickly as possible,
while lower priority or nice-to-have things may take a little longer to get approved.

To help facilitate a smoother and faster review, follow the guidelines [above](#pr-submission-guidelines).
Submissions which do not meet standards will be de-prioritised for review.

## ADRs (Architectural Decision Records)

:::note
Please feel free to skip [this](#adrs-architectural-decision-records) and [sub-section](#process) below, since they are only relevant to the [rancher-turtles](https://github.com/rancher/turtles) repository.
:::

Any impactful decisions to the architecture, design, development and behaviour
of rancher-turtles must be recorded in the form of an [ADR](https://engineering.atspotify.com/2020/04/14/when-should-i-write-an-architecture-decision-record/).

A template can be found at [`docs/adr/0000-template.md`][adr-template] of the repo,
with numerous examples of completed records in the same directory.

Contributors are also welcome to backfill ADRs if they are found to be missing.

### Process

1. Start a new [discussion](https://github.com/rancher/turtles/discussions/new?category=adr) using the `ADR` category.

1. Choose an appropriate clear and concise title (e.g. `ADR: Implement X in Go`).

1. Provide a context of the decision to be made. Describe
  the various options, if more than one, and explain the pros and cons. Highlight
  any areas which you would like the reviewers to pay attention to, or those on which
  you would specifically like an opinion.

1. Tag in the [maintainers](https://github.com/rancher/turtles/blob/main/CODEOWNERS) as the "Deciders", and invite them to
  participate and weigh in on the decision and its consequences.

1. Once a decision has been made, open a PR adding a new ADR to the [directory](https://github.com/rancher/turtles/blob/main/docs/adr).
  Copy and complete the [template][adr-template];
    - Increment the file number by one
    - Set the status as "Accepted"
    - Set the deciders as those who approved the discussion outcome
    - Summarise the decision and consequences from the discussion thread
    - Link back to the discussion from the ADR doc

[user-docs]: https://rancher.github.io/turtles-docs/
[adr-template]: https://github.com/rancher/turtles/blob/main/docs/adr/0000-template.md
