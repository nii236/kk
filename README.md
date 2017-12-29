# Pody

[![CircleCI](https://circleci.com/gh/JulienBreux/pody.svg?style=svg&circle-token=0a3523b14c7004814d4b057db4efe6840dc58e3a)](https://circleci.com/gh/JulienBreux/pody) [![Github issues](https://img.shields.io/github/issues/JulienBreux/pody.svg)](https://github.com/JulienBreux/pody/issues) [![License](https://img.shields.io/github/license/JulienBreux/pody.svg)](https://github.com/JulienBreux/pody/blob/master/LICENSE) [![Twitter](https://img.shields.io/twitter/follow/JulienBreux.svg)](https://twitter.com/JulienBreux)

👾 CLI app to manage your Pods in your Kubernetes cluster.

[![asciicast](https://asciinema.org/a/iMy1llucylhVslRIxZGrnmk9L.png)](https://asciinema.org/a/iMy1llucylhVslRIxZGrnmk9L)

## Getting started

- Download a latest release [here](https://github.com/JulienBreux/pody/releases).
- Run Pody `./pody`

## Key bindings
 Key combination | Description
---|---
<kbd>D</kbd>|Delete (pods)
<kbd>L</kbd>|Display logs (pods)
<kbd>PgUp</kbd>|Moves to the previous (pods, namespaces, containers)
<kbd>PgDn</kbd>|Moves to the next  (pods, namespaces, containers)
<kbd>Enter</kbd>|Select entry (namespaces)
<kbd>CTRL N</kbd>|Prompts the namespace to switch
<kbd>CTRL+C</kbd>|Exits the application

---

## Stargazers over time		

[![Stargazers over time](https://starcharts.herokuapp.com/julienbreux/pody.svg)](https://starcharts.herokuapp.com/julienbreux/pody)	

---

### GPG Signature

You can download Julien Breux's public key to verify the signature.

    gpg --keyserver hkp://pgp.mit.edu --recv-keys 951C3F93B6A8C22C

[Why sign commit?](https://julienbreux.uk/git-users-it-s-time-to-sign-your-commits-2eef5e51cce2)

---

### Credits
* [Kubernetes](https://kubernetes.io/) famous Kubernetes Go client
* [GOCUI](https://github.com/jroimartin/gocui) for the UI
* [Pad](https://github.com/willf/pad) for the string pad

---

### License

Licensed under the [MIT License](https://julienbreux.github.io/license/) by [Julien Breux](https://github.com/JulienBreux)
