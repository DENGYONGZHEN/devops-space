## variable

##### 1. 变量赋值

**语法：**

```shell
variable_name=value
```

**关键规则：**

- **等号 `=` 两边绝对不能有空格**。这是新手最常见的错误。
  - `正确： name="John"`
  - `错误： name = "John"` (这会被 Shell 解析为命令 `name` 和参数 `=`、`"John"`)
- 变量名只能包含字母、数字和下划线，且**不能以数字开头**。
- 如果值中包含空格、制表符或换行符，**必须使用引号**将其引起来。

**赋值示例：**

```shell
name="John Doe"        # 包含空格，需要双引号
count=10               # 赋值整数
pi=3.14                # 赋值浮点数 (但Bash本身不直接支持浮点运算)
PATH=/usr/bin:/bin     # 赋值路径字符串
greeting=Hello\ World  # 使用反斜杠转义空格 (不推荐，可读性差)
```

------

##### 2. 变量使用 (变量扩展)

要使用一个变量的值，你必须在变量名前加上美元符号 `$`。

**语法：**

```shell
$variable_name
# 或者
${variable_name}
```

**示例：**

```shell
name="Alice"
echo $name        # 输出： Alice
echo "Hello, $name" # 输出： Hello, Alice
echo "Hello, ${name}" # 输出： Hello, Alice (与上句等效)

count=5
echo $count       # 输出： 5
```

###### 为什么需要 `${}`？ (关键补充！)

花括号 `{}` 用于**明确界定变量名的边界**。当变量名后面紧跟其他字母、数字或下划线时，必须使用它。

```shell
fruit="apple"
echo "I have $fruits"      # 错误！Shell 会尝试寻找变量 `fruits`，其值为空。
echo "I have ${fruit}s"    # 正确！输出： I have apples

# 另一个例子
your_id=100
echo "Your ID is $your_id"     # 正确
echo "Your ID is $your_id_100" # 错误！寻找变量 your_id_100
echo "Your ID is ${your_id}_100" # 正确！输出： Your ID is 100_100
```

##### 3. 无类型特性详解

Shell 变量本质上都是**字符串**。即使你赋值为数字，它也被存储为字符串。但在算术上下文中，Shell 会自动将这些字符串解释为数字。

```shell
a=10        # 字符串 "10"
b=20        # 字符串 "20"

# 算术运算：需要放在 $(( ... )) 或 let 命令中
result=$((a + b))  # Shell 会将 a 和 b 作为数字处理
echo $result       # 输出： 30

# 字符串拼接：直接放在一起即可
full_name="$a $b"  # 字符串拼接
echo $full_name    # 输出： 10 20
```

##### 4. “可以使用在声明前” (关键补充！)

这个说法需要精确理解。更准确的说法是：**Shell 不会在脚本执行前检查变量是否存在**。

- **如果变量未被赋值，其值为空字符串 (`""`)**。
- 引用一个不存在的变量不会导致脚本崩溃，只会得到一个空值。

```shell
#!/bin/bash

echo "The value of UNKNOWN_VAR is: '$UNKNOWN_VAR'" 
# 输出： The value of UNKNOWN_VAR is: ''

# 然后我们再给这个变量赋值
UNKNOWN_VAR="Now I exist"
echo "The value of UNKNOWN_VAR is: '$UNKNOWN_VAR'"
# 输出： The value of UNKNOWN_VAR is: 'Now I exist'
```

##### 处理未设置变量的高级技巧 (极其重要！)

因为这种特性，Shell 提供了一系列操作符来安全地处理可能未设置的变量。

| 表达式                | 含义                                                         |
| :-------------------- | :----------------------------------------------------------- |
| `${var:-default}`     | 如果 `var` 未设置或为空，则使用 `default`，否则使用 `$var`。**`var` 的值不变**。 |
| `${var:=default}`     | 如果 `var` 未设置或为空，则将其**设置为** `default`，并使用这个值。 |
| `${var:?error_msg}`   | 如果 `var` 未设置或为空，则打印 `error_msg` 并**退出脚本**。用于强制参数检查。 |
| `${var:+replacement}` | 如果 `var` 已设置且不为空，则使用 `replacement`，否则使用空字符串。 |

**示例：**

```shell
echo "Your name is: ${username:-Guest}"  # 如果 username 为空，则使用 "Guest"
# username 变量本身没有被修改

readonly default_path="/usr/local/bin"
PATH="${PATH:=$default_path}" 
# 如果 PATH 为空，则将其设置为 /usr/local/bin，否则保持原样

input_file="${1:?Error: Please provide a filename as the first argument}"
# 如果第一个参数 $1 不存在，脚本会报错并退出

echo "Debug mode is: ${DEBUG:+ON}"
# 如果 DEBUG 变量已设置且不为空，则输出 "ON"，否则输出空。
```

##### 5. 变量作用域

- **默认是全局变量**： 在脚本中定义的变量，在整个脚本中都有效。
- **函数内的局部变量**： 使用 `local` 关键字在函数内创建局部变量，其作用域仅限于该函数。

```shell
#!/bin/bash

global_var="I am global"

my_function() {
    local local_var="I am local"
    global_var="Modified inside function"
    echo "Inside function: $local_var"
    echo "Inside function: $global_var"
}

my_function
echo "Outside function: $local_var"    # 输出： (空)
echo "Outside function: $global_var"   # 输出： Modified inside function
```

##### 6. 特殊变量

Shell 提供了一系列特殊的只读变量，用于访问脚本参数和环境等信息。

| 变量                 | 含义                                                         |
| :------------------- | :----------------------------------------------------------- |
| `$0`                 | 当前脚本的名称                                               |
| `$1`, `$2`, ... `$9` | 脚本的第1个到第9个参数                                       |
| `$#`                 | 传递给脚本的参数个数                                         |
| `$@`                 | 所有参数列表，每个参数都是一个独立的引用字符串 (例如 "`$1`" "`$2`" ...) |
| `$*`                 | 所有参数列表，所有参数被合并为一个字符串 (例如 "`$1 $2 ...`") |
| `$?`                 | 上一个命令的退出状态 (0 表示成功，非0 表示失败)              |
| `$$`                 | 当前 Shell 进程的进程ID (PID)                                |
| `$!`                 | 最后一个后台进程的 PID                                       |

**示例：**

```shell
#!/bin/bash
# 保存为 test.sh

echo "Script name: $0"
echo "First argument: $1"
echo "Number of arguments: $#"
echo "All arguments: $@"

ls /nonexistent_directory &> /dev/null
echo "Exit status of ls command: $?" # 会输出一个非0的错误码
```

运行：`./test.sh arg1 arg2`

------

##### 7. 只读变量和删除变量

- **只读变量**： 使用 `readonly` 声明，其值不可被修改。
- **删除变量**： 使用 `unset` 命令。注意，不能删除只读变量。

```shell
readonly my_constant="This cannot be changed"
my_constant="New value" # 会报错： bash: my_constant: readonly variable

temp_var="Temporary"
unset temp_var
echo $temp_var # 输出： (空)
```

##### 8. 字符串切片：

```shell
string="hello world"
echo ${string:0:5}  # 输出hello
```

##### 9.数组：

Shell也支持数组（一维数组）：

```shell
# 定义数组
fruits=("apple" "banana" "cherry")

# 访问数组元素
echo ${fruits[0]}  # 输出apple

# 访问所有元素
echo ${fruits[@]}

# 数组长度
echo ${#fruits[@]}
```
##### 使用 `read` 命令给变量赋值

###### 1. **基本读取输入**

```shell
#!/bin/bash

echo "请输入您的姓名："
read name
echo "您好, $name!"
```

###### 2. **带提示信息的读取**

```shell
# 使用 -p 选项直接显示提示
read -p "请输入您的年龄: " age
echo "您今年 $age 岁"
```

###### 3. **读取多个变量**

```shell
# 一次读取多个变量，用空格分隔
read -p "请输入姓名、年龄和城市: " name age city
echo "姓名: $name, 年龄: $age, 城市: $city"
```

###### 4. **静默输入（用于密码）**

```shell
# 使用 -s 选项隐藏输入内容
read -s -p "请输入密码: " password
echo  # 换行
echo "密码已接收（不显示）"
```

###### 5. **超时设置**

```shell
# 使用 -t 设置超时时间（秒）
read -t 10 -p "请在10秒内输入: " input
if [ -z "$input" ]; then
    echo "超时未输入！"
else
    echo "您输入了: $input"
fi
```

###### 6. **限制输入长度**

```shell
# 使用 -n 限制输入字符数
read -n 1 -p "是否继续? (y/n): " answer
echo  # 换行
echo "您的选择: $answer"
```

###### 7. **读取到数组**

```shell
# 使用 -a 将输入读取到数组
echo "请输入多个水果名称（空格分隔）:"
read -a fruits
echo "第一个水果: ${fruits[0]}"
echo "所有水果: ${fruits[@]}"
```

###### 8. **使用 IFS 自定义分隔符**

```shell
# 设置自定义分隔符
IFS=":" read -p "输入 姓名:年龄:城市: " name age city
echo "姓名: $name, 年龄: $age, 城市: $city"
```

###### 9. **读取文件内容**

```shell
# 从文件逐行读取
while read line; do
    echo "行内容: $line"
done < filename.txt

# 读取文件到多个变量
while IFS=: read user pass uid gid info home shell; do
    echo "用户: $user, UID: $uid, 家目录: $home"
done < /etc/passwd
```

###### 10. **带默认值的读取**

```shell
#!/bin/bash

read -p "请输入端口号 [默认: 8080]: " port
port=${port:-8080}  # 如果为空，使用默认值8080
echo "使用端口: $port"
```

###### 11. **输入验证**

```shell
#!/bin/bash

while true; do
    read -p "请输入数字 (1-100): " number
    # 检查是否为数字且在范围内
    if [[ "$number" =~ ^[0-9]+$ ]] && [ "$number" -ge 1 ] && [ "$number" -le 100 ]; then
        break
    else
        echo "输入无效，请重新输入！"
    fi
done
echo "您输入的数字: $number"
```

###### 12. **读取命令输出**

```shell
# 从命令输出读取
ls -l | while read perm links owner group size month day time filename; do
    echo "文件: $filename, 大小: $size"
done
```

###### 13. **REPLY 变量**

```shell
# 如果不指定变量名，输入会保存到 $REPLY
read -p "请输入任意内容: "
echo "您输入了: $REPLY"
```

###### 14. **实际应用示例**

```shell
#!/bin/bash

# 用户信息收集脚本
echo "=== 用户信息收集 ==="

read -p "姓名: " name
read -p "年龄: " age
read -p "邮箱: " email
read -s -p "密码: " password
echo

read -p "确认信息是否正确? (y/n): " -n 1 confirm
echo

if [ "$confirm" = "y" ]; then
    echo "信息已保存！"
    # 这里可以添加保存到文件或数据库的代码
else
    echo "请重新运行脚本输入正确信息。"
fi
```

###### 15. **高级用法：菜单选择**

```shell
#!/bin/bash

echo "请选择操作:"
echo "1) 备份文件"
echo "2) 恢复文件" 
echo "3) 退出"

read -p "输入选择 (1-3): " -n 1 choice
echo

case $choice in
    1) echo "执行备份操作..." ;;
    2) echo "执行恢复操作..." ;;
    3) echo "退出程序" ; exit 0 ;;
    *) echo "无效选择!" ;;
esac
```
