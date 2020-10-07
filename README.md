# 1120x

The project is irs sub project to support IRS Modernized e-File Forms 1120, 1120S, and 1120-F for Tax Year 2020 and beyond.

## Terminology

- Manifest

    ManifestXML is a manifest.xml file, which provides information about the Submission.

- Document

    Document is the submission.xml file with the submission data in XML format.<br/>
    example : https://s3.amazonaws.com/irs-form-990/201541349349307794_public.xml

- Submission

    An IRS submission consists of XML data and optional binary attachments (PDF Files), which are packaged and compressed into a zip file.
    In this project submission zip archive is called as submission

## Features

### File

Main focus is to convert JSON file and XML file (include manifest.xml and submission.xml)
So we should create structures for json and xml and implement validation logic, etc.

### PDF

The IRS publishes stylesheets that can be used to transform an XML document into HTML. Specifically, these XSLT (eXtensible Stylesheet Language Transformation) files are distributed each year by the IRS so that tax preparers can generate tools that submit tax filings in the proper format.

To generate pdf from stylesheets and document (submission.xml)

1. Generate formatted xml with stylesheets and xml<br/>
  (base library : https://github.com/jbowtie/ratago)<br/>
  (example : https://github.com/betson/irs-efile-viewer)
2. XML -> HTML -> PDF


After implement main features, we can merge https://github.com/moov-io/irs and https://github.com/moov-io/1120x.

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://docs.moov.io/) | Our project documentation available online.
 Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
 [GitHub Issue](https://github.com/moov-io) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
 [moov-io slack](https://slack.moov.io/) | Join our slack channel (`#irs`) to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! Checkout our [issues for first time contributors](https://github.com/moov-io/1120x/contribute) for something to help out with.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/1120x/releases/latest) as well. We highly recommend you use a tagged release for production.

### Test Coverage

Improving test coverage is a good candidate for new contributors while also allowing the project to move more quickly by reducing regressions issues that might not be caught before a release is pushed out to our users. One great way to improve coverage is by adding edge cases and different inputs to functions (or [contributing and running fuzzers](https://github.com/dvyukov/go-fuzz)).

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
