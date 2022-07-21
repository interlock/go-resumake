package packedcv

const education = `
{{ define "Education" }}
{{ if .JSON.Education }}
%==== Education ====%

\header{Education}
{{ range .JSON.Education }}
  {{- $line1 := "" -}}
  {{- $line2 := "" -}}
  {{- if .Institution -}}
	{{- $line1 = printf "%s\\textbf{%s}" $line1 (escape .Institution) }}
  {{- end -}}
  {{- if .Area -}}
  	{{- $line1 = printf "%s\\hfill %s" $line1 (escape .Area) -}}
  {{- end -}}
  {{- if .StudyType -}}
	{{- $line2 = printf "%s" (escape .StudyType) -}}
  {{- end -}}
  {{- if .Area -}}
	{{- if .StudyType -}}
	  {{- $line2 = printf "%s $\\cdot$ %s" $line2 (escape .Area) -}}
	{{- else -}}
	  {{- $line2 = printf "Degree in %s" (escape .Area) -}}
	{{- end -}}
  {{- end -}}
  {{- if .Gpa -}}
	{{- $line2 = printf "%s \\textit{GPA: %s}" $line2 .Gpa -}}
  {{- end -}}
  {{- if or .StartDate .EndDate -}}
	{{- $dates := newSlice }}
	{{- if .StartDate }}{{- $dates = appendSlice $dates (dateFormat .StartDate "2006") -}}{{- end -}}
	{{- if .EndDate }}{{- $dates = appendSlice $dates (dateFormat .EndDate "2006") -}}{{- end -}}
	{{- $line2 = printf "%s \\hfill %s" $line2 (join $dates " $\\cdot$ ") -}}
  {{- end -}}
  {{- if $line1 -}}
	  {{- $line1 = printf "%s\\\\" $line1 -}}
  {{- end -}}
  {{- if $line2 -}}
	  {{- $line2 = printf "%s\\\\" $line2 -}}
  {{- end -}}
{{ $line1 }}
{{ $line2 }}
\vspace{2mm}
{{- end -}}
{{- end -}}
{{- end -}}
`
