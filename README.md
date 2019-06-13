# super-hacker

<p align="left">
  <a href="https://godoc.org/github.com/briandowns/super-hacker"><img src="https://godoc.org/github.com/briandowns/super-hacker?status.svg" alt="GoDoc"></a>
  <a href="https://opensource.org/licenses/BSD-3-Clause"><img src="https://img.shields.io/badge/License-BSD%203--Clause-orange.svg?" alt="License"></a>
  <a href="https://github.com/briandowns/super-hacker/releases"><img src="https://img.shields.io/badge/version-0.1.0-green.svg?" alt="Version"></a>
</p>

super-hacker is a CLI application that will output beautifully written code to your terminal allowing you to astound your friends and family.  By default, the output buffer is set to 3 but is overridable.  To quit, press `esc` or `ctrl-c`.

## Examples

super-hacker with no flags runs it with Go as the default language and an output buffer of 3 bytes.

```sh
$ super-hacker
```

super-hacker with the `l` flag runs it with the given language. See the "Supported Languages" section below.

```sh
$ super-hacker -l python
```

super-hacker with the `b` flag runs it with the output buffer set to the given value.

```sh
$ super-hacker -b 16
```

## Supported Languages

- Go
- C
- Python
- Scala
- Haskell
- Java
- Javascript

## Installation

Binary releases can be obtained from the `releases` section of the repository.  To build, issue `make`.  This will build a binary for your platform.

## Adding More Languages and Templates

To add a new language to `super-hacker`, create a directory named after the language you're adding in the templates directory. In there, add a file to hold the code.  In that file, create a Go `const` that is exported from that new language package and paste in the code.  Make sure to include license headers if available.  If a license is needed, link to it in a comment in the file.

In `template.go`, add a new slice and reference the newly created constant.  Next, add an entry in the switch state in the `Random(lang string)` function.

## Contributing

Please feel free to open a PR!

## License

SuperHacker source code is available under the BSD 3 clause [License](/LICENSE).

## Contact

[@bdowns328](http://twitter.com/bdowns328)
