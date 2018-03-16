# TomlConfiguration

A simple configuration with Toml File.
supports default tag.

## Install

```bash
go get github.com/Yprolic/TomlConfiguration
```

## Usage and Examples

Lets define and struct that defines our configuration

```go
type Server struct {
	Name    string 
	Port    int    `default:"6060"`
	Enabled bool
	Users   []string
}
```

Load the configuration into multiconfig:

```go
m := TOMLLoader{Path: "conf/conf.toml"}
serverConf := &Server{}
err := m.Load(s)

serverConf.Port // by default 6060
serverConf.Name // "prolic"
```

## License

The MIT License (MIT) - see [LICENSE](/LICENSE) for more details