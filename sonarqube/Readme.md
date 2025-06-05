To run sonarqube code coverage:
   >- cd hedge
   >- make codecoverage
   >- docker run --network=host -it -e sonar_token_go=[BE GO authtoken] -e sonar_token_ui=[UI authtoken] -e sonar_token_py=[PY authtoken] -e sonar_token_all=[ALL authtoken] hedge-test-coverage

`make codecoverage` creates a docker image named _hedge-test-coverage_, and \
`docker run ...` above, runs this image to generate coverage files and apply it on our sonarqube server http://xxx.yyyy:9000 

\
\
\
===== Below is deprecated - replaced by simplified process above ===== \
\
To run sonarqube code coverage, a coverage file needs to be created and then passed to sonar-scanner cli

\
**For overall/combined coverage runs**
1. Run code analysis and coverage report with sonarqube (this uses the sonar-project.properties file under sonarqube dir)
   >- cd hedge
   >- docker run --network=host --name sonar_scanner -it --rm -v ".:/hedge" -v "./sonarqube/sonar-project.properties:/hedge/sonarqube/sonar-project.properties" sonarsource/sonar-scanner-cli
   >- rm sonar-project.properties


\
**For golang code coverage runs only**
1. Generate a code coverage file by running  
   >- cd hedge
   >- go test -json ./... -covermode=atomic -coverprofile=sonarqube/coverage.go.out

2. Run code analysis and coverage report with sonarqube (this uses the sonar-project.properties_go file under sonarqube dir)
   >- cd hedge
   >- docker run --network=host --name sonar_scanner -it --rm -v ".:/hedge" -v "./sonarqube/sonar-project.properties_go:/hedge/sonar-project.properties" sonarsource/sonar-scanner-cli
   >- rm sonar-project.properties


\
**For UI code coverage runs only**
1. Generate a UI code coverage file
   >- The Jenkins job creates it at hedge/ui/edge-portal/coverage/shell/lcov.info
   >- A sample coverage file is also kept at [sonar-project.properties_ui](sonar-project.properties_ui) if incase you need to run a scan outside of Jenkins.
   >- You can overwrite this sample file with an updated coverage file 

2. Run code analysis and coverage report with sonarqube (this uses the sonar-project.properties_ui file under sonarqube dir)
   >- cd hedge
   >- docker run --network=host --name sonar_scanner -it --rm -v ".:/hedge" -v "./sonarqube/sonar-project.properties_ui:/hedge/sonar-project.properties" sonarsource/sonar-scanner-cli
   >- rm sonar-project.properties
