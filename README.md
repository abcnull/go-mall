# 研发实际项目小疑问

传递的参数既有 file 又有 json，怎么写格式，都写成 form 吗

离线表和实时表区别

什么时候开始考虑使用 mq

什么时候开始考虑使用 es

上游 http 服务假如某个接口调用会很多，把数据就给到 kafka，kafka 数据被下游消费之后，下游做一些逻辑处理后怎么返回给最上游？

消费 kafka 的只能是 faas，不能是一个 rpc 服务吗，如果是一个 rpc 服务消费 kafka，消费的监听 kafka 的代码怎么写呢？

gorm 里头能否不拿到结构体直接拿到某个列字段的值？

如果多表查询，先查 A 表，再查 B 表，但是要依据 A 中是字段排序，在 B 中又要进行分页查询怎么办