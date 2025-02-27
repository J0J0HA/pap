package plugins

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/talwat/pap/internal/fs"
	"github.com/talwat/pap/internal/log"
	"github.com/talwat/pap/internal/net"
	"github.com/talwat/pap/internal/plugins/jenkins"
	"github.com/talwat/pap/internal/plugins/sources/paplug"
)

func PluginDownload(plugin paplug.PluginInfo) {
	for _, download := range plugin.Downloads {
		var url string

		log.Log("getting download url...")

		if download.Type == "url" {
			url = download.URL
		} else if download.Type == "jenkins" {
			url = jenkins.GetJenkinsURL(download)
		}

		url = SubstituteProps(plugin, url)
		path := filepath.Join("plugins", download.Filename)

		net.Download(
			url,
			fmt.Sprintf("%s not found, please report this to https://github.com/talwat/pap/issues", url),
			path,
			download.Filename,
			nil,
			fs.ReadWritePerm,
		)

		if strings.HasSuffix(path, ".zip") {
			log.Log("unzipping %s...", path)
			fs.Unzip(path, "plugins/")

			log.Log("cleaning up...")
			fs.DeletePath(path)
		}
	}
}
