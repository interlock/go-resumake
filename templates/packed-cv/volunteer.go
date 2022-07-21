package packedcv

const volunteer = `
{{ define "Volunteer" }}

%==== volunteer ====%

\header{Volunteering}
\begin{tabular}{ l p{.8\textwidth} }
{{ range .JSON.Volunteer -}}
	{{- $dates := newSlice }}
	{{- if or .StartDate .EndDate -}}
		{{- if .StartDate }}{{- $dates = appendSlice $dates (dateFormat .StartDate "2006") -}}{{- end -}}
		{{- if .EndDate }}{{- $dates = appendSlice $dates (dateFormat .EndDate "2006") -}}{{- end -}}
	{{- end -}}
	{{ printf "%s $\\cdot$ %s $\\cdot$ %s\\\\\n" (escape .Organization) (escape .Position) (join $dates " $\\cdot$ ") }}
{{- end -}}
\end{tabular}
\vspace{2mm}
{{- end -}}
`
