# egr

    go get -u github.com/codemodus/egr

egr is a CLI application which wraps commands for the purpose of pre-expanding 
globbing. This is particularly useful for go:generate directives due to the lack
of glob expansion. Command exit codes and output are passed through
transparently.

Simple usage:

    //go:generate egr protoc -I../ --go_out=. ../*.proto

Advanced usage:

    //go:generate -command pbgen egr protoc -I../ --go_out=. --longplugin_out=.
    //go:generate pbgen ../*.proto

