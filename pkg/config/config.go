package config

import "text/template"

/*set configuration so that once have template set. no need to load again
untill application restarts
this package can be imported by other pkg in application but it doesnt import anything else
from application it will help to avoid import cycle*/

//holds the application config
//any kind of config can add to struct that need for application
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
