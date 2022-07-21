package packedcv

const skills = `
{{ define "Skills" }}

%==== Skills ====%

\header{Skills}
\begin{tabular}{ l p{.8\textwidth} }
{{ range .JSON.Skills -}}
	{{ printf "%s:& %s\\\\\n" (escape .Name) (escape (join .Keywords ", ")) }}
{{- end -}}
\end{tabular}
\vspace{2mm}
{{- end -}}
`
