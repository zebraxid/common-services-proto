# common-services-proto

Common Services Protocol Buffer Repository

## Contribution
* Clone the repository, if already cloned, do git pull first
* Create your working branch from master branch and do update the proto file as the changes request
* Create pull request and ask others to review it
* When PR merged, it will trigger jenkins jobs and :
1. If proto file changed, jenkins will compile the go code and push back to this repository
2. If pom.xml file changed it will build maven package and deploy it to our github package. (Be carefull for the pom.xml versioning!!!!)
* Do taging to lock your changes version
