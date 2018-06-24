# SensitiveWords

## Get
github地址: https://github.com/TomatoMr/SensitiveWords.git

## Introduction
SensitiveWords是基于DFA算法编写的敏感词汇检测插件,可独立部署,也可以集成到项目中.

## Usage

### 独立部署

#### 1. 复制配置文件
`cd config
 cp config.toml.example config.toml
`

#### 2. 构建二进制包
`go build`

#### 3. 使用方法
`     -restart
            restart your http server, just like this: -restart or -restart=true|false.
      -start [-d]
            up your http server, just like this: -start or -start=true|false [-d or -d=true|false].
      -stop
            down your http server, just like this: -stop or -stop=true|false.
`

#### 4. Api
4.1 /check?content=xxx
`
作用:返回目标文本中,第一个敏感词汇
返回值:target:"", //第一个敏感词
      result:"", //是否含有敏感词
`
4.2 /all?content=xxx
`
作用:返回目标文本中,第一个敏感词汇
返回值:target:[
        word //敏感词
        i //相同的敏感词在原文本中的索引的数组
        l //该敏感词的长度
             ], 
`

### 插件方法
1. GetMap()
`获取SensitiveMap实例`
2. InitDictionary()
`初始化敏感词典,并获得实例`
3. CheckSensitive(text string)
`接受检测文本,并返回是否含有敏感词和第一个敏感词`
4. FindAllSensitive(text string)
`接受检测文本,并返回所有敏感词`
5. GetConfig()
`返回配置实例`

### 配置文件说明
`
DictionaryPath //敏感词典地址,根目录是本项目地址
Port //http server监听的web端口
PidFilePath //pid文件位置,用于命令行结束程序和重启程序,根目录是本项目地址
`

### 帮助
Q:重载词典?
A:修改config.toml->修改DictionaryPath->./SensitiveWords -restart