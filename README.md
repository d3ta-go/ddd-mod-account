# ddd-mod-account ( [![Go Report Card](https://goreportcard.com/badge/github.com/d3ta-go/ddd-mod-account)](https://goreportcard.com/report/github.com/d3ta-go/ddd-mod-account) )

DDD Module: Account (Generic Subdomain)

As a part of `Simple Implementation of Modular DDD Technical Architecture Patterns in Go`.

## Diagram v 0.2.2-Modular

![DDD-Technical-Architecture-Patterns-Golang-0.2.2-DDD Account Module](docs/img/DDD-Technical-Architecture-Patterns-Golang-0.2.2-DDD_Account_Module.png)

## Components

A. Interface Layer (None)

B. DDD Modules:

1. Account - using DDD Layered Architecture (Contract, Adapters) [ [d3ta-go/ddd-mod-account](https://github.com/d3ta-go/ddd-mod-account) ]

2. Email (Indirect) - using DDD Layered Architecture (Contract, GORM, SMTP) [ [d3ta-go/ddd-mod-email](https://github.com/d3ta-go/ddd-mod-email) ]

C. Common System Libraries [ [d3ta-go/system](https://github.com/d3ta-go/system) ]:

1. Configuration - using yaml
2. Identity & Securities - using JWT, Casbin (RBAC)
3. Initializer
4. Email Sender - using SMTP
5. Handler
6. Migrations
7. Utils

D. Databases

1. MySQL (tested)
2. PostgreSQL (untested)
3. SQLServer (untested)
4. SQLite3 (untested)

E. Persistent Caches

1. Session/Token/JWT Cache (Redis, File, DB, etc) [tested: Redis]

F. Messaging [to-do]

G. Logs [to-do]

### Development

**1. Clone**

```shell
$ git clone https://github.com/d3ta-go/ddd-mod-account.git
```

**2. Setup**

```
a. copy `conf/config-sample.yaml` to `conf/config.yaml`
b. copy `conf/data/test-data-sample.yaml` to `conf/data/test-data.yaml`
c. setup your dependencies/requirements (e.g: database, redis, smtp, etc.)
```

**3. Runing TDD on Development Stage**

Pre-requisites:

> TDD script in DDD Email Modul [ [d3ta-go/ddd-mod-email](https://github.com/d3ta-go/ddd-mod-email) ] section "3.2. TDD: Functionality Test (unit test)" has been executed.

3.1. TDD: DB Migration Test

```shell
$ cd ddd-mod-account
$ sh tdd/clean-testcache.sh

$ sh tdd/ut.db.migration.run-001.sh

$ sh tdd/ut.db.migration.rollback-001.sh
```

3.2. TDD: Functionality Test (unit test)

```shell
$ cd ddd-mod-account
$ sh tdd/clean-testcache.sh

$ sh tdd/ut.db.migration.run-001.sh

# Hold !!!
# ------------------------------------------------------
# To do:
# Set all `ClientKey` and `SecretKey` fields value in `conf/data/test-data.yaml`
# based on default client application in database table:
# `iam_sys_user_client_apps`
# ------------------------------------------------------

$ sh tdd/ut.pkg.infra-layer.repository-001.sh

$ sh tdd/ut.pkg.app-layer.service-001.sh

$ sh tdd/ut.pkg.app-layer.application-001.sh
```

**TDD Result Sample:**

- MySQL Database:

![MySQL Database Migration Result](docs/img/account-sample-db-migration-mysql.png)

- MailHog (SMTP Email Server):

![MailHog/SMTP Email Server Result](docs/img/sample-smtp-email-server-MailHog.png)
