package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenJSFile(data map[string]string) {
	const temp = `const url = "%s";

function Name() {
  return "%s";
}

function Search(wd, pg) {
  return comm.SearchVideo({ name: Name(), url, wd, pg });
}

function SearchByID(id) {
  return comm.SearchVideo({ name: Name(), url, id });
}

module.exports = {
  Name,
  Search,
  SearchByID,
};
	`
	pwd, _ := os.Executable()
	pwd = filepath.Dir(pwd)
	fpath := filepath.Join(pwd, "files")
	os.RemoveAll(fpath)
	os.MkdirAll(fpath, 0755)

	for k, v := range data {
		os.WriteFile(filepath.Join(fpath, v+".js"), []byte(fmt.Sprintf(temp, k, v)), 0644)
	}
}
