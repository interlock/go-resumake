## go-resumake

This is an alternate implementation of the [resumake.io](https://github.com/saadq/resumake.io) engine. The `awesome-cv` template is a nearly identical direct port of their base template.

## prereqs

You will need the appropriate tools to convert LaTeX to your desired format. `texlive` is what is tested and is generally avaliable
on most linux distributions. `pandoc` can target multiple output formats, but is untested.

## templates

There are currently two very similar templates:

- awesome-cv
- packed-cv

Templates are just go templates that render the parts of the resume in LaTeX. PRs will be accepted for new templates that can render correctly with
`texlive` to pdf.

## workflow

1. provide a valid `resume.json` file: https://jsonresume.org/schema/
2. compile your file to LaTex: `goresumake -input resume.json -template packed-cv -latex resume.tex`
3. compile your LaTex file to your output format of preference:
  a. PDF -> `pdflatex resume.tex`
  b. Docx -> `pandoc resume.tex resume.docx`
  c: htmk -> `pandoc resume.tex resume.html`