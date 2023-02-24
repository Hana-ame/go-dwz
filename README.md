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

## frontend