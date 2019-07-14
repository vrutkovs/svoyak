package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsceff1b12f5be500ebdc38a841eff4d5a095217bb = "<html>\n<head>\n  <title>Svoyak</title>\n</head>\n<body>\n  <a href=\"{{.Url}}/join?id={{.ID}}\">Join game</a><br/>\n  QR code goes here\n</body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"start-page.tmpl"}}, map[string]*assets.File{
	"/templates/start-page.tmpl": &assets.File{
		Path:     "/templates/start-page.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1563110276, 1563110276389652413),
		Data:     []byte(_Assetsceff1b12f5be500ebdc38a841eff4d5a095217bb),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1563110241, 1563110241702201044),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1563110207, 1563110207313753565),
		Data:     nil,
	}}, "")
