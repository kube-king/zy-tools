/*
Copyright 2024 Faw Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	_ "github.com/gabriel-vasile/mimetype"
	"io"
	"strings"
)

var mimeTypeMap = map[string]string{
	"3g2.3g2":            "video/3gpp2",
	"3gp.3gp":            "video/3gpp",
	"3mf.3mf":            "application/vnd.ms-package.3dmanufacturing-3dmodel+xml",
	"7z.7z":              "application/x-7z-compressed",
	"a.a":                "application/x-archive",
	"aac.aac":            "audio/aac",
	"aaf.aaf":            "application/octet-stream",
	"accdb.accdb":        "application/x-msaccess",
	"aiff.aiff":          "audio/aiff",
	"amf.amf":            "application/x-amf",
	"amr.amr":            "audio/amr",
	"ape.ape":            "audio/ape",
	"apng.png":           "image/vnd.mozilla.apng",
	"asf.asf":            "video/x-ms-asf",
	"atom.atom":          "application/atom+xml",
	"au.au":              "audio/basic",
	"avi.avi":            "video/x-msvideo",
	"avif.avif":          "image/avif",
	"avifsequence.avif":  "image/avif",
	"bmp.bmp":            "image/bmp",
	"bpg.bpg":            "image/bpg",
	"bz2.bz2":            "application/x-bzip2",
	"cab.cab":            "application/vnd.ms-cab-compressed",
	"cab.is.cab":         "application/x-installshield",
	"class.class":        "application/x-java-applet",
	"crx.crx":            "application/x-chrome-extension",
	"csv.csv":            "text/csv",
	"cpio.cpio":          "application/x-cpio",
	"dae.dae":            "model/vnd.collada+xml",
	"dbf.dbf":            "application/x-dbf",
	"dcm.dcm":            "application/dicom",
	"deb.deb":            "application/vnd.debian.binary-package",
	"djvu.djvu":          "image/vnd.djvu",
	"doc.doc":            "application/msword",
	"docx.1.docx":        "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"docx.docx":          "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"drpm.rpm":           "application/x-rpm",
	"dwg.1.dwg":          "image/vnd.dwg",
	"dwg.dwg":            "image/vnd.dwg",
	"eot.eot":            "application/vnd.ms-fontobject",
	"epub.epub":          "application/epub+zip",
	"exe.exe":            "application/vnd.microsoft.portable-executable",
	"fdf.fdf":            "application/vnd.fdf",
	"fits.fits":          "application/fits",
	"flac.flac":          "audio/flac",
	"flv.flv":            "video/x-flv",
	"gbr.gbr":            "image/x-gimp-gbr",
	"geojson.1.geojson":  "application/geo+json",
	"geojson.geojson":    "application/geo+json",
	"gif.gif":            "image/gif",
	"glb.glb":            "model/gltf-binary",
	"gml.gml":            "application/gml+xml",
	"gpx.gpx":            "application/gpx+xml",
	"gz.gz":              "application/gzip",
	"har.har":            "application/json",
	"hdr.hdr":            "image/vnd.radiance",
	"heic.single.heic":   "image/heic",
	"heif.heif":          "image/heif",
	"html.html":          "text/html; charset=utf-8",
	"html.iso88591.html": "text/html; charset=iso-8859-1",
	"html.svg.html":      "text/html; charset=utf-8",
	"html.usascii.html":  "text/html; charset=us-ascii",
	"html.utf8.html":     "text/html; charset=utf-8",
	"html.withbr.html":   "text/html; charset=utf-8",
	"ico.ico":            "image/x-icon",
	"ics.dos.ics":        "text/calendar",
	"ics.ics":            "text/calendar",
	"iso88591.txt":       "text/plain; charset=iso-8859-1",
	"jar.jar":            "application/jar",
	"jp2.jp2":            "image/jp2",
	"jpf.jpf":            "image/jpx",
	"jpg.jpg":            "image/jpeg",
	"jpm.jpm":            "image/jpm",
	"jxl.jxl":            "image/jxl",
	"jxr.jxr":            "image/jxr",
	"xpm.xpm":            "image/x-xpixmap",
	"js.js":              "application/javascript",
	"json.json":          "application/json",
	"json.lowascii.json": "application/json",
	// json.{int,float,string}.txt contain a single JSON value. They are valid JSON
	// documents, but they should not be detected as application/json. This mimics
	// the behaviour of the file utility and seems the correct thing to do.
	"json.int.txt":       "text/plain; charset=utf-8",
	"json.float.txt":     "text/plain; charset=utf-8",
	"json.string.txt":    "text/plain; charset=utf-8",
	"kml.kml":            "application/vnd.google-earth.kml+xml",
	"lit.lit":            "application/x-ms-reader",
	"ln":                 "application/x-executable",
	"lua.lua":            "text/x-lua",
	"lz.lz":              "application/lzip",
	"m3u.m3u":            "application/vnd.apple.mpegurl",
	"m4a.m4a":            "audio/x-m4a",
	"audio.mp4":          "audio/mp4",
	"lnk.lnk":            "application/x-ms-shortcut",
	"macho.macho":        "application/x-mach-binary",
	"mdb.mdb":            "application/x-msaccess",
	"midi.midi":          "audio/midi",
	"mkv.mkv":            "video/x-matroska",
	"mobi.mobi":          "application/x-mobipocket-ebook",
	"mov.mov":            "video/quicktime",
	"mp3.mp3":            "audio/mpeg",
	"mp3.v1.notag.mp3":   "audio/mpeg",
	"mp3.v2.5.notag.mp3": "audio/mpeg",
	"mp3.v2.notag.mp3":   "audio/mpeg",
	"mp4.1.mp4":          "video/mp4",
	"mp4.mp4":            "video/mp4",
	"mpc.mpc":            "audio/musepack",
	"mpeg.mpeg":          "video/mpeg",
	"mqv.mqv":            "video/quicktime",
	"mrc.mrc":            "application/marc",
	"msi.msi":            "application/x-ms-installer",
	"msg.msg":            "application/vnd.ms-outlook",
	"ndjson.xl.ndjson":   "application/x-ndjson",
	"ndjson.ndjson":      "application/x-ndjson",
	"nes.nes":            "application/vnd.nintendo.snes.rom",
	"elfobject":          "application/x-object",
	"odf.odf":            "application/vnd.oasis.opendocument.formula",
	"sxc.sxc":            "application/vnd.sun.xml.calc",
	"odg.odg":            "application/vnd.oasis.opendocument.graphics",
	"odp.odp":            "application/vnd.oasis.opendocument.presentation",
	"ods.ods":            "application/vnd.oasis.opendocument.spreadsheet",
	"odt.odt":            "application/vnd.oasis.opendocument.text",
	"ogg.oga":            "audio/ogg",
	"ogg.ogv":            "video/ogg",
	"ogg.spx.oga":        "audio/ogg",
	"otf.otf":            "font/otf",
	"otg.otg":            "application/vnd.oasis.opendocument.graphics-template",
	"otp.otp":            "application/vnd.oasis.opendocument.presentation-template",
	"ots.ots":            "application/vnd.oasis.opendocument.spreadsheet-template",
	"ott.ott":            "application/vnd.oasis.opendocument.text-template",
	"odc.odc":            "application/vnd.oasis.opendocument.chart",
	"owl2.owl":           "application/owl+xml",
	"pat.pat":            "image/x-gimp-pat",
	"pdf.pdf":            "application/pdf",
	"php.php":            "text/x-php",
	"pl.pl":              "text/x-perl",
	"png.png":            "image/png",
	"ppt.ppt":            "application/vnd.ms-powerpoint",
	"pptx.pptx":          "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"ps.ps":              "application/postscript",
	"psd.psd":            "image/vnd.adobe.photoshop",
	"p7s_pem.p7s":        "application/pkcs7-signature",
	"p7s_der.p7s":        "application/pkcs7-signature",
	"pub.pub":            "application/vnd.ms-publisher",
	"py.py":              "text/x-python",
	"qcp.qcp":            "audio/qcelp",
	"rar.rar":            "application/x-rar-compressed",
	"rmvb.rmvb":          "application/vnd.rn-realmedia-vbr",
	"rpm.rpm":            "application/x-rpm",
	"rss.rss":            "application/rss+xml",
	"rtf.rtf":            "text/rtf",
	"sample32.macho":     "application/x-mach-binary",
	"sample64.macho":     "application/x-mach-binary",
	"shp.shp":            "application/vnd.shp",
	"shx.shx":            "application/vnd.shx",
	"so.so":              "application/x-sharedlib",
	"sqlite.sqlite":      "application/vnd.sqlite3",
	"srt.srt":            "application/x-subrip",
	// not.srt.txt uses periods instead of commas for the decimal separators of
	// the timestamps.
	"not.srt.txt": "text/plain; charset=utf-8",
	// not.srt.2.txt does not specify milliseconds.
	"not.srt.2.txt":  "text/plain; charset=utf-8",
	"svg.1.svg":      "image/svg+xml",
	"svg.svg":        "image/svg+xml",
	"swf.swf":        "application/x-shockwave-flash",
	"tar.tar":        "application/x-tar",
	"tar.gnu.tar":    "application/x-tar",
	"tar.oldgnu.tar": "application/x-tar",
	"tar.posix.tar":  "application/x-tar",
	// tar.star.tar was generated with star 1.6.
	"tar.star.tar":  "application/x-tar",
	"tar.ustar.tar": "application/x-tar",
	"tar.v7.tar":    "application/x-tar",
	// tar.v7-gnu.tar is a v7 tar archive generated with GNU tar 1.29.
	"tar.v7-gnu.tar":  "application/x-tar",
	"tcl.tcl":         "text/x-tcl",
	"tcx.tcx":         "application/vnd.garmin.tcx+xml",
	"tiff.tiff":       "image/tiff",
	"torrent.torrent": "application/x-bittorrent",
	"tsv.tsv":         "text/tab-separated-values",
	"ttc.ttc":         "font/collection",
	"ttf.ttf":         "font/ttf",
	"tzfile":          "application/tzif",
	"utf16bebom.txt":  "text/plain; charset=utf-16be",
	"utf16lebom.txt":  "text/plain; charset=utf-16le",
	"utf32bebom.txt":  "text/plain; charset=utf-32be",
	"utf32lebom.txt":  "text/plain; charset=utf-32le",
	"utf8.txt":        "text/plain; charset=utf-8",
	"utf8ctrlchars":   "application/octet-stream",
	"vcf.dos.vcf":     "text/vcard",
	"vcf.vcf":         "text/vcard",
	"voc.voc":         "audio/x-unknown",
	"vtt.vtt":         "text/vtt",
	"vtt.space.vtt":   "text/vtt",
	"vtt.tab.vtt":     "text/vtt",
	"vtt.eof.vtt":     "text/vtt",
	"warc.warc":       "application/warc",
	"wasm.wasm":       "application/wasm",
	"wav.wav":         "audio/wav",
	"webm.webm":       "video/webm",
	"webp.webp":       "image/webp",
	"woff.woff":       "font/woff",
	"woff2.woff2":     "font/woff2",
	"x3d.x3d":         "model/x3d+xml",
	"xar.xar":         "application/x-xar",
	"xcf.xcf":         "image/x-xcf",
	"xfdf.xfdf":       "application/vnd.adobe.xfdf",
	"xlf.xlf":         "application/x-xliff+xml",
	"xls.xls":         "application/vnd.ms-excel",
	"xlsx.1.xlsx":     "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"xlsx.2.xlsx":     "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"xlsx.xlsx":       "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"xml.xml":         "text/xml; charset=utf-8",
	"xml.withbr.xml":  "text/xml; charset=utf-8",
	"xz.xz":           "application/x-xz",
	"zip.zip":         "application/zip",
	"zst.zst":         "application/zstd",
}

func GetFileMime(filePath string) (*mimetype.MIME, error) {
	mime, err := mimetype.DetectFile(filePath)
	if err != nil {
		return nil, err
	}
	return mime, nil
}

func CheckMimeByExt(filePath, ext string) bool {
	mime, err := GetFileMime(filePath)
	if err != nil {
		return false
	}
	targetMime, ok := mimeTypeMap[strings.Join([]string{ext, ext}, ".")]
	if !ok {
		return false
	}

	fmt.Println(targetMime)

	return mime.Is(targetMime)
}

func GetFileBufferMime(reader io.Reader) (*mimetype.MIME, error) {
	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return nil, err
	}
	return mime, nil
}

func CheckFileBufferMimeByExt(reader io.Reader, ext ...string) bool {
	mime, err := GetFileBufferMime(reader)
	if err != nil {
		return false
	}

	var isSuccess bool
	for _, e := range ext {
		targetMime, ok := mimeTypeMap[strings.Join([]string{e, e}, ".")]
		if !ok {
			return false
		}
		if mime.Is(targetMime) {
			isSuccess = true
			break
		}
	}
	return isSuccess
}
