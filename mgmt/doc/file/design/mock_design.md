##### 支持的Mock路由
- http://{IP}:{PORT}/mock/file/:name?lang=en 
- http://{IP}:{PORT}/mock/file?name=name.xml&lang=en
  - 可返回上传文件下name.xml的内容，如果name.xml中有占位符，会生成对应特征的数据,lang不传默认中文
  - 返回的文件支持.xml, .json, .yml, .txt等文本类型的文件
  - 返回文件模板示例: 
   - [JSON类型](../../../upload/create_template.json)
   - [XML类型](../../../upload/create_template.xml)
   - [TXT类型](../../../upload/create_template.txt)
   - [YAML类型](../../../upload/create_template.yml)
- http://{IP}:{PORT}/mock/data/quick  
  - 可返回JSON格式的数据，内容如下样例信息
  - 后续版本可能会废弃，与其他功能进行合并
- http://{IP}:{PORT}/mock/data/slow?sleep=N  
  - N为整数，表示等待的时间，模拟真实三方请求返回有延迟的情况，可返回JSON格式的数据
  - 后续版本可能会废弃，与其他功能进行合并
##### 样例信息，XX信息是随机的：
```
{
    addr: "XX市XX市XX路XX号XX小区XX单元XX室",
    age: "XX",
    bankCard: "XXXXX",
    bool: "XX",
    grade: "X",
    id_no: "XXXXX",
    mobile: "XX",
    name: "XXXXX",
    sex: "XXX",
    weight: "XX",
    yesOrNo: "XX"
}
```

- http://{IP}:{PORT}/mock/data/certid/:idno  传入大陆的身份证件号码，可返回基本信息
- http://{IP}:{PORT}/mock/data/certid?idno=XXX  传入大陆的身份证件号码，可返回基本信息
##### 生成的信息均为测试数据，不具有实际意义参考
```
{
     "city": "XX",
     "sex": "XX",
     "code": "XXXXX",
     "birthday": "XX",
     "district": "X",
     "province": "XXXXX",
     "address": "XX市XX市XX路XX号XX小区XX单元XX室",
     "country": "中国",
}
```
- http://{IP}:{PORT}/mock/systemParameter/:name?lang=en  传入系统参数中的名称，随机返回定义值列表中的一个值，支持多语种
- http://{IP}:{PORT}/mock/systemParameter?name=XXX&lang=en  传入系统参数中的名称，随机返回定义值列表中的一个值，支持多语种