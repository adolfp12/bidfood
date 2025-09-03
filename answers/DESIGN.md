- How would you manage high-concurrency in a Go microservice (thousands of requests per second)?
>>  Use caching to speed up data access. Choose libraries that use less memory and CPU. Using goroutine wisely, check if use cases are fit with goroutine and if needed, use channel to communicate between goroutines. Use gRPC for fast communication between services. Optimize your database by pooling, adding indexes, avoiding SELECT * to get data (use select efficiently), sharding data, and picking the right type of database. Take advantage of the cloud with horizontal scaling and containerization. Use asynchronous methods like messaging to handle tasks without waiting. Apply concurrency when needed to run tasks in parallel. Regularly perform load testing and improve your code based on the results. Use profiling tools to find and fix bottlenecks in your system.


- Recommended project structure for large Go services?
>> I choose to scale by separating the system into parts: handlers (APIs for external use), business logic, and data storage. Usually, I organize services by domain, like payment or product microservices. I focus on clean code, modular design, and easy-to-understand structure. For example I use it for answering question #1


- Approach to configuration management in production?
>> Should have separate configuration for production, staging and development (like a vault config such as Consul) and periodically changed due to security issue such as DB username/password, hashing key, encryption key, etc. Dont put/write production config on files and push to repository/git. Do not expose sensitive data or unwanted information in the code. Also minimize people accesing the config by separation duty of the person.


- Observability strategy (logging, metrics, tracing)?
>> logging needed for tracing error and monitoring the service. Logging need to maintain by rolling log daily, weekly or monthly by cron task or spesific service. Log must be standardize to make it ease to check.
>> Metric also need to define. Needed metric should define and push into proper monitoring tool. Metrics (technical metric and business metric) should be divided base on the stakeholder. For metric tool we can use open source tools such as Grafana by using Promotheus metrics or Kibana. Or with paid tool such as Data Dog or New relic. For Business metrics, can use open source like Grafana or Kibana. Or can use paid tools such as Tableau or PowerBI. 
>> Tracing needed to check errors, improve performance and flow of the codes. We can use existing GO tracing package, Jaeger


- Go API framework of choice (e.g., Gin, Chi) and why
>> Choosing a Go API framework should be based on our needs, such as the framework’s support, ease of use, available libraries, scalability and performance. It’s important to research and evaluate these factors by building a proof of concept (POC). Select the framework that best fits your requirements. For example (not on framework, but similar reason), I once switched HTTP libraries because the previous one was no longer supported.