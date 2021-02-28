# Slides

As a general rule of thumb, if you have slides within a go module, such as the case with this project, you may want to exclude the slide `.go` snippet files from the module's dependencies. To do this you can put your source code in a directory with the underscore prefix i.e `_slidesrc`. Go modules excludes these directories.