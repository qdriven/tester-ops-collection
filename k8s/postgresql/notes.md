Completed Notes:

```sh

NAME: postgresql-dev
LAST DEPLOYED: Wed Dec 29 15:34:14 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: postgresql
CHART VERSION: 10.14.0
APP VERSION: 11.14.0

** Please be patient while the chart is being deployed **

PostgreSQL can be accessed via port 5432 on the following DNS names from within your cluster:

    postgresql-dev.default.svc.cluster.local - Read/Write connection

To get the password for "postgres" run:

    export POSTGRES_PASSWORD=$(kubectl get secret --namespace default postgresql-dev -o jsonpath="{.data.postgresql-password}" | base64 --decode)

To connect to your database run the following command:

    kubectl run postgresql-dev-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:11.14.0-debian-10-r17 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host postgresql-dev -U postgres -d postgres -p 5432



To connect to your database from outside the cluster execute the following commands:

    kubectl port-forward --namespace default svc/postgresql-dev 5432:5432 &
    PGPASSWORD="$POSTGRES_PASSWORD" psql --host 127.0.0.1 -U postgres -d postgres -p 5432
```

https://github.com/bitnami/charts/tree/master/bitnami/postgresql