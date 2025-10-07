| Function                        | What it does                                    |
| ------------------------------- | ----------------------------------------------- |
| `reflect.TypeOf(i)`             | Returns the type info                           |
| `reflect.ValueOf(i)`            | Returns the runtime value                       |
| `t.Kind()`                      | Returns the category (struct, slice, map, etc.) |
| `t.NumField()` / `t.Field(i)`   | Get struct fields                               |
| `t.NumMethod()` / `t.Method(i)` | Get methods                                     |
| `v.FieldByName("Name")`         | Get field by name                               |
| `v.MethodByName("Speak")`       | Get method by name                              |
| `v.Call()`                      | Call method dynamically                         |
| `v.Interface()`                 | Convert `reflect.Value` back to `interface{}`   |
