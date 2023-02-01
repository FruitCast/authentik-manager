# authentik

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2022.11.3](https://img.shields.io/badge/AppVersion-2022.11.3-informational?style=flat-square)

Unofficial Authentik single-sign-on chart

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://../operator | akm | 0.1.0 |
| https://charts.bitnami.com/bitnami | postgresql | 12.1.2 |
| https://charts.bitnami.com/bitnami | redis | 17.3.11 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| akm.operator.enabled | bool | `true` |  |
| authentik.blueprints | list | `[]` |  |
| authentik.config.file | string | `"configuration.yml"` |  |
| authentik.config.generate | bool | `true` |  |
| authentik.config.name | string | `"authentik-config"` |  |
| authentik.config.noteFile | string | `"/notification.txt"` |  |
| authentik.config.path | string | `"/config"` |  |
| authentik.deployment.env[0].name | string | `"AUTHENTIK_LISTERN__HTTP"` |  |
| authentik.deployment.env[0].value | string | `"0.0.0.0:9000"` |  |
| authentik.deployment.env[10].name | string | `"AUTHENTIK_DEFAULT_USER_CHANGE_NAME"` |  |
| authentik.deployment.env[10].value | bool | `true` |  |
| authentik.deployment.env[11].name | string | `"AUTHENTIK_DEFAULT_USER_CHANGE_EMAIL"` |  |
| authentik.deployment.env[11].value | bool | `true` |  |
| authentik.deployment.env[12].name | string | `"AUTHENTIK_DEFAULT_USER_CHANGE_USERNAME"` |  |
| authentik.deployment.env[12].value | bool | `true` |  |
| authentik.deployment.env[13].name | string | `"AUTHENTIK_GDPR_COMPLIANCE"` |  |
| authentik.deployment.env[13].value | bool | `true` |  |
| authentik.deployment.env[14].name | string | `"AUTHENTIK_DEFAULT_TOKEN_LENGTH"` |  |
| authentik.deployment.env[14].value | int | `60` |  |
| authentik.deployment.env[15].name | string | `"AUTHENTIK_IMPERSONATION"` |  |
| authentik.deployment.env[15].value | bool | `true` |  |
| authentik.deployment.env[16].name | string | `"AUTHENTIK_WEB__WORKERS"` |  |
| authentik.deployment.env[16].value | int | `2` |  |
| authentik.deployment.env[17].name | string | `"AUTHENTIK_WEB__THREADS"` |  |
| authentik.deployment.env[17].value | int | `2` |  |
| authentik.deployment.env[1].name | string | `"AUTHENTIK_LISTERN__HTTPS"` |  |
| authentik.deployment.env[1].value | string | `"0.0.0.0:9443"` |  |
| authentik.deployment.env[2].name | string | `"AUTHENTIK_LISTERN__LDAP"` |  |
| authentik.deployment.env[2].value | string | `"0.0.0.0:3389"` |  |
| authentik.deployment.env[3].name | string | `"AUTHENTIK_LISTERN__LDAPS"` |  |
| authentik.deployment.env[3].value | string | `"0.0.0.0:6636"` |  |
| authentik.deployment.env[4].name | string | `"AUTHENTIK_LISTERN__METRICS"` |  |
| authentik.deployment.env[4].value | string | `"0.0.0.0:9300"` |  |
| authentik.deployment.env[5].name | string | `"AUTHENTIK_LISTERN__DEBUG"` |  |
| authentik.deployment.env[5].value | string | `"0.0.0.0:9900"` |  |
| authentik.deployment.env[6].name | string | `"AUTHENTIK_LOG_LEVEL"` |  |
| authentik.deployment.env[6].value | string | `"info"` |  |
| authentik.deployment.env[7].name | string | `"AUTHENTIK_DISABLE_UPDATE_CHECK"` |  |
| authentik.deployment.env[7].value | bool | `false` |  |
| authentik.deployment.env[8].name | string | `"AUTHENTIK_ERROR_REPORTING"` |  |
| authentik.deployment.env[8].value | bool | `false` |  |
| authentik.deployment.env[9].name | string | `"AUTHENTIK_AVATARS"` |  |
| authentik.deployment.env[9].value | string | `"gravatar"` |  |
| authentik.deployment.image | string | `"ghcr.io/goauthentik/server:2023.1.2"` |  |
| authentik.deployment.imagePullPolicy | string | `"Always"` |  |
| authentik.deployment.name | string | `"authentik"` |  |
| authentik.deployment.replicas | int | `3` |  |
| authentik.enabled | bool | `true` |  |
| authentik.labels[0].key | string | `"type"` |  |
| authentik.labels[0].value | string | `"auth"` |  |
| authentik.labels[1].key | string | `"app"` |  |
| authentik.labels[1].value | string | `"authentik"` |  |
| authentik.ports[0].containerPort | int | `9000` |  |
| authentik.ports[0].name | string | `"http"` |  |
| authentik.ports[0].protocol | string | `"TCP"` |  |
| authentik.ports[0].servicePort | int | `80` |  |
| authentik.ports[1].containerPort | int | `9443` |  |
| authentik.ports[1].name | string | `"https"` |  |
| authentik.ports[1].protocol | string | `"TCP"` |  |
| authentik.ports[1].servicePort | int | `443` |  |
| authentik.ports[2].containerPort | int | `3389` |  |
| authentik.ports[2].name | string | `"ldap"` |  |
| authentik.ports[2].protocol | string | `"TCP"` |  |
| authentik.ports[2].servicePort | int | `389` |  |
| authentik.ports[3].containerPort | int | `6636` |  |
| authentik.ports[3].name | string | `"ldaps"` |  |
| authentik.ports[3].protocol | string | `"TCP"` |  |
| authentik.ports[3].servicePort | int | `636` |  |
| authentik.ports[4].containerPort | int | `9300` |  |
| authentik.ports[4].name | string | `"metrics"` |  |
| authentik.ports[4].protocol | string | `"TCP"` |  |
| authentik.ports[4].servicePort | int | `300` |  |
| authentik.ports[5].containerPort | int | `9900` |  |
| authentik.ports[5].name | string | `"debug"` |  |
| authentik.ports[5].protocol | string | `"TCP"` |  |
| authentik.ports[5].servicePort | int | `900` |  |
| authentik.secrets.basePath | string | `"/secrets"` |  |
| authentik.secrets.lookup[0].env | string | `"AUTHENTIK_POSTGRESQL__PASSWORD"` |  |
| authentik.secrets.lookup[0].file | string | `"postgresql-pass"` |  |
| authentik.secrets.lookup[0].key | string | `"postgresUserPassword"` |  |
| authentik.secrets.lookup[1].env | string | `"AUTHENTIK_REDIS__PASSWORD"` |  |
| authentik.secrets.lookup[1].file | string | `"redis-pass"` |  |
| authentik.secrets.lookup[1].key | string | `"redisPassword"` |  |
| authentik.secrets.lookup[2].env | string | `"AUTHENTIK_SECRET_KEY"` |  |
| authentik.secrets.lookup[2].file | string | `"secret-key"` |  |
| authentik.secrets.lookup[2].key | string | `"authJwtToken"` |  |
| authentik.secrets.lookup[3].env | string | `"AUTHENTIK_EMAIL__PASSWORD"` |  |
| authentik.secrets.lookup[3].file | string | `"smtp-pass"` |  |
| authentik.secrets.lookup[3].key | string | `"smtpPassword"` |  |
| authentik.service.name | string | `"authentik"` |  |
| global.admin.email | string | `"somebody@pm.me"` |  |
| global.admin.name | string | `"somebody"` |  |
| global.domain.base | string | `"example.org"` |  |
| global.domain.full | string | `"auth.example.org"` |  |
| global.domain.ldap | string | `"ldap.example.org"` |  |
| ingress.authCertSecret | string | `"authentik-cert"` |  |
| ingress.clusterIssuer | string | `"letsencrypt-staging"` |  |
| ingress.enable | bool | `true` |  |
| ingress.ldapCertSecret | string | `"ldap-cert"` |  |
| ingress.maxBodySize | string | `"32M"` |  |
| ingress.name | string | `"auth"` |  |
| ingress.tls.enable | bool | `true` |  |
| ldap.deployment.env[0].name | string | `"LDAP_ADMIN_USERNAME"` |  |
| ldap.deployment.env[0].value | string | `"admin"` |  |
| ldap.deployment.image | string | `"docker.io/nitnelave/lldap:v0.4.1-alpine"` |  |
| ldap.deployment.imagePullPolicy | string | `"Always"` |  |
| ldap.deployment.name | string | `"ldap"` |  |
| ldap.deployment.replicas | int | `1` |  |
| ldap.enabled | bool | `false` |  |
| ldap.labels[0].key | string | `"type"` |  |
| ldap.labels[0].value | string | `"auth"` |  |
| ldap.labels[1].key | string | `"app"` |  |
| ldap.labels[1].value | string | `"ldap"` |  |
| ldap.persistence.accessMode | string | `"ReadWriteOnce"` |  |
| ldap.persistence.enabled | bool | `false` |  |
| ldap.persistence.mountPath | string | `"/data"` |  |
| ldap.persistence.name | string | `"ldap-pvc"` |  |
| ldap.persistence.size | string | `"8Gi"` |  |
| ldap.ports[0].containerPort | int | `3890` |  |
| ldap.ports[0].name | string | `"ldap"` |  |
| ldap.ports[0].protocol | string | `"TCP"` |  |
| ldap.ports[0].servicePort | int | `3890` |  |
| ldap.ports[1].containerPort | int | `17170` |  |
| ldap.ports[1].name | string | `"http"` |  |
| ldap.ports[1].protocol | string | `"TCP"` |  |
| ldap.ports[1].servicePort | int | `80` |  |
| ldap.secrets.basePath | string | `"/secrets"` |  |
| ldap.secrets.lookup[0].env | string | `"LLDAP_JWT_SECRET"` |  |
| ldap.secrets.lookup[0].key | string | `"authJwtToken"` |  |
| ldap.secrets.lookup[1].env | string | `"LLDAP_LDAP_USER_PASS"` |  |
| ldap.secrets.lookup[1].key | string | `"ldapAdminPassword"` |  |
| ldap.secrets.secretName | string | `"auth"` |  |
| ldap.service.name | string | `"ldap"` |  |
| pgadmin.deployment.env[0].name | string | `"PGADMIN_LISTEN_PORT"` |  |
| pgadmin.deployment.env[0].value | int | `80` |  |
| pgadmin.deployment.image | string | `"docker.io/dpage/pgadmin4:latest"` |  |
| pgadmin.deployment.imagePullPolicy | string | `"Always"` |  |
| pgadmin.deployment.name | string | `"pgadmin"` |  |
| pgadmin.deployment.replicas | int | `1` |  |
| pgadmin.enable | bool | `false` |  |
| pgadmin.labels[0].key | string | `"type"` |  |
| pgadmin.labels[0].value | string | `"auth"` |  |
| pgadmin.labels[1].key | string | `"app"` |  |
| pgadmin.labels[1].value | string | `"pgadmin"` |  |
| pgadmin.persistence.accessMode | string | `"ReadWriteOnce"` |  |
| pgadmin.persistence.enable | bool | `false` |  |
| pgadmin.persistence.mountPath | string | `"/var/lib/pgadmin/data"` |  |
| pgadmin.persistence.name | string | `"pgadmin-pvc"` |  |
| pgadmin.persistence.size | string | `"4Gi"` |  |
| pgadmin.ports[0].containerPort | int | `80` |  |
| pgadmin.ports[0].name | string | `"http-port"` |  |
| pgadmin.ports[0].protocol | string | `"TCP"` |  |
| pgadmin.ports[0].servicePort | int | `80` |  |
| pgadmin.secrets.lookup[0].env | string | `"PGADMIN_DEFAULT_PASSWORD"` |  |
| pgadmin.secrets.lookup[0].key | string | `"pgAdminPassword"` |  |
| pgadmin.secrets.secretName | string | `"auth"` |  |
| pgadmin.service.name | string | `"pgadmin"` |  |
| postgresql.fullnameOverride | string | `"postgres"` |  |
| postgresql.global.postgresql.auth.database | string | `"authentik"` |  |
| postgresql.global.postgresql.auth.existingSecret | string | `"auth"` |  |
| postgresql.global.postgresql.auth.secretKeys.adminPasswordKey | string | `"postgresPassword"` |  |
| postgresql.global.postgresql.auth.secretKeys.replicationPasswordKey | string | `"postgresReplicationPassword"` |  |
| postgresql.global.postgresql.auth.secretKeys.userPasswordKey | string | `"postgresUserPassword"` |  |
| postgresql.global.postgresql.auth.username | string | `"authentik"` |  |
| postgresql.global.postgresql.service.ports.postgresql | int | `5432` |  |
| postgresql.image.registry | string | `"docker.io"` |  |
| postgresql.image.repository | string | `"bitnami/postgresql"` |  |
| postgresql.image.tag | int | `15` |  |
| postgresql.primary.persistence.enabled | bool | `false` |  |
| postgresql.readReplicas.persistence.enabled | bool | `false` |  |
| redis.architecture | string | `"replication"` |  |
| redis.auth.enabled | bool | `true` |  |
| redis.auth.existingSecret | string | `"auth"` |  |
| redis.auth.existingSecretPasswordKey | string | `"redisPassword"` |  |
| redis.auth.sentinel | bool | `false` |  |
| redis.fullnameOverride | string | `"redis"` |  |
| redis.image.registry | string | `"docker.io"` |  |
| redis.image.repository | string | `"bitnami/redis"` |  |
| redis.image.tag | string | `"7.0"` |  |
| redis.master.persistence.enabled | bool | `false` |  |
| redis.master.persistence.size | string | `"8Gi"` |  |
| redis.master.service.ports.redis | int | `6379` |  |
| redis.replica.persistence.enabled | bool | `false` |  |
| redis.replica.persistence.size | string | `"8Gi"` |  |
| redis.replicaCount | int | `5` |  |
| secret.generate | bool | `true` |  |
| secret.name | string | `"auth"` |  |
| secret.randLength | int | `30` |  |
| smtp.enabled | bool | `false` |  |
| smtp.from | string | `"noreply@example.org"` |  |
| smtp.host | string | `"smtp.gmail.com"` |  |
| smtp.port | int | `587` |  |
| smtp.username | string | `"somebody@example.org"` |  |
