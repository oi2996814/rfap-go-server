
# rfap-go-server

![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/alexcoder04/rfap-go-server?include_prereleases)
![GitHub top language](https://img.shields.io/github/languages/top/alexcoder04/rfap-go-server)
![GitHub](https://img.shields.io/github/license/alexcoder04/rfap-go-server)
![GitHub issues](https://img.shields.io/github/issues/alexcoder04/rfap-go-server)
![GitHub pull requests](https://img.shields.io/github/issues-pr/alexcoder04/rfap-go-server)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/alexcoder04/rfap-go-server)
![GitHub contributors](https://img.shields.io/github/contributors-anon/alexcoder04/rfap-go-server)

The reference server implementation of the rfap protocol.

See [here](#related-projects) for protocol specifications and related projects.

## Installation

### Stable release

**Coming soon**

### Bleeding-edge

Make sure you have `git`, `make` and `go` installed.

```sh
git clone https://github.com/alexcoder04/rfap-go-server
cd rfap-go-server/src

make run     # start testing server
make linux   # compile executable, other possible arguments: windows/raspberry
go install . # compile and install executable to $GOPATH/bin
```

## Related projects

 - https://github.com/alexcoder04/rfap - general protocol specification
 - https://github.com/alexcoder04/librfap - Python library
 - https://github.com/BoettcherDasOriginal/rfap-cs-lib - C# library
 - https://github.com/alexcoder04/rfap-pycli - Python CLI client based on librfap

