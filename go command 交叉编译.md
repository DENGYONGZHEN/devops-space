## go command 交叉编译

### 1. Linux/Unix 系统

运行以下命令查看当前设备架构：

```bash
$ uname -m
```

常见输出：

- `x86_64` 表示 64 位 x86 架构（`amd64`）。
- `i686` 或 `i386` 表示 32 位 x86 架构（`386`）。
- `armv7l` 表示 32 位 ARM 架构（`arm`）。
- `aarch64` 表示 64 位 ARM 架构（`arm64`）。

### 2.设置目标平台和架构

Go 提供环境变量来指定编译目标：

- **`GOOS`**：目标操作系统（如 `linux`, `windows`, `darwin`）。
- **`GOARCH`**：目标架构（如 `amd64`, `arm`, `386` 等）

```powershell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o output_file_name
```

###### 注意事项

1. **需要支持的 C 库** 如果你的 Go 程序依赖了 C 库（通过 `cgo`），你需要在目标平台上安装这些库。对于交叉编译，还需要额外的工具链支持。可以通过禁用 `cgo` 来简化编译：

   ```powershell
   > set CGO_ENABLED=0
   ```

2. **ARM 架构的支持** 如果目标设备使用 ARM 架构（如树莓派），将 `GOARCH` 设置为 `arm` 或 `arm64`：

   ```powershell
   > set GOOS=linux
   > set GOARCH=arm64
   > go build -o output_file_name
   ```

3. **Go 环境变量检查** 确保你使用的是官方的 Go 编译器，而非第三方版本（如 MinGW 提供的工具链）。运行 `go version` 确保工具链版本兼容。

### 3.scp

scp [选项] [源路径] [目标路径]

- **`[源路径]`**：指定需要传输的文件或目录的路径，可以是本地路径或远程路径。

- **`[目标路径]`**：指定传输到的目标路径，也可以是本地路径或远程路径。

- 选项

  ：常见选项包括：

  - `-r`：递归复制整个目录。
  - `-P [端口号]`：指定 SSH 的端口号（默认是 22）。
  - `-C`：启用压缩，提高传输效率。
  - `-v`：显示详细的调试信息。

  ```powershell
  scp file.txt user@remote_host:/remote/directory/
  ```

  ### 4.设置可执行权限  执行

  ```bash
  $chmod +x countFlag
  $./countFlag
  ```

  