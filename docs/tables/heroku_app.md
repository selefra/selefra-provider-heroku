# Table: heroku_app

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| git_url | string | X | √ | Git repo URL of app. | 
| owner | json | X | √ | Identity of app owner. | 
| slug_size | int | X | √ | Slug size in bytes of app. | 
| stack | json | X | √ | Identity of app stack. | 
| name | string | X | √ | Unique name of app. | 
| buildpack_provided_description | string | X | √ | Description from buildpack of app. | 
| maintenance | bool | X | √ | Maintenance status of app. | 
| released_at | timestamp | X | √ | When app was released. | 
| updated_at | timestamp | X | √ | When app was updated. | 
| archived_at | timestamp | X | √ | When app was archived. | 
| created_at | timestamp | X | √ | When app was created. | 
| id | string | X | √ | Unique identifier of app. | 
| organization | json | X | √ | Identity of team. | 
| web_url | string | X | √ | Web URL of app. | 
| acm | string | X | √ | ACM status of this app. | 
| build_stack | json | X | √ | Identity of the stack that will be used for new builds. | 
| repo_size | int | X | √ | Git repo size in bytes of app. | 
| space | json | X | √ | Identity of space. | 
| team | json | X | √ | identity of team. | 
| internal_routing | bool | X | √ | Describes whether a Private Spaces app is externally routable or not. | 
| region | json | X | √ | Identity of app region. | 


