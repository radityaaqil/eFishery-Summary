// Professional is calm and decisive
// Pro SE test their code
// Unit, End to End, Coverage
// Collaboration is important, Inattention to Results, avoidance of accountability, lack of commitment, fear of conflict, absence if tryst
// Apprenticeship, flase idea that once you graduate you can code, bugs equal money loss, school does not and cannot teach the discipline
// When performance matters, professionals practice
// =================================================
// Path
// Associate --> Junior --> Senior --> Distinguished
// Associate --> Junior --> Engineering Manager --> VP Engineering
// Senior to Engineering Manager leadership training
// Berlatih, bersabar, berbagi

// =================================================
// SOFTWARE DEVELOPMENT LIFECYCLE
// Analysis, Design, Development, Testing, Deployment, Maintenance
// MVP : Minimum Viable Product
// SCRUM

// SOFTWARE ARCHITECTURE
// What is architecture?
// Function & Module
// Monolith
// Microservice

// Architecture : Akses db, layer2, http req handling, security, performance
// Function & Module : Module --> System inside architecture yang handle business logic
// Monolith Architecture : All business logic in one place and also the UI (MONO) pros:simple to develope, test e.g end to end with selenium, deploy and scale horizontally by running multiple copies behind a load balancer (nginx) cons: limitation in size and complexity, scaling is hard, CI/CD is difficult, reliability (memory leak cound bring down entire process), has a barrier in adopting new technology as changes in frameworks or languages will affect application in both time and cost.
// Microservices Architecture : pros:divide apps into a set of manageable services which are much faster to develop and much easier o understand and maintain, enables each service to be deveoped independently, reduces barrier of adopting new technologies, independent deployment, enables each service to be scaled independently, easy to maintain (one broken service would not affect the entire system) cons:more complex to deploy and test, partitioned db architecture, harder to implement changes in multiple services
// Master Slave Database (replication of main DB)
