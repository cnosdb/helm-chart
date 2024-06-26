{{/*
Expand the name of the chart.
*/}}
{{- define "cnosdb.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "cnosdb.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "cnosdb.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "cnosdb.labels" -}}
helm.sh/chart: {{ include "cnosdb.chart" . }}
{{ include "cnosdb.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "cnosdb.selectorLabels" -}}
app.kubernetes.io/name: {{ include "cnosdb.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "cnosdb.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "cnosdb.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
To Yaml
Usage:
{{ include "cnosdb.toYaml" ( dict "value" .Values.path.to.the.Value ) }}
{{ include "cnosdb.toYaml" ( dict "value" .Values.path.to.the.Value ) }}
*/}}
{{- define "cnosdb.toYaml" -}}
{{- $value := typeIs "string" .value | ternary .value (.value | toYaml) }}
{{- $value }}
{{- end -}}

{{/*
To Yaml
Usage:
{{ include "cnosdb.extraConfig" ( dict "value" .Values "conf" .Values.singleton.extraConf)}}
*/}}
{{- define "cnosdb.extraConfig" -}}
{{- if not (empty .value.license) }}
{{- set .conf "license_file" (printf "%s/license.json" (.value.licensePath | trimSuffix "/")) | toJson}}
{{- else -}}
{{- .conf | toJson }}
{{- end -}}
{{- end -}}