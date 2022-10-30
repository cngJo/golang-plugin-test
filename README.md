# Golang Plugin Example

This repository shows a very simple plugin system, based on golangs `plugin` library.

The main application that loads the plugins is located in `main-application`, the plugins are contained in `plugins/*`

## Compiling

**Main Application**
```console
pushd main-application

go build ./cmd/...

# NOTE: We expect plugins in ./plugins currently
mv plugin-loader ../
```

**Plugins**

```console
pushd plugins/first

go build --buildmode=plugin main.go
```

## Running the application

```console
./plugin-loader
```

