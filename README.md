## API

### go

```
GET /go?id=:id
```

redirect to the url specified by id


### Link

```
POST /link
```

create a link

should have

```
Url
Description
```


```
GET /link?id=:id
```
get link object by id

```
DELETE /link?id=:id
```
delete link object by id

### Tag

```
GET /tags/:tag
```

gat Links by tag

```
POST /tags/:tag?id=:id
```

add id to a tag

## Web-UI

排序方式之类的。
懒加载怎么弄啊。

TODO:
排序方式
动画
更好的排版
为Link做组件。

你这连增加tag都不行。
来个菜单？

历史记录
收藏

所有Links取出。

啊还有添加连接的页面，图了。
