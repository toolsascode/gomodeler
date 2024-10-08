## gomodeler version

Print the version number of Go Modeler

### Synopsis

All software has versions. This is Go Modeler's

```
gomodeler version [flags]
```

### Options

```
  -h, --help   help for version
```

### Options inherited from parent commands

```
  -c, --config string             Config file (default is $HOME/.gomodeler.yaml or ./.configs/gomodeler.yaml or .gomodeler.yaml)
  -e, --env key=value             List of environment variables in key=value format separated by commas.
  -f, --env-file List             List of environment variable files. Supported extensions JSON or YAML.
  -l, --log-level string          Log level [debug, info, warn, error, fatal, panic] (default "info")
  -m, --merge                     Merge the output files into a single file. (Default: false)
  -o, --output-file List          List of files that will be saved to disk after rendering. The order specified in the file list entry should be followed. templates.
      --output-path Location      Location where rendered files are saved. To use this functionality, use the following format in the template {file name}.{extension file}.{extension template file}.
  -t, --template-file List        List of template files for rendering separated by commas. Optional: filename.ext.exttmpl
      --template-path Directory   Directory of template files to be rendered. Optional: filename.ext.exttmpl
```

### SEE ALSO

* [gomodeler](gomodeler.md)	 - Go Modeler is a small CLI that brings the powerful features of the golang template into a simplified form.

###### Auto generated by spf13/cobra on 16-Sep-2024
