Struct to flags, inspired by [clap](https://github.com/fred1268/go-clap), but implemented on top of the standard flag package.

# Usage
Create a struct, pass it to clapr.Parse. Each exported field of the struct will be converted to a flag if the type is supported (bool, int, uint, string, time.Duration, []int and []string ). As it uses flag behind the scene, help (-h and -help) is automatically generated.

# Tags
Supported tags to customize the generated flags:
- *short*/*lon*g: short and long variant of the flag. if none is provided, the flag will be the field name in snake case (lowercase with dashes between "words")
- *desc*: description which will be displayed in the generated help.

