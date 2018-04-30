package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "createjoinevent.gtpl",
		FileModTime: time.Unix(1525031970, 0),
		Content:     string("<html>\n    <head>\n    <title></title>\n    </head>\n    <body>\n        <form action=\"/createevent\" method=\"POST\">\n            Event Name:<input type=\"text\" name=\"Name\">\n          \n            <input type=\"submit\" value=\"Create\">\n        </form>\n    </body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1525031970, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "createjoinevent.gtpl"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`static`, &embedded.EmbeddedBox{
		Name: `static`,
		Time: time.Unix(1525031970, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"createjoinevent.gtpl": file2,
		},
	})
}
