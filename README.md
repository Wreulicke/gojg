# gojg [![CircleCI](https://circleci.com/gh/wreulicke/gojg.svg?style=svg)](https://circleci.com/gh/wreulicke/gojg)

## TODO

- [x] Implement parser
- [ ] Command Line Interface
  - [ ] output
    - [ ] to stdout
    - [ ] to file
  - [ ] input as template file
    - [x] from file
    - [ ] from stdin
- [ ] Resolve template parameter
  - [x] from CLI parameter
  - [ ] from JSON
    - [ ] from file
    - [ ] from stdin
  - [ ] from javascript (using some golang implementation, [otto](https://github.com/robertkrimen/otto), [goja](https://github.com/dop251/goja), [v8eval](https://github.com/sony/v8eval))
    - [ ] from file
    - [x] Nothing in my roadmap (~~from stdin~~)
- [ ] Support JSON5 format
- [ ] Make more friendly message
- [ ] Implement Verbose parser.
    - [ ] refacotr lexer using channel?
- [ ] Release