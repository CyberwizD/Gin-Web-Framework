# Gin Web Framework

## Features
* **Fast**
Radix tree based routing, small memory foot print. No reflection. Predictable API performance.

* **Middleware support**
An incoming HTTP request can be handled by a chain of middlewares and the final action. For example: Logger, Authorization, GZIP and finally post a message in the DB.

* **Crash-free**
Gin can catch a panic occurred during a HTTP request and recover it. This way, your server will be always available. As an example - it’s also possible to report this panic to Sentry!

* **JSON validation**
Gin can parse and validate the JSON of a request - for example, checking the existence of required values.

* **Routes grouping**
Organize your routes better. Authorization required vs non required, different API versions… In addition, the groups can be nested unlimitedly without degrading performance.

* **Error management**
Gin provides a convenient way to collect all the errors occurred during a HTTP request. Eventually, a middleware can write them to a log file, to a database and send them through the network.