#!/bin/bash

# 定义源、目标文件路径
source_file="/opt/fcihpy/nohup.out"
target_file="/opt/fcihpy/bak.log"

# 设定文件最大值,单位是字节
max_size=100000

# 获取操作系统名称
os_name=$(uname -s)

# 根据不同的操作系统选择合适的命令获取文件大小
if [ "$os_name" = "Linux" ]; then
  file_size=$(stat -c %s "$source_file")
elif [ "$os_name" = "Darwin" ]; then
  file_size=$(stat -f%z "$source_file")
else
  echo "Unsupported operating system: $os_name"
  exit 1
fi

# 转换文件大小为更易读的单位
if [ "$file_size" -ge 1048576 ]; then
  file_size_readable=$(echo "scale=2; $file_size / 1048576" | bc)
  unit="MB"
elif [ "$file_size" -ge 1024 ]; then
  file_size_readable=$(echo "scale=2; $file_size / 1024" | bc)
  unit="KB"
else
  file_size_readable="$file_size"
  unit="bytes"
fi

# 判断是否超过限额
if [ "$file_size" -gt "$max_size" ]; then
    # 记录最后10条数据
    tail -n 100 "$source_file" > "$target_file"
    # 将最后10行记录下来，将执行清空操作
    echo -n "" > "$source_file"
else
    echo "File size $file_size_readable$unit is within the limit,no need to clear"
fi

