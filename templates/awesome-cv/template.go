package awesomecv

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"text/template"

	"github.com/interlock/go-resumake/jsonresume"
)

const document = `
\documentclass[a4paper]{article}
\usepackage{fullpage}
\usepackage{amsmath}
\usepackage{amssymb}
\usepackage{textcomp}
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\textheight=10in
\pagestyle{empty}

{{template "resumeDefinitions"}}

\begin{document}
\vspace*{-40pt}
{{ $data := . }}
{{ range .Sections }}
	{{- if eq . "Basics" -}}
	{{- template "Basics" $data -}}
	{{- end -}}

	{{- if eq . "Education" -}}
	{{- template "Education" $data -}}
	{{- end -}}

	{{- if eq . "Work" }}
	{{- template "Work" $data -}}
	{{ end -}}

	{{- if eq . "Skills" }}
	{{- template "Skills" $data -}}
	{{ end -}}
{{ end }}

%% WHITESPACE

\end{document}
`

const resumeDefinitions = `
{{ define "resumeDefinitions" }}
%\renewcommand{\encodingdefault}{cg}
%\renewcommand{\rmdefault}{lgrcmr}

\def\bull{\vrule height 0.8ex width .7ex depth -.1ex }

% DEFINITIONS FOR RESUME %%%%%%%%%%%%%%%%%%%%%%%

\newcommand{\area} [2] {
	\vspace*{-9pt}
	\begin{verse}
		\textbf{#1}   #2
	\end{verse}
}

\newcommand{\lineunder} {
	\vspace*{-8pt} \\\\
	\hspace*{-18pt} \hrulefill \\\\
}

\newcommand{\header} [1] {
	{\hspace*{-18pt}\vspace*{6pt} \textsc{#1}}
	\vspace*{-6pt} \lineunder
}

\newcommand{\employer} [3] {
	{ \textbf{#1} (#2)\\\\ \underline{\textbf{\emph{#3}}}\\\\  }
}

\newcommand{\contact} [3] {
	\vspace*{-10pt}
	\begin{center}
		{\Huge \scshape {#1}}\\\\
		#2 \\\\ #3
	\end{center}
	\vspace*{-8pt}
}

\newenvironment{achievements}{
	\begin{list}
		{$\bullet$}{\topsep 0pt \itemsep -2pt}}{\vspace*{4pt}
	\end{list}
}

\newcommand{\schoolwithcourses} [4] {
	\textbf{#1} #2 $\bullet$ #3\\\\
	#4 \\\\
	\vspace*{5pt}
}

\newcommand{\school} [4] {
	\textbf{#1} #2 $\bullet$ #3\\\\
	#4 \\\\
}
% END RESUME DEFINITIONS %%%%%%%%%%%%%%%%%%%%%%%
{{- end -}}
`

const basics = `
{{ define "Basics" }}
{{- if .JSON.Basics -}}
{{- with .JSON.Basics -}}

	{{- $address := "" -}}
	{{- if .Location -}}
	{{- $address = escape .Location.Address -}}
	{{- end -}}
	  
	{{- $line1 := "" -}} 
  {{- if .Name -}}
	{{- $line1 = printf "{\\Huge \\scshape {%s}}\\\\" (escape .Name) -}}
	{{- end -}}
	  
	{{- $line2 := newSlice -}}
	{{- if $address -}}
	{{- $line2 = appendSlice $line2 $address -}}
	{{- end -}}

	{{- if .Email -}}
	{{- $line2 = appendSlice $line2 (escape .Email) -}}
	{{- end -}}

	{{- if .Phone -}}
	{{- $line2 = appendSlice $line2 (escape .Phone) -}}
	{{- end -}}

	{{- if .Website -}}
	{{- $line2 = appendSlice $line2 (escape .Website) -}}
	{{- end -}}

	{{- $line2 = join $line2 " $\\cdot$ " -}}
	

%==== Profile ====%

\vspace*{-10pt}
\begin{center}
	{{ $line1 }}
	{{ $line2 }}
\end{center}
{{- end -}}
{{- else -}}
%% Basic skip
{{- end -}}
{{ end }}
`

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
	  {{- $line2 = printf "%s %s" $line2 (escape .Area) -}}
	{{- else -}}
	  {{- $line2 = printf "Degree in %s" (escape .Area) -}}
	{{- end -}}
  {{- end -}}
  {{- if .Gpa -}}
	{{- $line2 = printf "%s \\textit{GPA: %s}" $line2 .Gpa -}}
  {{- end -}}
  {{- if or .StartDate .EndDate -}}
	{{- $dates := newSlice }}
	{{- if .StartDate }}{{- $dates = appendSlice $dates .StartDate -}}{{- end -}}
	{{- if .EndDate }}{{- $dates = appendSlice $dates .EndDate -}}{{- end -}}
	{{- $line2 = printf "%s \\hfill %s" $line2 (join $dates ".") -}}
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

const work = `
{{ define "Work" }}
{{ if .JSON.Work }}
%==== Experience ====%

\header{Experience}
\vspace{1mm}
{{- range .JSON.Work -}}
	{{- $line1 := "" -}}
	{{- $line2 := "" -}}
	{{- $highlightLines := newSlice -}}
	
	{{- if .Company -}}
		{{- $line1 = printf "%s\\textbf{%s}" $line1 (escape .Company) -}}
  {{- end -}}

	{{- if .Position -}}
		{{- $line2 = printf "%s\\textit{%s}" $line2 (escape .Position) -}}
	{{- end -}}

	{{- if and .StartDate .EndDate -}}
		{{- $line2 = printf "%s \\hfill %s - %s" $line2 .StartDate .EndDate -}}
	{{- else -}}
		{{- if .StartDate -}}
			{{- $line2 = printf "%s \\hfill %s - Present" $line2 .StartDate -}}
		{{- end -}}
		{{- if .EndDate -}}
			{{- $line2 = printf "%s \\hfill %s" $line2 .EndDate -}}
		{{- end -}}
	{{- end -}}

	{{- if $line1 -}}{{- $line1 = printf "%s\\\\" $line1 -}}{{- end -}}
	{{- if $line2 -}}{{- $line2 = printf "%s\\\\" $line2 -}}{{- end -}}
	
	{{- if .Highlights -}}
		{{- $highlightLines = appendSlice $highlightLines "\\vspace{-1mm}" -}}
		{{- $highlightLines = appendSlice $highlightLines "\\begin{itemize} \\itemsep 1pt" -}}
		{{- range .Highlights -}}
			{{- $highlightLines = appendSlice $highlightLines (printf "\\item %s" (escape .)) -}}
		{{- end -}}
		{{- $highlightLines = appendSlice $highlightLines "\\end{itemize}" -}}
	{{- end -}}
	
	{{ $line1 }}
	{{ $line2 }}
	{{ join $highlightLines "\n" }}
{{- end -}}
{{- end -}}
{{- end -}}
`

const skills = `
{{ define "Skills" }}

%==== Skills ====%

\header{Skills}
\begin{tabular}{ l p{.8\textwidth} }
{{- range .JSON.Skills -}}
	{{- printf "%s: %s\\\\" (escape .Name) (escape (join .Keywords ", ")) -}}
{{- end -}}
\end{tabular}
\vspace{2mm}
{{- end -}}
`

func newSlice() []string {
	return make([]string, 0)
}

func appendSlice(in []string, v string) []string {
	return append(in, v)
}

// escapeLatex
// & % $ # _ { } ~ ^ \
func escapeLatex(s string) string {
	re := regexp.MustCompile(`[&$%#_\{\}~^\\]`)
	return re.ReplaceAllStringFunc(s, func(str string) string {
		if strings.Compare(str, "\\") == 0 {
			return "\\textbackslash{}"
		}
		if strings.Compare(str, "~") == 0 {
			// return "\\char`\\~"
			return "{\\raise.17ex\\hbox{$\\scriptstyle\\mathtt{\\sim}$}}"
		}
		if strings.Compare(str, "^") == 0 {
			return "\\char`\\^"
		}
		return fmt.Sprintf("\\%s", str)
	})
}

func Render(resume jsonresume.JSONResume, output io.Writer) error {
	funcs := template.FuncMap{
		"join":        strings.Join,
		"newSlice":    newSlice,
		"appendSlice": appendSlice,
		"escape":      escapeLatex,
	}

	documentTemplate := template.Must(template.New("document").Funcs(funcs).Parse(document))

	template.Must(documentTemplate.Parse(resumeDefinitions))
	template.Must(documentTemplate.Parse(basics))
	template.Must(documentTemplate.Parse(education))
	template.Must(documentTemplate.Parse(work))
	template.Must(documentTemplate.Parse(skills))

	sections := []string{"Basics", "Education", "Work", "Skills"}
	data := map[string]interface{}{
		"JSON":     resume,
		"Sections": sections,
	}
	err := documentTemplate.Execute(output, data)
	if err != nil {
		return err
	}

	return nil
}
