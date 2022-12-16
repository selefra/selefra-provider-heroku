# Table: heroku_account

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_login | timestamp | X | √ | When account last authorized with Heroku. | 
| updated_at | timestamp | X | √ | When account was updated. | 
| default_team | json | X | √ | Team selected by default. | 
| delinquent_at | timestamp | X | √ | When account became delinquent. | 
| id | string | X | √ | Unique identifier of an account. | 
| identity_provider | json | X | √ | Identity Provider details for federated users. | 
| italian_partner_terms | string | X | √ | Whether account has acknowledged the Italian provider terms of service. | 
| sms_number | string | X | √ | SMS number of account. | 
| name | string | X | √ | Full name of the account owner. | 
| acknowledged_msa | bool | X | √ | Whether account has acknowledged the MSA terms of service. | 
| allow_tracking | bool | X | √ | Whether to allow third party web activity tracking. | 
| created_at | timestamp | X | √ | When account was created. | 
| default_organization | json | X | √ | Team selected by default. | 
| acknowledged_msa_at | timestamp | X | √ | When account has acknowledged the MSA terms of service. | 
| beta | bool | X | √ | Whether allowed to utilize beta Heroku features. | 
| country_of_residence | string | X | √ | Country where account owner resides. | 
| italian_customer_terms | string | X | √ | Whether account has acknowledged the Italian customer terms of service. | 
| two_factor_authentication | bool | X | √ | Whether two-factor auth is enabled on the account. | 
| email | string | X | √ | Unique email address of account. | 
| federated | bool | X | √ | Whether the user is federated and belongs to an Identity Provider. | 
| suspended_at | timestamp | X | √ | When account was suspended. | 


