/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// WARNING: This code is generated by go-bindata.
// Anytime a translated string changes and this file
// is regenerated, the ASF header must be added
// manually and all TAB characters MUST be removed.
//
// sources:
// wski18n/resources/de_DE.all.json
// wski18n/resources/en_US.all.json
// wski18n/resources/es_ES.all.json
// wski18n/resources/fr_FR.all.json
// wski18n/resources/it_IT.all.json
// wski18n/resources/ja_JA.all.json
// wski18n/resources/ko_KR.all.json
// wski18n/resources/pt_BR.all.json
// wski18n/resources/zh_Hans.all.json
// wski18n/resources/zh_Hant.all.json

package wski18n

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
    gz, err := gzip.NewReader(bytes.NewBuffer(data))
    if err != nil {
        return nil, fmt.Errorf("Read %q: %v", name, err)
    }

    var buf bytes.Buffer
    _, err = io.Copy(&buf, gz)
    clErr := gz.Close()

    if err != nil {
        return nil, fmt.Errorf("Read %q: %v", name, err)
    }
    if clErr != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

type asset struct {
    bytes []byte
    info  os.FileInfo
}

type bindataFileInfo struct {
    name    string
    size    int64
    mode    os.FileMode
    modTime time.Time
}

func (fi bindataFileInfo) Name() string {
    return fi.name
}
func (fi bindataFileInfo) Size() int64 {
    return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
    return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
    return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
    return false
}
func (fi bindataFileInfo) Sys() interface{} {
    return nil
}

var _wski18nResourcesDe_deAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesDe_deAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesDe_deAllJson,
        "wski18n/resources/de_DE.all.json",
    )
}

func wski18nResourcesDe_deAllJson() (*asset, error) {
    bytes, err := wski18nResourcesDe_deAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/de_DE.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x97\xc1\x6e\xdb\x3c\x0c\xc7\xef\x79\x0a\xc2\x97\xf4\x03\x5a\x3f\x40\xbf\xd3\xb0\x05\xdb\x80\x62\x2b\xb6\x74\x97\x75\x07\xd5\x62\x1a\x61\x8e\xe8\x52\x72\xba\x36\xd0\xbb\x0f\x92\xe3\xae\x43\xed\x58\x56\x3c\xa0\xa7\x16\x8a\xf8\xe7\x4f\x24\x25\xd2\xdf\x67\x00\xbb\x19\x00\x40\xa6\x64\x76\x0e\xd9\x95\x16\x37\x25\x82\x25\x10\x52\x02\x53\x6d\x11\xa8\xb2\x8a\xb4\x81\xf9\x6e\x97\xef\xff\x77\x6e\x9e\x9d\x36\x76\x96\x85\x36\xa5\xf0\xcb\x03\x02\xe7\xf0\x5c\x20\x9b\x01\xb8\xd3\x7e\xff\x05\xa3\xb0\x08\x1f\x96\xcb\x4b\x60\xbc\xab\xd1\x58\x58\x11\xc3\xe5\xd5\x32\x90\x04\x69\xe7\xe6\x41\x15\x99\x9d\x1b\x24\x4a\x90\x4c\x84\x7c\xbf\x98\x1c\xf2\x80\x64\x22\xe4\xbb\xc5\xc5\x62\xb9\x98\x9a\xf3\xb0\x6a\x6a\xd2\x3f\x7f\x9d\x3e\xeb\x07\x34\x07\x30\x45\x55\xa1\x96\x3d\x17\xc3\x6f\xb8\xfa\x72\xb1\xaf\xfd\x44\xe8\xe3\x3d\xc4\x45\xba\x0d\x88\x97\xf3\x42\x35\x97\x49\xd1\x1d\xd4\xe9\xc4\xf9\xa8\xb7\xa2\x54\x32\x95\x22\xda\xbc\xd3\xf9\x82\x99\x18\x50\x17\x24\x95\xbe\x7d\x12\xb9\x21\xf9\x30\xe8\x39\xce\xf6\x80\x5b\xa5\x95\x55\xa2\x54\x8f\xcf\xcc\x23\xbd\x0e\x98\x0e\x95\xae\x94\x60\xd7\xfb\xeb\x20\x6a\xbb\x46\x6d\x55\x11\x5c\xc0\x1a\x85\x44\x1e\x53\xa5\x63\xc4\x3a\xc1\xde\xd4\x76\x4d\xac\x1e\x1b\x9b\x9f\xf8\x00\xca\x80\x26\x0b\x05\xe9\x95\xba\xad\x19\x25\x9c\x9c\x9d\x79\x6d\xff\x8b\x3f\xae\x62\x94\xff\xf5\xa0\x25\xcb\x75\xc2\xbd\xa5\xcd\x46\x68\x09\x2b\xa1\x4a\x94\x20\xeb\xe6\xd0\xba\x39\xaf\x5f\xad\x19\x7b\x48\xe2\x6c\x3b\xdd\x7e\x22\x50\xda\x22\xaf\x44\x81\x50\x31\x6d\x95\x44\xf9\x3f\x68\x6a\xdf\x30\x53\x91\x36\x18\xaa\x0d\xf0\x57\x85\x85\x45\xd9\x83\x91\xa6\x35\x2e\x1a\xc1\x81\x16\x65\x6a\x44\x5e\xd8\x77\xba\x5f\xae\x11\x56\x54\x96\x74\xef\xeb\x5e\x54\x55\xd9\x56\x1a\x86\x6b\x71\x2f\x7c\x3e\x0b\x54\x5b\x94\x83\x25\x9c\x28\xf6\xfa\x1e\xb1\x57\x7b\xb1\xff\x68\xf9\x03\x55\x82\x4d\xd3\xa9\xb6\xc8\x46\x91\x1e\xd7\x64\x22\x24\x8e\x1c\xd4\x52\xfb\x5e\xbc\x60\x3a\xe0\x34\x54\xf1\x28\x37\xb5\x2a\xff\xaa\xc7\x11\x00\x87\x6c\xe3\x22\xe0\x03\xf8\x82\xff\xb8\xd1\x6f\x8c\x64\x1c\xa4\xff\x62\x98\x18\x72\x8c\x64\x1c\xe4\x7e\x12\x9f\x98\x73\xa4\x6a\x64\x3c\xfd\x2c\x3e\x75\x40\xc7\x68\xf6\xbc\xec\xfb\xc6\x14\x9a\x42\x0e\xed\x53\x6d\xac\x78\xfa\xaa\x0d\x82\x61\xc1\xb9\x79\x0e\xdf\xc2\x86\x76\x6c\x17\x8c\x70\x9d\x89\xc2\xaa\x2d\x5e\x67\xe0\xbb\xdf\x75\xa6\x74\xbb\x90\xf7\xb6\x84\x7f\xed\x77\x20\x2b\xcd\x5b\xdb\x76\xa2\x84\x14\x0c\x0a\x0c\x01\x30\x15\x68\x4c\x50\xb8\xab\x91\x1f\x7a\xbe\x84\xc6\x20\x8d\x97\xec\x84\xdc\xed\xf2\x8d\xb9\x75\x0e\x4e\x0a\x92\xe8\x37\xfb\xbf\xce\xf5\x8d\xa4\xfd\xfb\x7b\xc7\x9c\x82\xb4\xc6\x22\x64\xb9\x19\x96\x4e\x81\x18\xac\xda\xa0\x04\xaa\x6d\x0e\x27\xa1\xac\x7d\xf6\x6b\x03\x71\x18\xc7\xeb\x7a\xdc\xd9\x8f\xd9\xef\x00\x00\x00\xff\xff\xe7\x84\x36\x0c\x33\x12\x00\x00")

func wski18nResourcesEn_usAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesEn_usAllJson,
        "wski18n/resources/en_US.all.json",
    )
}

func wski18nResourcesEn_usAllJson() (*asset, error) {
    bytes, err := wski18nResourcesEn_usAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/en_US.all.json", size: 4659, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesEs_esAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesEs_esAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesEs_esAllJson,
        "wski18n/resources/es_ES.all.json",
    )
}

func wski18nResourcesEs_esAllJson() (*asset, error) {
    bytes, err := wski18nResourcesEs_esAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/es_ES.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesFr_frAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesFr_frAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesFr_frAllJson,
        "wski18n/resources/fr_FR.all.json",
    )
}

func wski18nResourcesFr_frAllJson() (*asset, error) {
    bytes, err := wski18nResourcesFr_frAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/fr_FR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesIt_itAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesIt_itAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesIt_itAllJson,
        "wski18n/resources/it_IT.all.json",
    )
}

func wski18nResourcesIt_itAllJson() (*asset, error) {
    bytes, err := wski18nResourcesIt_itAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/it_IT.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesJa_jaAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesJa_jaAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesJa_jaAllJson,
        "wski18n/resources/ja_JA.all.json",
    )
}

func wski18nResourcesJa_jaAllJson() (*asset, error) {
    bytes, err := wski18nResourcesJa_jaAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/ja_JA.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesKo_krAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesKo_krAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesKo_krAllJson,
        "wski18n/resources/ko_KR.all.json",
    )
}

func wski18nResourcesKo_krAllJson() (*asset, error) {
    bytes, err := wski18nResourcesKo_krAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/ko_KR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesPt_brAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesPt_brAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesPt_brAllJson,
        "wski18n/resources/pt_BR.all.json",
    )
}

func wski18nResourcesPt_brAllJson() (*asset, error) {
    bytes, err := wski18nResourcesPt_brAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/pt_BR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesZh_hansAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesZh_hansAllJson,
        "wski18n/resources/zh_Hans.all.json",
    )
}

func wski18nResourcesZh_hansAllJson() (*asset, error) {
    bytes, err := wski18nResourcesZh_hansAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/zh_Hans.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

var _wski18nResourcesZh_hantAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
    return bindataRead(
        _wski18nResourcesZh_hantAllJson,
        "wski18n/resources/zh_Hant.all.json",
    )
}

func wski18nResourcesZh_hantAllJson() (*asset, error) {
    bytes, err := wski18nResourcesZh_hantAllJsonBytes()
    if err != nil {
        return nil, err
    }

    info := bindataFileInfo{name: "wski18n/resources/zh_Hant.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495933008, 0)}
    a := &asset{bytes: bytes, info: info}
    return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
    cannonicalName := strings.Replace(name, "\\", "/", -1)
    if f, ok := _bindata[cannonicalName]; ok {
        a, err := f()
        if err != nil {
            return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
        }
        return a.bytes, nil
    }
    return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
    a, err := Asset(name)
    if err != nil {
        panic("asset: Asset(" + name + "): " + err.Error())
    }

    return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
    cannonicalName := strings.Replace(name, "\\", "/", -1)
    if f, ok := _bindata[cannonicalName]; ok {
        a, err := f()
        if err != nil {
            return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
        }
        return a.info, nil
    }
    return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
    names := make([]string, 0, len(_bindata))
    for name := range _bindata {
        names = append(names, name)
    }
    return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
    "wski18n/resources/de_DE.all.json": wski18nResourcesDe_deAllJson,
    "wski18n/resources/en_US.all.json": wski18nResourcesEn_usAllJson,
    "wski18n/resources/es_ES.all.json": wski18nResourcesEs_esAllJson,
    "wski18n/resources/fr_FR.all.json": wski18nResourcesFr_frAllJson,
    "wski18n/resources/it_IT.all.json": wski18nResourcesIt_itAllJson,
    "wski18n/resources/ja_JA.all.json": wski18nResourcesJa_jaAllJson,
    "wski18n/resources/ko_KR.all.json": wski18nResourcesKo_krAllJson,
    "wski18n/resources/pt_BR.all.json": wski18nResourcesPt_brAllJson,
    "wski18n/resources/zh_Hans.all.json": wski18nResourcesZh_hansAllJson,
    "wski18n/resources/zh_Hant.all.json": wski18nResourcesZh_hantAllJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
    node := _bintree
    if len(name) != 0 {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        pathList := strings.Split(cannonicalName, "/")
        for _, p := range pathList {
            node = node.Children[p]
            if node == nil {
                return nil, fmt.Errorf("Asset %s not found", name)
            }
        }
    }
    if node.Func != nil {
        return nil, fmt.Errorf("Asset %s not found", name)
    }
    rv := make([]string, 0, len(node.Children))
    for childName := range node.Children {
        rv = append(rv, childName)
    }
    return rv, nil
}

type bintree struct {
    Func     func() (*asset, error)
    Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
    "wski18n": &bintree{nil, map[string]*bintree{
        "resources": &bintree{nil, map[string]*bintree{
            "de_DE.all.json": &bintree{wski18nResourcesDe_deAllJson, map[string]*bintree{}},
            "en_US.all.json": &bintree{wski18nResourcesEn_usAllJson, map[string]*bintree{}},
            "es_ES.all.json": &bintree{wski18nResourcesEs_esAllJson, map[string]*bintree{}},
            "fr_FR.all.json": &bintree{wski18nResourcesFr_frAllJson, map[string]*bintree{}},
            "it_IT.all.json": &bintree{wski18nResourcesIt_itAllJson, map[string]*bintree{}},
            "ja_JA.all.json": &bintree{wski18nResourcesJa_jaAllJson, map[string]*bintree{}},
            "ko_KR.all.json": &bintree{wski18nResourcesKo_krAllJson, map[string]*bintree{}},
            "pt_BR.all.json": &bintree{wski18nResourcesPt_brAllJson, map[string]*bintree{}},
            "zh_Hans.all.json": &bintree{wski18nResourcesZh_hansAllJson, map[string]*bintree{}},
            "zh_Hant.all.json": &bintree{wski18nResourcesZh_hantAllJson, map[string]*bintree{}},
        }},
    }},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
    data, err := Asset(name)
    if err != nil {
        return err
    }
    info, err := AssetInfo(name)
    if err != nil {
        return err
    }
    err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
    if err != nil {
        return err
    }
    err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
    if err != nil {
        return err
    }
    return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
    children, err := AssetDir(name)
    // File
    if err != nil {
        return RestoreAsset(dir, name)
    }
    // Dir
    for _, child := range children {
        err = RestoreAssets(dir, filepath.Join(name, child))
        if err != nil {
            return err
        }
    }
    return nil
}

func _filePath(dir, name string) string {
    cannonicalName := strings.Replace(name, "\\", "/", -1)
    return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

