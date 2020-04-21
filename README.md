# runGitDag

runGitDag is a way to run users job code from version control on your environment. It enables analysts and data
scientists to have more freedom and velocity by decoupling their work from data engineering and operations deploy cycle.

## Features

Connection:
  - HOST
  - PORT
  - TYPE (HIVE, IMPALA, PRESTO, etc)
  - OPTIONS
  
Service_User:
  - USERNAME
  - PASSWORD
  - CONNECTION
  
Job:
  - REPO_URL
  - REPO_PATH
  - (other configuration?)
  
Schedule:
  - CRON
  - (Maybe this should be part of a job? Or keep schedules and jobs seperate?)

User:
  - Username
  - Password
  - Type
  - 

## Stories
 What are some of the ideas for this application?
 
  - Admin:
    - Login
    - Create 'connection'
    - Create 'service_user'
    
  - Users:
    - Login
    - Create 'job'
    - Create 'schedule'

## TODOs

  - Write in the README about what we are doing
  - 