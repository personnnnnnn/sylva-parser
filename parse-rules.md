### `comments`

- `"--"` `...`

## [X] `variable`

- `atom` **(**
  **(** `"."` `SYMBOL` **)**
  **|** **(** `"["` `expr` `"]"` **)**
  **)** **+**
- `SYMBOL`

## [X] `statement`

- `let` `SYMBOL` `"="` `expr`
- `let` `SYMBOL` **(** `","` `SYMBOL` **)** **+** `"="` `expr`
- `let` `SYMBOL` **(** `","` `SYMBOL` **)** **\***
- `variable` `"="` `expr`
- `variable` **(** `","` `variable` **)** **+** `"="` `expr`
- `expr`

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
