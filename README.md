# Turtles Product Documentation

## Build the Documentation site

The repository uses [Antora Playbooks](https://docs.antora.org/antora/latest/) to locally or remotely build the AsciiDoc content into a static website.

### Prerequisites

#### git

You need git to get the source code of this repository. Run the command below to check whether git is installed on your machine.

```console
git --version
```

If you don't have git installed on your machine, download and install it for your operating system from the [git downloads](https://git-scm.com/downloads) page.

#### Node.js

Antora requires an active long term support (LTS) release of Node.js. Run the command below to check if you have Node.js installed, and which version. This command should return an [active Node.js LTS version number](https://nodejs.org/en/about/releases/)

```console
node -v
```

If you don't have Node.js installed on your machine, install it, preferably via [nvm](https://github.com/nvm-sh/nvm)

### Clone the Playbook repository

Run the git command to clone this repository.

```console
git clone https://github.com/rancher/turtles-product-docs.git
```

This playbook repository uses a [git submodule](https://git-scm.com/book/en/v2/Git-Tools-Submodules) to get the custom Antora supplemental files that provide a custom GUI theme for the documentation website. Run the command below to get the submodule.

```console
git submodule update --init
```

### Install node modules

Open a terminal at the root of the git repository. Run the command below.

```console
npm install
```

### Run Antora to build the static website

As a local example, run the command below to build the site:

```console
npx antora --fetch turtles-local-playbook.yml
```

Navigate to the `./build/site` directory and open the index.html file in your browser to view and navigate the documentation site.

Alternatively, run the below command first and then open `http://127.0.0.1:8080/` in your browser for a web server preview:

```console
make preview
```

### Run Antora to build the static website using the local documentation content

The command provided in the previous section fetches documentation content of the products from their respective remote GitHub repositories. If you want the playbook to use the documentation content from your local machine instead you can do so with `product-docs-playbook-local.yml`.

Clone all the individual product documentation Github repositories one level above the current playbook repository.

As an example, run the command below to use the local `turtles-local-playbook.yml` file.

```console
npx antora --fetch turtles-local-playbook.yml
```

### Makefile

The Makefile can also be used to set up your environment, build the local and remote static site, and clean your build output directory.

## How to report issues related to the SUSE Rancher Product Documentation

### If you are a SUSE Rancher Customer

It is recommended to report the issue via the [SUSE Customer Center](https://scc.suse.com/)

### If you are a SUSE Internal Employee

It is recommended to file a Jira ticket. If you do not have access to Jira then you can file a GitHub ticket on the respective product documentation repository.
