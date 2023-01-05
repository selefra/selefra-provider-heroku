# Table: heroku_apps

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| updated_at | timestamp | X | √ |  | 
| id | string | √ | √ |  | 
| internal_routing | bool | X | √ |  | 
| name | string | X | √ |  | 
| organization | json | X | √ |  | 
| region | json | X | √ |  | 
| released_at | timestamp | X | √ |  | 
| team | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| git_url | string | X | √ |  | 
| owner | json | X | √ |  | 
| archived_at | timestamp | X | √ |  | 
| repo_size | int | X | √ |  | 
| stack | json | X | √ |  | 
| web_url | string | X | √ |  | 
| acm | bool | X | √ |  | 
| build_stack | json | X | √ |  | 
| buildpack_provided_description | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| maintenance | bool | X | √ |  | 
| slug_size | int | X | √ |  | 
| space | json | X | √ |  | 


