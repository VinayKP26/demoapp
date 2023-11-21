{{- define "chart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "chart.labels" -}}
app: {{ include "chart.name" . }}
version: {{ .Chart.AppVersion }}
instance: {{ .Release.Name }}
{{- end }}

{{- define "chart.selectorLabels" -}}
app: {{ include "chart.name" . }}
{{- end }}