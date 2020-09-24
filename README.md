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

