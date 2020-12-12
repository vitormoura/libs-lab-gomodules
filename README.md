# libs-lab-gomodules

Personal go modules labs


### Regenerate samples file

Install go-bindata from https://github.com/go-bindata/go-bindata, then:

```shell
go-bindata.exe -o ./messages/phrases_samples.go -pkg messages ./data/phrases.txt
```