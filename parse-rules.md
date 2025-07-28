### `comments`

- `"--"` `...`

## [X] `variable`

- `atom` **(**
  **(** `"."` `SYMBOL` **)**
  **|** **(** `"["` `expr` `"]"` **)**
  **)** **+**
- `SYMBOL`

## [ ] `statement`

- [ ] **(** `pub` **)** **?** `let` `SYMBOL` `"="` `expr`
- [ ] **(** `pub` **)** **?** `let` `SYMBOL` **(** `","` `SYMBOL` **)** **+** `"="` `expr`
- [x] `variable` `"="` `expr`
- [x] `variable` **(** `","` `variable` **)** **+** `"="` `expr`
- [x] `expr`

## `expr`

- `logic`

## [X] `logic`

- `comparison` **(** **(** `"&&"` **|** `"||"` **)** `comparison` **)** **\***

## [X] `comparison`

- `concat` **(** **(** `"<"` **|** `">"` **|** `"=="` **|** `"!="` **|** `"<="` **|** `">="` **)** `concat` **)** **\***

## [X] `concat`

- `add-or-sub` **(** `".."` `add-or-sub` **)** **\***

## `add-or-sub`

- `mul-or-div` **(** **(** `"+"` **|** `"-"` **)** `mul-or-div` **)** **\***

## [X] `mul-or-div`

- `value` **(** **(** `"*"` **|** `"/"` **|** `"%"` **)** `value` **)** **\***

## [X] `value`

- **(** `"+"` **|** `"-"` **|** `"!"` **|** `"#"` **|** `"typeof"` **|** `"copyof"` **)** `value`
- `call`

## [X] `call`

- `literal` `arglist` **\***

## [X] `literal`

- `atom` **(**
  **(** `"."` `SYMBOL` **)**
  **|** **(** `"["` `expr` `"]"` **)**
  **)** **\***

## [X] `atom`

- `INT`
- `FLOAT`
- `TRUE` **|** `FALSE`
- `STRING`
- `SYMBOL`
- `"("` `expr` `")"`

## [X] `arglist`

- `"("` **(** `item` **(** `","` `item` **)** **\*** **(** `","` **)** **?** **)** **?** `")"`

## [X] `item`

- **(** `"..."` **)** **?** `expr`
