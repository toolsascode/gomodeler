# Go Modeler CLI and Library

The objective of this project is to help those people who need a small software to carry out their template compilations easily and quickly. The version is available to be used as a small CLI or a library in `golang`. We use Golang's templating system to perform the main demands of the application.

In this project we also have a set of our own functions and the phenomenal arrangement of the [Sprig](https://masterminds.github.io/sprig/) project to help you with the maximum resources available by the communities.


## How to install?

### Via Github
```shell
curl -sL https://api.github.com/repos/toolsascode/gomodeler/releases/latest | \
jq -r '.assets[] | select(.name? | match("gomodeler-v.*.tar.gz$")) | .browser_download_url'
```

_*Soon we will be making it available through other installation sources._

## How to use?
See full documentation at [gomodeler.md](./docs/gomodeler.md)

1. Create a structure.
    **Example:**
    ```shell
    root/
        /templates/
            file1.json.gotmpl
            file2.yaml.gotmpl
            file3.md.gotmpl
            file4.html.gotmpl
        /envFile.yaml

    ```

    *An important point is the file format, it is not mandatory, it avoids the need to inform each name for the output.

    - **Template file name format:** _{filename}.{extension to output}.{extension to template}_
    - **The following file will be generated automatically by the output:** _{filename}.{extension to output}_ 

2. Create templates using the following available resources:
    1. [Go template](https://pkg.go.dev/text/template#hdr-Actions)
    2. [GoModeler Funcs](./docs/functions.md)
    3. [Sprig](https://masterminds.github.io/sprig/)

3. Create a variables file in `Yaml` or `Json` format. It is also possible to pass the variables directly to the CLI in the `key=value` format.
    - **Example:** `envFile.yaml`
4. Run the command with the following options:

```shell
gomodeler -f root/envFile.yaml --template-path root/templates --output-path root/outputs
```

Just check the `root/outputs` folder where all files will be generated.

For more options run the `gomodeler --help` command or use `gomodeler [command] --help` for more information about a command. See [gomodeler.md](./docs/gomodeler.md)