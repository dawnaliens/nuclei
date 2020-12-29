package templates

import (
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/dns"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/http"
	"github.com/projectdiscovery/nuclei/v2/pkg/workflows"
)

// Template is a request template parsed from a yaml file
type Template struct {
	// ID is the unique id for the template
	ID string `yaml:"id"`
	// Info contains information about the template
	Info map[string]string `yaml:"info"`
	// RequestsHTTP contains the http request to make in the template
	RequestsHTTP []*http.Request `yaml:"requests,omitempty"`
	// RequestsDNS contains the dns request to make in the template
	RequestsDNS []*dns.Request `yaml:"dns,omitempty"`

	// Workflows is a yaml based workflow declaration code.
	*workflows.Workflow

	path          string
	totalRequests int
	executer      protocols.Executer
}

// GetPath returns the path of the template.
func (t *Template) GetPath() string {
	return t.path
}

// Requests returns the number of requests for the template
func (t *Template) Requests() int {
	return t.totalRequests
}
