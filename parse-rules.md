# Legend

- `"tokenString"` -> some token
- `TOKEN_NAME` -> some token
- `rule` -> some rule
- `rule[some-context]` -> some rule with the context type `some-context`
- `rule[none]` -> some rule with no context type, useful for `rule[a ? b : c]` syntax
- `rule[a ? b : c]` -> some rule with the context type `b` if the context is of type `a` else `c`
- `rule[context]` -> some rule with the current context type (yes, it has to be specifically `context`)
- `rule[context-a, context-b, ...]` -> some rule with multiple context types
- `[some-context] (...)` -> a pattern that is only available when the context is of type `some-context`

### `comments`

- `"--"` `...`

## `program`

- `rawblock[global-context]`

## `block`

- `"{"` `rawblock[context]` `"}"`

## `rawblock`

- **(** `statement[context]` **)** **\***

## `statement`

- `if-statement[context]`
- `let` `SYMBOL` **(** `","` `SYMBOL` **)** **\***
- `definition`
- `[global-context]` `pubdeclaration`
- `variable` `"="` `expr`
- `variable` **(** `","` `variable` **)** **+** `"="` `expr`
- `expr`

## `pubdeclaration`

- `"pub"` `definition`

## `definition`

- `let` `SYMBOL` `"="` `expr`
- `let` `SYMBOL` **(** `","` `SYMBOL` **)** **+** `"="` `expr`

## `variable`

- `atom` **(**
  **(** `"."` `SYMBOL` **)**
  **|** **(** `"["` `expr` `"]"` **)**
  **)** **+**
- `SYMBOL`

## `expr`

- `logic`

## `logic`

- `comparison` **(** **(** `"&&"` **|** `"||"` **)** `comparison` **)** **\***

## `comparison`

- `concat` **(** **(** `"<"` **|** `">"` **|** `"=="` **|** `"!="` **|** `"<="` **|** `">="` **)** `concat` **)** **\***

## `concat`

- `add-or-sub` **(** `".."` `add-or-sub` **)** **\***

## `add-or-sub`

- `mul-or-div` **(** **(** `"+"` **|** `"-"` **)** `mul-or-div` **)** **\***

## `mul-or-div`

- `value` **(** **(** `"*"` **|** `"/"` **|** `"%"` **)** `value` **)** **\***

## `value`

- **(** `"+"` **|** `"-"` **|** `"!"` **|** `"#"` **|** `"typeof"` **|** `"copyof"` **)** `value`
- `call`

## `call`

- `literal` `arglist` **\***

## `literal`

- `atom` **(**
  **(** `"."` `SYMBOL` **)**
  **|** **(** `"["` `expr` `"]"` **)**
  **)** **\***

## `atom`

- `INT`
- `FLOAT`
- `TRUE` **|** `FALSE`
- `STRING`
- `SYMBOL`
- `"("` `expr` `")"`

## `arglist`

- `"("` **(** `item` **(** `","` `item` **)** **\*** **(** `","` **)** **?** **)** **?** `")"`

## `item`

- **(** `"..."` **)** **?** `expr`

## [ ] `if-statement`

- `"if"` `"("` `expr` `")"` `block[context]` `"else"` `if-statement`
- `"if"` `"("` `expr` `")"` `block[context]` `"else"` `block[context]`
- `"if"` `"("` `expr` `")"` `block[context]`
