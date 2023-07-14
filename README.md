## go-test-endeus

## HOW TO RUN LOCALLY
- run test
  
      make test
      or
      make test-detail
    
- prepare database mysql for this test
- copy .env.example to .env file and fill all the environment
- running migration command
    
      make migrateup
      
- drop migration command

      make migratedown
      
- the last step is to run the program itself with command
  
      make run-dev
      
      
## API DOCS URL
- open this url too see all the list of api existing 

        http://localhost:<PORT>/swagger/index.html

