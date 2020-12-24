client 主板和内存信息收集需要依赖于 dmidecode 命令
client 硬盘信息收集优先使用 smartctl 命令，如果 smartctl 不存在则使用 ioctl 系统调用或 gopsutil
smartctl 尽量使用 7.0 版本
