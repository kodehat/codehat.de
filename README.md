# CodeHat's Website ![](https://img.shields.io/badge/license-MIT-blue.svg) [![Netlify Status](https://api.netlify.com/api/v1/badges/f7ed6cde-fcf8-4daf-a69d-ef422a607f8a/deploy-status)](https://app.netlify.com/sites/ch-website/deploys)

> My static website built with Hugo.

This repo contains my personal website and is built with [Hugo](https://gohugo.io/).

## Development

Linux, OS X & Windows:

`hugo` has to be installed.

```sh
hugo serve -D
```

## Deployment

Linux, OS X & Windows:

`hugo` has to be installed.

```sh
hugo --gc --minify --enableGitInfo
```

You'll find all required files in the `public` folder.

### Additions For Netlify

Set the following environment variables:

| Key                | Value      |
| :----------------- | :--------- |
| HUGO_ENABLEGITINFO | true       |
| HUGO_ENV           | production |
| HUGO_VERSION       | 0.74.3     |

## Usage example

My website acts as a demo, so just head over to <https://www.codehat.de> to see this repo in action.

## Meta

CodeHat – [@CodeHat](https://twitter.com/CodeHat)

Distributed under the MIT license. See ``LICENSE`` for more information.

[https://github.com/kodehat](https://github.com/kodehat)

## Contributing

1. Fork it (<https://github.com/kodehat/website/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
