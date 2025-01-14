package send

import (
	"bytes"
	"html/template"
	"os"
	"strings"
	testTemplate "text/template"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/utility"
	"github.com/pkg/errors"
)

var (
	funcMap template.FuncMap = template.FuncMap{
		"numberFormat": utility.NumberFormat,
		"add":          utility.Add,
	}
)

func ParseTemplate(extReq request.ExternalRequest, templateFileName, baseTemplateFileName string, templateData map[string]interface{}) (string, error) {
	var (
		outputBuffer bytes.Buffer
		t            *template.Template
	)
	templateData = AddMoreMailTemplateData(extReq, templateData)

	fileName, err := utility.FindTemplateFilePath(templateFileName, "/email")
	if err != nil {
		return "", err
	}

	if baseTemplateFileName != "" {
		baseFileName, err := utility.FindTemplateFilePath(baseTemplateFileName, "/email")
		if err != nil {
			return "", err
		}

		base, err := os.ReadFile(baseFileName)
		if err != nil {
			return "", err
		}
		t = template.New("base").Funcs(funcMap)
		t, err = t.Parse(string(base))
		if err != nil {
			return "", err
		}
		t, err = t.ParseFiles(fileName)
		if err != nil {
			return "", err
		}

	} else {
		filedata, err := os.ReadFile(fileName)
		if err != nil {
			return "", errors.Wrap(err, "template not found")
		}

		t, err = template.New("email_template").Funcs(funcMap).Parse(string(filedata))
		if err != nil {
			return "", err
		}
	}

	if err2 := t.Execute(&outputBuffer, templateData); err2 != nil {
		return "", err2
	}

	body := outputBuffer.String()

	return body, nil
}

func AddMoreMailTemplateData(extReq request.ExternalRequest, data map[string]interface{}) map[string]interface{} {

	return data
}

func ParseSMSTemplate(extReq request.ExternalRequest, templateFileName string, templateData map[string]interface{}) (string, error) {
	templateData = AddMoreMailTemplateData(extReq, templateData)
	fileName, err := utility.FindTemplateFilePath(templateFileName, "/sms")
	if err != nil {
		return "", err
	}

	filedata, er := os.ReadFile(fileName)
	if er != nil {
		return "", errors.Wrap(er, "template not found")
	}

	compl := string(filedata)

	t, err1 := testTemplate.New("sms_template").Funcs(funcMap).Parse(compl)
	if err1 != nil {
		return "", err1
	}

	buf := new(bytes.Buffer)
	if err2 := t.Execute(buf, templateData); err2 != nil {
		return "", err2
	}

	body := buf.String()

	return body, nil
}

func GeneratePDFFromTemplate(extReq request.ExternalRequest, templatePath, baseTemplatepath string, data map[string]interface{}) ([]byte, error) {
	var (
		tpl *template.Template
	)

	data = AddMoreMailTemplateData(extReq, data)
	templatePath, err := utility.FindTemplateFilePath(templatePath, "/email")
	if err != nil {
		return nil, err
	}

	if extReq.Test {
		return []byte("testing"), nil
	}

	if baseTemplatepath != "" {
		baseFileName, err := utility.FindTemplateFilePath(baseTemplatepath, "/email")
		if err != nil {
			return nil, err
		}

		base, err := os.ReadFile(baseFileName)
		if err != nil {
			return nil, err
		}
		tpl = template.New("base").Funcs(funcMap)
		tpl, err = tpl.Parse(string(base))
		if err != nil {
			return nil, err
		}
		tpl, err = tpl.ParseFiles(templatePath)
		if err != nil {
			return nil, err
		}

	} else {
		filedata, err := os.ReadFile(templatePath)
		if err != nil {
			return nil, errors.Wrap(err, "template not found")
		}

		tpl, err = template.New("pdf_template").Funcs(funcMap).Parse(string(filedata))
		if err != nil {
			return nil, err
		}
	}

	// tpl, err := template.ParseFiles(templatePath)
	// if err != nil {
	// 	return nil, err
	// }

	var renderedTemplate bytes.Buffer
	if err := tpl.Execute(&renderedTemplate, data); err != nil {
		return nil, err
	}

	html := renderedTemplate.String()

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	if err := pdfg.Create(); err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
