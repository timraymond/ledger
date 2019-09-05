Ledger
======

[![Build Status](https://dev.azure.com/tim0375/tim/_apis/build/status/timraymond.ledger?branchName=master)](https://dev.azure.com/tim0375/tim/_build/latest?definitionId=2&branchName=master)

This is an example of a package for parsing the Ledger plain text accounting
format (https://www.ledger-cli.org). The parser itself is found in the `parse`
package and it uses the [Pigeon](https://github.com/mna/pigeon) [Parsing Expression
Grammar (PEG)](https://en.wikipedia.org/wiki/Parsing_expression_grammar) parser
generator. That should be installed using `go get github.com/mna/pigeon` prior
to running `go generate`. This was written to prepare my lightning talk given
at GopherCon 2019 in San Diego.

**Note**: This is really just a demonstration. I expect there to be bugs, and
I'm not actually using this in anything at the moment. YMMV

Usage
-----

Package `parse` exposes three important functions for parsing ledger-formatted
text:

* [Parse](https://godoc.org/github.com/timraymond/ledger/parse#Parse) - parses the provided `[]byte`
* [ParseReader](https://godoc.org/github.com/timraymond/ledger/parse#ParseReader) - parses the `io.Reader` provided
* [ParseFile](https://godoc.org/github.com/timraymond/ledger/parse#ParseFile) - opens and parses the content of the filename provided
