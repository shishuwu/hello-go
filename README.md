# Go sample template (with mod)

# Code structure
    - go.mod
    - common [module]
    - server [module]
    - hellowolrd [do some code entrance]

## Command line

- init go repository 
  
    go mod init github.com/shishuwu/hello-go

- install current repo
  - go to main go file, run

    go install (安装 此代码到 gopath/bin 下)

  - the exe will be genereated under `%USERPROFILE%\go\bin\`


## Code
- import package inside module
  
    `import "github.com/shishuwu/hello-go/common/math"`

- import package outside
    ```
    import (
	    log "github.com/sirupsen/logrus" // log is the alias you could customize
    )
    ```

## Other
- [How to write Go code](https://golang.org/doc/code.html)

- [Go mod sample](https://mp.weixin.qq.com/s/TvTlz3uKIBqgg1FjlAItPQ)
    ```
    go mod init 创建了一个新的模块，初始化 go.mod 文件并且生成相应的描述
    go build, go test 和其它构建代码包的命令，会在需要的时候在 go.mod 文件中添加新的依赖项
    go list -m all 列出了当前模块所有的依赖项
    go get 修改指定依赖项的版本（或者添加一个新的依赖项）
    go mod tidy 移除模块中没有用到的依赖项。
    ```