package packedcv

const interests = `
{{ define "Interests" }}

%==== interests ====%

\header{Interests}
\begin{tabular}{ l p{.8\textwidth} }
{{ range .JSON.Interests -}}
	{{ printf "%s:& %s\\\\\n" (escape .Name) (escape (join .Keywords ", ")) }}
{{- end -}}
\end{tabular}
\vspace{2mm}
{{- end -}}
`
