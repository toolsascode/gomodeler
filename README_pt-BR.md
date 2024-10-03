# Go Modeler CLI and Library [![Go Reference](https://pkg.go.dev/badge/github.com/toolsascode/gomodeler.svg)](https://pkg.go.dev/github.com/toolsascode/gomodeler)

O objetivo desse projeto é ajudar aquelas pessoas que precisam de um pequeno software para realizarem suas compilações de templates com facilidade e agilidade. A versão está disponível para ser usada como um pequeno CLI ou uma biblioteca em `golang`. Nós usamos o sistema de templates do Golang para realizar as principais demandas da aplicação.

Nesse projeto contamos também com um conjunto de funções próprias e o arranjo fenomenal do projeto [Sprig](https://masterminds.github.io/sprig/) para ajudá-lo com o máximo de recursos disponíveis pelas comunidades.


## Como instalar

### Go Install

```shell
go install github.com/toolsascode/gomodeler@latest
```

### Via Github

- [Latest version](https://github.com/toolsascode/gomodeler/releases/latest)

```shell
curl -fLSs https://raw.githubusercontent.com/toolsascode/gomodeler/main/scripts/install.sh | bash
```

Or 

```shell
curl -fLSs https://raw.githubusercontent.com/toolsascode/gomodeler/main/scripts/install.sh | sudo bash
```

### Homebrew

```shell
brew install toolsascode/tap/gomodeler
```

### Scoop

1. Execute **PowerShell como Administrador**.
2. Adicione o bucket, executando `scoop bucket add gomodeler-scoop https://github.com/toolsascode/scoop-bucket`.
3. Para instalar, execute `scoop install gomodeler`.

_*Em breve estaremos disponibilizando por outras fontes de instalação._

## Como usar

Veja a documentação completa em [gomodeler.md](./docs/gomodeler.md)

1. Criar uma estrutura.
    **Exemplo:**

    ```shell
    root/
        /templates/
            file1.json.gotmpl
            file2.yaml.gotmpl
            file3.md.gotmpl
            file4.html.gotmpl
        /envFile.yaml

    ```

    *Um ponto importante é o formato do arquivo, não é obrigatório, evita que a necessidade de informar cada nome para o output.

    - **Formato de nome do arquivo de template:** _{filename}.{extension to output}.{extension to template}_
    - **Será gerado o seguinte arquivo automaticamente pelo output:** _{filename}.{extension to output}_ 

2. Criar os templates usando os seguintes recursos disponíveis:
    1. [Go template](https://pkg.go.dev/text/template#hdr-Actions)
    2. [GoModeler Funcs](./docs/functions.md)
    3. [Sprig](https://masterminds.github.io/sprig/)

3. Criar um arquivo de variáveis no formato `Yaml` ou `Json`. Também é possível passar as variáveis diretamente no CLI no format `chave=valor`.
    - **Exemplo:** `envFile.yaml`
4. Executar o comando com as seguintes opções:

```shell
gomodeler -f root/envFile.yaml --template-path root/templates --output-path root/outputs
```

5. Basta verificar a pasta `root/outputs` onde todos arquivos serão gerados.

## Exemplos:

Na pasta [examples](./examples/) é disponibilizado alguns exemplos e modelos de como usar o CLI.

- [Complete](./examples/complete/)
- [Summary](./examples/summary/)
- [Multiple templates](./examples/multiple-templates/)

Para mais informações execute o comando `gomodeler --help` ou use `gomodeler [command] --help` para mais informações sobre um comando especifico. Veja [gomodeler.md](./docs/gomodeler.md).
