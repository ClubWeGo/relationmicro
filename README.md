# relationmicor
社交模块


聚合message的方法到relation的生成代码中的方案
https://www.cloudwego.io/zh/docs/kitex/tutorials/code-gen/combine_service/

使用combine-service 合并多个分块的service
`kitex --combine-service  -service github.com/ClubWeGo/relationmicro -module github.com/ClubWeGo/relationmicro ./idl/relation.thrift`

`go get github.com/cloudwego/kitex@latest && go mod tidy`

注意生成的handler中是CombineServiceImpl