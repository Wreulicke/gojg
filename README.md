# gojg [![CircleCI](https://circleci.com/gh/wreulicke/gojg.svg?style=svg)](https://circleci.com/gh/wreulicke/gojg)

## TODO

- [x] Implement parser
- [x] Command Line Interface
  - [x] output
    - [x] to stdout
    - [x] to file
  - [x] input as template file
    - [x] from file
    - [x] from stdin
- [x] Resolve template parameter
  - [x] from CLI parameter
  - [x] from JSON
    - [x] from file
    - [x] ~~from stdin~~ (Remove from roadmap)
  - [ ] from javascript (using some golang implementation, [otto](https://github.com/robertkrimen/otto), [goja](https://github.com/dop251/goja), [v8eval](https://github.com/sony/v8eval))
    - [ ] from file
    - [x] Nothing in my roadmap (~~from stdin~~)
- [ ] Support JSON5 format
- [ ] Make more friendly message
- [ ] Implement Verbose parser.
    - [ ] refacotr lexer using channel?
- [ ] Release