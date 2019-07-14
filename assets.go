package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsf5c0efbde7635361885b1d248656b3f8c8dbe7a4 = "<html>\n<head>\n  <title>Svoyak</title>\n</head>\n<body>\n  <a href=\"{{.Url}}/join?id={{.ID}}\">Join game</a><br/>\n  QR code goes here\n</body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"session.tmpl"}}, map[string]*assets.File{
	"/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1563112186, 1563112186723115818),
		Data:     nil,
	}, "/templates/session.tmpl": &assets.File{
		Path:     "/templates/session.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1563112186, 1563112186725000000),
		Data:     []byte(_Assetsf5c0efbde7635361885b1d248656b3f8c8dbe7a4),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1563110241, 1563110241702201044),
		Data:     nil,
	}}, "")
