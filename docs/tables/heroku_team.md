# Table: heroku_team

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ | When the team was created. | 
| id | string | X | √ | Unique identifier of team. | 
| role | string | X | √ | Role in the team. | 
| updated_at | timestamp | X | √ | When the team was updated. | 
| membership_limit | float | X | √ | Upper limit of members allowed in a team. | 
| provisioned_licenses | bool | X | √ | Whether the team is provisioned licenses by salesforce. | 
| team_type | string | X | √ | Type of team. | 
| name | string | X | √ | Unique name of team. | 
| credit_card_collections | bool | X | √ | Whether charges incurred by the team are paid by credit card. | 
| is_default | bool | X | √ | Whether to use this team when none is specified. | 
| enterprise_account | json | X | √ | Enterprise Account associated with the team. | 
| identity_provider | json | X | √ | Identity Provider associated with the team. | 


