# 1120x

The project is irs sub project to support IRS Modernized e-File Forms 1120, 1120S, and 1120-F for Tax Year 2020 and beyond.

## Terminology

- Manifest

    ManifestXML is a manifest.xml file, which provides information about the Submission.

- Return

    Return is return schema for irs forms.
    example : Return990.xsd

- Submission

    An IRS submission consists of XML data and optional binary attachments (PDF Files), which are packaged and compressed into a zip file.
    In this project submission zip archive is called as submission

## Features

### Irs e-file logic

1120x project is to get transmission (submission) file structures for internet filing (IFA) and application-to-application (A2A).
Both channels use simple object access protocol (SOAP) with attachments.
We can get envelope and attachments of SOAP message from JSON or XML input file using 1120x project.

IrsTransmissionFile is interface for irs transmission file instance.
there are 2 main functions to get data.
- SOAPEnvelope
- SOAPAttachment

User can create SOAP message of return or state using 1120x package easily.


Main focus of this project is to convert form JSON file and XML file to irs e-file structure (raw data of SOAP envelope and attachments).

| Input      | Output     |
|------------|------------|
| JSON       | JSON       |
| XML        | XML        |
|            | Irs e-file |
|            | SQL        |

### Go Codes from XSD files
We should get golang codes from xsd files that provided irs.
There are many method to get go code from XSD, but perfect tool (application) for this don't exist. Because there are many problems about xml version, tags, etc.

To get go code from XSD files, we can use xsdgen (https://godoc.org/aqwari.net/xml/cmd/xsdgen) and following steps.  <br/> 
1. create a merged xsd file <br/>
xsdgen don't support include tag, so should create a merged xsd file to import other xsd file instead of include tag.  <br/>
2. remove repeated go structures <br/>
any go structures can be repeated in generated code, so the structures need to be removed manually.
3. change field value to pointer, add json tag <br/>
if field have omitempty tag, we need to change pointer the field, add json tag of omitempty. 
4. remove unnecessary xml namespace <br/>

### PDF

Other feature of the package is to create pdf file from XML and XSD files.

| Input      | Output     |
|------------|------------|
| XML        | PDF        |


### PDF Generator

The IRS publishes stylesheets that can be used to transform XML document into HTML. 
Specifically, these XSLT (eXtensible Stylesheet Language Transformation) files are distributed each year by the IRS so that tax preparers can generate tools that submit tax filings in the proper format.

To generate pdf from stylesheets and input xml (return.xml)

- XML -> HTML -> PDF<br/>
  we don't create pdf for any return from input xml and stylesheets.<br/>
  we should use html as temporary result.
   
XML -> HTML <br/>
There are 2 methods to get html in 1120x package.  <br/>

- To use github.com/jbowtie/ratago/xslt <br/>
the method used go package, but there is a problem that XSD parameters are invalid sometimes. <br/> 
- To use http://xmlsoft.org/XSLT/xsltproc.html  <br/>
the method should use application file (xsltproc), but above problem is fixed.

HTML -> PDF <br/>
Used go-wkhtmltopdf to convert pdf file. (based wkhtmltopdf)

After implement main features, we should merge https://github.com/moov-io/irs and https://github.com/moov-io/1120x.

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

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! Checkout our [issues for first time contributors](https://github.com/moov-io/1120x/issues) for something to help out with.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/1120x/releases/latest) as well. We highly recommend you use a tagged release for production.

### Test Coverage

Improving test coverage is a good candidate for new contributors while also allowing the project to move more quickly by reducing regressions issues that might not be caught before a release is pushed out to our users. One great way to improve coverage is by adding edge cases and different inputs to functions (or [contributing and running fuzzers](https://github.com/dvyukov/go-fuzz)).

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
