package web

type ContentType struct {
	ApplicationJson               any
	ApplicationXml                any
	ApplicationXWWWFormURLEncoded any
	ApplicationJavaScript         any
	ApplicationPdf                any
	ApplicationZip                any

	TextHtml       any
	TextPlain      any
	TextCss        any
	TextCsv        any
	TextJavaScript any

	ImageJpeg   any
	ImagePng    any
	ImageGif    any
	ImageSvgXml any
	ImageWebp   any

	AudioMpeg any
	AudioOgg  any
	AudioWav  any

	VideoMp4  any
	VideoMpeg any
	VideoOgg  any

	MultipartFormData any
}

func (c *ContentType) Iterate() func(func(contentType string, value any) bool) {
	return func(yield func(contentType string, value any) bool) {
		if c.ApplicationJson != nil {
			if !yield("application/json", c.ApplicationJson) {
				return
			}
		}
		if c.ApplicationXml != nil {
			if !yield("application/xml", c.ApplicationXml) {
				return
			}
		}
		if c.ApplicationXWWWFormURLEncoded != nil {
			if !yield("application/x-www-form-urlencoded", c.ApplicationXWWWFormURLEncoded) {
				return
			}
		}
		if c.ApplicationJavaScript != nil {
			if !yield("application/javascript", c.ApplicationJavaScript) {
				return
			}
		}
		if c.ApplicationPdf != nil {
			if !yield("application/pdf", c.ApplicationPdf) {
				return
			}
		}
		if c.ApplicationZip != nil {
			if !yield("application/zip", c.ApplicationZip) {
				return
			}
		}

		if c.TextHtml != nil {
			if !yield("text/html", c.TextHtml) {
				return
			}
		}
		if c.TextPlain != nil {
			if !yield("text/plain", c.TextPlain) {
				return
			}
		}
		if c.TextCss != nil {
			if !yield("text/css", c.TextCss) {
				return
			}
		}
		if c.TextCsv != nil {
			if !yield("text/csv", c.TextCsv) {
				return
			}
		}
		if c.TextJavaScript != nil {
			if !yield("text/javascript", c.TextJavaScript) {
				return
			}
		}

		if c.ImageJpeg != nil {
			if !yield("image/jpeg", c.ImageJpeg) {
				return
			}
		}
		if c.ImagePng != nil {
			if !yield("image/png", c.ImagePng) {
				return
			}
		}
		if c.ImageGif != nil {
			if !yield("image/gif", c.ImageGif) {
				return
			}
		}
		if c.ImageSvgXml != nil {
			if !yield("image/svg+xml", c.ImageSvgXml) {
				return
			}
		}
		if c.ImageWebp != nil {
			if !yield("image/webp", c.ImageWebp) {
				return
			}
		}

		if c.AudioMpeg != nil {
			if !yield("audio/mpeg", c.AudioMpeg) {
				return
			}
		}
		if c.AudioOgg != nil {
			if !yield("audio/ogg", c.AudioOgg) {
				return
			}
		}
		if c.AudioWav != nil {
			if !yield("audio/wav", c.AudioWav) {
				return
			}
		}

		if c.VideoMp4 != nil {
			if !yield("video/mp4", c.VideoMp4) {
				return
			}
		}
		if c.VideoMpeg != nil {
			if !yield("video/mpeg", c.VideoMpeg) {
				return
			}
		}
		if c.VideoOgg != nil {
			if !yield("video/ogg", c.VideoOgg) {
				return
			}
		}

		if c.MultipartFormData != nil {
			if !yield("multipart/form-data", c.MultipartFormData) {
				return
			}
		}
	}
}

func isContentType(t any) *ContentType {
	if t == nil {
		return nil
	}
	switch c := t.(type) {
	case ContentType:
		return &c
	case *ContentType:
		return c
	default:
		return nil
	}
}
